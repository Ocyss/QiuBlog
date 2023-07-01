<template>
  <n-space justify="space-between">
    <h1>留言板:</h1>
    <n-button type="success" @click="showModal = true">留言</n-button>
  </n-space>
  <n-divider />
  <n-grid x-gap="12" y-gap="12" cols="1 s:2 m:3 l:4" responsive="screen">
    <n-gi v-for="(item, index) in content" :key="item.id">
      <n-card>
        <n-thing>
          <template #avatar>
            <n-avatar round size="large" :src="item.qq
              ? `https://q.qlogo.cn/headimg_dl?dst_uin=${item.qq}&spec=640&img_type=jpg`
              : `https://api.multiavatar.com/${item.id}.png`
              " />
          </template>
          <template #header>{{ item.name }}</template>
          <template #header-extra>
            <n-button circle size="small" @click="() => like(item.id, index)">
              <template #icon>
                <ThumbsUpSharp />
              </template>
            </n-button>
          </template>
          <template #description>{{ item.email }}</template>
          {{ item.content }}
          <template #footer>
            <n-space justify="space-between">
              <TimerVue :t="item.created_at" />
              <span>点赞:{{ item.like }}</span>
            </n-space>
          </template>
        </n-thing>
      </n-card>
    </n-gi>
  </n-grid>
  <n-modal v-model:show="showModal" preset="dialog" title="发布留言：" content="你确认?" positive-text="确认" negative-text="算了"
    @positive-click="submitCallback">
    <n-space>
      <n-input v-model:value="data.name" :allow-input="noSpace" placeholder="*昵称"></n-input>
      <n-input v-model:value="data.email" :allow-input="noSpace" placeholder="Email"></n-input>
      <n-input v-model:value="data.qq" :allow-input="noSpace" placeholder="QQ"></n-input>
      <n-avatar round size="large" :src="data.qq
        ? `https://q.qlogo.cn/headimg_dl?dst_uin=${data.qq}&spec=640&img_type=jpg`
        : `https://api.multiavatar.com/${data.content}.png`
        " />
    </n-space>
    <n-divider />
    <n-input v-model:value="data.content" type="textarea" placeholder="*留言内容 最少10字" show-count minlength="10"
      maxlength="150" :allow-input="noSideSpace">
      <template #count="{ value }">
        <span style="font-size: 13px" :style="value.length < 10 ? `color: #FF0000` : ``">
          {{ value.length }} / 150
        </span>
      </template>
    </n-input>
    <GoCaptchaBtn class="go-captcha-btn" v-model:value="status" width="100%" height="80px" :image-base64="capt.image"
      style="margin-top: 20px;" :thumb-base64="capt.thumb" @confirm="handleConfirm" @refresh="handleRequestCaptCode" />
  </n-modal>
</template>

<script setup lang="ts">
import { ref, onMounted, onServerPrefetch } from "vue";
import { useMessage } from "naive-ui";
import api from "@/api";
import TimerVue from "@/components/Timer.vue";
import { ThumbsUpSharp } from "@vicons/ionicons5";
import lodash from 'lodash'
import GoCaptchaBtn from '@/components/admin/Captcha/GoCaptchaBtn.vue'

const message = useMessage();
const showModal = ref(false);
const content = ref([]);
const data = ref({
  name: "",
  qq: "",
  email: "",
  content: "",
  token: ""
});
const noSpace = (value) => !value || !/\s+/.test(value);
const noSideSpace = (value) => !value.startsWith(" ") && !value.endsWith(" ");
const capt = ref({
  image: "",
  thumb: "",
  key: "",
  autoRefreshCount: 0,
})
const status = ref("default")
function submitCallback() {
  if (data.value.content.length < 10) {
    message.error("字数不能少于10!!!");
    return false;
  } else if (data.value.name.split(" ").join("") == "") {
    message.error("昵称不可为空!!!");
    return false;
  } else if (!data.value.token) {
    message.error("辣么大个验证看不到!?!");
    return false;
  } else {

    api.message.addMessage(data.value).then((res) => {
      message.success("发布成功，等待审核！");
      data.value.name = "";
      data.value.qq = "";
      data.value.email = "";
      data.value.content = "";
    });
    data.value.token = ""
    return true;
  }
}
async function getMessage() {
  const params = { pagesize: 10, pagenum: 1 };
  const res = await api.message.getMessage(params);
  res.data.list.map((item) => {
    if (item.show == true || item.show == undefined) {
      content.value.push(item);
    }
  });
}

function handleRequestCaptCode() {
  api.utils.getCaptcha().then(res => {
    capt.value.image = res.data.image_base64
    capt.value.thumb = res.data.thumb_base64
    capt.value.key = res.data.captcha_key
  })
}
function handleConfirm(dots) {
  if (lodash.size(dots) <= 0) {
    message.warning(`请进行人机验证再操作`)
    return
  }

  let dotArr = []
  lodash.forEach(dots, (dot) => {
    dotArr.push(dot.x, dot.y)
  })

  api.utils.checkCaptcha(capt.value.key, dotArr.join(',')).then((res) => {

    message.success(`人机验证成功`)
    status.value = 'success'
    capt.value.autoRefreshCount = 0
    data.value.token = res.data

  }).catch(err => {
    message.warning(`人机验证失败`)
    if (capt.value.autoRefreshCount > 5) {
      capt.value.autoRefreshCount = 0
      status.value = 'over'
      return
    }
    capt.value.autoRefreshCount += 1
    status.value = 'error'
    console.log(capt.value);

  })
}

function like(id, index) {
  api.message.likeMessage("message", id).then(
    res => {
      if (res.data) {
        message.error(res.data)
      } else {
        content.value[index].like += 1
      }
    }
  )
}
onServerPrefetch(async () => {
  await getMessage();
});

onMounted(async () => {
  if (content.value.length == 0) {
    await getMessage();
  }
});
</script>

<style scoped lang="scss">
.card :deep(*) {
  padding: 5px;
}
</style>
