import {$http} from "@/apis/index";

export const upload_backApi = (data:{file: FormData,user_id:string}) => {
    return $http({
        method: "post",
        url: "http://39.107.70.238:10000/blinkable/user/upload/back",
        data: data, // Use FormData as the data
        headers: {
            'Content-Type': 'multipart/form-data',
        },
    });
};