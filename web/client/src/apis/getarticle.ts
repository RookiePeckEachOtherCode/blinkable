import {$http} from "@/apis/index";

export const getarticle= (data: {article_id:number})=>{
    return $http({
        method:"get",
        url:"http://127.0.0.1:10000/blinkable/article/get",
        params: data,
    });
}