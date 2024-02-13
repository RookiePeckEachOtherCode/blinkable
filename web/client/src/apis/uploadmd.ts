import {$http} from "@/apis/index";

export const uploadmd = (data:{content: string,user_id:string,title:string}) => {
    return $http({
        method: "post",
        url: "http://127.0.0.1:10000/blinkable/article/publish",
        data: data,
        headers: {
            'Content-Type': 'multipart/form-data',
        },
    });
};