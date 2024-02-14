<template>
  <el-dialog
    v-model="logvisable"
    :title="'给这篇文章一个标题吧'"
    width="30%"
    style="background-color:rgb(258,258,258)" >
    <el-input v-model="title">
    </el-input>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="logvisable = false">Cancel</el-button>
        <el-button type="primary" @click="logvisable = false;upload()">
          Upload
        </el-button>
      </span>
    </template>
  </el-dialog>
  <div style="display: flex;flex-direction: column ;gap: 50px">
  <div class="ed">
    <div class="box">
      <v-md-editor v-model="text" height="600px" style="position: absolute;" @save="saveinlocal" ></v-md-editor>
    </div>
    </div>
    <el-button type="primary" style="z-index: 1;margin-left: 500px;
            width: 135px;
            height: 55px;
            border-radius: 20px;
            font-size: 30px;
            font-family: ggbond;
            box-shadow:
                inset 0 -3em 3em rgba(0, 0, 0, 0.1),
                0 0 0 2px rgb(255, 255, 255);
            transition: all 0.2s;" size="large" @click="logvisable=true">发布</el-button>
  </div>

</template>

<script>
import Editor from "@/App.vue";
import {ElMessage} from "element-plus";
import {uploadmd} from "../../apis/uploadmd"
import { ref } from 'vue';
import {useUserInfoStore} from "../../stores/userinfo"
export default {
  components: { Editor },
  setup(props, context){
    const title=ref('')
    return{
      title,
    }
  },
  data() {
    return {
      text: '',
      logvisable:false,
    };
  },
  methods:{
    saveinlocal(){
      const content = this.text;
      const blob = new Blob([content], { type: 'text/markdown' });
      const url = window.URL.createObjectURL(blob);
      const link = document.createElement('a');
      link.href = url;
      link.download = 'templefile.md';
      document.body.appendChild(link);
      link.click();
      document.body.removeChild(link);
      window.URL.revokeObjectURL(url);
    },
   async upload(){
      const content = this.text;
     const formData = new FormData();
     formData.append('user_id', useUserInfoStore().getUserId());
     formData.append('content', this.text);
     formData.append('title',this.title);
     formData.append('token',useUserInfoStore().getToken());
     const response = await uploadmd(formData);
     if(response.status_code===0){
       ElMessage.success("上传成功")
       return;
     }
    },
  }
};
</script>

<style>
.ed {
  height: 600px;
  width: 100%;
  left: 50%;
  z-index: 1;
}
.box{
  position: absolute;
  left: 10%;
  width: 80%;
}
</style>
