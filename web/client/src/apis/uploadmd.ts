import {$http} from "@/apis/index";

export const uploadmd = (data:{content: string,user_id:string,title:string,token:string}) => {
    return $http({
        method: "post",
        url: "http://39.107.70.238:10000/blinkable/article/publish",
        data: data,
        headers: {
            'Content-Type': 'multipart/form-data',
        },
    });
};