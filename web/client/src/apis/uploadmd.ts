import {$http} from "@/apis/index";

export const uploadmd = (data:{file: FormData,user_id:string,title:string}) => {
    return $http({
        method: "post",
        url: "http://127.0.0.1:10000/blinkable/article/upload",
        data: data,
        headers: {
            'Content-Type': 'multipart/form-data',
        },
    });
};