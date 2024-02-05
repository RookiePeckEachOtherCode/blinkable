import axios from "axios";
import JSONbig from 'json-bigint';
import type { AxiosRequestConfig } from 'axios';
import { ElLoading, ElMessage } from "element-plus";
import { AxiosError } from "axios";

// 创建 Axios 实例
export const httpInstance = axios.create();

export type Status = {
    statuscode: number;
    statusmsg: string;
    succeed: boolean;
};

// 将 json-bigint 设置为 transformResponse
httpInstance.defaults.transformResponse = [
    function (data) {
        // 使用 json-bigint 解析响应数据
        try {
            return JSONbig.parse(data);
        } catch (error) {
            return data;
        }
    },
];

// 响应拦截器
export const $http = async (config: AxiosRequestConfig) => {
    const loadingInstance = ElLoading.service();
    try {
        const axiosResponse = await httpInstance(config);
        const status = axiosResponse.data;

       if (status?.succeed) {
            let errTitle: string = 'Error;';
            if (status.status_code === 401) {
                errTitle = 'Unauthorized';
                ElMessage.error('未授权');
                // 处理跳转
            } else if (status.status_code === 403) {
                errTitle = 'Forbidden';
            } else if (status.status_code === 500) {
                errTitle = "SeverError";
            }
            const err = new Error(status?.status_msg || 'Unknown');
            err.name = errTitle;
            throw err;
        }
        return status;
    } catch (err) {
        if (err instanceof AxiosError) {
            console.log(err)
            ElMessage.error('网络错误');
        }
         throw err;
    } finally {
        loadingInstance.close();
    }
};
