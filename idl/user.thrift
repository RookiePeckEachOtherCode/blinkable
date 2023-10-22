namespace go user

struct UserLoginRequsts{
    1: string username;
    2: string password;
}

struct UserLoginResponse{
    1: string token
    2: i32 user_id
    3: string status_msg 
    4: i32 status_code
}

struct UserRegisterRequest{
    1: string username;
    2: string password;
}

struct UseeRegisterResponse{
    1: string token
    2: i32 user_id
    3: string status_msg 
    4: i32 status_code
}

service UserService{
    UserLoginResponse UserLogin(1:UserLoginRequsts req);
    UseeRegisterResponse UserRegister(1:UserRegisterRequest req);
}