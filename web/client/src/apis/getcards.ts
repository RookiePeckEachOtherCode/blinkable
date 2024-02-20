import {$http} from "@/apis/index";

export const Getcards=()=>{
    return $http({
        method:"get",
        url:"https://211.149.141.23:17767/blinkable/homepage/get",
    });
}