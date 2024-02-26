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
                    <el-text style="position: absolute;font-size: 20px;color: white;font-family: ggbond;margin-top: 70px;margin-left: 10px;margin-right: 10px" @click="loadMarkdownFile()">{{creater.signature}}</el-text>
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
<el-row>
  <el-button  @click="deletearticle" v-if="this.user_id===this.creater.id" style="position: absolute;margin-left: 700px;width: 120px;margin-top: 30px;height: 50px;border-radius: 10px" color="LightSlateGray" ><svg t="1708088484800" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3496" width="32" height="32"><path d="M214.304 829.161c0 71.209-1.795 76.205 76.948 76.205h406.428c76.127 0 85.637-2.112 85.637-76.205v-411.525h-569.012v411.525z" fill="#CFCF5A" p-id="3497"></path><path d="M336.24 769.902c-7.472 0-13.53 6.058-13.53 13.53v27.119c0.397 7.173 6.313 12.841 13.55 12.841s13.152-5.664 13.549-12.811l0.001-27.154c-0.031-7.472-6.094-13.521-13.568-13.53zM336.24 471.838c-0.019 0-0.037 0-0.058 0-7.44 0-13.472 6.031-13.472 13.472 0 0.021 0 0.042 0 0.061v230.331c0 7.485 6.065 13.55 13.55 13.55s13.55-6.065 13.55-13.55v-230.317c0-0.003 0-0.007 0-0.020 0-7.472-6.058-13.53-13.53-13.53-0.007 0-0.025 0-0.042 0zM498.82 471.838c-0.006 0-0.023 0-0.037 0-7.462 0-13.509 6.052-13.509 13.509 0 0.004 0 0.007 0 0.021v325.165c0.286 7.261 6.237 13.036 13.54 13.036s13.254-5.775 13.54-13.010v-325.171c0-0.003 0-0.007 0-0.020 0-7.472-6.058-13.53-13.53-13.53 0 0 0 0 0 0zM661.383 471.838c-0.006 0-0.023 0-0.037 0-7.462 0-13.509 6.052-13.509 13.509 0 0.004 0 0.007 0 0.021v325.165c0 0.004 0 0.019 0 0.023 0 7.472 6.058 13.53 13.53 13.53 2.512 0 4.861-0.685 6.877-1.878 4.016-2.342 6.713-6.694 6.713-11.681 0 0 0-0.001 0-0.001v-325.148c-0.001-3.743-1.525-7.123-3.982-9.57-2.435-2.46-5.813-3.982-9.547-3.982-0.007 0-0.027 0-0.042 0z" fill="#848D2F" p-id="3498"></path><path d="M575.789 119.59h-162.585c-67.399 0-75.833 8.882-76.812 67.71h324.601c-2.597-60.78-20.385-67.71-85.209-67.71z" fill="#CFCF5A" p-id="3499"></path><path d="M913.156 220.063c-20.229-20.239-48.18-32.757-79.052-32.764h-118.949c-3.67-82.474-34.384-121.914-139.388-121.914h-162.585c-97.137 0-127.419 31.26-130.817 121.914h-105.318c0 0 0 0 0 0-61.732 0-111.779 50.045-111.779 111.779 0 0.004 0 0.007 0 0.021v6.793c0.023 55.713 40.844 101.887 94.205 110.274l0.626 412.992c0 112.287 33.507 130.41 131.127 130.41h406.428c105.278 0 139.835-26.809 139.835-130.41v-411.681c60.223-1.858 108.336-51.107 108.364-111.601v-6.795c0-0.033 0-0.076 0-0.117 0-30.817-12.498-58.715-32.703-78.899zM413.184 119.59h162.585c64.823 0 82.607 6.913 85.209 67.71h-324.58c0.975-58.829 9.39-67.71 76.794-67.71zM783.315 829.161c0 74.097-9.491 76.205-85.637 76.205h-406.428c-78.745 0-76.948-4.996-76.948-76.205v-411.525h569.034v411.525zM891.7 305.874c-0.019 15.908-6.461 30.31-16.874 40.75-10.403 10.377-24.793 16.81-40.681 16.81-0.004 0-0.019 0-0.023 0h-657.050c-15.903-0.001-30.295-6.433-40.734-16.835-10.406-10.432-16.842-24.833-16.844-40.738v-6.782c0-0.003 0-0.007 0-0.020 0-31.79 25.767-57.562 57.562-57.562 0.004 0 0.007 0 0.021 0h657.050c0 0 0 0 0 0 31.799 0 57.579 25.778 57.579 57.579v6.795z" fill="#504C23" p-id="3500"></path></svg></el-button>
  <el-button  style="margin-left: 60%;width: 120px;margin-top: 30px;height: 50px;border-radius: 10px" color="LightPink" @click="Like" >
      <svg t="1707210896266" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="6482" width="32" height="32"><path d="M510.671749 348.792894S340.102978 48.827055 134.243447 254.685563C-97.636714 486.565724 510.671749 913.435858 510.671749 913.435858s616.107079-419.070494 376.428301-658.749272c-194.095603-194.096626-376.428302 94.106308-376.428301 94.106308z" fill="#FF713C" p-id="6483"></path><path d="M510.666632 929.674705c-3.267417 0-6.534833-0.983397-9.326413-2.950192-16.924461-11.872399-414.71121-293.557896-435.220312-529.448394-5.170766-59.482743 13.879102-111.319341 56.643068-154.075121 51.043536-51.043536 104.911398-76.930113 160.095231-76.930114 112.524796 0 196.878996 106.48115 228.475622 153.195078 33.611515-45.214784 122.406864-148.20646 234.04343-148.20646 53.930283 0 105.46603 24.205285 153.210428 71.941496 45.063335 45.063335 64.954361 99.200326 59.133795 160.920016C935.306982 641.685641 536.758893 915.327952 519.80271 926.859589a16.205077 16.205077 0 0 1-9.136078 2.815116zM282.857183 198.75574c-46.25344 0-92.396363 22.682605-137.127124 67.413365-36.149315 36.157501-51.614541 78.120218-47.25321 128.291898 17.575284 202.089671 352.199481 455.119525 412.332023 499.049037 60.434417-42.86732 395.406538-289.147446 414.567947-492.458945 4.933359-52.344159-11.341303-96.465029-49.759288-134.88199-41.431621-41.423435-85.24243-62.424748-130.242319-62.424748-122.041544 0-220.005716 152.203494-220.989114 153.742547-3.045359 4.806469-8.53335 7.883551-14.101159 7.534603a16.257266 16.257266 0 0 1-13.736863-8.184403c-0.902556-1.587148-91.569532-158.081365-213.690893-158.081364z" fill="#885F44" p-id="6484"></path></svg>
    </el-button>
    <el-button  style="margin-left: 60px;width: 120px;margin-top: 30px;height: 50px;border-radius: 10px" color="LightSkyBlue"  @click="Commentbutton"><svg t="1707210801696" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2636" width="32" height="32"><path d="M520.533333 866.133333c-17.066667 0-29.866667-4.266667-38.4-17.066666l-38.4-38.4H166.4c-38.4 0-68.266667-29.866667-68.266667-68.266667V243.2c0-38.4 29.866667-68.266667 68.266667-68.266667h712.533333c38.4 0 68.266667 29.866667 68.266667 68.266667v494.933333c0 38.4-29.866667 68.266667-68.266667 68.266667h-277.333333l-38.4 38.4c-12.8 17.066667-25.6 21.333333-42.666667 21.333333zM166.4 234.666667c-4.266667 0-12.8 4.266667-12.8 12.8v494.933333c0 4.266667 4.266667 12.8 12.8 12.8H469.333333l55.466667 55.466667 51.2-55.466667h302.933333c4.266667 0 12.8-4.266667 12.8-12.8V243.2c0-4.266667-4.266667-12.8-12.8-12.8H166.4z" fill="#6A3906" p-id="2637"></path><path d="M797.866667 682.666667h-554.666667c-21.333333 0-34.133333-17.066667-34.133333-34.133334V337.066667c0-21.333333 17.066667-38.4 34.133333-38.4h554.666667c21.333333 0 34.133333 17.066667 34.133333 34.133333v311.466667c4.266667 21.333333-12.8 38.4-34.133333 38.4z" fill="#F5CB2B" p-id="2638"></path><path d="M708.266667 443.733333H337.066667c-17.066667 0-29.866667-12.8-29.866667-29.866666 0-17.066667 12.8-29.866667 29.866667-29.866667h371.2c17.066667 0 29.866667 12.8 29.866666 29.866667 0 17.066667-12.8 29.866667-29.866666 29.866666zM512 622.933333H337.066667c-17.066667 0-29.866667-12.8-29.866667-29.866666 0-17.066667 12.8-29.866667 29.866667-29.866667H512c17.066667 0 29.866667 12.8 29.866667 29.866667 0 17.066667-12.8 29.866667-29.866667 29.866666z" fill="#6A3906" p-id="2639"></path></svg></el-button>
</el-row>
  <ul class="infinite-list" style="overflow:auto; margin-bottom: 50px;margin-top: 50px; margin-left: 13%; width: 960px; z-index: 1; background-color: rgb(255,255,255,0.1);border-radius: 10px;box-shadow: 0 0 0 2px rgb(129,41,176)">
    <li v-for="comment in comments" :key="comment.comment_id" class="infinite-list-item" style="display: flex; border-bottom: 2px solid #ffffff;margin-top: 20px;width: 92%">
      <el-row>
      <el-avatar :size="60" style="margin-top: -30px" :src="comment.icon_url"></el-avatar>
      </el-row>
      <el-col >
        <el-row style="margin-top: -85px">
      <el-text style="margin-top: 50px; margin-left: 10px;color: white;font-size: 16px" size="default" >{{comment.name}}</el-text>
        </el-row>
        <el-text style="margin-left: 10px;margin-top: 45px;margin-bottom: 40px;white-space: normal;width: 800px;color: white;font-size: 18px" size="large">{{comment.context}}</el-text>
        <el-row>
          <el-text style="margin-bottom: -65px;margin-left: 600px ;color: white" >{{comment.create_time}}</el-text>
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
import {deletearticle} from "../../apis/deletearticle";

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
      says:'',
      user_id:'',
    };
  },
  created() {
    this.loadMarkdownFile(); // 在组件创建时执行
  },
  methods: {
    async loadMarkdownFile() {
      const n = this.$route.query.article_id
      this.aid=n
      let res = await getarticle({article_id: n})
      this.markdown = res.content
      this.creater.id = res.creater_id.toString()
      this.comments=res.Comments
      res = await getuserinfo({user_id: this.creater.id, token: useUserInfoStore().getToken()})
      if (res.succed === false) {
        ElMessage.error("获取作者信息失败")
      }
      this.user_id=useUserInfoStore().getUserId()
      this.creater.name = res.user.name
      this.creater.like = res.user.like_num
      this.creater.signature = res.user.signature
      this.creater.level = res.user.level
      this.creater.article_num = res.user.articles_num
      this.creater.icon_url=res.user.avatar
      for (let comment of this.comments) {
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
    async deletearticle(){
      const res=await deletearticle({user_id:useUserInfoStore().getUserId(),token:useUserInfoStore().getToken(),article_id:this.$route.query.article_id})
      if(res.succed===true){
        ElMessage.success("Accept")
      }else{
        ElMessage.error("删除失败")
      }
    },
    async Like(){
      const res= await Likeapi({from_user_id:useUserInfoStore().getUserId(),user_id:this.creater.id,token:useUserInfoStore().getToken()})
      if (res.succed===true){
        ElMessage.success("点赞成功")
      }
      else {
        ElMessage.error("点赞失败")
      }
},
    async AddComment(){
     // console.log(useUserInfoStore().getUserId())
      const res=await addComment({user_id:useUserInfoStore().getUserId(),article_id:this.aid,context:this.says})
      if(res.status_code===200){
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
