import request from "@/utils/request";

export const addMessage = (data) => {
  return request({ url: "/api/v1/message", method: "post", data });
};

export const addQuestion = (data) => {
  return request({ url: "/api/v1/question", method: "post", data });
};
export const getMessage = (params) => {
  return request({ url: "/api/v1/message", method: "get", params });
};

export const getQuestion = (params) => {
  return request({ url: "/api/v1/question", method: "get", params });
};
