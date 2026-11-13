<#
.SYNOPSIS
    查找本机所有 Java 进程，逐个触发一次 GC，并显示 GC 前后的堆占用。

.DESCRIPTION
    等价于 jconsole 里逐个点“执行 GC”，但一次跑完所有进程。
    底层用 JDK 自带的 jcmd GC.run（同步，返回即完成），jstat 读堆占用。

    说明：GC 只清理堆内对象，不会把内存还给操作系统。
          进程的 committed 内存由 JVM 自行管理，需要时会复用。

.EXAMPLE
    .\gc-all.ps1
#>

# 中文输出：统一用 UTF-8，避免默认 GB2312 控制台下乱码
try { [Console]::OutputEncoding = [System.Text.Encoding]::UTF8 } catch { }

# ---------- 定位 JDK 工具 ----------
function Find-JdkTool {
    param([string]$Name)

    $candidates = @()
    if ($env:JAVA_HOME) {
        $candidates += Join-Path $env:JAVA_HOME "bin\$Name.exe"
    }
    $cmd = Get-Command $Name -ErrorAction SilentlyContinue
    if ($cmd) { $candidates += $cmd.Source }

    # 兜底：从正在运行的 java 进程路径推断 JDK 目录
    foreach ($p in (Get-Process -Name java, javaw -ErrorAction SilentlyContinue)) {
        try {
            if ($p.Path) { $candidates += Join-Path (Split-Path $p.Path) "$Name.exe" }
        } catch { }
    }

    foreach ($c in $candidates) {
        if ($c -and (Test-Path -LiteralPath $c)) { return (Resolve-Path -LiteralPath $c).Path }
    }
    return $null
}

$jcmd  = Find-JdkTool 'jcmd'
$jstat = Find-JdkTool 'jstat'

if (-not $jcmd) {
    Write-Host "找不到 jcmd.exe。请安装 JDK，或设置 JAVA_HOME / 把 JDK 的 bin 加入 PATH。" -ForegroundColor Red
    exit 1
}

# ---------- 读取堆占用（KB） ----------
# jstat -gc 的列位置随 JDK / GC 变化，按表头名字定位；"-" 记 0。
function Get-HeapUsedKB {
    param([int]$ProcessId)

    if (-not $jstat) { return $null }
    try {
        $lines = & $jstat -gc $ProcessId 2>$null
        if ($LASTEXITCODE -ne 0 -or -not $lines) { return $null }
        $lines = @($lines | Where-Object { $_ -and $_.Trim() })
        if ($lines.Count -lt 2) { return $null }

        $header = $lines[0].Trim() -split '\s+'
        $values = $lines[-1].Trim() -split '\s+'
        if ($header.Count -ne $values.Count) { return $null }

        $used = 0.0
        foreach ($col in 'S0U', 'S1U', 'EU', 'OU') {
            $i = [Array]::IndexOf($header, $col)
            if ($i -lt 0) { continue }
            if ($values[$i] -match '^[\d.]+$') { $used += [double]$values[$i] }
        }
        return $used
    } catch { return $null }
}

function Format-Size {
    param($KB)
    if ($null -eq $KB) { return 'N/A' }
    $sign = if ($KB -lt 0) { '-' } else { '' }
    $v = [Math]::Abs($KB)
    if ($v -ge 1048576) { return '{0}{1:N2} GB' -f $sign, ($v / 1048576) }
    return '{0}{1:N1} MB' -f $sign, ($v / 1024)
}

# ---------- 枚举 Java 进程 ----------
# 只用 jcmd -l 会漏掉内嵌 JVM 的原生启动器进程（如 DBeaver / eclipse.exe，
# 进程名不是 java.exe）：这类进程 jps 能列出、jcmd -l 列不出，但按 PID 仍可 attach。
# 故取 jps -l 与 jcmd -l 的并集，按 PID 去重。
$jps = Find-JdkTool 'jps'

$seen = @{}
$targets = @()
foreach ($tool in @($jps, $jcmd)) {
    if (-not $tool) { continue }
    foreach ($line in (& $tool -l 2>$null)) {
        if ($line -notmatch '^\s*(\d+)\s*(.*)$') { continue }
        $procId = [int]$Matches[1]
        $desc   = $Matches[2].Trim()

        # 跳过 jps / jcmd 自身临时启动的进程
        if ($desc -match 'sun\.tools\.(jcmd\.JCmd|jps\.Jps)|jdk\.jcmd') { continue }
        if ($seen.ContainsKey($procId)) { continue }
        $seen[$procId] = $true

        $name = $desc.Split(' ')[0]
        if ($name -match '\.jar$') { $label = Split-Path $name -Leaf }
        elseif ($name)             { $label = ($name -split '\.')[-1] }
        else {
            # 描述为空（内嵌 JVM 常见）：退回用系统进程名，便于辨认
            $pn = (Get-Process -Id $procId -ErrorAction SilentlyContinue).ProcessName
            $label = if ($pn) { $pn } else { "PID $procId" }
        }
        if ($label.Length -gt 30) { $label = $label.Substring(0, 27) + '...' }

        $targets += [pscustomobject]@{ Id = $procId; Label = $label }
    }
}

if ($targets.Count -eq 0) {
    Write-Host "没有发现可 attach 的 Java 进程。" -ForegroundColor Yellow
    Write-Host "（加了 -XX:+DisableAttachMechanism 或 -XX:-UsePerfData 的进程不会出现在列表里）" -ForegroundColor DarkGray
    exit 0
}

Write-Host "发现 $($targets.Count) 个 Java 进程，开始触发 GC ..." -ForegroundColor Cyan
Write-Host ""

# ---------- 逐个 GC ----------
$results = foreach ($t in $targets) {
    $before = Get-HeapUsedKB -ProcessId $t.Id
    $status = '成功'
    $detail = ''

    try {
        $text = (& $jcmd $t.Id GC.run 2>&1 | Out-String)
        if ($LASTEXITCODE -ne 0 -or $text -match 'Exception|Error:|no such process') {
            $status = '失败'
            $detail = (($text -split "`n" | Where-Object { $_ -match 'Exception|Error' } |
                        Select-Object -First 1) -replace '^\s*\d+:\s*', '').Trim()
            if (-not $detail) { $detail = ($text.Trim() -split "`n")[0] }
        }
    } catch {
        $status = '失败'
        $detail = $_.Exception.Message
    }

    Start-Sleep -Milliseconds 150
    $after = Get-HeapUsedKB -ProcessId $t.Id

    # 负值 = GC 后占用反而上升（对象晋升、attach 触发类加载等），如实显示
    $freed = if ($null -ne $before -and $null -ne $after) { $before - $after } else { $null }

    [pscustomobject]@{
        PID    = $t.Id
        进程   = $t.Label
        GC前   = Format-Size $before
        GC后   = Format-Size $after
        释放   = Format-Size $freed
        状态   = $status
        Detail = $detail
        Freed  = $freed
    }
}

# ---------- 输出 ----------
$results | Format-Table PID, 进程, GC前, GC后, 释放, 状态 -AutoSize |
    Out-String -Width 200 | Write-Host

foreach ($r in $results | Where-Object { $_.状态 -eq '失败' }) {
    Write-Host ("  PID {0} 失败: {1}" -f $r.PID, $r.Detail) -ForegroundColor Red
}
if ($results | Where-Object { $null -ne $_.Freed -and $_.Freed -lt 0 }) {
    Write-Host "  注：释放为负值表示 GC 后堆占用反而上升，通常是存活对象晋升到老年代所致，属正常现象。" -ForegroundColor DarkGray
}

$ok = @($results | Where-Object { $_.状态 -eq '成功' }).Count
$totalFreed = ($results | Where-Object { $null -ne $_.Freed } | Measure-Object -Property Freed -Sum).Sum
Write-Host ""
Write-Host ("完成：成功 {0} / {1}，净释放 {2}" -f $ok, $results.Count, (Format-Size $totalFreed)) -ForegroundColor Green
