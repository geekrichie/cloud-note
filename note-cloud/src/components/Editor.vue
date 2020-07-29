<template>
    <div :style="{height:outsize}">
         <editor
         :id = "tinymceID"
       api-key="xjmfct767jqqoy263mz2swmxk18yzwwi2gsz3tc8zlpbakil"
        v-model="content"
       :init="{
         height: outsize,
         menubar: false,
         language: 'zh_CN',
         plugins: [
           'advlist autolink lists link image charmap print preview anchor',
           'searchreplace visualblocks code fullscreen',
           'insertdatetime media table paste code help wordcount'
         ],
        
         
         images_upload_handler: (blobInfo, success, failure) => {
            const img = 'data:image/jpeg;base64,' + blobInfo.base64()
            success(img)
         },
         
         
         
             //menubar: 'file edit view insert format tools table tc help',
        toolbar: 'undo redo | bold italic underline strikethrough | fontselect fontsizeselect formatselect | alignleft aligncenter alignright alignjustify | outdent indent |  numlist bullist checklist | forecolor backcolor casechange permanentpen formatpainter removeformat | pagebreak | charmap emoticons | fullscreen  preview save print | insertfile image media pageembed template link anchor codesample | a11ycheck ltr rtl | showcomments addcomment',
       }"
     />
        </div>
</template>

<script>
// import tinymce from 'tinymce/tinymce'
import Editor from '@tinymce/tinymce-vue'

 export default {
   components: {
     'editor': Editor
   },
   data() {
     return {
         content:"",
     }
   },
   props:[
       'outsize',
       'tinymceID'
   ]
 }
 </script>