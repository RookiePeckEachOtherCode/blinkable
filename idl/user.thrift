namespace go user

struct UserLoginRequsts{
    1: string username;
    2: string password;
}

struct UserLoginResponse{
    1: string token
    2: i32 userId
    3: string status_msg (go.tag="json:'status_msg'")
    4: i32 status_code(go.tag="json:'status_code'")
}

struct UserRegisterRequest{
    1: string username;
    2: string password;
}

struct UseeRegisterResponse{
    1: string token
    2: i32 userId
    3: string status_msg (go.tag="json:'status_msg'")
    4: i32 status_code(go.tag="json:'status_code'")
}

service UserService{
    UserLoginResponse UserLogin(1:UserLoginRequsts userLoginRequsts);
    UseeRegisterResponse UserRegister(1:UserRegisterRequest userRegisterRequest);
}