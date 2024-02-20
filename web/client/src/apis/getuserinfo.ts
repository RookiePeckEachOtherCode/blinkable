import {$http} from "@/apis/index";
export const getuserinfo= (data: {user_id:string;token:string})=>{
    return $http({
        method:"get",
        url:"https://211.149.141.23:17767/blinkable/user/info",
        params: data,
    });
}