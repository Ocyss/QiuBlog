import axios from "axios";
import { message } from "@/utils/client";

const baseURL = "http://127.0.0.1:16879";

let request = axios.create({
  baseURL,
  timeout: 8000, // 超时时间
});

// 添加请求拦截器
request.interceptors.request.use(
  function (config) {
    if (!import.meta.env.SSR) {
      let token = document.cookie.match(`[;\s+]?token=([^;]*)`)?.pop();
      if (token) {
        config.headers.set("Token", token);
      }
    }
    return config;
  },
  function (error) {
    return Promise.reject(error);
  }
);

// 添加响应拦截器
request.interceptors.response.use(
  function (res) {
    if (res.data.status == 200) {
      return res.data;
    } else if (res.data.status != undefined) {
      message.error(res.data.message);
      return Promise.reject(new Error(res.data.message));
    } else {
      return res;
    }
  },
  function (error) {
    message.error(error.response?.data.message);
    return Promise.reject(error);
  }
);

declare module "axios" {
  interface AxiosInstance {
    (config: AxiosRequestConfig): Promise<any>;
  }
}

export default request;
