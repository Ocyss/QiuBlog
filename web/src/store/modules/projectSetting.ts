import { defineStore } from "pinia";

export const useProjectSettingStore = defineStore({
  id: "app-project-setting",
  state: () => ({
    isMobile: false, // 是否处于移动端模式
    collapsed: false, // 侧栏是否折叠
    autuLoad: false, // 是否自动下一页
  }),
  persist: {
    enabled: true,
    strategies: [
      {
        key: "Setting",
        storage: localStorage,
        paths: ["collapsed", "autuLoad"],
      },
    ],
  },
});
