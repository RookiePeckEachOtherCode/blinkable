import {$http} from "@/apis/index";
export const getuserinfo= (data: number)=>{
    return $http({
        method:"get",
        url:"http://127.0.0.1:10000/blinkable/user/info",
        params: data,
    });
}