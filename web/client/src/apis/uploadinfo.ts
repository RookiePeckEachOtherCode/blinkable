import {$http} from "@/apis/index";
export const uploadinfo= (data: {user_id:string,token:string,username:string,signature:string,title:string})=>{
    return $http({
        method:"get",
        url:"http://127.0.0.1:10000/blinkable/user/update",
        params: data,
    });
}