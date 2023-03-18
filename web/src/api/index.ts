import * as category from "./v1/category";
import * as menuchild from "./v1/menuchild";
import * as tags from "./v1/tags";
import * as article from "./v1/article";
import * as user from "./v1/user";
import * as message from "./v1/message";
import * as statistics from "./v1/statistics";
import request from "@/utils/request";

function upload(files, type) {
  const formData = new FormData();
  formData.append("file", files);
  formData.append("class", type);
  return request({
    url: "/api/v1/upload/image",
    method: "post",
    data: formData,
    timeout: 20000,
  });
}

export default {
  request,
  upload,
  category,
  menuchild,
  tags,
  article,
  user,
  message,
  statistics,
};
