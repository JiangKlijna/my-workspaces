


function listfile([int]$col = 5) {
    Get-ChildItem | Format-Wide -Column $col -Property Name
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
