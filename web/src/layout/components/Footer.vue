<template>
  <div class="main">
    <div class="left">
      <div>2023 ©{{ $t("info.footer.copyright") }} Q</div>
      <div class="record">
        <img
          style="width: 20px; height: 20px"
          src="https://icp.gov.moe/images/ico64.png"
        />
        <a href="https://icp.gov.moe/?keyword=20231028" target="_blank">
          {{ $t("info.footer.icp.moe") }}
        </a>
      </div>

      <div>
        Go+Vue{{ $t("info.footer.info") }}:
        <a href="https://github.com/qiu-lzsnmb/QiuBlog">qiu-lzsnmb/QiuBlog</a>
      </div>
    </div>
    <div class="right"></div>
  </div>
</template>

<script setup lang="ts">
import { ref, inject } from "vue";
import { useDesignSettingStore } from "@/store/modules/designSetting";

const designStore = useDesignSettingStore();
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
