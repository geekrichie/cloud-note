<template>
    <div>我是主页</div>
    
</template>
<script>
import api from '@/api/login'
export default {
    created() {
      api.getHello().then(res=> {
          console.log(res.msg)
      })
    }
}
</script>