import enUS from "./enUS/index";
import zhCN from "./zhCN/index";
import { createI18n } from "vue-i18n";

export const i18n = createI18n({
  globalInjection: true, //全局生效$t
  locale: "zhCN",
  messages: {
    zhCN,
    enUS,
  },
  legacy: false,
});
