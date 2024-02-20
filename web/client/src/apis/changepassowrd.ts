import {$http} from "@/apis/index";

export const changepassowrd= (data: {user_id:string,old_password:string,new_password:string,token:string})=>{
    return $http({
        method:"post",
        url:"https://211.149.141.23:17767/blinkable/user/update/password",
        params: data,
    });
}