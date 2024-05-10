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
import IMHeader from "@/components/IMHeader.vue";
import IMLeftMenu from '@/components/IMLeftMenu.vue'
import IMRightContent from '@/components/IMRightContent.vue'

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
}

.main {
  display: flex;
  flex: 1;
  height: 100%;
}
</style>