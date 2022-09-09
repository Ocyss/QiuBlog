<template>
  <div class="logo">
    <img src="/Logo.png" alt="" />
    <h2 v-show="!collapsed" class="title">L4V3</h2>
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
  height: 64px;
  line-height: 64px;
  overflow: hidden;
  white-space: nowrap;

  img {
    width: auto;
    height: 45px;
  }

  .title {
    margin-bottom: 0;
  }
}
</style>
