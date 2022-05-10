<template>
  <header class="header">
    <el-link style="margin-left: 0.5rem" v-html="notice"></el-link>
    <span class="header-right">
      <el-icon @click="toggle" :size="20">
        <full-screen />
      </el-icon>
      <el-icon @click="showSettings" :size="20">
        <setting />
      </el-icon>
      <el-dropdown>
        <span class="el-dropdown-link">
          <el-avatar
            :size="30"
            src="https://img-static.mihoyo.com/communityweb/upload/222b847170feb3f2babcc1bd4f0e30dd.png"
          />
          <el-button type="text" class="ml-2">{{ userInfo.name }}</el-button>
          <el-icon> <arrow-down /></el-icon>
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item>
              <el-link @click="gotoRoute('profile')" :underline="false"
                >个人中心
              </el-link>
            </el-dropdown-item>
            <el-dropdown-item>
              <el-link
                href="https://github.com/yzqdev/slim-admin"
                target="_blank"
                :underline="false"
                >项目地址
              </el-link>
            </el-dropdown-item>
            <el-dropdown-item divided @click="logout"
              >退出登录</el-dropdown-item
            >
          </el-dropdown-menu>
        </template>
      </el-dropdown></span
    >
    <el-drawer
      size="20%"
      v-model="settingDraw"
      title="主题配置"
      direction="rtl"
    >
      <el-row>
        <el-col :span="24">
          <el-switch
            v-model="theme.contentPadding"
            size="large"
            active-text="显示边距"
          />
        </el-col><el-col :span="24">
          <el-switch
            v-model="theme.showFooter"
            size="large"
            active-text="显示footer"
          />
        </el-col>
      </el-row>
    </el-drawer>
  </header>
</template>

<script setup lang="ts">
import { ArrowDown, FullScreen, Setting } from "@element-plus/icons-vue";
import { useRouter } from "vue-router";
import { watch } from "vue";
import {   useThemeStore } from "@/store/themeConfig";
import {defaultTheme} from "@/constants/defaultTheme";

let router = useRouter();
let { setThemeConfig } = useThemeStore();
let userInfo = $ref({ name: "可莉" });
let notice = $ref(`你好,${userInfo.name},今天要炸鱼吗?`);
let settingDraw = $ref<boolean>(false);
const theme = $(
  useStorage("themeConfig", {
   ...defaultTheme
  })
);

const { isFullscreen, enter, exit, toggle } = useFullscreen();

function showSettings() {
  settingDraw = true;
}

function gotoRoute(name: string) {
  router.push({
    name: name,
  });
}

function logout() {
  localStorage.clear();
  router.push({ name: "login" });
}
watch(
  theme,
  (val , oldVal) => {
    setThemeConfig(val);
  },
  { immediate: true }
);
</script>

<style lang="scss" scoped>
.header {
  height: 3rem;
  padding: 0.5rem 1rem;
  display: flex;
  align-items: center;
  justify-content: space-between;

  .header-right {
    display: flex;
    align-items: center;

    .el-icon {
      margin-left: 1rem;

      &:hover {
        cursor: pointer;
        color: var(--el-color-primary);
      }
    }

    .el-dropdown-link {
      display: flex;
      align-items: center;
      padding: 0 2rem;
      cursor: pointer;
    }
  }
}
</style>
