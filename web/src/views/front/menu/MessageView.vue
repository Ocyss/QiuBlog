<template>
  <n-space justify="space-between">
    <h1>留言板:</h1>
    <n-button type="success" @click="showModal = true">留言</n-button>
  </n-space>
  <n-divider />
  <n-grid x-gap="12" y-gap="12" cols="1 s:2 m:3 l:4" responsive="screen">
    <n-gi v-for="item in content" :key="item.id">
      <n-card>
        <n-thing>
          <template #avatar>
            <n-avatar
              round
              size="large"
              :src="
                item.qq
                  ? `https://q.qlogo.cn/headimg_dl?dst_uin=${item.qq}&spec=640&img_type=jpg`
                  : `https://api.multiavatar.com/${item.content}.png`
              "
            />
          </template>
          <template #header>{{ item.name }}</template>
          <template #header-extra>
            <n-button circle size="small">
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
  <n-modal
    v-model:show="showModal"
    preset="dialog"
    title="发布留言："
    content="你确认?"
    positive-text="确认"
    negative-text="算了"
    @positive-click="submitCallback"
  >
    <n-space>
      <n-input
        v-model:value="data.name"
        :allow-input="noSpace"
        placeholder="*昵称"
      ></n-input>
      <n-input
        v-model:value="data.email"
        :allow-input="noSpace"
        placeholder="Email"
      ></n-input>
      <n-input
        v-model:value="data.qq"
        :allow-input="noSpace"
        placeholder="QQ"
      ></n-input>
      <n-avatar
        round
        size="large"
        :src="
          data.qq
            ? `https://q.qlogo.cn/headimg_dl?dst_uin=${data.qq}&spec=640&img_type=jpg`
            : `https://api.multiavatar.com/${data.content}.png`
        "
      />
    </n-space>
    <n-divider />
    <n-input
      v-model:value="data.content"
      type="textarea"
      placeholder="*留言内容 最少10字"
      show-count
      minlength="10"
      maxlength="150"
      :allow-input="noSideSpace"
    >
      <template #count="{ value }">
        <span
          style="font-size: 13px"
          :style="value.length < 10 ? `color: #FF0000` : ``"
        >
          {{ value.length }} / 150
        </span>
      </template>
    </n-input>
  </n-modal>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useMessage } from "naive-ui";
import api from "@/api";
import TimerVue from "@/components/Timer.vue";
import { ThumbsUpSharp } from "@vicons/ionicons5";
const message = useMessage();
const showModal = ref(false);
const content = ref([]);
const data = ref({
  name: "",
  qq: "",
  email: "",
  content: "",
});
const noSpace = (value) => !value || !/\s+/.test(value);
const noSideSpace = (value) => !value.startsWith(" ") && !value.endsWith(" ");
function submitCallback() {
  if (data.value.content.length < 10) {
    message.error("字数不能少于10!!!");
    return false;
  } else if (data.value.name.split(" ").join("") == "") {
    message.error("昵称不可为空!!!");
    return false;
  } else {
    api.message.addMessage(data.value).then((res) => {
      if (res.status == 200) {
        message.success("发布成功，等待审核！");
        data.value.name = "";
        data.value.qq = "";
        data.value.email = "";
        data.value.content = "";
      } else {
        message.error(res.message);
      }
    });
    return true;
  }
}

const params = { pagesize: 10, pagenum: 1 };
const res = await api.message.getMessage(params);
res.data.map((item) => {
  if (item.show == true || item.show == undefined) {
    content.value.push(item);
  }
});
</script>

<style scoped lang="scss">
.card :deep(*) {
  padding: 5px;
}
</style>
