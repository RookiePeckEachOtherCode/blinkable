import {$http} from "@/apis/index";

export const beadmin= (data: {user_id:string,key:string})=>{
    return $http({
        method:"post",
        url:"https://cn-cd-dx-tmp7.natfrp.cloud:17767/blinkable/user/beadmin",
        params: data,
    });
}