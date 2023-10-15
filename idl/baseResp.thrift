namespace go baseResp

struct BaseResponse{
    1: string status_msg (go.tag="json:'status_msg'")
    2: i32 status_code(go.tag="json:'status_code'")
}