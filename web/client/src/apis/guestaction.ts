import {$http} from "@/apis/index";

export const guestbookApi= (data: { from_user_id: string; user_id: string;context:string;token:string })=>{
    return $http({
        method:"post",
        url:"http://39.107.70.238:10000/blinkable/homepage/guestbook",
        params: data,
    });
}