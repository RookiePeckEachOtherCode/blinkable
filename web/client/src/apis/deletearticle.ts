import {$http} from "@/apis/index";

export const deletearticle= (data: {user_id:string,token:string,article_id:number})=>{
    return $http({
        method:"post",
        url:"http://127.0.0.1:10000/blinkable/article/delet",
        params: data,
    });
}