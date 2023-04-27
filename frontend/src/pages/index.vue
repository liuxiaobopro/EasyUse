<template>
    <div>
        <el-row :gutter="12" class="my-2" v-for="(item, index) in list" :key="index">
            <el-col :span="8" v-for="(item1, index1) in item" :key="index1" @click="$router.push(item1.path)">
                <el-card shadow="hover"> {{ item1.content }} </el-card>
            </el-col>
        </el-row>
    </div>
</template>

<script setup>
import { ref, reactive, getCurrentInstance } from "vue";
const { proxy } = getCurrentInstance();

const list = ref([
    [
        {
            content: "示例",
            path: "/demo"
        },
        {
            content: "图书",
            path: "/book"
        },
    ],
]);
</script>

<style scoped></style>