$ErrorActionPreference = "Stop"

if (Test-Path ".\dist") {
    Write-Host -ForegroundColor Yellow "Removing: .\dist"
    Remove-Item -Force -Recurse ".\dist"
}

$env:GOARCH="amd64"

$builds = @(
    @{GOOS="darwin"; Output="dist/bup-darwin"},
    @{GOOS="linux"; Output="dist/bup-linux"},
    @{GOOS="windows"; Output="dist/bup.exe"}
)

foreach ($build in $builds) {
    $env:GOOS = $build.GOOS
    Write-Host -ForegroundColor Cyan "Building: $env:GOOS - $($build.Output)"
    go build -o $build.Output
}

Write-Host -ForegroundColor Green "Finished Success"
