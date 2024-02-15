<template>
  <div style="width: 1500px;z-index: 1;position: absolute;left: 6%">
    <ul style="font-size: 35px;color: white;font-family: ggbond">
      <li v-for="backendArticles in paginatedArticles" :key="backendArticles.article_id"
          class="list" @click="goview(backendArticles.article_id)">
        <span>{{ backendArticles.title }}</span>
        <div>
        <span style="font-size: 20px;margin-top: 15px;margin-right: 70px;color:#7fb80e ">{{ backendArticles.create_time }}</span>
          <el-avatar :size="50" style="position: absolute;right: 1%"
              :src=backendArticles.icon_url></el-avatar>
        </div>
      </li>
    </ul>
    <el-pagination
        v-if="ArticleSum > 0"
        :current-page="currentPage"
        :page-sizes="[10, 20, 30, 40]"
        :page-size="pageSize"
        :total="ArticleSum"
        :background="true"
        @size-change="handleSizeChange"
        @current-change="handlePageChange"
        style="
         margin-top: 20px;
         margin-left: 60px;
"/>
  </div>
</template>
<script>
import {getarticlelist} from "../../apis/getarticlelist"
import {getarticlesum} from "../../apis/getarticlesum"
import {useUserInfoStore} from"../../stores/userinfo"
import {getuserinfo} from "../../apis/getuserinfo";

export default {
  computed: {
    paginatedArticles() {
      const startIndex = (this.currentPage - 1) * this.pageSize;
      const endIndex = startIndex + this.pageSize;
      return this.backendArticles
    },
  },
  data() {
    return {
      pageSize: 6,
      currentPage: 1,
      backendArticles: [],
      ArticleSum:0,
    };
  },
  created() {
    this.fetchArticles();
    this.getsum();
  },
  methods: {
    handleSizeChange(val) {
      this.pageSize = val;
      this.currentPage = 1;
    },
    async handlePageChange(val) {
      this.currentPage = val;
      await this.fetchArticles()
    },
    goview(id){
      this.$router.push({name:"paper-display",query:{article_id:id}})
    },
    async fetchArticles(){
      const startIndex = (this.currentPage - 1) * this.pageSize;
      const endIndex = startIndex + this.pageSize;
      const result=await getarticlelist({start:startIndex,end:endIndex});
      this.backendArticles=result.articles;
      for(let artciles of this.backendArticles){
        const userInfo = await getuserinfo({user_id: artciles.creater_id, token: useUserInfoStore().getToken()});
        console.log(userInfo)
        artciles.icon_url=userInfo.user.avatar;
      }
    },
    async getsum(){
      const result= await getarticlesum();
      this.ArticleSum=result.sum;
    }
  },
};
</script>
<style scoped>
  .list{
    margin-top: 30px;
    background-color: rgb(255,255,255,0.1);
    border-radius: 30px;
    padding: 10px;
    display: flex;
    justify-content: space-between; /* 将内容分散对齐 */
    align-items: center; /* 垂直居中对齐 */
    box-shadow:0 0 0 2px rgb(255, 255, 255);
  }
  .list:hover{
    transform: scale(1.02);
    transition: transform 0.2s; /* 添加过渡效果以使动画更平滑 */
    cursor: pointer;
  }
  @font-face {
    font-family: ggbond;
    src: url('../../assets/MFBoHeHaiYan_Noncommercial-Regular.otf')

  }
</style>
