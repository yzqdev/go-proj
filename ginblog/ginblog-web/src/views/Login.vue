<template>
  <div class="admin-login">
    <div class="admin-container">
      <h3 class="text-center">{{ title }}</h3>
      <el-form
        ref="loginFormRef"
        :model="loginForm"
        label-width="80px"
        label-position="top"
        :rules="loginRules"
      >
        <el-form-item prop="username" label="用户名">
          <el-input
            size="large"
            v-model="loginForm.username"
            placeholder="请输入用户名"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password" label="密码">
          <el-input
            size="large"
            @keydown.enter="login"
            placeholder="请输入密码"
            type="password"
            v-model="loginForm.password"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-tooltip class="item" effect="dark" placement="right">
            <template #content> 忘记密码请找管理员 </template>

            <el-link icon="el-icon-question" :underline="false"
              >忘记密码
            </el-link>
          </el-tooltip>
        </el-form-item>
      </el-form>

      <el-button size="large" type="primary" style="width: 100%" @click="login"
        >登录
      </el-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { loginApi, regApi } from "@/utils/apis";
import {
  defineComponent,
  onBeforeMount,
  reactive,
  ref,
  toRefs,
  watch,
} from "vue";
import {ElMessage, ElNotification} from "element-plus";
import { useRouter } from "vue-router";
import { useUserStore } from "@/store/user";

const router = useRouter();
let userStore = useUserStore();
import  'element-plus/theme-chalk/el-notification.css';
let activeName = $ref<string>("first");
let title = $ref<string>("用户登录");
let regForm = $ref<object>({ username: "", password: "", password2: "" });
let loginRules = $ref({
  username: [{ required: true, message: "请输入用户名" }],
  password: [{ required: true, message: "请输入密码" }],
});

let regRule = $ref({
  username: [{ required: true, message: "请输入用户名" }],
  password: [{ required: true, message: "请输入密码" }],
  password2: [{ required: true, message: "请输入确认密码" }],
});
let loginForm = $ref<any>({ username: "admin", password: "123456" });
let loginFormRef = ref(null);
let regFormRef = ref(null);

function login() {
  loginFormRef.value.validate((valid: boolean) => {
    if (valid) {
      localStorage.token = "this is token";
      userStore.setUserInfo(loginForm);
      ElNotification({
        title: '欢迎回来!',
        message: '这是一条欢迎语',
        type: 'success',
      })
      router.push({ name: "adminHome" });
      loginApi(loginForm).then((res) => {
        if (res.success) {
          userStore.setUserToken(res);
          localStorage.token = res.data;
          ElMessage({ message: "success", type: "success" });
          router.push({ name: "adminWelcome" });
        } else {
          ElMessage({ message: "登录失败", type: "error" });
        }
      });
    }
  });
}

onBeforeMount(() => {
  if (localStorage.token) {
    router.push({ name: "adminHome" });
  }
});
</script>

<style lang="scss" scoped>
.admin-login {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  background: #132ea0 url("../assets/login-bg.png") no-repeat 50%;

  .admin-container {
    background: #fff;

    border-radius: 5px;

    margin: -160px 0 0 -160px;
    width: 320px;
    padding: 22px 28px 28px 28px;
    border: 1px solid #eaeaea;
    //box-shadow: 0 0 25px #cac6c6;
    .text-center {
      text-align: center;
    }
  }
}
</style>
