<template>
  <n-tabs type="segment">
    <n-tab-pane name="chap1" tab="基本信息">
      <n-table striped>
        <tbody>
          <tr>
            <td>当前用户</td>
            <td>{{ "Admin" }}</td>
          </tr>
        </tbody>
      </n-table>
    </n-tab-pane>
    <n-tab-pane name="chap2" tab="全局设置">
      <n-table striped>
        <tbody>
          <tr>
            <td>网站时区</td>
            <td><n-select /></td>
          </tr>
          <tr>
            <td>默认语言</td>
            <td><n-select /></td>
          </tr>
          <tr>
            <td>默认主题</td>
            <td><n-select /></td>
          </tr>
          <tr>
            <td>开发模式</td>
            <td><n-switch :round="false" /></td>
          </tr>
          <tr>
            <td>安全增强</td>
            <td><n-switch :round="false" /></td>
          </tr>
          <tr>
            <td>GZip压缩</td>
            <td><n-switch :round="false" /></td>
          </tr>
          <tr>
            <td>代码高亮</td>
            <td><n-switch :round="false" /></td>
          </tr>
          <tr>
            <td>隐示公告</td>
            <td><n-switch :round="false" /></td>
          </tr>
        </tbody>
      </n-table>
    </n-tab-pane>
    <n-tab-pane name="chap3" tab="友链管理"></n-tab-pane>
    <n-tab-pane name="chap4" tab="分类管理">
      <menuList />
    </n-tab-pane>
  </n-tabs>
</template>

<script setup>
import menuList from "@/components/admin/other/menulist.vue";
</script>

<style scoped></style>
