<template>
  <adminVue>
    <div class="article" style="border: 1px solid #ccc">
      <n-space vertical>
        <n-space>
          <n-input
            v-model:value="content.title"
            type="text"
            placeholder="标题"
          />
          <n-input
            v-model:value="content.desc"
            type="text"
            placeholder="描述"
          />
          <n-cascader
            v-model:value="content.cid"
            placeholder="选择发布类别"
            :options="menuoptions"
            check-strategy="child"
            expand-trigger="hover"
            :show-path="true"
            :filterable="true"
            @update:value="handleUpdateValue"
          />
          <n-upload
            action="/api/v1/upload/image"
            :default-file-list="fileList"
            list-type="image"
            accept="image/*"
            :data="{ class: 'Article' }"
            :max="1"
            with-credentials
            :custom-request="customRequest"
          >
            <n-button>上传头图</n-button>
          </n-upload>
        </n-space>

        <n-dynamic-tags v-model:value="content.tags" />
        <Toolbar
          style="border-bottom: 1px solid #ccc"
          :editor="editorRef"
          :defaultConfig="toolbarConfig"
          :mode="mode"
        />
        <Editor
          style="height: 500px; overflow-y: hidden"
          v-model="content.content"
          :defaultConfig="editorConfig"
          :mode="mode"
          @onCreated="handleCreated"
        />
        <n-button type="success" @click="send">发布</n-button>
      </n-space>
    </div>
  </adminVue>
</template>

<script setup>
import "@wangeditor/editor/dist/css/style.css";
import { onBeforeUnmount, ref, shallowRef, onMounted } from "vue";
import { Editor, Toolbar } from "@wangeditor/editor-for-vue";
import adminVue from "@/layout/admin.vue";
import axios from "axios";
import { useMessage } from "naive-ui";
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
const content = ref({
  tags: [],
  cid: null,
  desc: "",
  title: "",
  content: "",
  img: "",
});
const fileList = ref([]);
const handleCreated = (editor) => {
  editorRef.value = editor; // 记录 editor 实例，重要！
};
const handleUpdateValue = (value, option) => {
  console.log(value, option);
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
  axios
    .post(action, formData)
    .then((res) => {
      if (res.data.errno == 0) {
        console.log(res.data);
        content.value.img = res.data.data.url;
        onFinish();
      } else {
        message.error(res.data.message);
        onError();
      }
    })
    .catch((error) => {
      message.error(error.message);
      onError();
    });
};
function send() {
  axios.post("/api/v1/article/add", content.value).then((res) => {
    console.log(res.data);
  });
}
onBeforeUnmount(() => {
  const editor = editorRef.value;
  if (editor == null) return;
  editor.destroy();
});

onMounted(() => {
  axios.get("/api/v1/menuchild").then((res) => {
    if (res.data.status == 200) {
      axios.get("/api/v1/category").then((res2) => {
        if (res2.data.status == 200) {
          res.data.data.map((item) => {
            menuoptions.value.push({
              value: item.name,
              label: item.name,
              children: res2.data.data
                .map((item2) => {
                  if (item2.mid == item.ID) {
                    return { value: item2.ID, label: item2.name };
                  }
                })
                .filter(Boolean),
            });
          });
        } else {
          message.error(res.data.message);
        }
      });
    } else {
      message.error(res.data.message);
    }
  });
});
</script>

<style lang="scss" scoped>
.article {
  width: 80%;
  margin: auto;
}
</style>
