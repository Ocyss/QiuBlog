import axios from "axios";

import { message } from "@/utils/client";

let request = axios.create({
  baseURL: "http://127.0.0.1:3000/",
  timeout: 8000,
});

// 添加请求拦截器
request.interceptors.request.use(
  function (config) {
    // 在发送请求之前做些什么
    return config;
  },
  function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
  }
);

// 添加响应拦截器
request.interceptors.response.use(
  function (res) {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
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
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
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
