if (!(Get-Command protoc -ErrorAction SilentlyContinue)) {
    Write-Error '错误：protoc命令未安装或者不在环境变量'
}

if (!(Get-Command kitex -ErrorAction SilentlyContinue)) {
    Write-Error '错误：kitex命令未安装或者不在环境变量'
    Write-Host '尝试安装kitex'
    go install github.com\cloudwego\kitex\tool\cmd\kitex@latest
}

if (!(Get-Command kitex -ErrorAction SilentlyContinue)) {
    Write-Error '安装kitex失败，请手动安装'
    exit 1
}

mkdir -p kitex_gen
kitex -module "blinkable" -I idl/ idl/"$1".thrift

mkdir -p cmd/"$1"
cd cmd/"$1" && kitex -module "blinkable" -service "$1" -use blinkable/kitex_gen/ -I ../../idl/ ../../idl/"$1".thrift

go mod tidy
