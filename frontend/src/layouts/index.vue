<template>
    <div>
        <div>
            <el-menu :default-active="$route.path" mode="horizontal">
                <el-menu-item index="/" @click="$router.push('/')">首页</el-menu-item>
                <el-menu-item index="/demo" @click="$router.push('/demo')">示例</el-menu-item>
                <el-menu-item index="/book" @click="$router.push('/book')">图书</el-menu-item>
            </el-menu>
        </div>
        <router-view></router-view>
    </div>
</template>

<style scope>
.el-menu {
    height: 39px;
}
</style>

