import { defineStore } from "pinia";
import axios from "axios";
export const globalData = defineStore({
  id: "app-global-data",
  state: () => {
    return { category: [] };
  },
  getters: {},
  actions: {
    getCategory(show = false) {
      if (this.category.length == 0) {
        axios.get("/api/v1/category?show=false").then((res) => {
          if (res.data.status == 200) {
            this.category = res.data.data;
          }
        });
      }
      if (show) {
        return this.category.filter((item) => {
          return item.homeshow;
        });
      } else {
        return this.category;
      }
    },
  },
  persist: {
    enabled: true,
    strategies: [
      {
        key: "globalData",
        storage: localStorage,
        paths: ["category"],
      },
    ],
  },
});
