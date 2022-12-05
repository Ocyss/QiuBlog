import request from "@/utils/request";

export const get = (show = false) => {
  return request({
    url: "/api/v1/category",
    method: "get",
    params: {
      show,
    },
  });
};

export const put = (data) => {
  return request({
    url: "/api/v1/category/list",
    method: "put",
    data,
  });
};

export const add = (data) => {
  return request({
    url: "/api/v1/category/add",
    method: "post",
    data,
  });
};
