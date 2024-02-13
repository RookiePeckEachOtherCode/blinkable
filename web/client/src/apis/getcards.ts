import {$http} from "@/apis/index";

export const Getcards=()=>{
    return $http({
        method:"get",
        url:"http://127.0.0.1:10000/blinkable/homepage/get",
    });
}