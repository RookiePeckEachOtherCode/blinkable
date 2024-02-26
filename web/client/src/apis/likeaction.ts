import {$http} from "@/apis/index";

export const Likeapi= (data: { from_user_id: string, user_id: string,token:string })=>{
    return $http({
        method:"post",
        url:"http://39.107.70.238:10000/blinkable/homepage/like",
        params: data,
    });
}