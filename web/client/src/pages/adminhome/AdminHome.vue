<template>
  <el-form :model="form" label-width="120px"
           style="background-color: rgb(255,255,255);
           width: 800px;z-index: 1;
           right: 27%;position: absolute;
           border-radius: 20px;

" :size="'large'">
    <el-form-item label="昵称" style="margin-top: 20px ;font-size: 30px">
      <el-input v-model="form.user_name" style="margin-right: 500px" />
    </el-form-item>
    <el-form-item label="称号">
    <el-input v-model="form.title" style="margin-right: 400px" :size="'large'"></el-input>
    </el-form-item>
    <el-form-item label="头像">
      <div style="display: flex; align-items: center;">
        <el-upload
            class="avatar-uploader"
            :auto-upload="true"
            action="http://127.0.0.1:10000/blinkable/user/update/avatar"
            :show-file-list="false"
            :before-upload="handleUpload_icon"
        >
          <img v-if="imageUrl" :src="imageUrl" class="avatar" alt=""/>
          <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
        </el-upload>
        <el-image :src=icon_url style="margin-left: 20px;width: 180px;height: 180px;border-radius: 5px"></el-image>
      </div>
    </el-form-item>
    <el-form-item label="背景图">
      <el-upload
          ref="upload"
          class="upload-demo"
          action="http://127.0.0.1:10000/blinkable/user/background"
          :limit="1"
          :on-exceed="handleExceed"
          :auto-upload="false"
          :data="{'user_id':user_id}"
      >
        <template #trigger>
          <el-button type="primary" color="#69D0D7" style="color: white">选择图片</el-button>
        </template>
        <el-button class="ml-3" type="success" @click="upbackgourd" style="margin-left: 20px;margin-bottom: 5px">
          上传
        </el-button>
      </el-upload>
    </el-form-item>
    <el-form-item label="签名">
      <el-input v-model="form.signature" type="textarea" style="margin-right: 200px" :size="'large'"/>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">Save</el-button>
      <el-button>Cancel</el-button>
    </el-form-item>
  </el-form>
</template>

<script lang="ts" setup>
import { reactive, onBeforeMount } from 'vue';
import { getuserinfo } from '../../apis/getuserinfo';
import { useUserInfoStore } from '../../stores/userinfo';
import {upload_iconApi} from "../../apis/uploadicon";
import {genFileId, UploadInstance, UploadProps} from "element-plus";
import {UploadRawFile} from "element-plus/lib/components";
import {ref} from "vue";
import {uploadinfo} from "../../apis/uploadinfo"
useUserInfoStore().authFromLocal();
const form = reactive({
  user_name: '',
  title: '',
  signature: '',
  icon_url:'',
  resource: '',
  desc: '',
  level:0,
  background_url:'',
});
let res;
const icon_url=useUserInfoStore().getIcon();
onBeforeMount(async () => {

  res = await getuserinfo(Number(useUserInfoStore().getUserId()));
  console.log(res);
  form.user_name=res.user_name
  form.icon_url=res.icon_url
  form.level=res.level
  form.signature=res.signature
  form.background_url=res.background_url
  form.title=res.title
});
const user_id=useUserInfoStore().getUserId()
const token=useUserInfoStore().getToken()
const handleUpload_icon=(file)=>{
   const result=upload_iconApi({"file":file,"user_id":user_id})
  return false
}
const upload = ref<UploadInstance>()
const handleExceed: UploadProps['onExceed'] = (files) => {
  upload.value!.clearFiles()
  const file = files[0] as UploadRawFile
  file.uid = genFileId()
  upload.value!.handleStart(file)
}
const onSubmit = () => {
  const result=uploadinfo({signature: form.signature, title: form.title, token: token, username: form.user_name, user_id:user_id})
};
const upbackgourd = (file) => {
  upload.value!.submit()
}
</script>

<style scoped>
.avatar-uploader .avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>

<style>
.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}
</style>