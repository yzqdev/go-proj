<template>
  <el-container class="home">
    <sidebar />

    <el-container class="main-wrap">
      <Header />
      <Content />
      <Footer />
    </el-container>
  </el-container>
</template>

<script lang="ts" setup>
import Sidebar from "./Sidebar.vue";
import Header from "./Header.vue";
import Content from "./Content.vue";
import Footer from "./Footer.vue";
import { onMounted } from "vue";
import { useRouter } from "vue-router";
import { useUserStore } from "@/store/user";
let router = useRouter();
let userStore = useUserStore();
onMounted(() => {
  let token = localStorage.getItem("token");
  if (token && token !== "undefined") {
    userStore.setUserToken(token);
    // store.commit("setUserToken", localStorage.getItem("token"));
  } else {
    router.push({ name: "adminHome" });
  }
});
</script>
<style lang="scss" scoped>
.home {
  display: flex;
  height: 100%;

  .main-wrap {
    height: 100%;

    display: flex;
    flex: 1;
    flex-direction: column;
  }
}
</style>
