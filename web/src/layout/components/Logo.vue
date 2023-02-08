<template>
  <div class="logo">
    <img
      src="/favicons/android-icon-192x192.png"
      alt=""
      :class="collapsed ? 'small' : 'big'"
    />
  </div>
</template>

<script>
export default {
  name: "logo",
  props: ["collapsed"],
};
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
