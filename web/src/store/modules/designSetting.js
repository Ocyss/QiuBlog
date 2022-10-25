import { defineStore } from "pinia";
import { store } from "@/store";
import designSetting from "@/settings/designSetting.js";

const { darkTheme, appTheme, appThemeList } = designSetting;

export const useDesignSettingStore = defineStore({
  id: "app-design-setting",
  state: () => ({
    darkTheme, //深色主题
    appTheme, //系统风格
    appThemeList, //系统内置风格
    locale: true, //国际化语音true中文,false英文
  }),
  getters: {
    getDarkTheme() {
      return this.darkTheme;
    },
    getAppTheme() {
      return this.appTheme;
    },
    getAppThemeList() {
      return this.appThemeList;
    },
    getLocale() {
      return this.locale;
    },
  },
  actions: {
    setDarkTheme(value) {
      this.darkTheme = value;
    },
    setLocale(value) {
      console.log(value);
      this.locale = value;
    },
  },
  persist: {
    enabled: true,
    strategies: [
      {
        key: "design-setting",
        storage: localStorage,
        paths: ["darkTheme", "appTheme", "appThemeList", "locale"],
      },
    ],
  },
});

// Need to be used outside the setup
export function useDesignSettingWithOut() {
  return useDesignSettingStore(store);
}
