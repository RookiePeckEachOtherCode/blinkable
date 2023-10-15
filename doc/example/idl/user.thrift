namespace go user

struct UserLoginRequests{//登录请求
    1: string username,
    2: string password,
}
struct UserLoginResponse{//登录响应
    1: string msg
}

service UserService{//用户服务
    UserLoginResponse UserLogin(1:UserLoginRequests req)
}