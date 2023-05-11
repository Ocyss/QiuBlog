<template>
  <div class="main">
    <div class="title">
      <h1>仪表盘</h1> <span @click="logOut">注销</span>
    </div>

    <div class="digitalData">
      <n-grid cols="1 s:2 m:4" responsive="screen" x-gap="12" y-gap="12">
        <n-grid-item>
          <n-card :title="`😎${designStore.getLocale ? '浏览量' : 'pageView'}`">
            卡片内容
          </n-card>
        </n-grid-item>
        <n-grid-item>
          <n-card :title="`📖${designStore.getLocale ? '文章数' : 'articleCount'}`">
            卡片内容
          </n-card>
        </n-grid-item>
        <n-grid-item>
          <n-card :title="`🧾${designStore.getLocale ? '总字数' : 'wordsTotal'}`">
            卡片内容
          </n-card>
        </n-grid-item>
        <n-grid-item>
          <n-card :title="`📨${designStore.getLocale ? '留言/问答' : 'message/Q&A'}`">
            卡片内容
          </n-card>
        </n-grid-item>
      </n-grid>
    </div>
    <div>
      <n-table striped>
        <tbody>
          <tr>
            <td>
              <h2>网站信息</h2>
            </td>
          </tr>
          <tr>
            <td>当前用户</td>
            <td>{{ "Admin" }}</td>
            <td>当前版本</td>
            <td>{{ "1.6.6" }}</td>
          </tr>
          <tr>
            <td>文章总数</td>
            <td>...</td>
            <td>分类总数</td>
            <td>...</td>
          </tr>
          <tr>
            <td>标签总数</td>
            <td>...</td>
            <td>浏览总数</td>
            <td>...</td>
          </tr>
          <tr>
            <td>留言总数</td>
            <td>...</td>
            <td>问答总数</td>
            <td>...</td>
          </tr>
        </tbody>
      </n-table>
    </div>
  </div>
</template>

<script setup lang="ts">
import { inject, ref } from "vue";
import { useDesignSettingStore } from "@/store/modules/designSetting";
import { VueCookies } from "vue-cookies";
import { useRouter } from "vue-router";
const $cookie = inject<VueCookies>("$cookies");
const designStore = useDesignSettingStore();
const router = useRouter()
function logOut() {
  $cookie.remove("token")
  router.push({ name: "home" })
}
</script>

<style scoped lang="scss">
.main>div {
  margin: 15px 0;
}

.title {
  display: flex;
  align-items: center;
  justify-content: space-between;

  span {
    cursor: pointer;

    &:hover {
      color: rgb(42, 114, 248);
    }
  }
}
</style>
