import {$http} from "@/apis/index";

export const uploadApi = (data:{file: FormData,user_id:string}) => {
    return $http({
        method: "post",
        url: "/user/upload",
        data: data, // Use FormData as the data
        headers: {
            'Content-Type': 'multipart/form-data', // Set the correct Content-Type header
        },
    });
};
