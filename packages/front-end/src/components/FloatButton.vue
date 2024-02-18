<script setup lang="ts">
import {Message} from "@vicons/tabler";

defineEmits(["click"]);

</script>
<template>
  <button
    class="fixed bottom-8 right-8 bg-blue-500 rounded-full p-2 shadow-lg hover:shadow-xl"
    @click="$emit('click')"
  >
    <Message
      style="width:50px"
      class="text-gray-200"
    />
  </button>
</template>

<style scoped lang="scss">
button {
  z-index: 1000;
}
</style>
