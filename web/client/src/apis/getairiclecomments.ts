import {$http} from "@/apis/index";

export const getarticlecomments= (data: {article_id:number})=>{
    return $http({
        method:"get",
        url:"http://127.0.0.1:4523/m1/3417810-0-default/blinkable/comment",
        params: data,
    });
}