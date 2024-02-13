<template>
  <el-dialog
      v-model="logvisable"
      :title="'想说点什么?'"
      width="30%"
      style="background-color:rgb(258,258,258)" >
    <el-input v-model="says">
    </el-input>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="logvisable = false">Cancel</el-button>
        <el-button type="primary" @click="logvisable = false;AddComment()">
          Upload
        </el-button>
      </span>
    </template>
  </el-dialog>
<div style="display: flex;position: absolute; flex-direction: column;width: 88%;margin-left: 15%;">
  <div class="ed" >
    <div class="box" style="margin-top: 50px">
  <v-md-editor :model-value="markdown" mode="preview" >

  </v-md-editor>

      <div style="background-color: rgb(255,255,255,0.1);width: 520px;height: 200px;border-radius: 15px;margin-top: 50px;box-shadow: inset 0 -3em 3em rgba(0, 0, 0, 0.1),
                0 0 0 2px rgb(255, 255, 255),
                0.5em 0.5em 3em rgba(255, 255, 255, 0.3);">
        <el-row :gutter="20">
        <el-col>
          <el-avatar
              :src="creater.icon_url" :size="80" style="margin-left: 20px;margin-top: 20px"/>
                    <el-text size="large" style="font-size: 30px;color: white;position: absolute;margin-top: 20px;margin-left: 10px" >{{creater.name}}</el-text>
                    <el-text style="position: absolute;font-size: 20px;color: white;font-family: ggbond;margin-top: 70px;margin-left: 10px" @click="loadMarkdownFile()">{{creater.signature}}</el-text>
        </el-col>
        </el-row>

        <el-row style="margin-top: 30px">
          <el-text style="color: white;font-size: 25px;margin-left: 20px" size="large">获赞:
          <el-text style="color:#f1606a;font-size: 25px">
            {{creater.like}}
          </el-text></el-text>
          <el-text style="color: white;font-size: 25px;margin-left: 20px" size="large">等级:
            <el-text style="color:#7faff6;font-size: 25px">
              {{creater.level}}
            </el-text></el-text>
          <el-text style="color: white;font-size: 25px;margin-left: 20px" size="large">文章数:
            <el-text style="color:#79d79c;font-size: 25px">
              {{creater.article_num}}
            </el-text></el-text>
        </el-row>

      </div>

    </div>

  </div>
  <div class="butbox" style="display: flex;flex-direction: row;">
    <el-button style="margin-left: 60%;width: 120px;margin-top: 30px;height: 50px;border-radius: 10px" color="LightPink" @click="Like" >
      <svg t="1707210896266" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="6482" width="32" height="32"><path d="M510.671749 348.792894S340.102978 48.827055 134.243447 254.685563C-97.636714 486.565724 510.671749 913.435858 510.671749 913.435858s616.107079-419.070494 376.428301-658.749272c-194.095603-194.096626-376.428302 94.106308-376.428301 94.106308z" fill="#FF713C" p-id="6483"></path><path d="M510.666632 929.674705c-3.267417 0-6.534833-0.983397-9.326413-2.950192-16.924461-11.872399-414.71121-293.557896-435.220312-529.448394-5.170766-59.482743 13.879102-111.319341 56.643068-154.075121 51.043536-51.043536 104.911398-76.930113 160.095231-76.930114 112.524796 0 196.878996 106.48115 228.475622 153.195078 33.611515-45.214784 122.406864-148.20646 234.04343-148.20646 53.930283 0 105.46603 24.205285 153.210428 71.941496 45.063335 45.063335 64.954361 99.200326 59.133795 160.920016C935.306982 641.685641 536.758893 915.327952 519.80271 926.859589a16.205077 16.205077 0 0 1-9.136078 2.815116zM282.857183 198.75574c-46.25344 0-92.396363 22.682605-137.127124 67.413365-36.149315 36.157501-51.614541 78.120218-47.25321 128.291898 17.575284 202.089671 352.199481 455.119525 412.332023 499.049037 60.434417-42.86732 395.406538-289.147446 414.567947-492.458945 4.933359-52.344159-11.341303-96.465029-49.759288-134.88199-41.431621-41.423435-85.24243-62.424748-130.242319-62.424748-122.041544 0-220.005716 152.203494-220.989114 153.742547-3.045359 4.806469-8.53335 7.883551-14.101159 7.534603a16.257266 16.257266 0 0 1-13.736863-8.184403c-0.902556-1.587148-91.569532-158.081365-213.690893-158.081364z" fill="#885F44" p-id="6484"></path></svg>
    </el-button>
    <el-button  style="margin-left: 60px;width: 120px;margin-top: 30px;height: 50px;border-radius: 10px" color="LightSkyBlue"  @click="Commentbutton"><svg t="1707210801696" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2636" width="32" height="32"><path d="M520.533333 866.133333c-17.066667 0-29.866667-4.266667-38.4-17.066666l-38.4-38.4H166.4c-38.4 0-68.266667-29.866667-68.266667-68.266667V243.2c0-38.4 29.866667-68.266667 68.266667-68.266667h712.533333c38.4 0 68.266667 29.866667 68.266667 68.266667v494.933333c0 38.4-29.866667 68.266667-68.266667 68.266667h-277.333333l-38.4 38.4c-12.8 17.066667-25.6 21.333333-42.666667 21.333333zM166.4 234.666667c-4.266667 0-12.8 4.266667-12.8 12.8v494.933333c0 4.266667 4.266667 12.8 12.8 12.8H469.333333l55.466667 55.466667 51.2-55.466667h302.933333c4.266667 0 12.8-4.266667 12.8-12.8V243.2c0-4.266667-4.266667-12.8-12.8-12.8H166.4z" fill="#6A3906" p-id="2637"></path><path d="M797.866667 682.666667h-554.666667c-21.333333 0-34.133333-17.066667-34.133333-34.133334V337.066667c0-21.333333 17.066667-38.4 34.133333-38.4h554.666667c21.333333 0 34.133333 17.066667 34.133333 34.133333v311.466667c4.266667 21.333333-12.8 38.4-34.133333 38.4z" fill="#F5CB2B" p-id="2638"></path><path d="M708.266667 443.733333H337.066667c-17.066667 0-29.866667-12.8-29.866667-29.866666 0-17.066667 12.8-29.866667 29.866667-29.866667h371.2c17.066667 0 29.866667 12.8 29.866666 29.866667 0 17.066667-12.8 29.866667-29.866666 29.866666zM512 622.933333H337.066667c-17.066667 0-29.866667-12.8-29.866667-29.866666 0-17.066667 12.8-29.866667 29.866667-29.866667H512c17.066667 0 29.866667 12.8 29.866667 29.866667 0 17.066667-12.8 29.866667-29.866667 29.866666z" fill="#6A3906" p-id="2639"></path></svg></el-button>
  </div>
  <ul class="infinite-list" style="overflow:auto; margin-top: 50px; margin-left: 13%; width: 960px; z-index: 1; background-color: rgb(255,255,255,0.1);border-radius: 10px;box-shadow: 0 0 0 2px rgb(129,41,176)">
    <li v-for="comment in comments" :key="comment.comment.comment_id" class="infinite-list-item" style="display: flex; border-bottom: 2px solid #ffffff;margin-top: 20px;width: 92%">
      <el-row>
      <el-avatar :size="60" style="margin-top: -30px" :src="comment.icon_url"></el-avatar>
      </el-row>
      <el-col >
        <el-row style="margin-top: -85px">
      <el-text style="margin-top: 50px; margin-left: 10px;color: white;font-size: 16px" size="default" >{{comment.name}}</el-text>
        </el-row>
        <el-text style="margin-left: 10px;margin-top: 45px;margin-bottom: 40px;white-space: normal;width: 800px;color: white;font-size: 18px" size="large">{{comment.comment.context}}</el-text>
        <el-row>
          <el-text style="margin-bottom: -65px;margin-left: 680px ;color: white" >{{comment.comment.create_time}}</el-text>
        </el-row>
      </el-col>
    </li>
  </ul>

</div>
</template>

<script>
import { getarticle } from "../../apis/getarticle";
import {getuserinfo} from "../../apis/getuserinfo";
import {useUserInfoStore} from "../../stores/userinfo";
import {ElMessage} from "element-plus";
import{Likeapi} from"../../apis/likeaction"
import {addComment} from "../../apis/addComment";

export default {
  data() {
    return {
      markdown: "",
      t: 6,
      aid:0,
      creater: {
        id: "",
        name: '',
        signature: '',
        like: '',
        article_num: 0,
        level: 0
      },
      comments: [],
      logvisable:false,
      says:''
    };
  },
  created() {
    this.loadMarkdownFile(); // 在组件创建时执行
  },
  methods: {
    async loadMarkdownFile() {
      const n = this.$route.query.article_id
      this.aid=n
      var res = await getarticle({article_id: n})
      this.markdown = res.content
      this.creater.id = res.creater_id
      this.comments=res.comments
      res = await getuserinfo({user_id: this.creater.id, token: useUserInfoStore().getToken()})
      if (res.succed === false) {
        ElMessage.error("获取作者信息失败")
      }
      this.creater.name = res.user.name
      this.creater.like = res.user.like_num
      this.creater.signature = res.user.signature
      this.creater.level = res.user.level
      this.creater.article_num = res.user.articles_num
      this.creater.icon_url=res.user.avatar
      for (const comment of this.comments) {
        const userInfo = await getuserinfo({user_id: comment.user_id, token: useUserInfoStore().getToken()});
        if (userInfo.succed) {
          comment.user_id=userInfo.user.id.toString()
          comment.name = userInfo.user.name;
          comment.icon_url=userInfo.user.avatar;
        } else {
          comment.name = "未知用户";
        }
      }
    },
    async Like(){
      const res=Likeapi({user_id:useUserInfoStore().getUserId(),admin_id:this.creater.id})
      if (res.succed===true){
        ElMessage.success("点赞成功")
      }
      else {
        ElMessage.error("点赞失败")
      }
},
    async AddComment(){
      const res=addComment({user_id:useUserInfoStore().getUserId(),article_id:this.aid,context:this.says})
      if(res.succed===true){
        ElMessage.success("Accepted")
      }
      else{
        ElMessage.error("评论失败")
      }
    },
    Commentbutton(){
      this.logvisable=!this.logvisable
    }
  },
}
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

  position: relative;
  left: 10%;
  width: 100%;
  gap: 50px;
  display: flex;
  margin-top: 10px;
}
.infinite-list {
  height: 300px;
  width: 100%;
  padding: 0;
  margin: 0;
  list-style: none;
}
.infinite-list .infinite-list-item {
  display: flex;
  align-items: center;
  height: 100px;
  margin: 10px;
}
.infinite-list .infinite-list-item + .list-item {
  margin-top: 10px;
}
</style>
