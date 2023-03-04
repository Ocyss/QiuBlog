<template>
  <div class="comment">
    <n-space class="space">
      <div class="spaceEl">
        用户名：
        <n-input v-model:value="user.name" type="text" placeholder="必填" />
      </div>
      <div class="spaceEl">
        QQ：
        <n-input v-model:value="user.qq" type="text" placeholder="二选一" />
      </div>
      <div class="spaceEl">
        邮箱：
        <n-input v-model:value="user.qq" type="text" placeholder="二选一" />
      </div>
    </n-space>
    <n-button style="margin: 20px" type="info">保存</n-button>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from "vue";
import { useDesignSettingStore } from "@/store/modules/designSetting";
const designStore = useDesignSettingStore();
const user = ref({ name: "", qq: "", email: "" });
const login = ref(false);
</script>

<style lang="scss" scoped>
.space {
  width: 100%;
  justify-content: space-around !important;
  align-items: center;
  > div {
    display: flex;
    width: 33%;
    align-items: center;
  }
  .spaceEl {
    white-space: nowrap;
  }
}
.comment {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
}
</style>
