<template>
  <div>
    <n-space>
      <n-button type="primary" @click="router.push({ name: 'article-create' })">
        文章发布
      </n-button>
    </n-space>
    <n-divider />
    <n-data-table :columns="colsReactive" :data="data" />
  </div>
</template>

<script setup>
import { useRouter } from "vue-router";
import { defineComponent, h, reactive } from "vue";
import {
  NButton,
  NSpace,
  NTag,
  NTooltip,
  NTime,
  NImage,
  NEllipsis,
} from "naive-ui";
import { PawOutline, SearchOutline } from "@vicons/ionicons5";
import axios from "axios";
const router = useRouter();
const data = ref([]);

const timer = (t) => {
  return h(
    NTooltip,
    { trigger: "hover" },
    {
      default: () =>
        h(NTime, {
          timeZone: "Asia/Shanghai",
          time: new Date(t),
          format: "yyyy-MM-dd HH:mm:ss",
        }),
      trigger: () =>
        h(NTime, {
          timeZone: "Asia/Shanghai",
          time: new Date(t),
          type: "relative",
        }),
    }
  );
};

const colsReactive = reactive([
  { title: "ID", key: "id" },
  {
    title: "发布时间",
    key: "created_at",
    render(row) {
      return timer(row.created_at);
    },
  },
  {
    title: "更新时间",
    key: "updated_at",
    render(row) {
      return timer(row.updated_at);
    },
  },
  { title: "标题", key: "title" },
  {
    title: "简介",
    key: "desc",
    width: 200,
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: "标签",
    key: "tags",
    render(row) {
      const tags = row.tags.map((item) => {
        return h(
          NTag,
          {
            style: {
              marginRight: "6px",
            },
            type: "info",
            bordered: false,
          },
          {
            default: () => item.name,
          }
        );
      });
      return tags;
    },
  },
  {
    title: "头图",
    key: "img",
    render(row) {
      return h(NImage, { src: row.img });
    },
  },
]);

axios
  .get("/api/v1/article/list", {
    params: { mid: -1 },
  })
  .then((res) => {
    if (res.data.status == 200) {
      data.value.push(...res.data.data);
    }
  });
</script>

<style lang="scss" scoped></style>
