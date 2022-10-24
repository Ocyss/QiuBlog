import { defineStore } from "pinia";

export const projectSetting = defineStore({
  id: "app-project-setting",
  state: () => {
    return {
      // 是否处于移动端模式
      isMobile: false,
      // 侧栏是否折叠
      collapsed: false,
    };
  },
  getters: {
    getIsMobile() {
      return this.isMobile;
    },
    getCollapsed() {
      return this.collapsed;
    },
  },
  actions: {
    setIsMobile(value) {
      this.isMobile = value;
    },
    setCollapsed(value) {
      this.collapsed = value;
    },
  },
  persist: {
    enabled: true,
    strategies: [
      {
        key: "Setting",
        storage: localStorage,
        paths: ["collapsed", "darkMode"],
      },
    ],
  },
});
