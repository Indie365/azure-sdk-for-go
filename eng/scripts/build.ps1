#Requires -Version 7.0
param([string]$filter, [switch]$clean, [switch]$vet, [switch]$generate, [switch]$skipBuild, [string]$config = "autorest.md", [string]$outputFolder)

. $PSScriptRoot/meta_generation.ps1
. $PSScriptRoot/get_module_dirs.ps1


function Process-Sdk ($path) {
    if ($clean) {
        Write-Host "##[command]Executing go clean -v ./... in " $path
        go clean -v ./...
    }

    if ($generate) {
        Write-Host "##[command]Executing autorest.go in " $path
        $autorestPath = $path + "/" + $config

        if (ShouldGenerate-AutorestConfig $autorestPath) {
            Generate-AutorestConfig $autorestPath
            $removeAutorestFile = $true
        }

        $autorestVersion = "@autorest/go@4.0.0-preview.23"
        if ($outputFolder -eq '') {
            $outputFolder = $path
        }
        autorest --use=$autorestVersion --go --track2 --go-sdk-folder=$root --output-folder=$outputFolder --file-prefix="zz_generated_" --clear-output-folder=false $autorestPath
        if ($removeAutorestFile) {
            Remove-Item $autorestPath
        }
    }

    if (!$skipBuild) {
        Write-Host "##[command]Executing go build -v ./... in " $path
        go build -x -v ./...
        Write-Host "##[command]Build Complete!"

    }

    if ($vet) {
        Write-Host "##[command]Executing go vet ./... in " $path
        go vet ./...
    }

}

$startingDirectory = Get-Location
$root = Resolve-Path ($PSScriptRoot + "/../..")
Set-Location $root
$sdks = @{};

foreach ($sdk in (Get-ModuleDirs 'sdk/...')) {
    $name = $sdk | split-path -leaf
    $sdks[$name] = @{
        'path' = $sdk;
    }
}

$keys = $sdks.Keys | Sort-Object;
if (![string]::IsNullOrWhiteSpace($filter)) {
    Write-Host "Using filter: $filter"
    $keys = $keys.Where( { $_ -match $filter })
}

try {
    $keys | ForEach-Object { $sdks[$_] } | ForEach-Object {
        Push-Location $_.path
        Process-Sdk $_.path
        Pop-Location
    }
}
finally {
    Set-Location $startingDirectory
}