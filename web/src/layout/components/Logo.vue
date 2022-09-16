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
  }
  .big {
    height: 120px;
  }
  .small {
    height: 60px;
  }
}
</style>
