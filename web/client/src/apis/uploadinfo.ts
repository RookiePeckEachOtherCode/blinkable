import {$http} from "@/apis/index";
export const uploadinfo= (data: {user_id:string,token:string,username:string,signature:string,title:string,git_url:string})=>{
    return $http({
        method:"post",
        url:"https://211.149.141.23:17767/blinkable/user/update",
        params: data,
    });
}