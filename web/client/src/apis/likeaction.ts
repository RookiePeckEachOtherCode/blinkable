import {$http} from "@/apis/index";

export const Likeapi= (data: { user_id: string, admin_id: string })=>{
    return $http({
        method:"post",
        url:"http://127.0.0.1:10000/blinkable/homepage/like",
        params: data,
    });
}