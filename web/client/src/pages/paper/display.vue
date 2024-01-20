<template>
  <div class="ed" >
    <div class="box">
  <v-md-editor :model-value="markdown" mode="preview" ></v-md-editor>
      <div style="background-color: rgb(255,255,255,0.1);width: 500px;height: 200px;border-radius: 15px;margin-top: 50px;box-shadow: inset 0 -3em 3em rgba(0, 0, 0, 0.1),
                0 0 0 2px rgb(255, 255, 255),
                0.5em 0.5em 3em rgba(255, 255, 255, 0.3);">
        <el-row :gutter="20">
        <el-col>
          <el-avatar
              src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" :size="80" style="margin-left: 20px;margin-top: 20px"/>
                    <el-text size="large" style="font-size: 30px;color: white;position: absolute;margin-top: 20px;margin-left: 10px" >ggbond</el-text>
                    <el-text style="position: absolute;font-size: 20px;color: white;font-family: ggbond;margin-top: 70px;margin-left: 10px">哼哼哼啊啊啊啊</el-text>
        </el-col>
        </el-row>
        <el-row style="margin-top: 30px">
          <el-text style="color: white;font-size: 25px;margin-left: 20px" size="large">获赞:
          <el-text style="color:#bb505d;font-size: 25px">
            {{t}}
          </el-text></el-text>
          <el-text style="color: white;font-size: 25px;margin-left: 20px" size="large">等级:
            <el-text style="color:#bb505d;font-size: 25px">
              {{t}}
            </el-text></el-text>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      markdown: '## Quick Start\n' +
          '\n' +
          '```js\n' +
          'import Vue from \'vue\';\n' +
          'import VueMarkdownEditor from \'@kangc/v-md-editor\';\n' +
          'import \'@kangc/v-md-editor/lib/style/base-editor.css\';\n' +
          'import vuepressTheme from \'@kangc/v-md-editor/lib/theme/vuepress.js\';\n' +
          '\n' +
          'VueMarkdownEditor.use(vuepressTheme);\n' +
          '\n' +
          'Vue.use(VueMarkdownEditor);\n' +
          '```',
      t:6,
    };
  },
  mounted() {
    this.loadMarkdownFile();
  },
  methods: {
    async loadMarkdownFile() {

    },
  },
};
</script>

<style scoped lang="scss">
.ed {
  height: auto;
  width: 100%;
  left: 50%;
  z-index: 1;
  display: flex;
  flex-direction: row;
}
.box{
  position: absolute;
  left: 10%;
  width: 85%;
  gap: 50px;
  display: flex;
}
</style>
