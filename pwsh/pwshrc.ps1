

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


Remove-Item alias:ls
Set-Alias ll Get-ChildItem
Set-Alias ls Show-File
