<template>
  <div id="chart1" style="width: 40rem; height: 30rem"></div>

</template>

<script setup lang="ts">
import * as echarts from "echarts";
import { onMounted } from "vue";

type EChartsOption = echarts.EChartsOption;
onMounted(() => {
  let chartDom = document.getElementById("chart1")!;
  let myChart = echarts.init(chartDom);
  let option: EChartsOption;

  option = {
    title: {
      text: "可莉被禁闭次数",

      left: "center",
      textStyle: {
        fontSize: 25,
      },
      subtextStyle: {
        fontSize: 20,
      },
    },
    xAxis: {
      type: "category",
      data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
    },
    yAxis: {
      type: "value",
    },
    series: [
      {
        data: [150, 230, 224, 218, 135, 147, 260],
        type: "line",
      },
    ],
  };

  myChart.setOption(option);
});
</script>

<style scoped></style>
