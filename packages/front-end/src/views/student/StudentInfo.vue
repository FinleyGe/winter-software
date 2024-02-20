<script setup lang="ts">
import Pic from "@components/Pic.vue";
import { ref } from "vue";
import type { Form } from "@api/type";

const student_info = ref<Form | null>();

student_info.value = {
  id: 2018000000,
  name: "张三",
};

</script>
<template>
  <div class="container my-2 mx-auto w-5/6 bg-gray-50 p-2 rounded-xl shadow-xl">
    <h1 class="text-2xl font-bold text-center">
      学生信息
    </h1>
    <div class="flex justify-between">
      <div class="text-blue-500 text-xl font-bold">
        学号: {{ student_info?.id }}
      </div>
      <div class="text-blue-500 text-xl font-bold">
        姓名: {{ student_info?.name }}
      </div>
      <Pic />
    </div>
  </div>
</template>

<style scoped lang="scss">

</style>
