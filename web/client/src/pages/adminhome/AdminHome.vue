<template>
  <el-form :model="form" label-width="120px"
           style="background-color: rgb(255,255,255);
           width: 800px;z-index: 1;
           right: 27%;position: absolute;
           border-radius: 20px;

" :size="'large'" :rules="rules" ref="formRef">
    <el-form-item label="昵称" style="margin-top: 20px ;font-size: 30px" prop="user_name">
      <el-input  v-model="form.user_name" style="margin-right: 500px" />
    </el-form-item>
    <el-form-item label="称号" prop="title">
    <el-input v-model="form.title" style="margin-right: 400px" :size="'large'"></el-input>
    </el-form-item>
    <el-form-item label="头像" >
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
    <el-form-item label="git主页地址" prop="git_url">
      <el-input v-model="form.git_url" style="margin-right: 400px" :size="'large'"></el-input>
    </el-form-item>
    <el-form-item label="签名" prop="signature">
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
import {ElMessage, type FormInstance, FormRules, genFileId, UploadInstance, UploadProps} from "element-plus";
import {UploadRawFile} from "element-plus/lib/components";
import {ref} from "vue";
import {uploadinfo} from "../../apis/uploadinfo"
import {sync} from "rimraf";
useUserInfoStore().authFromLocal();
const formRef=ref<FormInstance>()
interface Form{
  user_name:string,
  title:string,
  signature: string,
  icon_url:string,
  resource: string,
  level:number,
  git_url:string,
  background_url:string,
}
const form = reactive<Form>({
  user_name: '',
  title: '',
  signature: '',
  icon_url:'',
  resource:'',
  level:0,
  git_url:'',
  background_url:'',
});
const rules=reactive<FormRules<Form>>({
  user_name:[
    { required: true, message: 'WHO ARE YOU?', trigger: 'blur' },
    { min: 2, max: 12, message: '别太荒谬，2-12个字符够了吧', trigger: 'blur' },
  ],
  title:[
    { required: true, message: '巨佬别装蒟蒻', trigger: 'blur' },
    { min: 2, max: 12, message: '这真的合理吗？', trigger: 'blur' },
  ],
  signature:[
    { required: true, message: '随便写个艹也行', trigger: 'blur' },
    { min: 1, max: 114514, message: '?真的有人写超过114514个字符？ ', trigger:'blur' },
  ],
  git_url:[
    { required: true, message: '写前端的太拉了，不写交不了表单', trigger: 'blur' },
    { min: 1, max: 114514, message: '?谁家网页这么离谱？', trigger:'blur' },
  ]
})
let res;
const icon_url=useUserInfoStore().getIcon();
onBeforeMount(async () => {
  const user_id=useUserInfoStore().getUserId()
  const token=useUserInfoStore().getToken()
  console.log(user_id,token)
  res = await getuserinfo({user_id,token})
  console.log(res);
  form.user_name=res.user.name
  form.icon_url=res.user.avatar
  form.level=res.user.level
  form.signature=res.user.signature
  form.background_url=res.user.background_img
  form.title=res.user.title
  form.git_url=res.user.github_url
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
const onSubmit = async() => {
  const result= await uploadinfo({signature: form.signature, title: form.title, token: token, username: form.user_name, user_id:user_id,git_url:form.git_url})
  if(result.succed==true){
    ElMessage.success("修改成功")
  }
  else{
    ElMessage.error("修改失败，请检表单是否填满")
  }
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