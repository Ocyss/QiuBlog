<template>
  <div class="main">
    <div class="up">
      <KunVue style="width: 200px; height: 200px" />
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import KunVue from "./other/Kun.vue";
</script>

<style scoped lang="scss">
.main {
  width: 100%;
  height: 500px;
  display: flex;
  justify-content: center;
  align-content: center;
}

.up {
}
</style>
