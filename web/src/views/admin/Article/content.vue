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
            <n-input
              v-model:value="imageurl"
              type="text"
              placeholder="图片url"
            />
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
      <md-editor
        v-model="content.content"
        :theme="designStore.getDarkTheme ? `dark` : `light`"
        showCodeRowNumber
        @onUploadImg="onUploadImg"
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

<script setup lang="ts">
import { ref, onMounted, inject, computed, onServerPrefetch } from "vue";
import { useMessage } from "naive-ui";
import { useRouter, useRoute } from "vue-router";
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import api from "@/api";
import { useDesignSettingStore } from "@/store/modules/designSetting";

const designStore = useDesignSettingStore();
const router = useRouter();
const route = useRoute();
const message = useMessage();
const menuoptions = ref([]);
const tags = ref([]);
const fileList = ref([]);
const uploadref = ref();
const imageurl = ref();

MdEditor.config({});

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

function send() {
  api.article.add(content.value).then((res) => {
    message.success("发布成功！！！");
    uploadref.value.clear();
    tags.value = [];
    content.value = {
      tags: null,
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

async function onUploadImg(files, callback) {
  const res = await Promise.all(
    files.map((file) => {
      return new Promise((rev, rej) => {
        api
          .upload(file, "Article")
          .then((res) => {
            if (res.status == 200) {
              rev(res.data);
            } else {
              rej(res.message);
            }
          })
          .catch((error) => rej(error));
      });
    })
  );
  callback(res);
}

function customRequest({
  file,
  data,
  headers,
  withCredentials,
  action,
  onFinish,
  onError,
}) {
  const formData = new FormData();
  if (data) {
    Object.keys(data).forEach((key) => {
      formData.append(key, data[key]);
    });
  }
  formData.append("file", file.file);
  api.request
    .post(action, formData)
    .then((res) => {
      if (res.status == 200) {
        content.value.img = res.data;
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
}
async function getMC() {
  //请求菜单项和分类;
  const mres = await api.menuchild.gets();
  const cres = await api.category.get(false);
  mres.data.map((item) => {
    menuoptions.value.push({
      value: item.name,
      label: item.name,
      children: cres.data
        .map((item2) => {
          if (item2.mid == item.id) {
            return { value: item2.id, label: item2.name };
          }
        })
        .filter((item) => typeof item !== "undefined"),
    });
  });
}

onServerPrefetch(() => {
  getMC();
});

onMounted(() => {
  if (!menuoptions.value) {
    getMC();
  }
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
.md-editor-dark {
  --md-bk-color: #333 !important;
  --md-color: rgba(255, 255, 255, 0.82) !important;
}
</style>
