<template>
    <div class="main">
        <sidebar ></sidebar>
    </div>
    
</template>
<script>
import sidebar from './sidebar'
export default {
   components :{
       sidebar
   }
}
</script>
<style scoped>
.main {
    text-align:left;
    position: absolute;
    top: 50px;
    left: 0;
    right: 0;
    bottom: 0px;
}

</style>