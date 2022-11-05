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
