import {$http} from "@/apis/index";

export const upload_iconApi = (data:{file: FormData,user_id:string}) => {
    return $http({
        method: "post",
        url: "https://211.149.141.23:17767/blinkable/user/upload/icon",
        data: data, // Use FormData as the data
        headers: {
            'Content-Type': 'multipart/form-data',
        },
    });
};