<template>
  <el-aside width="15rem" :native-scrollbar="false" bordered class="sidebar">
    <router-link to="/" custom>
      <span class="logo">
        <img v-if="false" class="logo-icon" src="../assets/logo.png" />
        <span>Slim</span>
      </span>
    </router-link>
    <el-menu
      :default-active="defaultActive"
      class="side-menu"
      background-color="#001529"
      text-color="#fff"
      @open="handleOpen"
      :collapse-transition="false"
      @close="handleClose"
      ><el-menu-item index="adminHome" @click="gotoRoute('adminHome')">
        <el-icon>
          <home-filled />
        </el-icon>
        <template #title>首页</template>
      </el-menu-item>
      <el-sub-menu index="settings">
        <template #title
          ><el-icon><setting /></el-icon><span>管理</span></template
        >
        <el-menu-item index="notice" @click="gotoRoute('notice')">
          <el-icon>
            <icon-menu />
          </el-icon>
          <template #title>通知公告</template>
        </el-menu-item>
        <el-menu-item index="profile" @click="gotoRoute('profile')">
          <el-icon>
            <credit-card />
          </el-icon>
          <template #title>天命既定</template>
        </el-menu-item>
        <el-menu-item index="userManage" @click="gotoRoute('userManage')">
          <el-icon>
            <user />
          </el-icon>
          <template #title>用户管理</template>
        </el-menu-item>
      </el-sub-menu>
      <el-sub-menu index="others">
        <template #title
          ><el-icon> <icon-menu /></el-icon><span>功能组件</span></template
        >
        <el-menu-item index="charts" @click="gotoRoute('charts')">
          <el-icon>
            <pie-chart />
          </el-icon>
          <template #title>图表</template>
        </el-menu-item>
        <el-menu-item index="icons" @click="gotoRoute('icons')">
          <el-icon>
            <alarm-clock />
          </el-icon>
          <template #title>图标</template>
        </el-menu-item> <el-menu-item index="directives" @click="gotoRoute('directives')">
          <el-icon>
            <arrow-down />
          </el-icon>
          <template #title>指令</template>
        </el-menu-item>
      </el-sub-menu>
    </el-menu>
  </el-aside>
</template>

<script lang="ts" setup>
import { computed, watch } from "vue";
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
  CreditCard,
  HomeFilled,
  User,
  PieChart,
  AlarmClock, ArrowDown,
} from "@element-plus/icons-vue";

import { useRoute, RouterLink, useRouter } from "vue-router";
import { Menu } from "@/components/type";
import {  useThemeStore } from "@/store/themeConfig";

let defaultActive = $ref<string>("adminHome");
const router = useRouter();
const route = useRoute();
let themeStore = useThemeStore();
function gotoRoute(params: string) {
  router.push({ name: params });
}
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};
const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};
let themeConfig = $computed(() => {
  return themeStore.$state;
});
watch(
  themeConfig,
  (val, oldVal) => {
    if (val.dark) {
      console.log(`%cdark`, `color:red;font-size:16px;background:transparent`);
    } else {
    }
  },
  { immediate: true }
);
watch(
  route,
  (val, oldVal) => {
    defaultActive = val.name;
    console.log("move to ", val.name);
  },
  { immediate: true }
);
</script>

<style lang="scss" scoped>
.sidebar {
  width: 15rem;
  background-color: #001529;
  height: 100%;
  .logo {
    position: sticky;
    top: 0;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 12px 20px;
    text-align: center;
    font-size: 1.8em;
    font-weight: 600;
    line-height: 1;
    text-decoration: none;
    background: #001529;
    color: white;
    .logo-icon {
      width: 2rem;
      height: 2rem;
    }
  }
  :deep(.el-menu) {
    .el-menu {
      background-color: #000c17;
    }
    .el-menu-item.is-active {
      background: var(--el-color-primary);
      color: white;
    }
  }
}

.side-menu {
  height: 90vh;
}
</style>
