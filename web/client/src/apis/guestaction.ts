import {$http} from "@/apis/index";

export const guestbookApi= (data: { from_user_id: string; user_id: string;context:string;token:string })=>{
    return $http({
        method:"post",
        url:"https://211.149.141.23:17767/blinkable/homepage/guestbook",
        params: data,
    });
}