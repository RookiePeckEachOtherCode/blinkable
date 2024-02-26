import {$http} from "@/apis/index";

export const Getcards=()=>{
    return $http({
        method:"get",
        url:"http://39.107.70.238:10000/blinkable/homepage/get",
    });
}