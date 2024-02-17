<script setup lang="ts">
import NoticeProvider from "@components/NoticeProvider.vue";
import {RouterView} from "vue-router";

</script>

<template>
  <NoticeProvider>
    <router-view />
  </NoticeProvider>
</template>

<style scoped lang="scss">
</style>
