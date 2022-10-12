import { defineStore } from "pinia";
import axios from "axios";
export const globalData = defineStore({
  id: "app-global-data",
  state: () => {
    return { category: [{ id: -1, name: "全部", homeshow: true }] };
  },
  getters: {
    getCategory() {
      if (this.category.length == 1) {
        this.askCategory(false);
      }
      return this.category;
    },
  },
  actions: {
    askCategory(show) {
      axios
        .get("/api/v1/category", {
          params: { show: show },
        })
        .then((res) => {
          if (res.data.status == 200) {
            this.category.push(...res.data.data);
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
        paths: [""],
      },
    ],
  },
});
