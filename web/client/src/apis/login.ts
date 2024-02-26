import {$http} from "@/apis/index";

/*
*登录
 */
export const loginApi=(data:{password:string;username:string})=>{
    return $http({
        method:"post",
        url:"http://39.107.70.238:10000/blinkable/user/login",
        params: data,
    });
}