import { defineStore } from "pinia";
const appThemeList: Array<string> = [
  "#FC9D99",
  "#D24D57",
  "#26A65B",
  "#EB7347",
  "#AEDD81",
  "#00CCFF",
  "#D0D0D0",
  "#ff9800",
  "#FF3D68",
  "#00C1D4",
  "#71EFA3",
  "#78DEC7",
  "#1768AC",
  "#FB9300",
  "#FC5404",
  "#3B99D4",
  "#8ED14B",
  "#F06B49",
  "#ECC2F1",
  "#82C7C3",
  "#E3698A",
  "#1776EB",
  "#F5B2AC",
];

export const useDesignSettingStore = defineStore({
  id: "app-design-setting",
  state: () => ({
    darkTheme: false, //深色主题
    appTheme: "#FC9D99", //系统风格
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
