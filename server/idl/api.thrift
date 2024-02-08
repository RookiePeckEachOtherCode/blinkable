namespace go api

include "base.thrift"

//========================================================UserServer====================================================================//
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
    6: string github_url (api.query = "github_url")
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

//======================================================HomepageServer======================================================================//
struct like_request {
    1: i64    user_id (api.query = "user_id")
    2: i64    from_user_id (api.query = "from_user_id")
    3: string token (api.query = "token")
}

struct like_response {
    1: i32    status_code;
    2: string status_msg;
    3: bool   succed;
}

struct get_homepage_request {
    1: string token (api.query = "token")
}

struct get_homepage_response {
    1: list<base.User> users;
    2: i32             status_code;
    3: string          status_msg;
    4: bool            succed;
}

struct add_guestbook_request {
    1: i64    user_id (api.query = "user_id")
    2: i64    from_user_id (api.query = "from_user_id")
    3: string context (api.query = "context")
    4: string token (api.query = "token")
}

struct add_guestbook_response {
    1: i32    status_code;
    2: string status_msg;
    3: bool   succed;
}
//======================================================ArticleServer======================================================================//
struct get_articlesum_request{
    1: string token (api.query = "token")
}
struct get_articlesum_response{
    1: i32    status_code;
    2: string status_msg;
    3: bool   succed;
    4: i64    sum;
}
struct get_articlelist_request{
    1:i32 start (api.query="start")
    2:i32 end (api.query="end")
}
struct get_articlelist_response{
    1: i32    status_code;
    2: string status_msg;
    3: bool   succed;
    4:list<base.ArticleMsg> articles
}
struct get_article_request{
    1:i32 article_id (api.query = "article_id")
}
struct get_article_response{
    1: i32    status_code;
    2: string status_msg;
    3: bool   succed;
    4: list<base.Comment> Comments
    5:string content
    6:i64 creater_id
}
struct publish_article_request{
    1:i64 user_id(api.qeury="user_id")
    2:binary file (api.query="file")
}
struct publish_article_response{
    1: i32    status_code;
    2: string status_msg;
    3: bool   succed;
}

struct add_comment_request{
    1:i64 user_id(api.query="user_id")
    2:i32 article_id (api.query = "article_id")
    3:string context (api.query="context")
}
struct add_comment_response{
    1: i32    status_code;
    2: string status_msg;
    3: bool   succed;
}

struct delete_article_request{
    1:i64 user_id(api.query="user_id")
    2:i32 article_id (api.query = "article_id")
}
struct delete_article_response{
    1: i32    status_code;
    2: string status_msg;
    3: bool   succed;
}

//=======================================================================================================================================//

service ApiService {
    user_login_response UserLogin(1: user_login_request req) (api.post = "/blinkable/user/login")
    user_register_response UserRegister(1: user_register_request req) (api.post = "/blinkable/user/register")
    get_user_info_response GetUserInfo(1: get_user_info_request req) (api.get = "/blinkable/user/info")
    update_user_info_response UpdateUserInfo(1: update_user_info_request req) (api.post = "/blinkable/user/update")
    update_user_password_response UpdateUserPassword(1: update_user_password_request req) (api.post = "/blinkable/user/update/password")
    get_homepage_response GetHomePage(1: get_homepage_request req) (api.get = "/blinkable/homepage/get")
    add_guestbook_response AddGuestbook(1: add_guestbook_request req) (api.post = "/blinkable/homepage/guestbook")
    like_response LikeAction(1: like_request req) (api.post = "blinkable/homepage/like")
    get_articlesum_response GetArticleSum(1:get_articlesum_request req)(api.get="blinkable/article/sum")
    get_articlelist_response GetArtcleList(1:get_articlelist_request req)(api.get="blinkable/article/list")
    get_article_response GetArticle(1:get_article_request req)(api.get="blinkable/article/get")
    publish_article_response PublishArticle(1:publish_article_request req)(api.post="blinkable/article/publish")
    add_comment_response AddComment(1:add_guestbook_request req)(api.post="blinkable/article/comment")
    delete_article_response DeleteArticle(1:delete_article_request req)(api.post="blinkable/article/delet")
}
