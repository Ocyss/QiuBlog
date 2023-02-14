<template>
  <div class="logo" @mousedown="start" @mouseup="stop">
    <img
      src="/favicons/android-icon-192x192.png"
      alt=""
      :class="collapsed ? 'small' : 'big'"
    />
  </div>
</template>

<script setup>
import { useRouter } from "vue-router";

const router = useRouter();
let presstimer = null;
const props = defineProps(["collapsed"]);
function start() {
  if (presstimer != null) {
    clearTimeout(presstimer);
    presstimer = null;
  }
  presstimer = setTimeout(() => {
    router.push({ name: "admin" });
  }, 3000);
}

function stop() {
  if (presstimer != null) {
    clearTimeout(presstimer);
    presstimer = null;
  }
}
</script>

<style lang="scss" scoped>
.logo {
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  white-space: nowrap;
  width: 100%;
  img {
    margin-top: 1rem;
    width: auto;
    height: 160px;
  }
  .big {
    animation-name: gbig;
    animation-duration: 0.3s;
    animation-fill-mode: forwards;
  }
  .small {
    animation-name: gsmall;
    animation-duration: 0.3s;
    animation-fill-mode: forwards;
  }
}

@keyframes gsmall {
  0% {
    transform: scale(1);
  }
  100% {
    transform: scale(0.4);
  }
}
@keyframes gbig {
  0% {
    transform: scale(0.4);
  }
  100% {
    transform: scale(1);
  }
}
</style>
