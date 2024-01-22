import {$http} from "@/apis/index";

export const getarticlelist= (data: {start:number,end:number})=>{
    return $http({
        method:"post",
        url:"http://127.0.0.1:4523/m2/3417810-0-default/137820055",
        params: data,
    });
}