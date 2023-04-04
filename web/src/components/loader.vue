<template>
  <div class="loader-wrapper">
    <div class="loader"><img src="/static/img/avatar.png" alt="" /></div>
    <div class="loader-text">
      <div>L</div>
      <div>O</div>
      <div>A</div>
      <div>D</div>
      <div>I</div>
      <div>N</div>
      <div>G</div>
      <div></div>
      <div></div>
      <div></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
</script>

<style scoped lang="scss">
.loader-wrapper {
  /* 固定定位 */
  position: fixed;
  left: 0;
  top: 0;
  z-index: 999;
  width: 100%;
  height: 100%;
  /* 溢出隐藏 */
  overflow: hidden;
  /* 渐变背景 */
  background: linear-gradient(
    45deg,
    rgb(90, 54, 148) 0%,
    rgb(19, 189, 206) 33%,
    rgb(0, 148, 217) 66%,
    rgb(111, 199, 181) 100%
  );
  background-size: 400%;
  background-position: 0% 100%;
  /* 执行背景渐变动画：动画名 时长 加速后减速 无限次播放 */
  animation: gradient 7.5s ease-in-out infinite;
}
/* 旋转loading的外圈 */
.loader {
  width: 150px;
  height: 150px;
  border: 3px solid transparent;
  border-top-color: #fff;
  /* 相对定位 居中 */
  position: relative;
  left: 50%;
  top: 50%;
  margin-left: -75px;
  margin-top: -75px;
  z-index: 2;
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  /* 执行旋转动画 */
  animation: spin 1.7s linear infinite;
}
/* 旋转loading的中圈 */
.loader::before {
  content: "";
  /* 绝对定位 */
  position: absolute;
  top: 5px;
  left: 5px;
  bottom: 5px;
  right: 5px;
  border-radius: 50%;
  border: 3px solid transparent;
  border-top-color: #fff;
  /* 执行反向的旋转动画 reverse:反向 */
  animation: spin 0.6s linear infinite reverse;
}
/* 旋转loading的内圈 */
.loader::after {
  content: "";
  /* 绝对定位 */
  position: absolute;
  top: 15px;
  left: 15px;
  bottom: 15px;
  right: 15px;
  border-radius: 50%;
  border: 3px solid transparent;
  border-top-color: #fff;
  /* 执行旋转动画 */
  animation: spin 1s linear infinite;
}
/* logo */
.loader img {
  width: 55%;
  height: 55%;
  border-radius: 50%;
  /* 执行反向的旋转动画,时长必须和外圈的动画一致(不让其跟着旋转) */
  animation: spin 1.7s linear infinite reverse;
}
.loader-text {
  width: 50%;
  height: 36px;
  /* 绝对定位 水平居中 */
  position: absolute;
  top: 72%;
  left: 50%;
  transform: translateX(-50%);
  z-index: 3;
  /* 防止选取 */
  user-select: none;
}
.loader-text div {
  width: 30px;
  height: 36px;
  color: #fff;
  font-size: 32px;
  margin: 0 20px;
  /* 绝对定位 */
  position: absolute;
  /* 默认隐藏+旋转180度 */
  opacity: 0;
  transform: rotate(180deg);
  /* 执行文字移动动画 */
  animation: move 2s linear infinite;
}
/* 最后面的三个圆 */
.loader-text div:nth-child(8)::before,
.loader-text div:nth-child(9)::before,
.loader-text div:nth-child(10)::before {
  content: "";
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background-color: #fff;
  position: absolute;
  left: 0;
  bottom: 0;
}
/* 文字下的投影 */
.loader-text div::after {
  content: "";
  width: 10px;
  height: 5px;
  border-radius: 50%;
  background-color: rgba(255, 255, 255, 0.15);
  position: absolute;
  bottom: -40px;
  left: 50%;
  margin-left: -5px;
}
/* 最后面的三个圆的投影 */
.loader-text div:nth-child(8)::after,
.loader-text div:nth-child(9)::after,
.loader-text div:nth-child(10)::after {
  left: 0;
  margin-left: 0;
}
/* 接下来为各个文字设置动画延迟时间 */
.loader-text div:nth-child(2) {
  animation-delay: 0.2s;
}
.loader-text div:nth-child(3) {
  animation-delay: 0.4s;
}
.loader-text div:nth-child(4) {
  animation-delay: 0.6s;
}
.loader-text div:nth-child(5) {
  animation-delay: 0.8s;
}
.loader-text div:nth-child(6) {
  animation-delay: 1s;
}
.loader-text div:nth-child(7) {
  animation-delay: 1.2s;
}
.loader-text div:nth-child(8) {
  animation-delay: 1.4s;
}
.loader-text div:nth-child(9) {
  animation-delay: 1.6s;
}
.loader-text div:nth-child(10) {
  animation-delay: 1.8s;
}

/* 定义动画 */
/* 背景渐变动画 */
@keyframes gradient {
  50% {
    background-position: 100% 0%;
  }
}
/* 旋转动画 */
@keyframes spin {
  0% {
    transform: rotate(0);
  }
  100% {
    transform: rotate(360deg);
  }
}
/* 文字移动动画 */
@keyframes move {
  0% {
    right: 0;
    opacity: 0;
  }
  35% {
    right: 41%;
    opacity: 1;
    transform: rotate(0);
  }
  65% {
    right: 59%;
    opacity: 1;
    transform: rotate(0);
  }
  100% {
    right: 100%;
    transform: rotate(-180deg);
  }
}
</style>
