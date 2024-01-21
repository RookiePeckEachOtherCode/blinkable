namespace go base

struct base_response {
    1: i32    status_code
    2: string status_msg
    3: bool   succed
}

struct User {
    1: i64    id
    2: string name
    3: string avatar
    4: i32    articles_num
    5: i32    experience
    6: string background_img
    7: i32    level
    8: string signature
    9: string title
}
