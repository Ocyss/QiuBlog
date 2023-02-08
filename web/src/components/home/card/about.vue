<template>
  <n-card class="about">
    <div class="img">
      <img class="avatar" src="/img/avatar.png" alt="" />
    </div>
    <div class="text">
      <div class="name">- - -</div>
      <div class="word">故事很短，满是遗憾。</div>
    </div>
  </n-card>
</template>

<script setup></script>

<style lang="scss" scoped>
.img {
  width: 100%;
  display: flex;
  align-items: center;
  flex-direction: column;
  overflow: hidden;
  .avatar {
    width: 50%;
    height: auto;
    transition: all 0.8s;
  }
  .avatar:hover {
    transform: rotate(360deg);
  }
}
.text {
  width: 100%;
  display: flex;
  align-items: center;
  flex-direction: column;
  .name {
    font-size: 2rem;
  }
  .word {
    font-size: 1rem;
    white-space: nowrap;
    text-shadow: 0px 0px 5px #00ffff, 0px 0px 15px #00ffff, 0px 0px 25px #00ffff;
  }
}
</style>
