<script setup lang="ts">
const value = defineModel<boolean>();
defineProps<{
  label: string;
}>();
</script>
<template>
  <div class="flex items-center">
    <input
      v-model="value"
      type="checkbox"
      class="h-5 w-5"
    >
    <label class="ml-2">{{ label }}</label>
  </div>
</template>

<style scoped lang="scss">

</style>
