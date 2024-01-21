import {$http} from "@/apis/index";

export const getarticle= (data: {article_id:number})=>{
    return $http({
        method:"post",
        url:"http://127.0.0.1:4523/m2/3417810-0-default/116777645",
        params: data,
    });
}