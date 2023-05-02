<template>
  <div>
    <el-button type="primary" size="default" @click="demo">click</el-button>
    <el-button type="primary" size="default" @click="doGreeting('1111')">click1</el-button>
  </div>
</template>

<script setup>
import { ref, reactive, getCurrentInstance } from "vue";
const { proxy } = getCurrentInstance();
import { Greet, NewComputerEssentials } from "/wailsjs/go/main/App";


const demo = () => {
  console.log("demo");
  proxy.$elNotification("demo")
};


function doGreeting(name) {
  NewComputerEssentials().then(res => {
    console.log(res);
  })
}
</script>

<style scoped></style>