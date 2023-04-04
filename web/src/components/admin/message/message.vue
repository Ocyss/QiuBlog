<template>
  <n-data-table
    size="small"
    max-height="68vh"
    style="max-width: 93%; margin: auto"
    :columns="cols"
    :data="Data"
    :pagination="{
      pageSize: 10,
    }"
  />
</template>

<script setup lang="ts">
import api from "@/api";
import { ref, h, onMounted } from "vue";
import { timeControl } from "@/utils";
import { NAvatar, NButton, NSwitch } from "naive-ui";
import { railStyle } from "@/utils";
import moment from "moment";
import { useMessage, useDialog } from "naive-ui";
import type { DataTableColumns } from "naive-ui";
type Song = {
  id: number;
  created_at: string;
  updated_at: string;
  name: string;
  qq: string;
  email: string;
  content: string;
  like: number;
  check: boolean;
  show: boolean;
};
const Data = ref([]);
const message = useMessage();
const dialog = useDialog();

let loading = false;
const cols: DataTableColumns<Song> = [
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
      moment(row1.created_at).unix() - moment(row2.created_at).unix(),
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
        [
          row.name,
          h(NAvatar, {
            round: true,
            size: "small",
            src: row.qq
              ? `https://q.qlogo.cn/headimg_dl?dst_uin=${row.qq}&spec=640&img_type=jpg`
              : `https://api.multiavatar.com/${row.content}.png`,
          }),
        ]
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
    title: "操作",
    key: "operation",
    width: 150,
    render(row, index) {
      return h("div", { class: "operation" }, [
        h("div", { class: "switch" }, [
          h(
            NSwitch,
            {
              railStyle: railStyle,
              onUpdateValue: (val) => upShow(val, row),
              value: row.show,
              loading: loading,
              rubberBand: false,
            },
            {
              checked: () => "显示",
              unchecked: () => "隐藏",
            }
          ),
          h(
            NSwitch,
            {
              railStyle: railStyle,
              onUpdateValue: (val) => upCheck(val, row),
              value: row.check,
              loading: loading,
              rubberBand: false,
            },
            {
              checked: () => "过审",
              unchecked: () => "不过",
            }
          ),
        ]),
        h(
          NButton,
          {
            type: "error",
            size: "small",
            onClick: () => delmessage(row, index),
          },
          () => "删除"
        ),
      ]);
    },
  },
];
const params = { pagesize: 10, pagenum: 1 };

function upShow(val, row) {
  loading = true;
  api.message.upMessage(row.id, val, true, true).then((res) => {
    message.success(res.message);
    row.show = !row.show;
  });
  loading = false;
}

function upCheck(val, row) {
  loading = true;
  if (val) {
    api.message.upMessage(row.id, val, false, true).then((res) => {
      message.success(res.message);
      row.check = !row.check;
    });
  } else {
    message.error("不可以反审核");
  }

  loading = false;
}

function delmessage(row, index) {
  dialog.warning({
    title: "确认？",
    content: "你确定？删了就找不回来了哦！",
    positiveText: "确定",
    negativeText: "不确定",
    maskClosable: false,
    closeOnEsc: false,
    onPositiveClick: () => {
      api.message.delMessage(row.id, true).then((res) => {
        message.success(res.message);
        Data.value.splice(index, 1);
      });
    },
  });
}
onMounted(() => {
  api.message.getMessage(params).then((res) => {
    Data.value.push(...res.data);
  });
});
</script>

<style scoped lang="scss"></style>
