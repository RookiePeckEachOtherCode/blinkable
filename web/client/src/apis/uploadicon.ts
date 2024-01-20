import {$http} from "@/apis/index";

export const upload_iconApi = (data:{file: FormData,user_id:string}) => {
    return $http({
        method: "post",
        url: "http://127.0.0.1:10000/blinkable/user/update/avatar",
        data: data, // Use FormData as the data
        headers: {
            'Content-Type': 'multipart/form-data',
        },
    });
};