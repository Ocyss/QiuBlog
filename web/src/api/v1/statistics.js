import request from "@/utils/request";

export const mainuv = () => {
  return request({ url: "/api/v1/statistics/set/mainuv", method: "post" });
};

export const statistics = () => {
  return request({ url: "/api/v1/statistics", method: "get" });
};
