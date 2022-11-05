import axios from "axios";
import { createDiscreteApi } from "naive-ui";
const { message } = createDiscreteApi(["message"]);

export const request = axios.create({
  timeout: 5000,
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
    } else {
      message.error(res.data.message);
      return Promise.reject(new Error(res.data.message));
    }
  },
  function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    return Promise.reject(error);
  }
);
