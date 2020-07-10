<template>
    <div class="sidebar">
           <div class="folder-menu">
                <el-menu
                    class="el-menu-vertical-demo"        
                    background-color="#545c64"
                    text-color="#fff"
                    :default-active="$router.path"
                    router
                    >  
                    <el-menu-item class="add-note">   
                        <div class="el-icon-plus"></div>
                        <span slot="title">新增笔记</span> 
                    </el-menu-item>
                    <el-submenu index="2">
                        <template slot="title">我的文件夹</template>  
                        <el-menu-item v-for="(folder, index) in navMenu" :key="index" :index="folder.path" 
                        :class="$route.path==folder.path?'is-active':''">
                            <span slot="title">{{folder.name}}</span>
                        </el-menu-item>   
                    </el-submenu>
                </el-menu>
           </div>
           <div class="article-list">
                <el-menu
                    class="el-menu-vertical-article"        
                    background-color="#ffffff"
                    text-color="#000">  
                    <el-menu-item v-for="(file, index) in selectDocument"
                    :key ="index" :index="file.file_id" @click="articleChange(file.file_id)">   
                        <span slot="title">{{ file.title }}</span> 
                    </el-menu-item>
                </el-menu>
           </div>
           <div class="article-editor">
               <editor outsize="100%"></editor>
               </div>
    </div>
</template>
<script>
import res from "../assets/json/interface.json"
import Editor from "./Editor"
export default {
    components:{
       Editor
    },
    data() {
         return {
            structure:[],
            documents:[],
            navMenu:[],
            outsize:{
                height:"100%"
            }
         }
    },
    computed: {
        selectDocument(){
            let folderid = this.$route.params.fid
            let all_files = res.data.file
            let ret = []
            for(var i=0;i<all_files.length;i++) {
                if(all_files[i].folder_id == folderid) {
                    ret.push(all_files[i])
                }
            }
            return ret;
        }
    },
    methods:{
          articleChange(id) {
            //  let path = this.$route.path;
             let noteid = this.$route.params.noteid
             let fid = this.$route.params.fid
             if( id != noteid) {
                 this.$router.push({path:`/file/${fid}/note/${id}`})
             }
          }
    },
    created() {
        this.structure = res.data.structure
        for(var i=0;i<this.structure.length;i++) {
            let navItem = {} 
            let id = this.structure[i]["folder_id"]
            let name = this.structure[i]["name"]
            let files = this.structure[i]["files"]
            let path = "";
            if (files == null || files.length == 0) {
                path= "/file/"+ id + "/note/"+"empty"
            }else {
                let file = files[0];
                path= "/file/"+ id + "/note/"+file
            }
            navItem.id = id;
            navItem.path = path;
            navItem.name = name;
            this.navMenu.push(navItem)
        }
    }
}
</script>
<style scoped>
.sidebar {
    width:300px;
    height:100%;
}
.el-menu-item.is-active {
    background-color:#d3d3d3
}
.el-icon-plus {
    background-image:url("../assets/images/plus.png")
}
.article-list {
    position:absolute;
    top:0px;
    left:301px;
    width:300px;
    bottom :0px;
}
.article-editor{
    position:absolute;
    top:0px;
    left:601px;
    bottom :0px;
    right:50px;
}
ul,.folder-menu {
    height:100%;
}
</style>