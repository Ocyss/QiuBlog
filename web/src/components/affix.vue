<template>
  <n-el tag="div" class="affix">
    <Transition>
      <div id="affixContent"></div>
    </Transition>
    <n-back-top
      to="#affixContent"
      ref="backTopRef"
      v-if="scrollableEl"
      :listen-to="scrollableEl"
      style="z-index: 5000"
    />
  </n-el>
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

<style scoped lang="scss">
.affix {
  position: fixed;
  right: 15%;
  height: 30vh;
  bottom: 0;
  z-index: 999;
}
</style>

<style lang="scss">
#affixContent {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
  div {
    width: 44px;
    height: 44px;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: var(--popover-color);
    border-radius: 50%;
    margin: 5px 0px;
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
