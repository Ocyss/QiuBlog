import { zhCN } from "naive-ui";
import config from "../config";

export default {
  main: {
    title: "💘  Ocyss 的博客",
    name: "Ocyss",
    email: "qiudie@88.com",
    motto: "故事很短，满是遗憾。",
  },
  info: {
    title: {
      load: "加载中 ~,亲亲耐心等等哇",
      hidden: "啊💔怎么离开了呢💔怎么会?怎么会呢!",
    },
    card: {
      title: {
        message: "最新留言",
        popular: "近期热门",
        statistics: "统计信息",
        tagCloud: "标签云",
      },
      statistics: {
        mainUv: "浏览量",
        wordsTotal: "总字数",
        articleCount: "文章数量",
        lastUpdated: "最后更新于",
        elapsedTime: "已稳点运行",
        date: "{d}天{h}时{m}分{s}秒",
      },
    },
    footer: {
      icp: { moe: `萌ICP备${config.icp.moe}号` },
      info: "驱动 开源地址",
      copyright: "版权所有",
    },
  },
  component: {
    chain: {
      name: "友链",
    },
    about: { author: "作者介绍", project: "项目介绍" },
    header: {
      darkTheme: { dark: "下班", bright: "上班" },
    },
  },
};
