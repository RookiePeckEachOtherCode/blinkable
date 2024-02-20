import {$http} from "@/apis/index";

export const deletearticle= (data: {user_id:string,token:string,article_id:number})=>{
    return $http({
        method:"post",
        url:"https://211.149.141.23:17767/blinkable/article/delet",
        params: data,
    });
}