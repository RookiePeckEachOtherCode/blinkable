import {$http} from "@/apis/index";

export const uploadmd = (data:{content: string,user_id:string,title:string,token:string}) => {
    return $http({
        method: "post",
        url: "https://211.149.141.23:17767/blinkable/article/publish",
        data: data,
        headers: {
            'Content-Type': 'multipart/form-data',
        },
    });
};