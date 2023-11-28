import {$http} from "@/apis/index";

export const registerApi=(data:{password:string;username:string})=>{
    return $http({
        method:"post",
        url:"/user/register",
        data,
    });
}