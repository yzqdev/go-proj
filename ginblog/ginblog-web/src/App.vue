<script setup lang="ts">
import zhCn from "element-plus/lib/locale/lang/zh-cn";

let locale = $ref(zhCn);
</script>

<template>
  <el-config-provider :locale="locale">
    <router-view />
  </el-config-provider>
</template>

<style lang="scss">
#app {
  display: flex;
  flex-direction: column;
  height: 100%;
}

html,
body {
  height: 100%;
  margin: 0;
  padding: 0;
}
</style>
