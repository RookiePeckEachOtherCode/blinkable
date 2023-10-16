import {$http} from "@/apis/index";

/*
*登录
 */
export const loginApi=(data:{password:string;username:string})=>{
    return $http({
        method:"post",
        url:"/admin/login",
        data,
    });
}