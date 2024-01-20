import {$http} from "@/apis/index";

export const Likeapi= (data: { admin_id: number; user_id: number })=>{
    return $http({
        method:"post",
        url:"/Main/like",
        params: data,
    });
}