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

export const upMessage = (id, val, show, message) => {
  return request({
    url: "/api/v1/message/updata",
    method: "put",
    data: { id, val, show, message },
  });
};

export const replyQuestion = (id, content) => {
  return request({
    url: "/api/v1/message/reply",
    method: "put",
    data: { id, content },
  });
};

export const delMessage = (id, message) => {
  return request({
    url: "/api/v1/message/del",
    method: "delete",
    data: { id, message },
  });
};
