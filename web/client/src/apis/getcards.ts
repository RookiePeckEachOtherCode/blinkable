import {$http} from "@/apis/index";

export const Getcards=()=>{
    return $http({
        method:"get",
        url:"http://127.0.0.1:4523/m2/3417810-0-default/120338492",
    });
}