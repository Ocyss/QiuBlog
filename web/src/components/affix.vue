<template>
  <n-back-top
    to="#affixContent"
    ref="backTopRef"
    v-if="scrollableEl"
    :listen-to="scrollableEl"
    style="z-index: 5000"
  />
</template>

<script setup lang="ts">
import { inject, Ref, ref, provide, readonly } from "vue";

const backTopRef = ref(void 0);
const scrollableEl: Ref<HTMLElement> = inject("scrollableEl");

function backTop() {
  backTopRef.value?.handleClick();
}

defineExpose({
  backTop,
});
</script>

<style lang="scss">
#affixContent {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
  position: fixed;
  right: 30%;
  height: 30vh;
  bottom: 0;
  z-index: 999;

  div {
    width: 44px;
    height: 44px;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: var(--n-color);
    border-radius: 50%;
    margin: 5px 0px;
    box-shadow: 0 2px 8px 0px rgba(0, 0, 0, 0.12);
    cursor: pointer;
    font-size: 20px;
    &:hover {
      color: var(--primary-color);
    }
    position: static;
    transition: all 0.3s;
  }
}
</style>
