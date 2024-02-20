<script setup lang="ts">
import Input from "@components/Input.vue";
import { useNotice } from "@components/Notice";
import router from "@/router";
import {ref} from "vue";

const originalPassword = ref<string>("");
const newPassword = ref<string>("");
const confirmPassword = ref<string>("");
const notice = useNotice();

function handelPwd() {
  // TODO: This is a fake login function
  if (newPassword.value === "" || originalPassword.value === "") {
    notice({
      type: "error",
      message: "密码不能为空"
    });
    return;
  }if (newPassword.value !=confirmPassword.value) {
    notice({
      type: "error",
      message: "新密码输入不一致"
    });
    return;
  }
  notice({
    type: "success",
    message: "密码修改成功，请重新登录"
  });
  router.push("/login");
}

</script>
<template>
  <div class="form">
    <h1 class="text-2xl">
      修改密码
    </h1>
    <span class="text-gray-400 my-2 text-sm"> 初始密码为 "zjut" + 工号后六位,忘记密码请联系管理员</span>
    <Input
      v-model="originalPassword"
      label="原密码"
      placeholder="请输入原密码"
      type="password"
    />
    <Input
      v-model="newPassword"
      label="新密码"
      placeholder="请输入新密码"
      type="password"
    />
    <Input
      v-model="confirmPassword"
      label="重输新密码"
      placeholder="请重新输入新密码"
      type="password"
    />
    <button
      class="bg-blue-500 text-white py-2 px-4 rounded-md my-4 hover:bg-blue-300 transition-all duration-100"
      style="width: 100%;"
      @click="handelPwd"
    >
      登录
    </button>
  </div>
</template>

<style scoped lang="scss">

.form {
  width: fit-content;
  margin-inline: auto;
  border-radius: 0.25rem;
  background-color: white;
  padding: 40px;
  text-align: left;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
}

</style>
