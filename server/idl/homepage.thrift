namespace go Homepage

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

struct LikeRequest{
 1:i32 admin_id;
 2:i32 user_id;
}

struct GetHomepageRequest{

}

struct AddGuestbookRequest{
 1:i32 user_id;
 2:i32 admin_id;
 3:string context;

}

struct GetHomepageResponse{
1:list<Users> users;
2:i32 status_code;
3:string status_msg;
4:bool succed;
}

struct LikeResponse{
1:i32 status_code;
2:string status_msg;
3:bool succed;
}

struct AddGuestbookResponse{
1:i32 status_code;
2:string status_msg;
3:bool succed;
}




service HomepageService{
 LikeResponse LikeAction(1:LikeRequest req);
 GetHomepageResponse GetMainview(1:GetHomepageRequest req);
 AddGuestbookResponse AddGuestbook(1:AddGuestbookRequest req);
}
