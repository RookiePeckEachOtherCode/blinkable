import {$http} from "@/apis/index";

export const Getcards=()=>{
    return $http({
        method:"get",
        url:"/Main",
    });
}