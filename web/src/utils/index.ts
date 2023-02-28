import { h, unref, CSSProperties } from "vue";
import { NIcon, NTag, NTooltip, NTime } from "naive-ui";
import { cloneDeep } from "lodash-es";
import { Ref, inject } from "vue";
import type { Config } from "@/types";

async function getConfig(): Promise<Config> {
  let config: Config;
  let file = await fetch("static/config.json5");
  let res = await file.text();
  config = new Function("return " + res)();
  //console.log(config);
  return config;
}
//Nicon封装
export function renderIcon(icon) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

function addLight(color, amount) {
  const cc = parseInt(color, 16) + amount;
  const c = cc > 255 ? 255 : cc;
  return c.toString(16).length > 1 ? c.toString(16) : `0${c.toString(16)}`;
}

export function lighten(color, amount) {
  color = color.indexOf("#") >= 0 ? color.substring(1, color.length) : color;
  amount = Math.trunc((255 * amount) / 100);
  return `#${addLight(color.substring(0, 2), amount)}${addLight(
    color.substring(2, 4),
    amount
  )}${addLight(color.substring(4, 6), amount)}`;
}

//判断是否url
export function isUrl(url) {
  return /^(http|https):\/\//g.test(url);
}
//随机颜色
export const randomRgb = (min, max, opacity) => {
  let R = Math.floor(Math.random() * (max - min)) + min;
  let G = Math.floor(Math.random() * (max - min)) + min;
  let B = Math.floor(Math.random() * (max - min)) + min;
  return "rgba(" + R + "," + G + "," + B + ", " + opacity + ")";
};
//时间组件封装
export const timeControl = (t) => {
  return h(
    NTooltip,
    { trigger: "hover" },
    {
      default: () =>
        h(NTime, {
          timeZone: "Asia/Shanghai",
          time: new Date(t),
          format: "yyyy-MM-dd HH:mm:ss",
        }),
      trigger: () =>
        h(NTime, {
          timeZone: "Asia/Shanghai",
          time: new Date(t),
          type: "relative",
        }),
    }
  );
};

export const railStyle = (
  { focused, checked },
  bg = "#d03050",
  bs = "#d0305040",
  ubg = "#1d46ff",
  ubs = "#2080f040"
) => {
  const style: CSSProperties = {};
  if (checked) {
    style.background = bg;
    if (focused) {
      style.boxShadow = `0 0 0 2px ${bs}`;
    }
  } else {
    style.background = ubg;
    if (focused) {
      style.boxShadow = `0 0 0 2px ${ubs}`;
    }
  }
  return style;
};

//修改网页标题
export const setTitle = (t: any) => {
  getConfig().then((config) => {
    window.document.title = `${t} ${config.userInfo.title}`;
  });
};
