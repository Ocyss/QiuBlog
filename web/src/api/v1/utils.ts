import request from "@/utils/request";

export const getCaptcha = () => {
    return request({url: "/api/v1/captcha", method: "get"});
};

export const checkCaptcha = (key, dots) => {
    return request({
        url: "/api/v1/captcha",
        method: "post",
        data: {key, dots},
    });
};
