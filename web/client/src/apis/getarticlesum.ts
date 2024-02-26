import {$http} from "@/apis/index";

export const getarticlesum= ()=>{
    return $http({
        method:"get",
        url:"http://39.107.70.238:10000/blinkable/article/sum",
    });
}