import { defineStore } from "pinia";
export const globalData = defineStore({
  id: "app-global-data",
  state: () => {
    return {};
  },
  getters: {},
  actions: {},
  persist: {
    enabled: true,
    strategies: [
      {
        key: "globalData",
        storage: localStorage,
        paths: [],
      },
    ],
  },
});
