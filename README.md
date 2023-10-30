# blinkable

```
.
├── LICENSE
├── README.md
├── add-service.sh          //服务生成脚本
├── cmd                     //微服务
│   ├── api                //api服务
│   └── user               //user服务
├── common                  //公共模块，如一些常量
│   ├── consts             //常量
│   ├── errno              //错误信息
│   └── response           //响应模型 
├── configs                 //配置文件
├── dal                     //数据访问层
│   └── db
│       ├── gen            // gorm gen模式
│       ├── init.go        //初始化
│       ├── migrate        //根据model自动创建表
│       ├── model          //model
│       └── query          // gorm gen生成的代码
│
├── doc                     //文档
├── docker-compose.yml  
├── go.mod 
├── go.sum
├── idl                     //接口定义
├── kitex_gen
├── pkg                     //公共模块
│   ├── hash               //密码加密
│   ├── jwt                //token
│   ├── minio              //对象存储
│   ├── viper              //加载配置文件
│   └── zap                //打印日志
└── web                     //web服务
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
