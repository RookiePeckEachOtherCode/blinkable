# blinkable

```
.
├── LICENSE
├── README.md
├── add-service.sh      //服务生成脚本
├── cmd                 //微服务
├── common              //公共模块，如一些常量
├── configs             //配置文件
├── dal                 //数据访问层
├── doc                 //文档
├── docker-compose.yml  
├── go.mod 
├── go.sum
├── idl                 //接口定义
├── kitex_gen
├── pkg                 //公共模块
└── web                 //web服务
```

定义完接口文件后，使用add-service.sh生成服务代码，如：
>windown下用 git bash 作为终端就能运行 bash 脚本了
```bash
sh ./add-service.sh 服务名
```
服务名就是idl/文件夹下定义的文件名

例如
定义了 idl/user.thrift 文件

添加user服务
```bash
sh ./add-service.sh user
```
