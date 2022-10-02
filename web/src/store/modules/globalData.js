import { defineStore } from "pinia";
import axios from "axios";
export const globalData = defineStore({
  id: "app-global-data",
  state: () => {
    return { category: [{ id: -1, name: "全部" }] };
  },
  getters: {
    getCategory() {
      return this.category;
    },
  },
  actions: {
    askCategory() {
      axios
        .get("/api/v1/category", {
          params: { show: true },
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
