<template>
  <div class="main">
    <el-card shadow="never">
      <section class="main-header">
        <article class="main-left">
          <el-avatar
            src="https://img-static.mihoyo.com/communityweb/upload/222b847170feb3f2babcc1bd4f0e30dd.png"
          ></el-avatar>
        </article>
        <article class="content1">早上好啊,要和可莉一起去炸鱼吗?</article>
        <article class="content2">今天琴团长不在哦</article>
        <article class="main-right"></article>
      </section>
    </el-card>
    <br />
    <el-row gutter="10">
      <el-col :span="8">
        <el-card shadow="never">
          <template #header> 炸鱼计划1 </template>
          <div v-for="o in 4" :key="o">{{ o + ".xxxxxxxxxxxx" }}</div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card shadow="never">
          <template #header> 炸鱼计划2 </template>
          <div v-for="o in 4" :key="o">{{ o + ".xxxxxxxxxxxx" }}</div>
        </el-card>
      </el-col>
      <el-col :span="8">
        <el-card shadow="never">
          <template #header> 炸鱼计划3 </template>
          <div v-for="o in 4" :key="o">{{ o + ".xxxxxxxxxxxx" }}</div>
        </el-card>
      </el-col>
    </el-row>
    <el-card shadow="never" style="margin-top: 1rem">
      <div id="effects" style="width: 100%; height: 25rem"></div>
    </el-card>
  </div>
</template>

<script setup lang="ts">
import * as echarts from "echarts";
onMounted(() => {
  type EChartsOption = echarts.EChartsOption;

  let chartDom = document.getElementById("effects")!;
  let myChart = echarts.init(chartDom);
  let option: EChartsOption;
  myChart.on("click", () => {
    myChart.resize();
  });
  option = {
    title: {
      text: "炸鱼成果",

      left: "left",
      textStyle: {
        fontSize: 25,
      },
      subtextStyle: {
        fontSize: 20,
      },
    },
    color: ["#003366", "#006699", "#4cabce", "#e5323e"],
    dataset: {
      source: [
        ["type", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
        ["甜甜花鳉", 320, 132, 301, 334, 390, 156, 345],
        ["金赤假龙", 220, 182, 291, 234, 290, 176, 258],
        ["长生仙", 150, 232, 201, 154, 190, 356, 154],
        ["雷鸣仙", 98, 77, 101, 99, 40, 65, 44],
      ],
    },
    legend: {},
    xAxis: {
      type: "category",
      data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
      axisTick: {
        show: false,
      },
    },
    yAxis: {},
    series: [
      {
        type: "bar",
        seriesLayoutBy: "row",
      },
      {
        type: "bar",
        seriesLayoutBy: "row",
      },
      {
        type: "bar",
        seriesLayoutBy: "row",
      },
      {
        type: "bar",
        seriesLayoutBy: "row",
      },
    ],
  };

  option && myChart.setOption(option);
  myChart.resize();
});
</script>

<style lang="scss" scoped>
.main {
  .main-header {
    display: grid;
    grid-template-columns: 1fr 10fr 1fr;
    grid-template-areas: "left content1 right" "left content2 right";

    .main-left {
      grid-area: left;
      display: flex;
      justify-content: center;
    }

    .content1 {
      grid-area: content1;
    }

    .content2 {
      grid-area: content2;
      color: #00000073;
    }

    .main-right {
      grid-area: right;
    }
  }
}
</style>
