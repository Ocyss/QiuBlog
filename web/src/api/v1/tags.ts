import request from "@/utils/request";

export const get = (data) => {
  return request({ url: "/api/v1/tags", method: "get" });
};
