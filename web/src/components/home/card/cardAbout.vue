<template>
  <n-card class="about">
    <div class="img">
      <img class="avatar" src="/static/img/avatar.png" alt="" />
    </div>
    <div class="text">
      <div class="name">{{ $t("main.name") }}</div>
      <div class="motto">
        {{ $t("main.motto") }}
      </div>
    </div>
  </n-card>
</template>

<script setup lang="ts"></script>

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
  .motto {
    font-size: 1rem;
    white-space: nowrap;
    text-shadow: 0px 0px 5px #00ffff, 0px 0px 15px #00ffff, 0px 0px 25px #00ffff;
  }
}
</style>
