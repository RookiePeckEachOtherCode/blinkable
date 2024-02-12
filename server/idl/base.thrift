namespace go base

struct base_response {
    1: i32    status_code
    2: string status_msg
    3: bool   succed
}

struct Guestbook {
    1: i64    id;
    2: i64    user_id;
    3: string context;
    4: i64    from_user_id;
    5: string create_time;
}

struct User {
    1:  i64    id
    2:  string name
    3:  string avatar
    4:  i32    articles_num
    5:  i32    experience
    6:  string background_img
    7:  i32    level
    8:  string signature
    9:  string title
    10: i64    like_num
    11: string github_url
    12: list<Guestbook> guestbooks;
}

struct ArticleMsg{
    1: string create_time
    2: string update_time
    3: i64 creater_id
    4: i64 article_id
    5:string title
}
struct Comment{
    1:i32 comment_id
    2:i32 user_id
    3:string context
    4:string create_time
}
