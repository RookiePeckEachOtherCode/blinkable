import {$http} from "@/apis/index";

export const beadmin= (data: {user_id:string,key:string})=>{
    return $http({
        method:"post",
        url:"http://39.107.70.238:10000/blinkable/user/beadmin",
        params: data,
    });
}