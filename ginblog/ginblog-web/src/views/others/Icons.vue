<template>
  <icon-selector @selected="onSelectedIcon"/>
  <h4>下面是可用的图标</h4>
  <div class="icons-container">

   <span class="icons-wrap" v-for="item in iconList" @click="selectIcon(item)">
      <i class="iconfont" :class="[item]" style="font-size: 2rem"></i>
   </span>
  </div>
</template>

<script setup lang="ts">
import {ElMessage} from "element-plus";
import "element-plus/theme-chalk/el-message.css";
import {onMounted} from "vue";
import axios from "axios";
import {iconCss} from "@/constants";

function onSelectedIcon(item: any) {

  ElMessage({type: "success", message: `选择了${item}`});
}

let iconPrefix = 'icon-'
let iconList = $ref<string[]>([])
function selectIcon(item:string){
  let iconCode=`<i class="iconfont ${item}" style="font-size: 1.5rem;"></i>`
  navigator.clipboard.writeText(iconCode).then((res) => {
    ElMessage({type:'success',message:`选择了${iconCode}`})
  })

}
onMounted(() => {
  axios
      .get(iconCss)
      .then(({data}) => {
        const regExp = new RegExp(`\\n\\.(${iconPrefix}.*?):before`, "g");
        let result;

        while ((result = regExp.exec(data))) {
          iconList.push(result[1]);
        }

      });
});
</script>

<style lang="scss" scoped>
.icons-container {
  margin-top: 1rem;
  display:flex;
  flex-wrap: wrap;
  gap: 1rem;

  .icons-wrap {
    display: grid;
    align-items: center;
    justify-content: center;
padding: 0.5rem;
    cursor: pointer;
    border: 1px solid var(--el-border-color);
  }
}

</style>
