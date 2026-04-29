<script setup lang="ts">
// App.vue 根组件
// 只负责渲染路由视图（RouterView），实际页面由子路由组件接管
import { RouterView } from 'vue-router'
</script>

<template>
  <RouterView />
</template>

<style>
#app {
  width: 100%;
  height: 100vh;
}
</style>