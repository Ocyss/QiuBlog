<template>
  <n-space justify="space-between">
    <h1>问答:</h1>
    <n-button type="success" @click="showModal = true">提问</n-button>
  </n-space>
  <n-divider />
  <n-grid x-gap="12" y-gap="12" cols="1 m:2 l:3" responsive="screen">
    <n-gi v-for="(item, index) in content" :key="item.id">
      <n-card>
        <n-thing>
          <template #avatar>
            <n-avatar round size="large" v-if="item.name" :src="item.qq
              ? `https://q.qlogo.cn/headimg_dl?dst_uin=${item.qq}&spec=640&img_type=jpg`
              : `https://api.multiavatar.com/${item.id}.png`
              "></n-avatar>
            <n-icon v-else size="38">
              <LogoSnapchat />
            </n-icon>
          </template>
          <template #header>{{ item.name ? item.name : "匿名" }}</template>
          <template #header-extra>
            <n-button circle size="small" @click="() => like(item.id, index)">
              <template #icon>
                <ThumbsUpSharp />
              </template>
            </n-button>
          </template>
          <template #description>{{ item.email }}</template>
          {{ item.question }}
          <n-divider></n-divider>
          {{ item.reply ? `回复：` + item.reply : "未回复" }}
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
  <n-modal v-model:show="showModal" preset="dialog" title="灵魂提问：" content="你确认?" positive-text="确认" negative-text="算了"
    @positive-click="submitCallback">
    <div style="margin-bottom: 20px">
      <n-switch v-model:value="switchData.sf" :rail-style="railStyle">
        <template #unchecked>公开</template>
        <template #checked>匿名</template>
        <template #icon>{{ switchData.sf ? `😈` : `😇` }}</template>
      </n-switch>
      <n-switch v-model:value="switchData.ty" :rail-style="railStyle" style="margin-left: 20%"
        @update:value="switchData.tyupdate">
        <template #unchecked>大众</template>
        <template #checked>私密</template>
        <template #icon>{{ switchData.ty ? `🙈` : `😊` }}</template>
      </n-switch>
    </div>

    <n-space v-if="!switchData.sf">
      <n-input v-model:value="data.name" :allow-input="noSpace" placeholder="*昵称"></n-input>
      <n-input v-model:value="data.email" :allow-input="noSpace" placeholder="Email"></n-input>
      <n-input v-model:value="data.qq" :allow-input="noSpace" placeholder="QQ"></n-input>
      <n-avatar round size="large" :src="data.qq
        ? `https://q.qlogo.cn/headimg_dl?dst_uin=${data.qq}&spec=640&img_type=jpg`
        : `https://api.multiavatar.com/${data.question}.png`
        " />
    </n-space>
    <n-divider />
    <n-input v-model:value="data.question" type="textarea" placeholder="*提问内容 最少10字" show-count minlength="10"
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
import { ThumbsUpSharp, LogoSnapchat } from "@vicons/ionicons5";
import { railStyle } from "@/utils";
import lodash from 'lodash'
import GoCaptchaBtn from '@/components/admin/Captcha/GoCaptchaBtn.vue'

const status = ref("default")
const message = useMessage();
const showModal = ref(false);
const content = ref([]);
const data = ref({
  name: "",
  qq: "",
  email: "",
  question: "",
  token: ""
});
const capt = ref({
  image: "",
  thumb: "",
  key: "",
  autoRefreshCount: 0,
})
const switchData = ref({
  sf: false,
  ty: false,
  tyupdate: (value) => {
    message.error("需要登陆完成，暂不实现");
  },
});
const noSpace = (value) => !value || !/\s+/.test(value);
const noSideSpace = (value) => !value.startsWith(" ") && !value.endsWith(" ");

function submitCallback() {
  if (data.value.question.length < 10) {
    message.error("字数不能少于10!!!");
    return false;
  } else if (
    !switchData.value.sf &&
    data.value.name.split(" ").join("") == ""
  ) {
    message.error("昵称不可为空!!!");
    return false;
  } else if (!data.value.token) {
    message.error("辣么大个验证看不到!?!");
    return false
  } else {
    api.message.addQuestion(data.value).then((res) => {
      if (res.status == 200) {
        message.success("发布成功，等待审核与回答！");
        data.value.name = "";
        data.value.qq = "";
        data.value.email = "";
        data.value.question = "";
      } else {
        message.error(res.message);
      }
    });
    return true;
  }
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
  api.message.likeMessage("question", id).then(
    res => {
      if (res.data) {
        message.error(res.data)
      } else {
        content.value[index].like += 1
      }
    }
  )
}
async function getQuestion() {
  const params = { pagesize: 10, pagenum: 1 };
  const res = await api.message.getQuestion(params);
  res.data.list.map((item) => {
    if (item.show == true || item.show == undefined) {
      content.value.push(item);
    }
  });
}

onServerPrefetch(async () => {
  await getQuestion();
});

onMounted(async () => {
  if (content.value.length == 0) {
    await getQuestion();
  }
});
</script>

<style scoped lang="scss"></style>
