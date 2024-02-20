import {$http} from "@/apis/index";

export const Likeapi= (data: { user_id: string, admin_id: string,token:string })=>{
    return $http({
        method:"post",
        url:"https://211.149.141.23:17767/blinkable/homepage/like",
        params: data,
    });
}