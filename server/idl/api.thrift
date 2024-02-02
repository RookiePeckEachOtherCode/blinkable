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

struct Guestbook {
 1:i32 book_id;
 2:i32 user_id;
 3:string context;
 4:i32 formuser_id;
 5:string create_time;
}

struct Users {
 1:i32 admin_id;
 2:string signature;
 3:i32 likes;
 4:string title;
 5:i32 comments;
 6:list<Guestbook> guestbooks;
 7:string icon_url;
 8:string image_url;
 9:string git_url;
}

struct like_request{
 1:i32 admin_id;
 2:i32 user_id;
}

struct get_homepage_request{

}

struct add_guestbook_request{
 1:i32 user_id;
 2:i32 admin_id;
 3:string context;

}

struct get_homepage_response{
1:list<Users> users;
2:i32 status_code;
3:string status_msg;
4:bool succed;
}

struct like_response{
1:i32 status_code;
2:string status_msg;
3:bool succed;
}

struct add_guestbook_response{
1:i32 status_code;
2:string status_msg;
3:bool succed;
}

service ApiService {
    user_login_response UserLogin(1: user_login_request req) (api.post = "/blinkable/user/login")
    user_register_response UserRegister(1: user_register_request req) (api.post = "/blinkable/user/register")
    get_user_info_response GetUserInfo(1: get_user_info_request req) (api.get = "/blinkable/user/info")
    update_user_info_response UpdateUserInfo(1: update_user_info_request req) (api.post = "/blinkable/user/update")
    update_user_password_response UpdateUserPassword(1: update_user_password_request req) (api.post = "/blinkable/user/update/password")
    get_homepage_response GetHomePage(1:get_homepage_request req) (api.get="/blinkable/homepage/get")
    add_guestbook_response AddGuestbook(1:add_guestbook_request req)(api.post="/blinkable/homepage/guesybook")
    like_response Like(1:like_request req)(api.post="blinkable/homepage/like")

}
