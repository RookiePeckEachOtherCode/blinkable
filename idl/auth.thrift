namespace go auth

struct UserRegisterRequests{
    1: string username
    2: string password
}

struct UserRegisterResponse{
    1: i32 status_code
    2: string status_msg
    3: string token
}

