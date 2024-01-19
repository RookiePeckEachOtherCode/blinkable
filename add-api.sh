if ! command -v hz &>/dev/null; then
    echo '错误: hz 未安装或不在环境变量'
    echo '尝试安装 hz'
    go install github.com/cloudwego/hertz/cmd/hz@latest
fi
if ! command -v hz &>/dev/null; then
	echo '安装 hz 失败，请手动安装'
	exit 1
fi

cd ./server/service/api/
hz update -idl ../../idl/api.thrift
go mod tidy

