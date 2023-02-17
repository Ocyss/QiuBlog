<template>
  <div class="main">
    <n-gradient-text type="info">网页去外太空旅行了</n-gradient-text>
    <div id="lottie_box" class="animation" />
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import lottie from "lottie-web";
import err404 from "@/assets/exception/404.json";
//https://www.iconfont.cn/lotties/detail?spm=a313x.7781069.1998910419.d9df05512&cid=39334
onMounted(() => {
  const jsonAnimation = lottie.loadAnimation({
    container: document.querySelector("#lottie_box"), //进行播放的元素
    renderer: "svg",
    loop: true, //循环播放
    autoplay: true, //自动播放
    animationData: err404, //要播放的文件
  });
  jsonAnimation.play();
});
</script>

<style scoped lang="scss">
.main {
  width: 100%;
  min-height: 95vh;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  font-size: 32px;
}
.animation {
  height: 70vh;
}
</style>
