$ErrorActionPreference = "Stop"

$RepoOwner = "lakshya-8000cr"
$RepoName = "scout"
$Version = "v1.0.0"

$InstallDir = "$env:USERPROFILE\Tools"
$BinaryName = "scout.exe"
$DownloadUrl = "https://github.com/$RepoOwner/$RepoName/releases/download/$Version/$BinaryName"
$BinaryPath = Join-Path $InstallDir $BinaryName

Write-Host "Installing Scout..." -ForegroundColor Cyan

if (!(Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir | Out-Null
}

Write-Host "Downloading Scout from $DownloadUrl"
Invoke-WebRequest -Uri $DownloadUrl -OutFile $BinaryPath

$userPath = [Environment]::GetEnvironmentVariable("Path", "User")

if ($userPath -notlike "*$InstallDir*") {
    [Environment]::SetEnvironmentVariable(
        "Path",
        "$userPath;$InstallDir",
        "User"
    )

    Write-Host "Added Scout to PATH." -ForegroundColor Green
    Write-Host "Please open a new terminal before running scout." -ForegroundColor Yellow
} else {
    Write-Host "Scout is already in PATH." -ForegroundColor Green
}

Write-Host "Scout installed successfully!" -ForegroundColor Green
Write-Host "Run: scout pods" -ForegroundColor Cyan