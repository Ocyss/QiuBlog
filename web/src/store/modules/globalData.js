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
      axios.get("/api/v1/category?show=false").then((res) => {
        if (res.data.status == 200) {
          this.category = res.data.data;
          if (show) {
            return res.data.data.filter((item) => {
              return item.homeshow;
            });
          } else {
            return res.data.data;
          }
        }
      });
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
