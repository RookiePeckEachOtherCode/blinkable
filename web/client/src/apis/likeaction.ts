import {$http} from "@/apis/index";

export const Likeapi= (data: { user_id: string, admin_id: string })=>{
    return $http({
        method:"post",
        url:"/Main/like",
        params: data,
    });
}