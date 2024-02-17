<script setup lang="ts">
import Input from "@components/Input.vue";
import Checkbox from "@components/Checkbox.vue";
import { useNotice } from "@components/Notice";
import {ref} from "vue";

const user = ref<string>("");
const password = ref<string>("");
const remember = ref<boolean>(false);
const notice = useNotice();
function handelLogin(){
  notice({
    type: "success",
    message: "登录成功"
  });
  notice({
    type: "info",
    message: "登录成功"
  });
  notice({
    type: "error",
    message: "登录成功"
  });
}
</script>
<template>
  <div class="min-h-svh w-full bg-gray-50 base">
    <div class="left flex items-center">
      <img
        src="/login-pic.png"
        style="width: 400px;"
        alt="logo"
        class="mx-auto"
      >
    </div>
    <div class="right flex items-center">
      <div class="container w-fit mx-auto flex flex-col px-8">
        <h1 class="text-2xl">
          登录
        </h1>
        <span class="text-gray-400 my-2 text-sm"> 默认帐号为学号，初始密码为 "zjut" + 学号后六位</span>
        <Input
          v-model="user"
          label="帐号"
          placeholder="请输入学号"
          type="text"
        />
        <Input
          v-model="password"
          label="密码"
          placeholder="请输入密码"
          type="password"
        />
        <Checkbox
          v-model="remember"
          label="记住密码"
        />
        <button
          class="bg-blue-500 text-white py-2 px-4 rounded-md my-4 hover:bg-blue-300 transition-all duration-100"
          @click="handelLogin"
        >
          登录
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.base{
  display: grid;
  @media (width < 768px) {
    grid-template-columns: 1fr;
    .left {
      display: none;
    }
  }
  @media (width >= 768px) {
    grid-template-columns: 1fr 1fr;
  }

  .left {
    background-color: blue;
  }
}
</style>
