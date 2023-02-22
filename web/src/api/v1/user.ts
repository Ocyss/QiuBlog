import request from "@/utils/request";

export const register = (data) => {
  return request({ url: "/api/v1/user/register", method: "post", data });
};

export const login = (data) => {
  return request({ url: "/api/v1/user/login", method: "post", data });
};
