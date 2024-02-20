import {$http} from "@/apis/index";

export const addComment= (data: {user_id:string,context:string,article_id:number})=>{
    return $http({
        method:"post",
        url:"https://211.149.141.23:17767/blinkable/article/comment",
        params: data,
    });
}