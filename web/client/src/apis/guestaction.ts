import {$http} from "@/apis/index";

export const guestbookApi= (data: { admin_id: number; user_id: number;context:string })=>{
    return $http({
        method:"post",
        url:"/Main/guestbook",
        params: data,
    });
}