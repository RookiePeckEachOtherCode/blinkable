import {$http} from "@/apis/index";

export const addComment= (data: {user_id:string,context:string,article_id:number})=>{
    return $http({
        method:"post",
        url:"http://39.107.70.238:10000/blinkable/article/comment",
        params: data,
    });
}