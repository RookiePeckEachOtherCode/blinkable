import { defineStore } from "pinia";
import { httpInstance } from "@/apis";
import router from "@/router";

export const useUserInfoStore = defineStore("userinfo-store", () => {
    const setAuth = (token: string, userId: string) => {
        httpInstance.defaults.headers.common.Authorization = token;
        localStorage.setItem("token", token);
        localStorage.setItem("user_id", userId);
    }; // 存放token和user_id

    const authFromLocal = () => {
        const token = localStorage.getItem("token");
        const userId = localStorage.getItem("user_id");
        if (token && userId) {
            setAuth(token, userId);
            return true;
        }
        return false;
    }; // 本地拿token和user_id检测登录
    const removeAuth = () => {
        delete httpInstance.defaults.headers.common.Authorization;
        localStorage.removeItem("token");
        localStorage.removeItem("user_id");
        router.push("/login");
    }; // 清除登录并跳转登录

    const getUserId = () => {
        return localStorage.getItem("user_id");
    }; // 获取用户ID

    return {
        setAuth,
        authFromLocal,
        removeAuth,
        getUserId,
    };
});
