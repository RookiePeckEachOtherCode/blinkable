import {$http} from "@/apis/index";

export const guestbookApi= (data: { admin_id: number; user_id: number;context:string })=>{
    return $http({
        method:"post",
        url:"http://127.0.0.1:10000/blinkable/homepage/guestbook",
        params: data,
    });
}