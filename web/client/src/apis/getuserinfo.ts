import {$http} from "@/apis/index";
export const getuserinfo= (data: {user_id:string;token:string})=>{
    return $http({
        method:"get",
        url:"http://39.107.70.238:10000/blinkable/user/info",
        params: data,
    });
}