<script setup lang="ts">
import Nav from "@components/Nav.vue";
import { ref } from "vue";
import type {NavItem} from "@components/Nav";

const index = ref<number>(0);
const items = <NavItem[]>[
  {name:"个人中心", path:"/student"},
  {name:"导师信息", path:"/student/teacher"},
];

</script>
<template>
  <div class="base bg-blue-300 flex">
    <Nav
      v-model="index"
      :items="items"
    />
  </div>
</template>

<style scoped lang="scss">
.base{
  height: 50px;
}
</style>
