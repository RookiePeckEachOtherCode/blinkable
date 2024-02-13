import {$http} from "@/apis/index";

export const getarticlelist= (data: {start:number,end:number})=>{
    return $http({
        method:"get",
        url:"http://127.0.0.1:10000/blinkable/article/list",
        params: data,
    });
}