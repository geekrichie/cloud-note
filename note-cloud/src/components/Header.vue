<template>
      <div class="main-header">
          <div class="main_title">个人云笔记</div>
          <div class="info">
            <div class="center" v-if="!isLogin">
               <router-link :to="'login'" class="navigation">登录</router-link>
               <router-link :to="'register'" class="navigation">注册</router-link>
            </div>
             <div v-if="isLogin" class="center">
                 <el-dropdown>
                    <span class="el-dropdown-link drop-title">
                        {{ username  }}<i class="el-icon-arrow-down el-icon--right"></i>
                    </span>
                    <el-dropdown-menu slot="dropdown">
                        <el-dropdown-item><el-button  type = "text" @click="logout">退出登录</el-button></el-dropdown-item>
                    </el-dropdown-menu>
                </el-dropdown>
            </div>
          </div>
          <div class="clearfix"></div>
      </div>
</template>
<script>
import { mapState } from 'vuex'
export default {
    computed: {
        ...mapState('login',[
            'isLogin',
            'username'
        ])
    },
    methods: {
       logout() {
           this.$store.dispatch("login/logout").then(()=>{
               this.$router.push('/login')
           })
       }
    }
}
</script>
<style scoped>
a {
    text-decoration: none;
}
a:hover {
    text-decoration: none;
}
.info {
    height: 100%;
    position: absolute;
    top: 0;
    right: 50px;
}
.navigation {
   color:white;
   margin-left:20px;
}
.clearfix {
    clear: both;
}
.main_title {
    color:white;
    font-size: 1.8em;
    font-weight: 400;
    height: 100%;
    position: absolute;
    top: 0;
    left: 18px;
    line-height: 50px;
}
.main-header {
    position: relative;
    width: 100%;
    height: 50px;
    z-index: 100;
    background: #398dee;
}
.center {
    line-height:50px;
}
.drop-title {
    font-size:20px;
     color:white;
}
</style>