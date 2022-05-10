<template>
  <div>用户管理</div>
  <el-table :data="tableData">
    <el-table-column prop="name" label="姓名"></el-table-column>
    <el-table-column prop="sex" label="性别"></el-table-column>
    <el-table-column prop="age" label="年龄"></el-table-column>
    <el-table-column prop="birthday" label="出生日期">
      <template #default="{ row }">
        <span>{{ formatDate(row.birthday) }}</span>
      </template>
    </el-table-column>
    <el-table-column label="操作">
      <template #default="{ row }">
        <el-button type="primary" @click="showDialog(row)">信息</el-button>
        <el-button type="danger" @click="deleteRow(row)">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
  <el-dialog v-model="dialogVisible">
    <template #title> 用户信息 </template>
    <p>{{ cur.id }}</p>
    <p>{{ cur.name }}</p>
    <p>{{ cur.sex }}</p>
  </el-dialog>
</template>

<script setup lang="ts">
import dayjs from "dayjs";
import { ElMessage } from "element-plus";

interface User {
  id: string;
  name: string;
  sex: string;
  age: number;
  birthday: string;
}

let dialogVisible = $ref<boolean>(false);
let cur = $ref<User>();
let tableData = $ref<User[]>([
  { id: "1", name: "小明", sex: "女", age: 12, birthday: "1922-12-12" },
  { id: "2", name: "小蓝", sex: "女", age: 56, birthday: "2122-03-12" },
  { id: "3", name: "凝光", sex: "女", age: 34, birthday: "2022-03-12" },
  { id: "4", name: "小白", sex: "女", age: 34, birthday: "2013-03-12" },
  { id: "5", name: "刻晴", sex: "女", age: 34, birthday: "2011-03-12" },
  { id: "6", name: "雷军", sex: "女", age: 34, birthday: "2000-03-12" },
]);

function formatDate(date: string) {
  return dayjs(date).format("YYYY-MM-DD HH:mm:ss");
}

function showDialog(row: User) {
  cur = row;
  dialogVisible = true;
}

function deleteRow(row: User) {
  for (let index = tableData.length - 1; index >= 0; index--) {
    if (tableData[index] && tableData[index].id === row.id) {
      tableData.splice(index, 1);
    }
  }
  ElMessage({ type: "success", message: `删除成功!${row.name}` });
}
</script>

<style scoped></style>
