import {$http} from "@/apis/index";

export const guestbookApi= (data: { from_user_id: string; user_id: string;context:string;token:string })=>{
    return $http({
        method:"post",
        url:"http://127.0.0.1:10000/blinkable/homepage/guestbook",
        params: data,
    });
}