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
import { ref, h, resolveDynamicComponent } from "vue";
import { timeControl } from "@/utils";
import { NAvatar, NIcon, NSwitch, NButton } from "naive-ui";
import { railStyle } from "@/utils";
import { LogoSnapchat } from "@vicons/ionicons5";
import { useMessage, useDialog, NInput, NDivider } from "naive-ui";
import type { DataTableColumns } from "naive-ui";
import moment from "moment";
const dialog = useDialog();
const message = useMessage();
const Data = ref([]);
type Song = {
  id: number;
  created_at: string;
  updated_at: string;
  name: string;
  qq: string;
  email: string;
  question: string;
  reply: string;
  like: number;
  check: boolean;
  show: boolean;
};

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
  { title: "问题", key: "question", resizable: true },
  { title: "回答", key: "reply", resizable: true },
  {
    title: "昵称",
    key: "name",
    resizable: true,
    width: 100,
    render(row) {
      if (row.name) {
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
              src: `https://q.qlogo.cn/headimg_dl?dst_uin=${row.qq}&spec=640&img_type=jpg`,
            }),
          ]
        );
      } else {
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
            "匿名",
            h(NIcon, {
              round: true,
              size: 30,
              component: resolveDynamicComponent(LogoSnapchat),
            }),
          ]
        );
      }
    },
  },
  {
    title: "QQ | Email",
    key: "QE",
    width: 200,
    resizable: true,
    render(row) {
      return `${row.qq ? row.qq : "空"} | ${row.email ? row.email : "空"}`;
    },
  },
  {
    title: "",
    key: "reply",
    width: 80,
    render(row, index) {
      return h(
        NButton,
        {
          type: row.reply ? "info" : "success",
          size: "small",
          onClick: () => reply(row, index),
        },
        () => (row.reply ? "编辑" : "回复")
      );
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
  api.message.upMessage(row.id, val, true, false).then((res) => {
    message.success(res.message);
    row.show = !row.show;
  });
  loading = false;
}

function upCheck(val, row) {
  loading = true;
  if (val) {
    api.message.upMessage(row.id, val, false, false).then((res) => {
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
      api.message.delMessage(row.id, false).then((res) => {
        message.success(res.message);
        Data.value.splice(index, 1);
      });
    },
  });
}

function reply(row, index) {
  const content = ref(row.reply);
  dialog.warning({
    title: `回复:${row.id}`,
    content: () => {
      return h("div", {}, [
        row.question,
        h(NDivider),
        h(NInput, {
          type: "textarea",
          value: content.value,
          onUpdateValue: (v) => {
            content.value = v;
          },
        }),
      ]);
    },
    positiveText: "确定",
    negativeText: "取消",
    maskClosable: false,
    closeOnEsc: false,
    onPositiveClick: () => {
      api.message.replyQuestion(row.id, content.value).then((res) => {
        message.success(res.message);
        Data.value[index].reply = content.value;
      });
    },
  });
}
api.message.getQuestion(params).then((res) => {
  Data.value.push(...res.data);
});
</script>
