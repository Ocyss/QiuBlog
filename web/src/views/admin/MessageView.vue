<template>
  <n-tabs
    default-value="signin"
    size="large"
    justify-content="space-evenly"
    class="main"
  >
    <n-tab-pane name="signin" tab="留言">
      <n-data-table
        size="small"
        max-height="68vh"
        style="max-width: 93%; margin: auto"
        :columns="colsM"
        :data="mData"
        :pagination="{
          pageSize: 10,
        }"
        :row-props="rowProps"
      />
    </n-tab-pane>
    <n-tab-pane name="signup" tab="问答">
      <!-- <n-data-table
        size="small"
        max-height="68vh"
        :columns="colsQ"
        :data="qData"
        :pagination="{
          pageSize: 10,
        }"
        :row-props="rowProps"
      /> -->
      问答！
    </n-tab-pane>
  </n-tabs>
</template>

<script setup>
import api from "@/api";
import { ref, h } from "vue";
import { timeControl } from "@/utils";
import { NButton, NAvatar, NSwitch } from "naive-ui";
const mData = ref([]);
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

const colsM = reactive([
  {
    title: "ID",
    key: "id",
    width: 50,
    sorter: (row1, row2) => row1.id - row2.id,
  },
  {
    title: "留言时间",
    key: "created_at",
    resizable: true,
    width: 100,
    render(row) {
      return timeControl(row.created_at);
    },
    sorter: (row1, row2) =>
      new Date(row1.created_at) - new Date(row2.created_at),
  },
  { title: "内容", key: "content", resizable: true },
  {
    title: "昵称",
    key: "name",
    width: 100,
    render(row) {
      return h(
        "div",
        {
          style: {
            display: "flex",
            alignItems: "center",
            maxWidth: "80px",
          },
        },
        row.name,
        h(NAvatar, {
          round: true,
          size: "small",
          src: `http://q.qlogo.cn/headimg_dl?dst_uin=${row.qq}&spec=640&img_type=jpg`,
        })
      );
    },
  },
  {
    title: "QQ | Email",
    key: "QE",
    width: 100,
    render(row) {
      return `${row.qq ? row.qq : "空"} | ${row.email ? row.email : "空"}`;
    },
  },
  {
    title: "是否显示",
    key: "show",
    width: 100,
    render(row) {
      return h(
        NSwitch,
        {
          railStyle: ({ focused, checked }) => {
            const style = {};
            if (checked) {
              style.background = "#d03050";
              if (focused) {
                style.boxShadow = "0 0 0 2px #d0305040";
              }
            } else {
              style.background = "#1d46ff";
              if (focused) {
                style.boxShadow = "0 0 0 2px #2080f040";
              }
            }
            return style;
          },
          defaultValue: row.show,
        },
        {
          checked: () => "显示",
          unchecked: () => "隐藏",
        }
      );
    },
  },
  {
    title: "审核状态",
    key: "check",
    width: 100,
    render(row) {
      return h(
        NSwitch,
        {
          railStyle: ({ focused, checked }) => {
            const style = {};
            if (checked) {
              style.background = "#d03050";
              if (focused) {
                style.boxShadow = "0 0 0 2px #d0305040";
              }
            } else {
              style.background = "#1d46ff";
              if (focused) {
                style.boxShadow = "0 0 0 2px #2080f040";
              }
            }
            return style;
          },
          defaultValue: row.check,
        },
        {
          checked: () => "已审核",
          unchecked: () => "未审核",
        }
      );
    },
  },
]);
const params = { pagesize: 10, pagenum: 1 };
api.message.getMessage(params).then((res) => {
  mData.value.push(...res.data);
});
</script>

<style scoped lang="scss"></style>
