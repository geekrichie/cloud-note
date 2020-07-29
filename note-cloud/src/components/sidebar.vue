<template>
    <div class="sidebar">
           <div class="folder-menu">
                <el-menu
                    class="el-menu-vertical-demo"        
                    background-color="#545c64"
                    text-color="#fff"
                    :default-active="defaultActive"
                    router
                    >  
                    <el-menu-item class="add-note">   
                        <div class="el-icon-plus"></div>
                        <span slot="title">新增笔记</span> 
                    </el-menu-item>
                    <el-submenu index="2">
                        <template slot="title">我的文件夹</template>  
                        <el-menu-item v-for="(folder, index) in navMenu" :key="index" :index="folder.id" 
                        :route="folder.path"
                       >
                            <span slot="title">{{folder.name}}</span>
                        </el-menu-item>   
                    </el-submenu>
                </el-menu>
           </div>
           <div class="article-list">
                <el-menu
                    class="el-menu-vertical-article"        
                    background-color="#ffffff"
                    :default-active="$route.params.noteid"
                    text-color="#000">  
                    <el-menu-item v-for="(file, index) in selectDocument"
                    :key ="index" :index="file.file_id.toString()" @click="articleChange(file.file_id)">   
                        <span slot="title">{{ file.title }}</span> 
                    </el-menu-item>
                </el-menu>
           </div>
    </div>
</template>
<script>
import api from "../api/login"
export default {
    data() {
         return {
            structure:[],
            documents:[],
            navMenu:[],
            outsize:{
                height:"100%"
            },
            selectDocument: []
         }
    },
    computed: {
         defaultActive() {
            return this.$route.path.split('/')[2];
        }
    },
    watch: {
       $route(to, from) {
           
             // 对路由变化作出响应..
             let newFid = to.params.fid
             let oldFid = from.params.fid
             this.foldersOwnArticle(newFid, oldFid)
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
          },
          foldersOwnArticle(newFid, oldFid) {
            if(newFid != undefined && newFid != "" && newFid != oldFid) {
                 let params = {}
                 params.folderid = newFid;
                 this.selectDocument = [];
                api.getFilesByFolder(params).then((res)=> {
                   let data = res.data
                   for(let i = 0;i<data.length;i++) {
                       let item = {}
                       item.file_id = data[i].ID 
                       item.title = data[i].Title
                       item.content = data[i].Content
                       this.selectDocument.push(item)
                   }
                })
             }
          },
          onLoadFn() {
             let fid = this.$route.params.fid
             this.foldersOwnArticle(fid, "")
          }
    },
    created() {
        api.getUserFolders().then((res)=> {
            this.structure = res.data;
            for(var i=0;i<this.structure.length;i++) {
                let navItem = {} 
                let id = this.structure[i]["folder_id"]
                let name = this.structure[i]["name"]
                let files = this.structure[i]["files"]
                let path = "";
                if (files == null || files.length == "") {
                    path= "/file/"+ id + "/note/"+"empty"
                }else {
                    let file = files;
                    path= "/file/"+ id + "/note/"+file
                }
                navItem.id = id;
                navItem.path = path;
                navItem.name = name;
                this.navMenu.push(navItem)
            }
            this.onLoadFn()
        })

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