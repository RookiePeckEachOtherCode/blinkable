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
      <el-upload
          class="avatar-uploader"
          action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
          :show-file-list="false"
      >
        <img v-if="imageUrl" :src="imageUrl" class="avatar" />
        <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
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
import { reactive } from 'vue'

const form = reactive({
  user_name: '',
  title: '',
  signature: '',
  date2: '',
  delivery: false,
  type: [],
  resource: '',
  desc: '',
})

const onSubmit = () => {
  console.log('submit!')
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