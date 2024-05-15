<template>
  <div id="app" class="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
.app {
  height: 100%;
  overflow: hidden; /* 防止内容溢出产生滚动条，实际上会溢出，待解决 */
}
</style>
