import axios from "axios";
import type {  AxiosRequestConfig } from 'axios';
import {ElLoading, ElMessage} from "element-plus";
import {AxiosError} from "axios";
export const httpInstance=axios.create();
export type Status={
    status_code:number;
    status_msg:string;
    succeed:boolean;
};
//设置请求根路径
httpInstance.defaults.baseURL=import.meta.env.VITE_BASEURL;

//响应拦截器
export const $http=async(config:AxiosRequestConfig)=>{
    const loadingInstance= ElLoading.service();
    try{
        const axiosResponse=await httpInstance<Status>(config)
        const status=axiosResponse.data

        if(!status?.succeed){
            let errTitle:string='Error;' ;
            if(status.status_code===401){
                errTitle='Unauthorized';
                ElMessage.error('未授权')
                //跳转处理
            }else if(status.status_code===403){
                errTitle='Forbidden';
            }else if(status.status_code===500){
                errTitle="SeverError"
            }else {
                errTitle='Unknown'
            }
            const err=new Error(status?.status_msg||'Unknown');
            err.name=errTitle;
            throw err;

        }
        return status;
    }catch (err){
        if(err  instanceof AxiosError){
            ElMessage.error('网络错误');
        }
        // throw err;
    }finally {
        loadingInstance.close()
    }
}