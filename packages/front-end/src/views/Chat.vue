<script setup lang="ts">
import { X,Send } from "@vicons/tabler";

const emit = defineEmits<{
  (e: "close"): void
}>();

function handleClose() {
  emit("close");
}

</script>
<template>
  <div class="mask fixed top-0 left-0 h-svh w-svw z-10 flex items-center align-middle justify-center bg-transparent">
    <div class="container fixed z-50 w-2/3 h-4/6 m-auto bg-white rounded-xl shadow-xl flex flex-col">
      <header class="h-10 bg-blue-500 flex flex-row justify-between items-center">
        <h2 class="text-white px-2">
          信息
        </h2>
        <button
          class="bg-white w-8 h-8 rounded-full flex items-center justify-center hover:bg-gray-200 transition-all duration-300 ease-in-out"
          @click="handleClose"
        >
          <X class="w-4" />
        </button>
      </header>
      <main class="h-full">
        TODO: main
        <!-- TODO: Chat -->
      </main>
      <footer class="h-20 flex flex-row items-center justify-center">
        <textarea
          class="h-10 border-2 border-gray-300 rounded-md w-10/12"
        />
        <button class="bg-blue-500 p-2 rounded-full mx-2">
          <Send class="w-6 h-6 text-white" />
        </button>
      </footer>
    </div>
  </div>
</template>

<style scoped lang="scss">
</style>
