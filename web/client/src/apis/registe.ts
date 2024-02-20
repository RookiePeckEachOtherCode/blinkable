import {$http} from "@/apis/index";

export const registerApi=(data:{password:string;username:string})=>{
    return $http({
        method:"post",
        url:"https://211.149.141.23:17767/blinkable/user/register",
        params:data,
    });
}