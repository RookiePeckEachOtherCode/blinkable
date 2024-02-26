import {$http} from "@/apis/index";

export const deletearticle= (data: {user_id:string,token:string,article_id:number})=>{
    return $http({
        method:"post",
        url:"http://39.107.70.238:10000/blinkable/article/delet",
        params: data,
    });
}