# generate-mtls.ps1
param(
    [string]$CN = "localhost",
    [string]$OutDir = "mtls-certs",
    [int]$Days = 3650,
    [string]$OpenSSL = "openssl"
)

if (-not (Get-Command $OpenSSL -ErrorAction SilentlyContinue)) {
    Write-Error "OpenSSL not found: $OpenSSL"
    exit 1
}

New-Item -ItemType Directory -Force -Path $OutDir | Out-Null
Set-Location $OutDir
$OutDirFull = (Get-Location).Path
$Years = [math]::Round($Days / 365, 1)

Write-Host "========================================"
Write-Host "  mTLS Certificate Generator"
Write-Host "========================================"
Write-Host ""
Write-Host "CN: $CN"
Write-Host "Days: $Days (~ $Years years)"
Write-Host "Output: $OutDirFull"
Write-Host ""

# 1. CA
Write-Host "[1/6] Generating CA..."
& $OpenSSL genrsa -out ca.key 4096 2>$null
& $OpenSSL req -new -x509 -key ca.key -out ca.crt -days $Days -subj "/CN=JiangCA/O=JiangOrg" 2>$null
Write-Host "      Done: ca.key, ca.crt"

# 2. Server
Write-Host "[2/6] Generating server cert..."
& $OpenSSL genrsa -out server.key 2048 2>$null
& $OpenSSL req -new -key server.key -out server.csr -subj "/CN=$CN/O=JiangOrg" 2>$null

Set-Content -Path "server.ext" -Value "basicConstraints=CA:FALSE"
Add-Content -Path "server.ext" -Value "keyUsage = digitalSignature, keyEncipherment"
Add-Content -Path "server.ext" -Value "extendedKeyUsage = serverAuth"
Add-Content -Path "server.ext" -Value "subjectAltName = @alt_names"
Add-Content -Path "server.ext" -Value ""
Add-Content -Path "server.ext" -Value "[alt_names]"

$ipAddr = $null
$isIP = [System.Net.IPAddress]::TryParse($CN, [ref]$ipAddr)

if ($isIP) {
    Add-Content -Path "server.ext" -Value "DNS.1 = localhost"
    Add-Content -Path "server.ext" -Value "IP.1 = $CN"
    Add-Content -Path "server.ext" -Value "IP.2 = 127.0.0.1"
    Add-Content -Path "server.ext" -Value "IP.3 = ::1"
} else {
    Add-Content -Path "server.ext" -Value "DNS.1 = $CN"
    Add-Content -Path "server.ext" -Value "DNS.2 = localhost"
    Add-Content -Path "server.ext" -Value "IP.1 = 127.0.0.1"
    Add-Content -Path "server.ext" -Value "IP.2 = ::1"
}

& $OpenSSL x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days $Days -extfile server.ext 2>$null
Write-Host "      Done: server.key, server.crt"

# 3. Client
Write-Host "[3/6] Generating client cert..."
& $OpenSSL genrsa -out client.key 2048 2>$null
& $OpenSSL req -new -key client.key -out client.csr -subj "/CN=client1/O=JiangOrg" 2>$null

Set-Content -Path "client.ext" -Value "basicConstraints=CA:FALSE"
Add-Content -Path "client.ext" -Value "keyUsage = digitalSignature"
Add-Content -Path "client.ext" -Value "extendedKeyUsage = clientAuth"

& $OpenSSL x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days $Days -extfile client.ext 2>$null
Write-Host "      Done: client.key, client.crt"

# 4. PFX with random password
Write-Host "[4/6] Packing client PFX..."

$chars = @{
    Upper = 65..90 | ForEach-Object { [char]$_ }
    Lower = 97..122 | ForEach-Object { [char]$_ }
    Digit = 48..57 | ForEach-Object { [char]$_ }
    Special = @('!','@','#','$','%','^','&','*','-','+','=','?')
}
$password = @()
$password += $chars.Upper | Get-Random -Count 4
$password += $chars.Lower | Get-Random -Count 4
$password += $chars.Digit | Get-Random -Count 4
$password += $chars.Special | Get-Random -Count 4
$password += ($chars.Upper + $chars.Lower + $chars.Digit + $chars.Special) | Get-Random -Count 8
$password = $password | Get-Random -Count $password.Count
$plainPwd = -join $password

& $OpenSSL pkcs12 -export -in client.crt -inkey client.key -certfile ca.crt -out client.pfx -name "JiangClientCert" -passout "pass:$plainPwd" 2>$null
$plainPwd | Out-File -FilePath "client.pfx.pwd" -NoNewline -Encoding UTF8
Write-Host "      Done: client.pfx, client.pfx.pwd"

# 5. Chain
Write-Host "[5/6] Creating server chain..."
Copy-Item server.crt server-chain.crt
Add-Content server-chain.crt (Get-Content ca.crt)
Write-Host "      Done: server-chain.crt"

# Cleanup
Remove-Item -Force server.csr, client.csr, server.ext, client.ext, ca.srl -ErrorAction SilentlyContinue

Write-Host ""
Write-Host "========================================"
Write-Host "  Done!"
Write-Host "========================================"
Write-Host ""
Write-Host "Files:"
Get-ChildItem -Path $OutDirFull -File | ForEach-Object {
    $sizeKB = [math]::Round($_.Length / 1KB, 1)
    Write-Host ("  {0,-20} {1,8} KB" -f $_.Name, $sizeKB)
}
Write-Host ""
Write-Host "PFX Password: $plainPwd"
Write-Host ""
Write-Host "File Usage:"
Write-Host "  ca.key              CA Private Key [SECRET]"
Write-Host "  ca.crt              CA Certificate [Public]"
Write-Host "  server.key          Server Private Key [SECRET]"
Write-Host "  server.crt          Server Certificate [Public]"
Write-Host "  server-chain.crt    Server Chain [Public]"
Write-Host "  client.key          Client Private Key [SECRET]"
Write-Host "  client.crt          Client Certificate [Public]"
Write-Host "  client.pfx          Client PFX Package [Restricted]"
Write-Host "  client.pfx.pwd      PFX Password [SECRET]"
Write-Host ""
Write-Host "Browser Setup:"
Write-Host "  1. Import ca.crt to 'Trusted Root Certification Authorities'"
Write-Host "  2. Import client.pfx to 'Personal' (password in client.pfx.pwd)"
Write-Host ""
Write-Host "Nginx Config:"
Write-Host "  ssl_certificate      server-chain.crt;"
Write-Host "  ssl_certificate_key  server.key;"
Write-Host "  ssl_verify_client    on;"
Write-Host "  ssl_client_certificate ca.crt;"
Write-Host ""