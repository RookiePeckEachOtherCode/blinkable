moduleName="blinkable"

if ! command -v thriftgo &>/dev/null; then
	echo '错误：protoc命令未安装或者不在环境变量'
	echo '尝试安装thriftgo'
	go install github.com/cloudwego/thriftgo@latest
fi

if ! command -v kitex &>/dev/null; then
	echo '错误：kitex命令未安装或者不在环境变量'
	echo '尝试安装kitex'
	go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
fi

if ! command -v kitex &>/dev/null; then
	echo '安装kitex失败，请手动安装'
	exit 1
fi

mkdir -p kitex_gen
kitex -module ${moduleName} -I idl/ idl/"$1".thrift

mkdir -p cmd/"$1"
cd cmd/"$1" && kitex -module ${moduleName} -service "$1" -use blinkable/kitex_gen/ -I ../../idl/ ../../idl/"$1".thrift

go mod tidy
