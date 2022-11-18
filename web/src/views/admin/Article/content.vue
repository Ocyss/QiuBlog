<template>
  <div class="content" style="border: 1px solid #ccc">
    <n-space vertical>
      <n-button class="back" @click="router.push({ name: 'article' })">
        返回
      </n-button>
      <n-grid x-gap="12" y-gap="12" cols="2">
        <n-gi span="2">
          <n-input
            size="large"
            v-model:value="content.title"
            type="text"
            placeholder="标题"
          />
        </n-gi>
        <n-gi>
          <n-input
            v-model:value="content.desc"
            type="textarea"
            placeholder="描述"
          />
        </n-gi>
        <n-gi>
          <n-space vertical>
            <n-cascader
              v-model:value="content.cid"
              placeholder="选择发布类别"
              :options="menuoptions"
              check-strategy="child"
              expand-trigger="hover"
              :show-path="true"
              :filterable="true"
            />
            <!-- <n-input
              v-model:value="imageurl"
              type="text"
              placeholder="图片url"
            /> -->
          </n-space>
        </n-gi>
        <n-gi span="2">
          <n-space>
            <n-upload
              ref="uploadref"
              action="/api/v1/upload/image"
              :default-file-list="fileList"
              list-type="image-card"
              accept="image/*"
              :data="{ class: 'Article' }"
              :max="1"
              with-credentials
              :custom-request="customRequest"
            >
              <n-button>上传头图</n-button>
            </n-upload>
            <n-dynamic-tags v-model:value="tags" />
          </n-space>
        </n-gi>
      </n-grid>

      <Toolbar
        class="weTool"
        :editor="editorRef"
        :defaultConfig="toolbarConfig"
        :mode="mode"
      />
      <Editor
        class="WeEditor"
        style="height: 500px; overflow-y: hidden"
        v-model="content.content"
        :defaultConfig="editorConfig"
        :mode="mode"
        @onCreated="handleCreated"
      />

      <n-button
        v-if="route.name == 'article-updata'"
        type="success"
        @click="save"
      >
        保存
      </n-button>
      <n-button v-else type="success" @click="send">发布</n-button>
    </n-space>
  </div>
</template>

<script setup>
import "@wangeditor/editor/dist/css/style.css";
import { onBeforeUnmount, ref, shallowRef, onMounted } from "vue";
import { Editor, Toolbar } from "@wangeditor/editor-for-vue";
import request from "@/utils/request";
import { useMessage } from "naive-ui";
import { useRouter, useRoute } from "vue-router";
import api from "@/api";
const router = useRouter();
const route = useRoute();
const message = useMessage();
const editorRef = shallowRef();
const mode = ref("default");
const menuoptions = ref([]);
const toolbarConfig = {};
const editorConfig = {
  placeholder: "请输入内容...",
  MENU_CONF: {
    uploadImage: {
      fieldName: "file",
      // 上传图片的配置
      server: "/api/v1/upload/image",
      maxNumberOfFiles: 15,
      maxFileSize: 10 * 1024 * 1024, // 10M
      meta: {
        class: "Article",
      },
      withCredentials: true,
      timeout: 5 * 1000, // 5 秒
      base64LimitSize: 5 * 1024, // 5kb
    },
  },
};
const tags = ref([]);
const fileList = ref([]);
const uploadref = ref();
// const imageurl = computed({
//   get: () => {
//     return fileList.value.length == 0 ? "" : fileList.value[0].url;
//   },
//   set: (val) => {
//     fileList.value = [
//       {
//         id: val,
//         status: "finished",
//         url: val,
//       },
//     ];
//   },
// });
const content = ref({
  tags: computed(() => {
    return tags.value.map((item) => {
      return { name: item };
    });
  }),
  cid: "",
  desc: "",
  title: "",
  content: "",
  img: "",
});

const handleCreated = (editor) => {
  editorRef.value = editor; // 记录 editor 实例，重要！
};
const customRequest = ({
  file,
  data,
  headers,
  withCredentials,
  action,
  onFinish,
  onError,
}) => {
  const formData = new FormData();
  if (data) {
    Object.keys(data).forEach((key) => {
      formData.append(key, data[key]);
    });
  }
  formData.append("file", file.file);
  request
    .post(action, formData)
    .then((res) => {
      if (res.errno == 0) {
        content.value.img = res.data.url;
        onFinish();
      } else {
        message.error(res.message);
        onError();
      }
    })
    .catch((error) => {
      message.error(error.message);
      onError();
    });
};

function send() {
  api.article.add(content.value).then((res) => {
    message.success("发布成功！！！");
    uploadref.value.clear();
    tags.value = [];
    content.value = {
      tags: computed({
        get: () =>
          tags.value.map((item) => {
            return { name: item };
          }),
      }),
      cid: "",
      desc: "",
      title: "",
      content: "",
      img: "",
    };
  });
}

function save() {
  api.article.put(route.params.id, content.value).then((res) => {
    message.success("保存成功！！！");
    router.push({ name: "article" });
  });
}
onBeforeUnmount(() => {
  //关闭前注销编辑器组件
  const editor = editorRef.value;
  if (editor == null) return;
  editor.destroy();
});

onMounted(() => {
  //请求菜单项和分类;
  api.menuchild.gets().then((res) => {
    api.category.get(false).then((res2) => {
      res.data.map((item) => {
        menuoptions.value.push({
          value: item.name,
          label: item.name,
          children: res2.data
            .map((item2) => {
              if (item2.mid == item.id) {
                return { value: item2.id, label: item2.name };
              }
            })
            .filter((item) => typeof item !== "undefined"),
        });
      });
    });
  });

  console.log(menuoptions.value);
  //判断是不是修改帖子
  if (route.name == "article-updata") {
    api.article.get(route.params.id).then((res) => {
      content.value.cid = res.data.cid;
      content.value.desc = res.data.desc;
      content.value.title = res.data.title;
      content.value.content = res.data.content;
      res.data.tags.map((item) => {
        tags.value.push(item.name);
      });
      fileList.value.push({
        id: 0,
        status: "finished",
        url: res.data.img,
      });
    });
  }
});
</script>

<style lang="scss" scoped>
.content {
  width: 80%;
  margin: auto;
  border: 0px !important;
}

.back {
  position: absolute;
  top: 6px;
  left: 4px;
}
</style>

<style lang="scss">
:root,
:host {
  // 编辑器
  --w-e-textarea-bg-color: #333;
  --w-e-textarea-color: #fff;
  //工具栏
  --w-e-toolbar-color: #fff;
  --w-e-toolbar-bg-color: #333;
  --w-e-toolbar-border-color: #666; //工具栏分割线
  --w-e-toolbar-active-bg-color: #666; //选中颜色
  --w-e-toolbar-active-color: #fff; //悬浮提示颜色
}
</style>
