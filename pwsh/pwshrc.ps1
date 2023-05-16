

# (Invoke-WebRequest -Uri "https://raw.githubusercontent.com/JiangKlijna/my-workspaces/master/pwsh/pwshrc.ps1").Content > $profile

function Show-File([int]$col = 5) {
    Get-ChildItem | Format-Wide -Column $col -Property Name
}

function Get-Voice {
    param([string]$text)
    (New-Object -ComObject Sapi.SpVoice).Speak($text) | out-null
}

function Get-Goto([string]$text) {
    Start-Process -FilePath $text
}

function Get-Baidu([string]$text) {
    goto https://www.baidu.com/s?wd=$text
}

function Get-Bing([string]$text) {
    goto https://cn.bing.com/search?q=$text
}

function Get-Github([string]$text) {
    goto https://github.com/search?q=$text
}

function Get-Google([string]$text) {
    goto https://www.google.com.hk/search?q=$text
}

function Get-P([string]$link) {
    $uri=([uri]$link)
    #$uri=[System.Uri]::new($link)
    if ($uri.IsAbsoluteUri) {
        $urihost = $uri.Host
    } else {
        $urihost = $link
    }
    echo $urihost
    ping $urihost
}

function Get-UUID([int]$size = 1, [string]$format="N") {
    for ($i = 1; $i -le $size; $i++) {
        [guid]::NewGuid().ToString($format)
    }
}

function Merge-Files {
    param (
        [string]$FolderPath,
        [string]$FileSuffix,
        [string]$FilePrefix
    )
    $currentDirectory = Get-Location
    $searchPattern = "$FilePrefix*$FileSuffix"
    $files = Get-ChildItem $FolderPath -Filter $searchPattern

    if ($files.Count -eq 0) {
        Write-Host "No matching files found."
        return
    }
    $outputFileNamePrefix = (Get-Item -Path $folderPath).Name
    $filePaths = Get-ChildItem -Path $folderPath -Filter "$filePrefix*$fileSuffix" | Select-Object -ExpandProperty FullName
    $newFilePath = Join-Path -Path $currentDirectory -ChildPath "$outputFileNamePrefix.$FileSuffix"
    $content = Get-Content -Path $filePaths | Out-String

    Set-Content -Path $newFilePath -Value $content
}

function Get-HashSum {
    param(
        [string] $Path,
        [string] $Al = "md5"
    )
    $hash = Get-FileHash -Path $Path -Algorithm $Al | Select-Object -ExpandProperty Hash
    return $hash
}


function Get-Hash {
    param(
        [string] $Al,
        [string] $InputString
    )
    $bytes = [System.Text.Encoding]::UTF8.GetBytes($InputString)
    $hashAlgorithm = [System.Security.Cryptography.HashAlgorithm]::Create($Al.ToLower())
    $hash = [System.BitConverter]::ToString($hashAlgorithm.ComputeHash($bytes)).Replace('-', '').ToLower()
    return $hash
}


Remove-Item alias:ls
Set-Alias ll Get-ChildItem
Set-Alias ls Show-File
