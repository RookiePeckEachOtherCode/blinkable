import {$http} from "@/apis/index";

export const getarticle= (data: {article_id:number})=>{
    return $http({
        method:"get",
        url:"http://39.107.70.238:10000/blinkable/article/get",
        params: data,
    });
}