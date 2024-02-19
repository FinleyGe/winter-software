<script setup lang="ts">
import { useRouter } from "vue-router";
import type { NavItem } from "Nav";

defineProps<{
  items: Array<NavItem>,
}>();

const index = defineModel<number>();
index.value = 0;
const router = useRouter();

const handleClick = (item: NavItem, i: number) => {
  if (item.onClick) item.onClick();
  index.value = i;
  if (item.path)
    router.push(item.path);
};

</script>
<template>
  <div class="flex my-1">
    <div
      v-for="item, i in items"
      :key="item.name"
      class="item"
      :class="{'chosen': i === index}"
      @click="handleClick(item, i)"
    >
      {{ item.name }}
    </div>
  </div>
</template>

<style scoped lang="scss">
.item {
  padding: 0.5rem;
  margin-inline: 0.5rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  border-radius: 0.5rem;
  &:hover {
    background-color: #8ee4ef;
  }

  &.chosen {
    background-color: #00abe9;
  }
}
</style>
