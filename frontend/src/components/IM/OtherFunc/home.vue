<template>
  <h2>
    Hello, {{ user }}
  </h2>
<p>
  home.vue
</p>
</template>

<script>
export default {
  name: 'Home',
  data() {
    return {
      user: localStorage.getItem('user')
    }
  }
}
</script>
