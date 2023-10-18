 <template>
    <div class="home">
      <el-container>
        <el-header>
          <HomeHeader></HomeHeader>
        </el-header>
        <el-main><RouterView></RouterView></el-main>
        <el-footer>Footer</el-footer>
      </el-container>
    </div>
  </template>
 <script setup lang="ts">
 import HomeHeader from "@/components/layout/HomeHeader.vue";
 import {useUserInfoStore} from "@/stores/userinfo";
 let loginstate=useUserInfoStore().authFromLocal()
 </script>
<style lang="scss" scoped>
 .home {
   min-height: 100vh;
   background-size: 100% auto;
   background-image: url("https://ts1.cn.mm.bing.net/th/id/R-C.7c4dae09992d663a81de74c71e402bad?rik=ov3oBD%2fpUEVSqA&riu=http%3a%2f%2fpic1.bbzhi.com%2ffengjingbizhi%2fshijielvyoumingshengzhilvouzhoupian%2fnature_window7_europe_wallpapers_1920x1200_43316_16.jpg&ehk=AHv7P57icUvWXKFOIJkB2poG6WgYhCXFSv6vYT1o8NQ%3d&risl=&pid=ImgRaw&r=0");
   .el-header {
     background-color: transparent;
     padding: 0;
   }

   .el-footer {
     background-color: darkgrey;
     padding: 0;
   }

   .el-main {
     height: 1500px;
     display: flex;
     flex-direction: column;
     width: 100%;
   }
 }
</style>