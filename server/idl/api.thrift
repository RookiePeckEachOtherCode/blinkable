namespace go api

include "base.thrift"

struct user_login_request {
    1: string username (api.query = "username", api.vd = "len($)>0 && len($)<33")
    2: string password (api.query = "password", api.vd = "len($)>0 && len($)<33")
}

struct user_login_response {
    1: i32    status_code
    2: string status_msg
    3: bool   succed
    4: string token
    5: i64    user_id
}

struct user_register_request {
    1: string username (api.query = "username", api.vd = "len($)>0 && len($)<33")
    2: string password (api.query = "password", api.vd = "len($)>0 && len($)<33")
}

struct user_register_response {
    1: i32    status_code
    2: string status_msg
    3: bool   succed
    4: string token
    5: i64    user_id
}

struct get_user_info_request {
    1: i64    user_id (api.query = "user_id")
    2: string token (api.query = "token")
    3: string user_name (api.query = "user_name")
    4: i32    tp (api.query = "tp")
}

struct get_user_info_response {
    1: i32       status_code
    2: string    status_msg
    3: bool      succed
    4: base.User user
}

struct update_user_info_request {
    1: i64    user_id (api.query = "user_id")
    2: string token (api.query = "token")
    3: string username (api.query = "username", api.vd = "len($)>0 && len($)<33")
    4: string signature (api.query = "signature", api.vd = "len($)>0 && len($)<100")
    5: string title (api.query = "title", api.vd = "len($)>0 && len($)<21")
}

struct update_user_info_response {
    1: i32    status_code
    2: string status_msg
    3: bool   succed
}

struct update_user_password_request {
    1: i64    user_id (api.query = "user_id")
    2: string token (api.query = "token")
    3: string new_passwd (api.query = "new_password")
    4: string old_passwd (api.query = "old_password")
}

struct update_user_password_response {
    1: i32    status_code
    2: string status_msg
    3: bool   succed
}

service ApiService {
    user_login_response UserLogin(1: user_login_request req) (api.post = "/blinkable/user/login")
    user_register_response UserRegister(1: user_register_request req) (api.post = "/blinkable/user/register")
    get_user_info_response GetUserInfo(1: get_user_info_request req) (api.get = "/blinkable/user/info")
    update_user_info_response UpdateUserInfo(1: update_user_info_request req) (api.post = "/blinkable/user/update")
    update_user_password_response UpdateUserPassword(1: update_user_password_request req) (api.post = "/blinkable/user/update/password")
}
