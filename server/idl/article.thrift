namespace go Article
include "base.thrift"
struct GetArticleSumRequest{
    1: string token
}
struct GetArticleSumResponse{
    1: i32    status_code;
    2: string status_msg;
    3: bool   succed;
    4: i64    sum;
}
struct GetArticleListRequest{
    1:i32 start
    2:i32 end
}
struct GetArticleListResponse{
    1: i32    status_code;
    2: string status_msg;
    3: bool   succed;
    4:list<base.ArticleMsg> articles
}
struct GetArticleRequest{
    1:i32 article_id
}
struct GetArticleResponse{
    1: i32    status_code;
    2: string status_msg;
    3: bool   succed;
    4: list<base.Comment> Comments
    5:string content
    6:i64 creater_id
    7:string title
}
struct PublishArticleRequest{
    1:i64 user_id
    2:string content
    3:string title
}
struct PublishArticleResponse{
    1: i32    status_code;
    2: string status_msg;
    3: bool   succed;
}

struct AddCommentRequest{
    1:i64 user_id
    2:i32 article_id
    3:string context
}
struct AddCommentResponse{
    1: i32    status_code;
    2: string status_msg;
    3: bool   succed;
}

struct DeleteArticleRequest{
    1:i64 user_id
    2:i32 article_id
}
struct DeleteArticleResponse{
    1: i32    status_code;
    2: string status_msg;
    3: bool   succed;
}
service ArticleService{
    GetArticleSumResponse GetArticleSum(1:GetArticleListRequest req);
    GetArticleListResponse GetArticleList(1:GetArticleListRequest req);
    GetArticleResponse GetArticle(1:GetArticleRequest req);
    PublishArticleResponse PublishArticle(1:PublishArticleRequest req);
    AddCommentResponse AddComment(1:AddCommentRequest req);
    DeleteArticleResponse DeleteArticle(1:DeleteArticleRequest req);
}