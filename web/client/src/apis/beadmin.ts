import {$http} from "@/apis/index";

export const beadmin= (data: {user_id:string,key:string})=>{
    return $http({
        method:"post",
        url:"http://127.0.0.1:10000/blinkable/user/beadmin",
        params: data,
    });
}