import { defineStore } from "pinia";

export const useProjectSettingStore = defineStore({
  id: "app-project-setting",
  state: () => {
    return {
      // 是否处于移动端模式
      isMobile: false,
      // 侧栏是否折叠
      collapsed: false,
      // 是否自动下一页
      autuLoad: false,
    };
  },
  getters: {},
  actions: {},
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
