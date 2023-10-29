namespace go user

struct UserLoginRequest{
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

struct UserInfoRequest{
     1: string token
     2: i32 user_id
}

struct UserInfoResponse{
     1: i32 user_id
     2: string username
     3: string avatar
     4: string background_img
     5: string signature
     6: i32 status_code
     7: string status_msg
     8: i32 level
     9: i32 experience
     10: i32 articles_num
}


struct UserInfoUpdateRequest{
    1: i32 user_id
    2: string token
    3: binary  avatar
    4: string avatar_type
    5: binary  background_img
    6: string background_img_type

}

struct UserInfoUpdateResponse{
    1: string status_msg 
    2: i32 status_code
}

service UserService{
    UserLoginResponse UserLogin(1:UserLoginRequest req);
    UseeRegisterResponse UserRegister(1:UserRegisterRequest req);
    UserInfoResponse UserInfo(1: UserInfoRequest req);
    UserInfoUpdateResponse UserInfoUpdate(1: UserInfoUpdateRequest req)
}