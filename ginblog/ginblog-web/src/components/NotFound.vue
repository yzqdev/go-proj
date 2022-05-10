<style lang="scss" scoped>
.exception {
  width: 100%;
}
</style>
<template>
  <div class="exception">
    <el-result :status="type" :title="type" :subTitle="config[type].desc">
      <template #extra>
        <el-button type="primary" @click="handleToHome">返回首页</el-button>
      </template>
    </el-result>
  </div>
</template>

<script>
import types from "./type";
export default {
  name: "Exception",
  props: {
    type: {
      type: String,
      default: "404",
    },
  },
  data() {
    return {
      config: types,
    };
  },
  methods: {
    handleToHome() {
      this.$router.push({ name: "adminHome" });
    },
  },
};
</script>
