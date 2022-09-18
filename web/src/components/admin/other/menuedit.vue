<template>
  <n-card size="huge" class="dialog">
    <div>
      <n-input round placeholder="中文名" v-model:value="props.value.name" />
      <n-input round placeholder="英文名" v-model:value="props.value.ename" />
    </div>
    <div>
      <n-input round placeholder="Logo" v-model:value="props.value.logo" />
      <n-input round placeholder="跳转地址" v-model:value="props.value.link" />
    </div>
    <div>
      <n-icon size="40" v-html="props.value.logo"></n-icon>
    </div>
  </n-card>
</template>

<script setup>
const props = defineProps(["value"]);
</script>

<style scoped>
.dialog div {
  display: flex;
  align-items: center;
  margin-top: 5px;
}
</style>
