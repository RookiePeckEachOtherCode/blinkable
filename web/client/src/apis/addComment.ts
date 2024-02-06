import {$http} from "@/apis/index";

export const addComment= (data: {user_id:string,context:string,article_id:number})=>{
    return $http({
        method:"post",
        url:"http://127.0.0.1:4523/m1/3417810-0-default/blinkable/comment",
        params: data,
    });
}