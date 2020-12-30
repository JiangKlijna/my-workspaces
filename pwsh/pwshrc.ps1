

# (Invoke-WebRequest -Uri "https://raw.githubusercontent.com/JiangKlijna/my-workspaces/master/pwsh/pwshrc.ps1").Content > $profile

function Show-File([int]$col = 5) {
    Get-ChildItem | Format-Wide -Column $col -Property Name
}

function Start-Voice {
    param([string]$text)
    (New-Object -ComObject Sapi.SpVoice).Speak($text) | out-null
}

function Get-Files {
    param([string]$Path)
    $dirs = Get-ChildItem -Path $Path -Directory -Name
    if ($dirs) {
        if ($dirs.GetType() -eq "".GetType()) {
            Get-Files("$Path\$dirs")
        }
        else {
            for ($i = 0; $i -lt $dirs.Count; $i++) {
                $dir = $dirs[$i]
                Get-Files("$Path\$dir")
            }
        }
    }
    $fs = Get-ChildItem -Path $Path -File -Name
    if ($fs) {
        if ($fs.GetType() -eq "".GetType()) {
            Write-Output "$Path\$fs"
        }
        else {
            for ($i = 0; $i -lt $fs.Count; $i++) {
                $f = $fs[$i]
                Write-Output "$Path\$f"
            }
        }
    }
}

Remove-Item alias:ls
Set-Alias ll Get-ChildItem
Set-Alias ls Show-File
