# 快速入门
上例子
- 定义一个用户登陆服务

```
namespace go user
//固定格式 namespace go 包名，也就是生成的代码的包名，例如package user

struct UserLoginRequests {//定义登陆请求结构体
//语法
//  n: 类型 变量名

    1: string username
    2: string password
//  3: .......
}

struct UserLoginResponse{//定义登陆响应结构体
    1: string msg       //返回信息
}


server UserService{//定义服务
    UserLoginResponse UserLogin(1:UserLoginRequests req) //定义服务方法
}
```


现在重新定一个用户服务，包含登陆注册

```
namespace go user

struct UserLoginRequests {//定义登陆请求结构体
    1: string username
    2: string password
}

struct UserLoginResponse{//定义登陆响应结构体
    1: string msg       //返回信息
}

struct UserRegisterRequests {//定义注册请求结构体
    1: string username
    2: string password
}

struct UserRegisterResponse{//定义注册响应结构体
    1: string msg       //返回信息
}

server UserService{//定义服务
    UserLoginResponse UserLogin(1:UserLoginRequests req) //定义服务方法
    UserRegisterResponse UserRegister(1:UserRegisterRequests req) //定义服务方法
}
```


## 具体的数据类型
### 基本类型：
- bool: 布尔值
- byte: 8位有符号整数
- i16: 16位有符号整数
- i32: 32位有符号整数
- i64: 64位有符号整数
- double: 64位浮点数
- string: UTF-8编码的字符串
- binary: 二进制串
### 结构体类型：
- struct: 定义的结构体对象
### 容器类型：
- list: 有序元素列表
- set: 无序无重复元素集合
- map: 有序的key/value集合
### 异常类型：
- exception: 异常类型
### 服务类型：
- service: 具体对应服务的类





