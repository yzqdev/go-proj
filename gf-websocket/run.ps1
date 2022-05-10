

param (

    [String]$Type = "run"
)
$name="bpass-server"
$curDir= Split-Path -Parent $MyInvocation.MyCommand.Definition
Write-Host "当前路径"$curDir   -ForegroundColor Yellow
if ($Type -eq "run") {
    Write-Host "运行" -ForegroundColor Cyan

        gf run  main.go -p bin


}
elseif ($Type -eq "build") {
    Write-Host "编译" -ForegroundColor Red


        gf build main.go


}