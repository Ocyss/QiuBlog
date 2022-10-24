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
  },
  actions: {
    setDarkTheme(value) {
      this.darkTheme = value;
    },
  },
  persist: {
    enabled: true,
    strategies: [
      {
        key: "design-setting",
        storage: localStorage,
        paths: ["darkTheme", "appTheme", "appThemeList"],
      },
    ],
  },
});

// Need to be used outside the setup
export function useDesignSettingWithOut() {
  return useDesignSettingStore(store);
}
