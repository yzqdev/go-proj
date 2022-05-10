<template>
  <el-popover trigger="click" placement="bottom" width="400">
    <template #reference>
      <el-button>
        <span v-if="selectItem">
          <i class="iconfont" :class="[selectItem]" style="font-size: 20px"></i
        ></span>
        {{ selectItem }}
      </el-button>
    </template>
    <el-scrollbar class="grid-wrapper">
      <section style="height: 300px" class="grid-icon">
        <div v-for="item of icons" :key="item">
          <div class="icons" @click="onIconClick(item)">
            <i class="iconfont" :class="[item]" style="font-size: 20px"></i>
          </div>
        </div>
      </section>
    </el-scrollbar>
    <div class="page">
      <el-pagination
        :page="currentPage"
        :page-size="pageSize"
        :total="itemCount"
        @current-change="onUpdatePage"
      />
    </div>
  </el-popover>
</template>

<script setup lang="ts">
import { computed, onMounted, toRef } from "vue";
import axios from "axios";
import {iconCss} from "@/constants";

let props = defineProps({
  onUpdateIcon: Function,
});

let emit = defineEmits(["selected"]);
let iconPrefix = "icon-";
let iconList = $ref<string[]>([]);
let icons = $ref<string[]>([]);
onMounted(() => {
  axios
    .get(iconCss)
    .then(({ data }) => {
      const regExp = new RegExp(`\\n\\.(${iconPrefix}.*?):before`, "g");
      let result;

      while ((result = regExp.exec(data))) {
        iconList.push(result[1]);
      }
      icons = iconList.slice(0, 40);
    });
});
const pageSize = $ref<number>(40);

let currentPage = $ref(1);
let itemCount = computed(() => iconList.length);
let selectItem = $ref<string>("点击我选择图标");

function onUpdatePage(page: number) {
  currentPage = page;
  icons.length = 0;
  const start = (currentPage - 1) * pageSize;
  icons.push(...iconList.slice(start, start + pageSize));
}

const { onUpdateIcon } = props;

function onIconClick(item: any) {
  selectItem = item;
  if (onUpdateIcon) {
    onUpdateIcon(item);
  } else {
    emit("selected", item);
  }
}
</script>
<style lang="scss" scoped>
.grid-wrapper {
  .icon-wrapper {
    cursor: pointer;
    border: 1px solid #f5f5f5;
  }
  .grid-icon {
    display: grid;
    grid-template-columns: repeat(6, 1fr);
  }

  .icons {
    display: flex;
    align-items: center;
    justify-content: center;
    border: 1px solid #f5f5f5;
    cursor: pointer;
  }
}
</style>
