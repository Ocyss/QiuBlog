import { defineStore, StoreDefinition } from "pinia";
const appThemeList = [
  "#2d8cf0",
  "#0960bd",
  "#0084f4",
  "#009688",
  "#536dfe",
  "#ff5c93",
  "#ee4f12",
  "#0096c7",
  "#9c27b0",
  "#ff9800",
  "#FF3D68",
  "#00C1D4",
  "#71EFA3",
  "#171010",
  "#78DEC7",
  "#1768AC",
  "#FB9300",
  "#FC5404",
];

export const useDesignSettingStore = defineStore({
  id: "app-design-setting",
  state: () => ({
    darkTheme: false, //深色主题
    appTheme: "#2d8cf0", //系统风格
    appThemeList, //系统内置风格
    locale: true,
  }),
  getters: {
    getDarkTheme(): boolean {
      return this.darkTheme;
    },
    getAppTheme(): string {
      return this.appTheme;
    },
    getAppThemeList(): string[] {
      return this.appThemeList;
    },
    getLocale(): boolean {
      return this.locale;
    },
    judgeLocale(): number {
      return this.locale ? 0 : 1;
    },
  },
  actions: {
    setDarkTheme(value: boolean): void {
      this.darkTheme = value;
    },
    setLocale(value: boolean): void {
      this.locale = value;
    },
  },
  persist: true,
});
