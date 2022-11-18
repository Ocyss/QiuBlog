<template>
  <div>
    <n-button
      style="margin-bottom: 10px"
      type="primary"
      @click="router.push({ name: 'article-create' })"
    >
      文章发布
    </n-button>

    <n-data-table
      size="small"
      max-height="68vh"
      :columns="colsReactive"
      :data="data"
      :pagination="{
        pageSize: 10,
      }"
      :row-props="rowProps"
    />
    <n-dropdown
      placement="bottom-start"
      trigger="manual"
      :x="dropdown.x"
      :y="dropdown.y"
      :options="dropdown.options"
      :show="dropdown.showDropdown"
      :on-clickoutside="onClickoutside"
      @select="handleSelect"
    />
  </div>
</template>

<script setup>
import { useRouter } from "vue-router";
import { h, reactive, nextTick } from "vue";
import { NButton, NTag, NImage } from "naive-ui";
import { timeControl } from "@/utils";
import api from "@/api";
const router = useRouter();
const data = ref([]);
const dropdown = ref({
  x: 0,
  y: 0,
  showDropdown: false,
  options: [
    {
      label: "编辑",
      key: "edit",
    },
    {
      label: () => h("span", { style: { color: "red" } }, "删除"),
      key: "delete",
    },
  ],
  row: { id: -1 },
});
const onClickoutside = () => {
  dropdown.value.showDropdown = false;
};
const handleSelect = (key) => {
  dropdown.value.showDropdown = false;
  if (key == "edit") {
    router.push({
      name: "article-updata",
      params: {
        id: dropdown.value.row.id,
      },
    });
  }
};
const rowProps = (row) => {
  return {
    onContextmenu: (e) => {
      e.preventDefault();
      dropdown.value.showDropdown = false;
      nextTick().then(() => {
        dropdown.value.showDropdown = true;
        dropdown.value.x = e.clientX;
        dropdown.value.y = e.clientY;
        dropdown.value.row = row;
      });
    },
  };
};

const colsReactive = reactive([
  {
    title: "ID",
    key: "id",
    width: 50,
    sorter: (row1, row2) => row1.id - row2.id,
  },
  {
    title: "发布时间",
    key: "created_at",
    resizable: true,
    width: 100,
    render(row) {
      return timeControl(row.created_at);
    },
    sorter: (row1, row2) =>
      new Date(row1.created_at) - new Date(row2.created_at),
  },
  {
    title: "更新时间",
    key: "updated_at",
    resizable: true,
    width: 100,
    render(row) {
      return timeControl(row.updated_at);
    },
    sorter: (row1, row2) =>
      new Date(row1.updated_at) - new Date(row2.updated_at),
  },
  { title: "标题", key: "title", minWidth: 150, resizable: true },
  {
    title: "简介",
    key: "desc",
    ellipsis: {
      tooltip: true,
    },
    resizable: true,
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
              marginBottom: "5px",
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
    resizable: true,
  },
  {
    title: "头图",
    key: "img",
    render(row) {
      return h(NImage, { width: 80, src: row.img });
    },
  },
]);

api.article.getList({ mid: -2 }).then((res) => {
  res.data.map((item) => {
    data.value.push(item);
  });
});
</script>

<style lang="scss" scoped></style>
