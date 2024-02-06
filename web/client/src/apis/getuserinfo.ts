import {$http} from "@/apis/index";
export const getuserinfo= (data: {user_id:string;token:string})=>{
    return $http({
        method:"get",
        url:"http://127.0.0.1:4523/m2/3417810-0-default/116777224",
        params: data,
    });
}