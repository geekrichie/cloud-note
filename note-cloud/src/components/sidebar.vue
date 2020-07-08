<template>
    <div class="sidebar">
         <el-menu
            class="el-menu-vertical-demo"        
            background-color="#545c64"
            text-color="#fff">
            <el-menu-item index="1">
                <div class="el-icon-plus"></div>
                <span slot="title">新增笔记</span>
            </el-menu-item> 
        </el-menu>
    </div>
</template>
<script>
export default {
    
}
</script>
<style scoped>
.sidebar {
    width:300px;
    height:100%;
}
.el-icon-plus {
    background-image:url("../assets/plus.png")
}
ul {
    height:100%;
}
</style>