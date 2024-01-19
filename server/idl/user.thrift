namespace go user
include "base.thrift"

struct user_login_request{
    1: string username;
    2: string password;
}

struct user_login_response{
    1: base.base_response base_resp
    2: i64 user_id
    3: string token
}

struct user_register_request{
    1: string username;
    2: string password;
}

struct user_register_response{
    1: base.base_response base_resp
    2: i64 user_id
    3: string token
}

struct get_user_info_request{
     1: string user_name
     2: i64 user_id
}

struct get_user_info_response{
     1: base.base_response base_resp
     2: base.User user_info
}

service UserService{
    user_login_response UserLogin(1:user_login_request req);
    user_register_response UserRegister(1:user_register_request req);
    get_user_info_response GetUserInfo(1: get_user_info_request req);
}
