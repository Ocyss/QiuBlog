import request from "@/utils/request";

export const set = (data) => {
  return request({ url: "/api/v1/menuchild/set", method: "put", data });
};

export const get = (params = {}) => {
  return request({ url: "/api/v1/menuchilds", method: "get", params });
};
