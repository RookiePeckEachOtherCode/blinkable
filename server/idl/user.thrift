namespace go user
include "base.thrift"

struct UserLoginRequest{
    1: string username;
    2: string password;
}

struct UserLoginResponse{
    1: base.base_response base_resp
    2: i64 user_id
    3: string token
}

struct UserRegisterRequest{
    1: string username;
    2: string password;
}

struct UserRegisterResponse{
    1: base.base_response base_resp
    2: i64 user_id
    3: string token
}

struct GetUserInfoRequest{
     1: string user_name
     2: i64 user_id
}

struct GetUserInfoResponse{
     1: base.base_response base_resp
     2: base.User user_info
}

service UserService{
    UserLoginResponse UserLogin(1:UserLoginRequest req);
    UserRegisterResponse UserRegister(1:UserRegisterRequest req);
    GetUserInfoResponse GetUserInfo(1: GetUserInfoRequest req);
}
