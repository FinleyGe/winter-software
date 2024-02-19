<script setup lang="ts">
import Nav from "@components/Nav.vue";
import { ref } from "vue";
import type {NavItem} from "@components/Nav";

const index = ref<number>(0);
const items = <NavItem[]>[
  {name:"申请处理", path:"/teacher"},
  {name:"学生名单", path:"/teacher/slist"},
  {name:"密码修改", path:"/teacher/cipher"},
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
