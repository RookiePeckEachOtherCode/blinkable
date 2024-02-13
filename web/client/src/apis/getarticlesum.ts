import {$http} from "@/apis/index";

export const getarticlesum= ()=>{
    return $http({
        method:"get",
        url:"http://127.0.0.1:10000/blinkable/article/sum",
    });
}