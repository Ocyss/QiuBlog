import request from "@/utils/request";

export const get = (pid) => {
  return request({ url: `/api/v1/article/${pid}`, method: "get" });
};

export const add = (data) => {
  return request({ url: "/api/v1/article/add", method: "post", data });
};

export const getList = (params) => {
  return request({ url: "/api/v1/article/list", method: "get", params });
};

export const put = (id, data) => {
  return request({ url: `/api/v1/article/${id}`, method: "put", data });
};
