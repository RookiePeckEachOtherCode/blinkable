import {$http} from "@/apis/index";

export const addComment= (data: {user_id:string,context:string,article_id:number})=>{
    return $http({
        method:"post",
        url:"http://127.0.0.1:10000/blinkable/article/comment",
        params: data,
    });
}