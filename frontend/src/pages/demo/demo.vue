<template>
    <div>
        <el-button type="primary" size="default" @click="demo">click</el-button>
    </div>
</template>

<script setup>
import { ref, reactive, getCurrentInstance } from "vue";
const { proxy } = getCurrentInstance();

const demo = () => {
    console.log("demo");
    proxy.$elNotification("demo")
};
</script>

<style scoped></style>