<!-- src/views/IMView.vue -->

<template>
    <div class="container">
        <IMHeader></IMHeader>
        <div class="main">
            <IMLeftMenu @selectRoute="selectRoute"></IMLeftMenu>
            <IMRightContent></IMRightContent>
        </div>
    </div>
</template>

<script>
import IMHeader from "@/components/IM/Frame/IMHeader.vue";
import IMLeftMenu from '@/components/IM/Frame/IMLeftMenu.vue'
import IMRightContent from '@/components/IM/Frame/IMRightContent.vue'

export default {
    name: 'IMView',
    components: {
        IMHeader,
        IMLeftMenu,
        IMRightContent
    },
    methods: {
        selectRoute(route) {
            this.$router.push('/im/' + route)
        }
    }
}
</script>

<style scoped>
.container {
    display: flex;
    flex-direction: column;
    height: 100%;
    overflow: hidden;
    /* 防止内容溢出产生滚动条 */
}

.main {
    display: flex;
    flex: 1;
    height: calc(100% - 60px);
    /* 占据整个可视区域的高度 */
}
</style>