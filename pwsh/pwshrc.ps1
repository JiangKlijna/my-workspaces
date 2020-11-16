

# (Invoke-WebRequest -Uri "https://raw.githubusercontent.com/JiangKlijna/my-workspaces/master/pwsh/pwshrc.ps1").Content > $profile

function listfile([int]$col = 5) {
    Get-ChildItem | Format-Wide -Column $col -Property Name
}

function voice {
    param([string]$text)
    (New-Object -ComObject Sapi.SpVoice).Speak($text) | out-null
}

function getjson($f) {
    (Get-Content $f) | ConvertFrom-Json
}

function tojson($obj) {
    ConvertTo-Json -InputObject $obj
}

Remove-Item alias:ls
Set-Alias ll Get-ChildItem
Set-Alias ls listfile
