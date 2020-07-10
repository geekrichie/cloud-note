<template>
    <div class="main">
        <sidebar ></sidebar>
        
       <div class="article-editor">
            <editor outsize="100%" tinymceID="main-editor" ></editor>
        </div>
    </div>
    
</template>
<script>
import sidebar from './sidebar'
import Editor from "./Editor"
export default {
   components :{
       sidebar,
       Editor
   },
   data() {
       return  {
            outsize:{
                height:"100%"
            }
       }
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
.article-editor{
    position:absolute;
    top:0px;
    left:601px;
    bottom :0px;
    right:50px;
}

</style>