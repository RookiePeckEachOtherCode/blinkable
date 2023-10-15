namespace go user

struct BaseResponse{
    1: string status_msg (go.tag="json:'status_msg'")
    2: i32 status_code(go.tag="json:'status_code'")
}
struct UserLoginRequsts{
    1: string username;
    2: string password;
}

struct UserLoginResponse{
    1: string token
    2: i32 userId
    3: BaseResponse baseResponse
}

struct UserRegisterRequest{
    1: string username;
    2: string password;
}

struct UseeRegisterResponse{
    1: string token
    2: i32 userId
    3: BaseResponse baseResponse
}

service UserService{
    UserLoginResponse UserLogin(1:UserLoginRequsts userLoginRequsts);
    UseeRegisterResponse UserRegister(1:UserRegisterRequest userRegisterRequest);
}