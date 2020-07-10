<template>
    <div :style="{height:outsize}">
         <editor
       api-key="xjmfct767jqqoy263mz2swmxk18yzwwi2gsz3tc8zlpbakil"
       :init="{
         height: outsize,
         menubar: false,
          language: 'zh_CN',
         plugins: [
           'advlist autolink lists link image charmap print preview anchor',
           'searchreplace visualblocks code fullscreen',
           'insertdatetime media table paste code help wordcount'
         ],
         
             //menubar: 'file edit view insert format tools table tc help',
             toolbar: 'undo redo | bold italic underline strikethrough | fontselect fontsizeselect formatselect | alignleft aligncenter alignright alignjustify | outdent indent |  numlist bullist checklist | forecolor backcolor casechange permanentpen formatpainter removeformat | pagebreak | charmap emoticons | fullscreen  preview save print | insertfile image media pageembed template link anchor codesample | a11ycheck ltr rtl | showcomments addcomment',
       }"
     />
        </div>
</template>

<script>
 import Editor from '@tinymce/tinymce-vue'

 export default {
   components: {
     'editor': Editor
   },
   props:[
       'outsize'
   ]
 }
 </script>