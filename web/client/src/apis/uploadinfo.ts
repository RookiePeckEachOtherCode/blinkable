import {$http} from "@/apis/index";
export const uploadinfo= (data: {user_id:string,token:string,username:string,signature:string,title:string,git_url:string})=>{
    return $http({
        method:"post",
        url:"http://39.107.70.238:10000/blinkable/user/update",
        params: data,
    });
}