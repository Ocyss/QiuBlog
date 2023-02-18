<template>
  <div class="main">
    <div class="left">
      <div>2023 ©{{ designStore.getLocale ? "版权所有" : "Copyright" }} Q</div>
      <div class="record">
        <img
          style="width: 20px; height: 20px"
          src="https://icp.gov.moe/images/ico64.png"
        />
        <a href="https://icp.gov.moe/?keyword=20231028" target="_blank">
          {{ designStore.getLocale ? "萌ICP备20231028号" : "MoeICP:20231028" }}
        </a>
      </div>

      <div>
        Go+Vue{{ designStore.getLocale ? "驱动 开源地址" : " Address" }}:
        <a href="https://github.com/qiu-lzsnmb/QiuBlog">qiu-lzsnmb/QiuBlog</a>
      </div>
    </div>
    <div class="right"></div>
  </div>
</template>

<script setup>
import { ref } from "vue";
const designStore = inject("designStore");
</script>

<style scoped lang="scss">
.main {
  width: 100%;
  min-height: 100px;
  display: flex;
  justify-content: flex-end;
  align-content: center;
}
.left {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
}
.record {
  display: flex;
  justify-content: center;
  align-items: center;
}
a {
  text-decoration: none;
  color: var(--n-item-text-color);
  background-color: transparent;
}
.right {
  max-width: 320px;
  width: 30%;
}

@media screen and (max-width: 700px) {
  .right {
    width: 0px;
  }
}
</style>
