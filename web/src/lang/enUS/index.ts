import { enUS } from "naive-ui";
import config from "../config";

export default {
  main: {
    title: "💘  Ocyss's Blog",
    name: "Ocyss",
    email: "qiudie@88.com",
    motto: "Now is better than never.",
  },
  info: {
    title: {
      load: "Loading ~,be patient and wait!",
      hidden: "Ahh 💔 how did it leave 💔 how? How could it be!",
    },
    card: {
      title: {
        message: "latest message",
        popular: "recently popular",
        statistics: "statistics",
        tagCloud: "tag cloud",
      },
      statistics: {
        mainUv: "Page_View",
        wordsTotal: "Words_Total",
        articleCount: "Article_Count",
        lastUpdated: "Last_Updated",
        elapsedTime: "Run_Time",
        date: "{d} d {h} h {m} m {s} s",
      },
    },
    footer: {
      icp: { moe: `MoeICP:${config.icp.moe}` },
      info: "Address",
      copyright: "Copyright",
    },
  },
  component: {
    chain: {
      name: "FriendlyLinks",
    },
    about: { author: "author_introduce", project: "project_introduce" },
    header: {
      darkTheme: { dark: "LowerClass", bright: "UpperClass" },
    },
  },
};
