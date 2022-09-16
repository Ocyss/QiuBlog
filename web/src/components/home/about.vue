<template>
  <div class="about">
    <div class="img">
      <img class="avatar" src="/avatar.png" alt="" />
    </div>
    <div class="text">
      <div class="name">- - -</div>
      <div class="word">故事很短，满是遗憾。</div>
    </div>
  </div>
</template>

<script setup></script>

<style lang="scss" scoped>
.about {
  width: 95%;
  height: 220px;
  display: flex;
  align-items: center;
  flex-direction: column;
  border-radius: 8px;
  border-style: solid;
  border-width: 2px;
  border-color: rgb(236, 236, 236);
}
.img {
  width: 100%;
  display: flex;
  align-items: center;
  flex-direction: column;
  overflow: hidden;
  .avatar {
    width: 35%;
    height: auto;
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
    font-size: 1.1rem;
  }
}
</style>
