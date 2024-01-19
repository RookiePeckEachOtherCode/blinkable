module_name="blinkable"
server_dir="server/service"
kitex_dir="server/kitex_gen"
idl_dir="server/idl"
if ! command -v thriftgo &>/dev/null; then
	echo '错误：thriftgo 未安装或者不在环境变量'
	echo '尝试安装 thrift go'
	go install github.com/cloudwego/thriftgo@latest
fi
if ! command -v thriftgo &>/dev/null; then
	echo '安装 kitex 失败，请手动安装'
	exit 1
fi
if ! command -v kitex &>/dev/null; then
	echo '错误：kitex 命令未安装或者不在环境变量'
	echo '尝试安装 kitex'
	go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
fi
if ! command -v kitex &>/dev/null; then
	echo '安装 kitex 失败，请手动安装'
	exit 1
fi
mkdir -p $kitex_dir
mkdir -p $server_dir/"$1"
cd $server_dir/"$1" && kitex -module $module_name -service "$1" -gen-path ../../kitex_gen   -I ../../idl/ ../../idl/"$1".thrift
go mod tidy
