import {$http} from "@/apis/index";

/*
*登录
 */
export const loginApi=(data:{password:string;username:string})=>{
    return $http({
        method:"post",
        url:"https://211.149.141.23:17767/blinkable/user/login",
        params: data,
    });
}