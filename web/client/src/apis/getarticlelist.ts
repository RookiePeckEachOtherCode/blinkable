import {$http} from "@/apis/index";

export const getarticlelist= (data: {start:number,end:number})=>{
    return $http({
        method:"get",
        url:"https://cn-cd-dx-tmp7.natfrp.cloud:17767/blinkable/article/list",
        params: data,
    });
}