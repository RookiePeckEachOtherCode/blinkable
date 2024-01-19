namespace go api

include "base.thrift"

struct user_login_request {
    1: string username (api.query = "username", api.vd = "len($)>0 && len($)<33")
    2: string password (api.query = "password", api.vd = "len($)>0 && len($)<33")
}

struct user_login_response {
    1: string             token
    2: i64                user_id
    3: base.base_response base_resp
}

struct user_register_request {
    1: string username (api.query = "username", api.vd = "len($)>0 && len($)<33")
    2: string password (api.query = "password", api.vd = "len($)>0 && len($)<33")
}

struct user_register_response {
    1: string             token
    2: i64                user_id
    3: base.base_response base_resp
}

struct get_user_info_request {
    1: i64    user_id
    2: string token
}

struct get_user_info_response {
    1: base.base_response base_resp
    2: base.User          user
}

service ApiService {
    user_login_response UserLogin(1: user_login_request req) (api.post = "/blinkable/user/login")
    user_register_response UserRegister(1: user_register_request req) (api.post = "/blinkable/user/register")
    get_user_info_response GetUserInfo(1: get_user_info_request req) (api.get = "/blinkable/user/info")
}
