import {$http} from "@/apis/index";
export const getuserinfo= (data: {user_id:string;token:string})=>{
    return $http({
        method:"get",
        url:"http://127.0.0.1:10000/blinkable/user/info",
        params: data,
    });
}