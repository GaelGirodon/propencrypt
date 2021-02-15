#
# package.ps1
#
# Build and package propencrypt for Windows
#

New-Item -Name "dist" -ItemType "Directory" -Force

# Build
$env:GOARCH = "amd64"
$env:GOOS = "windows"
go build -ldflags="-s -w" -o "dist/propencrypt.exe" "cmd/propencrypt.go"

# Package
Push-Location "dist"
$archive = "propencrypt_windows_amd64.zip"
Compress-Archive "propencrypt.exe" -DestinationPath "$archive" -CompressionLevel "Optimal"
$hash = (Get-FileHash -Algorithm "SHA256" "$archive").Hash.ToLower()
Set-Content -Path "$archive.sha256" -Value "$hash`n" -Encoding "ascii" -NoNewline
Pop-Location
