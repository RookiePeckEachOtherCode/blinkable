import {$http} from "@/apis/index";

export const registerApi=(data:{password:string;username:string})=>{
    return $http({
        method:"post",
        url:"http://127.0.0.1:10000/blinkable/user/register",
        params:data,
    });
}