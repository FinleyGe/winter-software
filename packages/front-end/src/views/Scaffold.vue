<script setup lang="ts">

</script>
<template>
  <div>
    <header
      style="height: 50px"
    >
      header
    </header>
    <main style="min-height: calc(100vh - 50px - 30px);">
      <router-view />
    </main>
    <footer
      style="height: 30px"
    >
      footer
    </footer>
  </div>
</template>

<style scoped lang="scss">

</style>
