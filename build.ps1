# Root-relative paths
$cliPath = ".\cli"
$outputPath = ".\packages\umbr-zip\bin"

# Check if CLI folder exists
if (-Not (Test-Path $cliPath)) {
    Write-Error "Error: CLI folder '$cliPath' does not exist."
    exit 1
}

# Ensure output directory exists
if (-Not (Test-Path $outputPath)) {
    New-Item -ItemType Directory -Path $outputPath -Force | Out-Null
}

# Platforms to build for
$platforms = @{
    "windows" = ".exe"
    "linux"   = ""
    "darwin"  = ""
}

# Build for each platform
foreach ($platform in $platforms.Keys) {
    $ext = $platforms[$platform]
    $outputFile = "go-zip-$platform$ext"
    Write-Host "Building for $platform..."

    # Set environment variables and run go build
    $env:GOOS = $platform
    $env:GOARCH = "amd64"
    go build -o "$cliPath\$outputFile" $cliPath

    # Move the binary to output directory
    Move-Item "$cliPath\$outputFile" -Destination "$outputPath" -Force
}

Write-Host "Build completed. Binaries are in $outputPath"
