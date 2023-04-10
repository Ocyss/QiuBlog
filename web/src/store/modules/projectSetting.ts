import { defineStore } from "pinia";
import api from "@/api";

export const useProjectSettingStore = defineStore({
  id: "app-project-setting",
  state: () => ({
    isMobile: false, // 是否处于移动端模式
    collapsed: false, // 侧栏是否折叠
    autuLoad: false, // 是否自动下一页
    allTags: undefined, //所有标签
    isMainUV: false, // 是否统计过
  }),
  actions: {
    async getAllTags(): Promise<Array<any>> {
      if (this.allTags == undefined) {
        let res = await api.tags.get();
        this.allTags = res.data;
      }
      return this.allTags;
    },
    mainUV(): void {
      if (!this.isMainUV) {
        api.statistics.mainuv().then(() => {
          this.isMainUV = true;
        });
      }
    },
  },
  getters: {},
  persist: {
    paths: ["collapsed", "autuLoad"],
  },
});
