import {$http} from "@/apis/index";

export const getarticle= (data: {article_id:number})=>{
    return $http({
        method:"get",
        url:"https://211.149.141.23:17767/blinkable/article/get",
        params: data,
    });
}