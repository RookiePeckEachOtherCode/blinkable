namespace go user

include "base.thrift"

struct user_login_request {
    1: string username;
    2: string password;
}

struct user_login_response {
    1: base.base_response base_resp
    2: i64                user_id
    3: string             token
}

struct user_register_request {
    1: string username;
    2: string password;
}

struct user_register_response {
    1: base.base_response base_resp
    2: i64                user_id
    3: string             token
}

struct get_user_info_request {
    1: string user_name
    2: i64    user_id
    3: string token
    4: i32    tp
}

struct get_user_info_response {
    1: base.base_response base_resp
    2: base.User          user_info
}

struct update_user_info_request {
    1: i64    user_id
    2: string token
    3: string username
    4: string signature
    5: string title
    6: string github_url
}

struct update_user_info_response {
    1: base.base_response base_resp
}

struct update_user_password_request {
    1: i64    user_id
    2: string token
    3: string newpassword
    4: string oldpassword
}

struct update_user_password_response {
    1: base.base_response base_resp
}

struct upload_user_icon_request{
    1:i64 user_id
    2:string icon_url
}

struct upload_user_icon_response{
    1:base.base_response base_resp
}

struct upload_user_back_request{
    1:i64 user_id
    2:string back_url
}

struct upload_user_back_response{
    1:base.base_response base_resp
}
struct be_admin_request{
    1:i64 user_id
}

struct be_admin_response{
    1:base.base_response base_resp
}

service UserService {
    user_login_response UserLogin(1: user_login_request req);
    user_register_response UserRegister(1: user_register_request req);
    get_user_info_response GetUserInfo(1: get_user_info_request req);
    update_user_info_response UpdateUserInfo(1: update_user_info_request req);
    update_user_password_response UpdateUserPassword(1: update_user_password_request req);
    upload_user_icon_response UploadUserIcon(1:upload_user_icon_request req);
    upload_user_back_response UploadUserBack(1:upload_user_back_request req);
    be_admin_response BeAdmin(1:be_admin_request req);

}
