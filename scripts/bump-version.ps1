param(
    [ValidateSet("minor", "major", "patch")]
    [string]$kind = "minor"
)

$version = Get-Content -Path "VERSION"

$versionParts = $version.Split(".")
if ($kind -eq "major") {
    $versionParts[0] = [int]$versionParts[0] + 1
    $versionParts[1] = 0
    $versionParts[2] = 0
}
elseif ($kind -eq "minor") {
    $versionParts[1] = [int]$versionParts[1] + 1
    $versionParts[2] = 0
}
elseif ($kind -eq "patch") {
    $versionParts[2] = [int]$versionParts[2] + 1
}
$version = $versionParts -join "."
$version | Out-File "VERSION"