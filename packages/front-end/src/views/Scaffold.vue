<script setup lang="ts">
import FloatButton from "@components/FloatButton.vue";
import Chat from "@views/Chat.vue";
import { ref } from "vue";

const isChat = ref<boolean>(false);

function handleFloatButtonClick() {
  isChat.value = !isChat.value;
}

</script>
<template>
  <div>
    <header
      style="height: 50px"
    >
      <router-view name="header" />
    </header>
    <main style="min-height: calc(100vh - 50px);">
      <router-view />
    </main>
    <!-- <footer -->
    <!--   style="height: 30px" -->
    <!-- > -->
    <!--   <router-view name="footer" /> -->
    <!-- </footer> -->
    <FloatButton @click="handleFloatButtonClick" />
    <Chat
      v-show="isChat"
      @close="isChat = false"
    />
  </div>
</template>

<style scoped lang="scss">
main {
  background: linear-gradient(to bottom, #b2dffb, #E6E6FA);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
</style>
