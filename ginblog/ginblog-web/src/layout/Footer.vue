<template>

   <div v-if=" showFooter" class="footer">
   <span>Copyright © 2021-2022 yzqdev</span></div>
</template>

<script setup>
import { useThemeStore } from "@/store/themeConfig";

let themeStore = useThemeStore();
let showFooter = $computed(() => {
  return themeStore.showFooter;
});
</script>

<style lang="scss" scoped>
.footer{
  height: 3rem;
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
