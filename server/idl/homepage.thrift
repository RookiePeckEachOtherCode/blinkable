namespace go Homepage

include "base.thrift"

struct GetHomepageRequest {
}
struct GetHomepageResponse {
    1: list<base.User> users;
    2: i32             status_code;
    3: string          status_msg;
    4: bool            succed;
}

struct LikeRequest {
    1: i64 user_id;
    2: i64 from_user_id;
}

struct LikeResponse {
    1: i32    status_code;
    2: string status_msg;
    3: bool   succed;
}

struct AddGuestbookRequest {
    1: i64    user_id;
    2: i64    from_user_id;
    3: string context;
}

struct AddGuestbookResponse {
    1: i32    status_code;
    2: string status_msg;
    3: bool   succed;
}

service HomepageService {
    LikeResponse LikeAction(1: LikeRequest req);
    GetHomepageResponse GetMainview(1: GetHomepageRequest req);
    AddGuestbookResponse AddGuestbook(1: AddGuestbookRequest req);
}
