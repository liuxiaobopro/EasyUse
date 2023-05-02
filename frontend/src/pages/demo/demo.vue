<template>
    <div>
        <el-button type="primary" size="default" @click="demo">click</el-button>
        <el-button type="primary" size="default" @click="doGreeting('1111')">click1</el-button>
    </div>
</template>

<script setup>
import { ref, reactive, getCurrentInstance } from "vue";
const { proxy } = getCurrentInstance();
import { Greet } from "/wailsjs/go/main/App";


const demo = () => {
    console.log("demo");
    proxy.$elNotification("demo")
};


function doGreeting(name) {
  Greet(name).then((result) => {
    // Do something with result
    console.log(result);
  });
}
</script>

<style scoped></style>