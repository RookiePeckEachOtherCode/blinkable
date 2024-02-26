import {$http} from "@/apis/index";

export const changepassowrd= (data: {user_id:string,old_password:string,new_password:string,token:string})=>{
    return $http({
        method:"post",
        url:"http://39.107.70.238:10000/blinkable/user/update/password",
        params: data,
    });
}