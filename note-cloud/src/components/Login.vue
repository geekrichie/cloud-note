<template>
      <div class="login-background">
            <el-card class="box-card login">
                <div slot="header" class="clearfix">
                    <span>登录</span>
                </div>
               <div>
                  <div class="pad"><label> 用户名或者邮箱</label></div>
                  <el-input v-model="username"></el-input>
               </div>
               <div>
                    <div class="pad"><label> 密码</label></div>
                    <el-input v-model="password" show-password></el-input>
                </div>
                <div class="pad">
                    <el-button class = "submit-btn" type="success" @click="login">登录</el-button>
                </div>
            </el-card>
        </div>
</template>
<script>
import { mapState, mapActions } from 'vuex'
export default {
    data() {
          return {
              username : "",
              password : "",
          }
    },
    computed :{
         ...mapState('login', [
             'isLogin'
         ])
    },
    methods: {
        ...mapActions('login',[
            'userLogin'
        ]),
        async login() {
            let userInfo = new FormData();
            userInfo.append("username",this.username)
            userInfo.append("password",this.password)
            await this.userLogin(userInfo)
            if(this.isLogin) {
                this.$router.push("/")
            }
        },
    }
}
</script>
<style scoped>
.login {
    width:400px;
    margin:0 auto;
    text-align:left;
}
.clearfix {
    text-align:center;
}
.pad {
    margin-top:5px;
    margin-bottom:5px;
}
.login-background {
    padding-top:20px;
    background:#e9ebee;
    height:400px;
}
.submit-btn {
    width:100%;
}
label {
 height:20px;
}
.tabs {
    margin:0 auto;
}
</style>