import { defineStore } from "pinia";
import { useOsTheme } from "naive-ui";
const osThemeRef = useOsTheme();

export const useDesignSettingStore = defineStore({
  id: "app-design-setting",
  state: () => ({
    darkTheme: osThemeRef.value === "dark", //深色主题
    appTheme: "#FC9D99", //系统风格
    locale: true,
  }),
  getters: {
    getDarkTheme(): boolean {
      return this.darkTheme;
    },
    getAppTheme(): string {
      return this.appTheme;
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
