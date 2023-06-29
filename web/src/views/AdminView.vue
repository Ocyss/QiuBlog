<template>
  <adminVue>
    <RouterView style="padding: 20px" />
  </adminVue>
</template>

<script setup lang="ts">
import api from "@/api";
import adminVue from "@/layout/admin.vue";
import { onUpdated, inject } from "vue";
import { VueCookies } from "vue-cookies";
import { useRouter } from "vue-router";
const $cookie = inject<VueCookies>("$cookies");
const router = useRouter()

onUpdated(async () => {
  const res = await api.user.chech()
  if (res.data.code == 500) {
    $cookie.remove("token")
    router.push({ name: "login" })
  }
})
</script>
