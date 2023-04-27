<template>
    <div>
        <el-row :gutter="20" class="mt-2">
            <el-col :span="10" style="text-align: end;">
                <el-icon @click="drawer = true" class="text-blue-500 cursor-pointer" style="font-size: 30px;">
                    <Menu />
                </el-icon>
            </el-col>
            <el-col :span="10">
                <el-button size="default" @click="prevPage" class="mx-1">上一页</el-button>
                <el-button size="default" @click="nextPage" class="mx-1">下一页</el-button>
            </el-col>
        </el-row>

        <div id="viewer">
            <div v-if="loading" style="text-align: center;">正在拼命加载中...</div>
        </div>

        <el-drawer v-model="drawer" title="目录" :with-header="false" direction="ltr">
            <ul>
                <li v-for="(item, index) in booknavigation" :key="index">
                    <a class="cursor-pointer" @click="goto(item.href)">{{ item.label }}</a>
                </li>
            </ul>
        </el-drawer>
    </div>
</template>

<script setup>
import { ref, getCurrentInstance } from "vue";
const { proxy } = getCurrentInstance();
import epub from 'epubjs'

//#region init
const loading = ref(true)
const drawer = ref(false)
// const bookUrl = 'book/c++ primer plus.epub'
const bookUrl = 'public/book/vue.js.epub'
const booknavigation = ref(null)
const book = new epub(bookUrl)
const rendition = book.renderTo('viewer', {
    width: window.innerWidth,
    height: window.innerHeight - 100
})
book.loaded.navigation.then(nav => {
    booknavigation.value = nav.toc
})
rendition.display().then(() => {
    console.log('displayed');
    loading.value = false
})
//#endregion

const prevPage = () => {
    rendition.prev()
}

const nextPage = () => {
    rendition.next()
}

const goto = (href) => {
    rendition.display(href)
}
</script>

<style scoped></style>
