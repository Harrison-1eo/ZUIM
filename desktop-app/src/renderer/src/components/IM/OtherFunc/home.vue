<template>
	<div>
		<h1>首页</h1>
		<p>{{ message }}</p>
		<button @click="updateMessage">更新消息</button>
	</div>
</template>

<script>
export default {
	name: 'Home',
	data() {
		return {
			message: '欢迎来到首页'
		};
	},
	methods: {
		updateMessage() {
			this.message = '消息已更新';
		}
	}
};
</script>

<style scoped>
/* 这里可以添加组件的样式 */
</style>
