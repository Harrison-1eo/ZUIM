<template>
    <div>
        <h1>设置</h1>
        <p>通知设置：<input type="checkbox" v-model="notificationEnabled"></p>
    </div>
</template>

<script>
export default {
    name: 'Setting',
    data() {
        return {
            notificationEnabled: true
        };
    }
};
</script>

<style scoped>
/* 这里可以添加组件的样式 */
</style>
