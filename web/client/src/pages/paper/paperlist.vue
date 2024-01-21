<template>
  <div style="width: 1500px;z-index: 1;position: absolute;left: 6%">
    <ul style="font-size: 35px;color: white;font-family: ggbond">
      <li v-for="backendArticles in paginatedArticles" :key="backendArticles.article_id"
          class="list" @click="goview(backendArticles.article_id)">
        <span>{{ backendArticles.title }}</span>
        <div>
        <span style="font-size: 20px;margin-top: 15px;margin-right: 70px;color:#7fb80e ">{{ backendArticles.time }}</span>
          <el-avatar :size="50" style="position: absolute;right: 1%"
              :src=backendArticles.user.icon_rul></el-avatar>
        </div>
      </li>
    </ul>
    <el-pagination
        v-if="totalArticles > 0"
        :current-page="currentPage"
        :page-sizes="[10, 20, 30, 40]"
        :page-size="pageSize"
        :total="totalArticles"
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

export default {
  computed: {
    totalArticles() {
      return this.backendArticles.length;
    },
    paginatedArticles() {
      const startIndex = (this.currentPage - 1) * this.pageSize;
      const endIndex = startIndex + this.pageSize;
      return this.backendArticles.slice(startIndex, endIndex);
    },
  },
  data() {
    return {
      pageSize: 6,
      currentPage: 1,
      backendArticles: [      { article_id: 1, title: 'Vue 3 教程', content: '这是Vue 3的入门教程。',time:'2023.11.4' ,user:{icon_rul:"https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0"} },
        { article_id: 2, title: 'Element Plus 使用指南', content: '学习如何使用Element Plus构建漂亮的界面。' ,time:'2023.11.4',user:{icon_rul:"https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0"} },
        { article_id: 3, title: 'Vue Router 详解', content: '深入了解Vue Router的使用和配置。' ,time:'2023.11.4',user:{icon_rul:"https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0"} },
        { article_id: 3, title: 'Vue Router 详解', content: '深入了解Vue Router的使用和配置。' ,time:'2023.11.4',user:{icon_rul:"https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0"} },
        { article_id: 4, title: 'Vue Router 详解', content: '深入了解Vue Router的使用和配置。' ,time:'2023.11.4',user:{icon_rul:"https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0"} },
        { article_id: 5, title: 'Vue Router 详解', content: '深入了解Vue Router的使用和配置。' ,time:'2023.11.4',user:{icon_rul:"https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0"} },
        { article_id: 6, title: 'Vue Router 详解', content: '深入了解Vue Router的使用和配置。' ,time:'2023.11.4',user:{icon_rul:"https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0"} },
        { article_id: 7, title: 'Vue Router 详解', content: '深入了解Vue Router的使用和配置。' ,time:'2023.11.4',user:{icon_rul:"https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0"} },
        { article_id: 8, title: 'Vue Router 详解', content: '深入了解Vue Router的使用和配置。' ,time:'2023.11.4',user:{icon_rul:"https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0"} },
        { article_id: 9, title: 'Vue Router 详解', content: '深入了解Vue Router的使用和配置。' ,time:'2023.11.4',user:{icon_rul:"https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0"} },
        { article_id: 10, title: 'Vue Router 详解', content: '深入了解Vue Router的使用和配置。' ,time:'2023.11.4',user:{icon_rul:"https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0"} },
        { article_id: 11, title: 'Vue Router 详解', content: '深入了解Vue Router的使用和配置。' ,time:'2023.11.4',user:{icon_rul:"https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0"} },
        { article_id: 12, title: 'Vue Router 详解', content: '深入了解Vue Router的使用和配置。' ,time:'2023.11.4',user:{icon_rul:"https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0"} },
        { article_id: 13, title: 'Vue Router 详解', content: '深入了解Vue Router的使用和配置。' ,time:'2023.11.4',user:{icon_rul:"https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0"} },
        { article_id: 14, title: 'Vue Router 详解', content: '深入了解Vue Router的使用和配置。' ,time:'2023.11.4',user:{icon_rul:"https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0"} },

      ], // 从后端获取这个数组
    };
  },
  methods: {
    handleSizeChange(val) {
      this.pageSize = val;
      this.currentPage = 1;
    },
    handlePageChange(val) {
      this.currentPage = val;
    },
    goview(id){
      this.$router.push({name:"paper-display",query:{article_id:id}})
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
