<template>
      <div>
          <div class="main_title">个人云笔记</div>
          <div class="info">
            <div v-if="!isLogin">
                <el-button>登录</el-button>
                <el-button>注册</el-button>
            </div>
             <div v-if="isLogin">
                 <label>{{ username  }}</label>
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
    }
}
</script>
<style scoped>
.info {
    float: right;
    padding: 8px 24px 8px 0;
}
.clearfix {
    clear: both;
}
.main_title {
    text-align: left;
    font-size: 1.8em;
    font-weight: 400;
    padding: 8px 24px 8px 0;
    float: left;
}
</style>