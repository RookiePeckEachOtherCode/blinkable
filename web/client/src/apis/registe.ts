import {$http} from "@/apis/index";

export const registerApi=(data:{password:string;username:string})=>{
    return $http({
        method:"post",
        url:"http://39.107.70.238:10000/blinkable/user/register",
        params:data,
    });
}