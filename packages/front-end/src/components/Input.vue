<script setup lang="ts">

const value = defineModel<string>();
defineProps<{
  placeholder: string;
  label: string;
  type: string;
}>();

</script>
<template>
  <div class="flex flex-col my-2">
    <label class="text-sm text-gray-500"> {{ label }} </label>
    <input
      v-model="value"
      class="border-b-2 border-gray-300 focus:border-blue-500 focus:outline-none transition-all duration-300 my-1 px-2 py-1 w-full"
      :type="type"
      :placeholder="placeholder"
    >
  </div>
</template>

<style scoped lang="scss">

</style>
