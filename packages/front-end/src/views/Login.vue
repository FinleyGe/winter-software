<script setup lang="ts">
import Input from "@components/Input.vue";
import Checkbox from "@components/Checkbox.vue";
import { useNotice } from "@components/Notice";
import { useUserStore } from "@/store";
import router from "@/router";
import {onMounted, ref} from "vue";

const user = ref<string>("");
const password = ref<string>("");
const remember = ref<boolean>(false);
const notice = useNotice();
const userStore = useUserStore();

function handelLogin(){
  // TODO: This is a fake login function
  if (user.value === "" || password.value === "") {
    notice({
      type: "error",
      message: "帐号或密码不能为空"
    });
    return;
  }
  if (remember.value) {
    localStorage.setItem("user", user.value);
    localStorage.setItem("password", password.value);
  }
  notice({
    type: "success",
    message: "登录成功"
  });
  userStore.isLogin = true;
  userStore.role = "student";
  router.push("/student");
}

onMounted(() => {
  const u = localStorage.getItem("user");
  const p = localStorage.getItem("password");
  if (u && p) {
    user.value = u;
    password.value = p;
    remember.value = true;
  }
});

</script>
<template>
  <div class="min-h-svh w-full bg-gray-50 base">
    <div class="left flex items-center">
      <img
        src="/login.jpg"
        style="width: 100%;height: 100vh;"
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
