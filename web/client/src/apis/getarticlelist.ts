import {$http} from "@/apis/index";

export const getarticlelist= (data: {start:number,end:number})=>{
    return $http({
        method:"get",
        url:"http://39.107.70.238:10000/blinkable/article/list",
        params: data,
    });
}