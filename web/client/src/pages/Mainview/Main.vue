  <script  lang="ts">
  import $ from 'jquery';
  import {Lock ,Message,Notebook,Comment} from "@element-plus/icons-vue";
  import {Getcards} from "@/apis/getcards";
  import {Likeapi} from "@/apis/likeaction";
  import {guestbookApi} from "@/apis/guestaction";
  import {useUserInfoStore} from "@/stores/userinfo";
  import router from "@/router";
  import {ElMessage} from "element-plus";
  import { ref, defineProps } from 'vue';
  export interface Admin {
    user_id: number;
    image_url: string;
    signature: string;
    icon_url: string;
    title: string;
    user_name:string;
    git_url:string;
    guestbooks:Guestbook[];
  }
  export interface Guestbook {
    admin_id: number;
    book_id: number;
    context: string;
    avatar_url:string
    create_time: string;
    user_id: number;
    user_name:string;
  }

  export  interface Guestbookfrom{
    user_id:number;
    admin_id:number;
    context:string;
  }

  export default {
    setup(props, context) {
      const guestbookfrom = ref({
        user_id: Number(useUserInfoStore().getUserId()),
        admin_id: 1,
        context: 'cin',
      });
      const whoguest="admin"
      return {
        whoguest,
        guestbookfrom,
      };
    },
    computed: {
      Notebook() {
        return Notebook;
      },
      Comment() {
        return Comment;
      },
    },
    methods: {
      hideButtons(id: number) {
        var type=false;
        $('.hide').each(function () {
          if (this.id.startsWith(id + '-'+'1')){
            if(this.style.display!='none')type=true;
          }
        });
        $('.hide').each(function () {
          if (this.id.startsWith(id + '-')) {

            if (type) {
              $(this).fadeOut('slow');
            } else $(this).fadeIn('slow');
          }
        })
      },
      hitarrow(){
        // 获取目标元素
        var targetElement = document.querySelector('.card');
        // 检查目标元素是否存在
        if (targetElement) {
          // 获取目标元素的位置信息
          var targetPosition = targetElement.getBoundingClientRect().top;
          // 调整目标位置（减去 100px)
          var adjustedPosition = targetPosition - 150;
          // 使用 scrollIntoView 方法滚动到目标元素
          window.scrollTo({
            top: adjustedPosition,
            behavior: 'smooth',
            block: 'start',
            inline: 'nearest'
          });
        }
      },
      tableview(id:number){
        if (id == 1) {
          console.log(this.table1);
          this.table1 = !this.table1;
        } else if (id == 2) {
          this.table2 = !this.table2;
        } else if (id == 3) {
         this.table3 = !this.table3;
        }
      },
       async likebut(adimin_id:number){
        const user_id=Number(useUserInfoStore().getUserId());
        if(!user_id){
          router.push("/login");
        return;
        }
        const form={
          admin_id:adimin_id,
          user_id:user_id,
        }
        const res=await Likeapi(form)
       if(res.status_code===0)  ElMessage.success("点赞成功")
         return;
      },
      gogit(admin_id:number){
        window.location.href=this.admins[admin_id].git_url;
      },
      async guestbookbut(adimin_id:number){
        const user_id=Number(useUserInfoStore().getUserId());
        if(!user_id){
          router.push("/login");
          return;
        }
        const form={
          admin_id:adimin_id,
          user_id:user_id,
        }
        const res=await guestbookApi(this.guestbookfrom);
        if(res.status_code==0)ElMessage.success("留言成功")
        return;
      },
      async getcard() {
         const res=await Getcards();
         let p:Admin[];
         if(res){
           p=res.users;
         }
        return p as Admin[];
      },
    },
      data() {
       return {
         admins: [] as Admin[],
         table1:false,
         table2:false,
         table3:false,
         dialogVisible:false,

      };
    },
     async created() {
       this.admins = await this.getcard();
       return {admins: [] as Admin[],

       }
     },
  };


  </script>

  <template>
    <div v-if="admins.length">
      <header class="header">
        <div class="header-text">
        <h1 class="heading-primary">

            <span class="heading-primary-main">菜 鸟 营</span>
            <span class="header-primary-sub">Rookiable Coven</span>
        </h1>
        </div>
        <div class="header-arrow" @click="hitarrow"> </div>
      </header>
      <el-dialog
          v-model="dialogVisible"
          :title="'cin>>'+whoguest+'>>endl;'"
          width="30%"
          style="background-color:rgb(258,258,258)"
      >
        <el-input  v-model="guestbookfrom.context"></el-input>
        <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="dialogVisible = false;guestbookbut()">
          Confirm
        </el-button>
      </span>
        </template>
      </el-dialog>
    <div class="card" :style="{ 'background-image': 'url('+admins[0].image_url+')' }" >
      <div class="xbox">
      <el-image :src="admins[0].icon_url" fit="fill" :lazy="true" style="bottom: -36px" >
      </el-image>
      <div class="name-box" >{{admins[0].user_name}}</div>
        <div class="good-box">
          <svg  width="45" height="45" viewBox="0 0 20 21" fill="none" ><path fill-rule="evenodd" clip-rule="evenodd" d="M8.95 2.563l-1.726 3.98-1.158.008v13.748h-.648l6.462.047c3.079-.297 5.037-1.813 5.807-4.473.805-3.99 1.317-6.347 1.538-7.075.366-1.207-.032-2.178-1.235-2.178h-4.699l-.03-3.273c0-1.552-.77-2.358-2.207-2.358-.98 0-1.714.574-2.105 1.574zM4.765 20.294V6.56l-2.14.015a1.8 1.8 0 00-1.788 1.8v10.105a1.8 1.8 0 001.787 1.8l2.14.015z" fill="#FF6880"/></svg>
        <p>114514</p>
        </div>
        <div class="good-box">
          <svg width="45" height="45" viewBox="0 0 19 18" fill="none" :style="{cursor:'pointer'}" id="0" @click="tableview(1)"><path clip-rule="evenodd" d="M.733 2.8a2 2 0 012-2h13.2a2 2 0 012 2v6.015a6 6 0 01-6 6H5.534a4 4 0 00-2.189.652L1.507 16.67a.5.5 0 01-.774-.418V2.8z" stroke="#333" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/><path d="M4.5 5.62h5" stroke="#333" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
          <p> 114514</p>
        </div>
        </div>
      <div class="xbox">
        <div class="but-box">
        <div class="desi-box">{{admins[0].title}}</div>
          <el-button type="primary"  @click="hideButtons(1)" class="show" color="#2e5496">撅一下</el-button>
        </div>
      <div class="sing-box">{{admins[0].signature}}</div>
      </div>
      <div class="xbox" style="width: 400px">
        <el-button type="primary" class="hide" id="1-1" style="display: none" color="#2e5496" @click="gogit(0)"><svg t="1705830951101" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5500" width="32" height="32"><path d="M347.8 794.8c0 4-4.6 7.2-10.4 7.2-6.6 0.6-11.2-2.6-11.2-7.2 0-4 4.6-7.2 10.4-7.2 6-0.6 11.2 2.6 11.2 7.2z m-62.2-9c-1.4 4 2.6 8.6 8.6 9.8 5.2 2 11.2 0 12.4-4s-2.6-8.6-8.6-10.4c-5.2-1.4-11 0.6-12.4 4.6z m88.4-3.4c-5.8 1.4-9.8 5.2-9.2 9.8 0.6 4 5.8 6.6 11.8 5.2 5.8-1.4 9.8-5.2 9.2-9.2-0.6-3.8-6-6.4-11.8-5.8zM505.6 16C228.2 16 16 226.6 16 504c0 221.8 139.6 411.6 339 478.4 25.6 4.6 34.6-11.2 34.6-24.2 0-12.4-0.6-80.8-0.6-122.8 0 0-140 30-169.4-59.6 0 0-22.8-58.2-55.6-73.2 0 0-45.8-31.4 3.2-30.8 0 0 49.8 4 77.2 51.6 43.8 77.2 117.2 55 145.8 41.8 4.6-32 17.6-54.2 32-67.4-111.8-12.4-224.6-28.6-224.6-221 0-55 15.2-82.6 47.2-117.8-5.2-13-22.2-66.6 5.2-135.8 41.8-13 138 54 138 54 40-11.2 83-17 125.6-17s85.6 5.8 125.6 17c0 0 96.2-67.2 138-54 27.4 69.4 10.4 122.8 5.2 135.8 32 35.4 51.6 63 51.6 117.8 0 193-117.8 208.4-229.6 221 18.4 15.8 34 45.8 34 92.8 0 67.4-0.6 150.8-0.6 167.2 0 13 9.2 28.8 34.6 24.2C872.4 915.6 1008 725.8 1008 504 1008 226.6 783 16 505.6 16zM210.4 705.8c-2.6 2-2 6.6 1.4 10.4 3.2 3.2 7.8 4.6 10.4 2 2.6-2 2-6.6-1.4-10.4-3.2-3.2-7.8-4.6-10.4-2z m-21.6-16.2c-1.4 2.6 0.6 5.8 4.6 7.8 3.2 2 7.2 1.4 8.6-1.4 1.4-2.6-0.6-5.8-4.6-7.8-4-1.2-7.2-0.6-8.6 1.4z m64.8 71.2c-3.2 2.6-2 8.6 2.6 12.4 4.6 4.6 10.4 5.2 13 2 2.6-2.6 1.4-8.6-2.6-12.4-4.4-4.6-10.4-5.2-13-2z m-22.8-29.4c-3.2 2-3.2 7.2 0 11.8 3.2 4.6 8.6 6.6 11.2 4.6 3.2-2.6 3.2-7.8 0-12.4-2.8-4.6-8-6.6-11.2-4z" p-id="5501"></path></svg></el-button>
        <el-button type="primary" class="hide" id="1-2" style="display: none" :icon="Notebook" color="#2e5496"></el-button>
        <el-button type="primary" class="hide" id="1-3" style="display: none" :icon="Comment" color="#2e5496" @click="dialogVisible=true;whoguest=admins[0].title;guestbookfrom.user_id=admins[0].user_id"></el-button>
        <el-button type="primary" class="hide" id="1-4" style="display: none" color="#2e5496" @click="likebut(admins[0].user_id)"><svg width="24" height="25.2" viewBox="0 0 20 21" fill="none">
          <path fill-rule="evenodd" clip-rule="evenodd" d="M8.95 2.563l-1.726 3.98-1.158.008v13.748h-.648l6.462.047c3.079-.297 5.037-1.813 5.807-4.473.805-3.99 1.317-6.347 1.538-7.075.366-1.207-.032-2.178-1.235-2.178h-4.699l-.03-3.273c0-1.552-.77-2.358-2.207-2.358-.98 0-1.714.574-2.105 1.574zM4.765 20.294V6.56l-2.14.015a1.8 1.8 0 00-1.788 1.8v10.105a1.8 1.8 0 001.787 1.8l2.14.015z" fill="#FF6880"/>
        </svg></el-button>
      </div>
    </div>
      <el-table
          v-if="table1"
          :data="admins[0].guestbooks"
          style="width: 860px; margin: 0 auto; display: flex; flex-direction: row;margin-left: -35px;
          border: 1px solid white;
      box-shadow : 6px 6px 3px rgba(0, 0, 4, 0.3);
      border-radius: 10px;
                   "
          :show-header="false"
          id="gb-1"
          class="table"
          max-height="250"
      >
        <el-table-column label="头像" prop="avatar">
          <template #default="scope">
            <div style="
        display: flex;
        margin-left: 20px;
        flex-direction: column;
        width: auto;
        height: auto;

      ">
              <div style="display: flex;width: auto;height: auto;" >
              <el-image :src="scope.row.avatar_url" fit="fill" :lazy="true" style="
          width: 50px;
          height: 50px;
          border-radius: 50%;
          box-shadow: 0 0 4px 2px;
          overflow: hidden;
          margin-top: 10px;
          margin-right:10px;
          margin-bottom: 10px;
" >
              </el-image>
              <div class="name" style="font-size: 18px; width: auto; white-space: normal;display: flex;flex-direction: row">{{ scope.row.user_name }}
                <br>
                {{scope.row.create_time}}
              </div>
              </div>
              <div style="margin-left: 60px; width: auto; white-space: normal;font-size: 20px;margin-top: -20px">{{ scope.row.context }}</div>
            </div>
          </template>
        </el-table-column>
      </el-table>



      <div class="card" :style="{ 'background-image': 'url('+admins[1].image_url+')' }">   <div class="xbox">
      <el-image :src="admins[1].icon_url" fit="fill" :lazy="true" style="bottom: -36px" >
      </el-image>
      <div class="name-box">{{admins[1].user_name}}</div>
      <div class="good-box">
        <svg  width="45" height="45" viewBox="0 0 20 21" fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M8.95 2.563l-1.726 3.98-1.158.008v13.748h-.648l6.462.047c3.079-.297 5.037-1.813 5.807-4.473.805-3.99 1.317-6.347 1.538-7.075.366-1.207-.032-2.178-1.235-2.178h-4.699l-.03-3.273c0-1.552-.77-2.358-2.207-2.358-.98 0-1.714.574-2.105 1.574zM4.765 20.294V6.56l-2.14.015a1.8 1.8 0 00-1.788 1.8v10.105a1.8 1.8 0 001.787 1.8l2.14.015z" fill="#FF6880"/></svg>
        <p>114514</p>
      </div>
      <div class="good-box">
        <svg width="45" height="45" viewBox="0 0 19 18" fill="none" :style="{cursor:'pointer'}" id="1" @click="tableview(2)"><path clip-rule="evenodd" d="M.733 2.8a2 2 0 012-2h13.2a2 2 0 012 2v6.015a6 6 0 01-6 6H5.534a4 4 0 00-2.189.652L1.507 16.67a.5.5 0 01-.774-.418V2.8z" stroke="#333" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/><path d="M4.5 5.62h5" stroke="#333" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
        <p> 114514</p>
      </div>
    </div>
      <div class="xbox">
        <div class="but-box">
          <div class="desi-box">{{admins[1].title}}</div>
          <el-button type="primary"  @click="hideButtons(2)" class="show" color="#5f962e">撅一下</el-button>
        </div>
        <div class="sing-box">{{admins[1].signature}}</div>
      </div>
      <div class="xbox">
        <el-button type="primary" class="hide" id="2-1" style="display: none" color="#5f962e" @click="gogit(1)"><svg t="1705830951101" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5500" width="32" height="32"><path d="M347.8 794.8c0 4-4.6 7.2-10.4 7.2-6.6 0.6-11.2-2.6-11.2-7.2 0-4 4.6-7.2 10.4-7.2 6-0.6 11.2 2.6 11.2 7.2z m-62.2-9c-1.4 4 2.6 8.6 8.6 9.8 5.2 2 11.2 0 12.4-4s-2.6-8.6-8.6-10.4c-5.2-1.4-11 0.6-12.4 4.6z m88.4-3.4c-5.8 1.4-9.8 5.2-9.2 9.8 0.6 4 5.8 6.6 11.8 5.2 5.8-1.4 9.8-5.2 9.2-9.2-0.6-3.8-6-6.4-11.8-5.8zM505.6 16C228.2 16 16 226.6 16 504c0 221.8 139.6 411.6 339 478.4 25.6 4.6 34.6-11.2 34.6-24.2 0-12.4-0.6-80.8-0.6-122.8 0 0-140 30-169.4-59.6 0 0-22.8-58.2-55.6-73.2 0 0-45.8-31.4 3.2-30.8 0 0 49.8 4 77.2 51.6 43.8 77.2 117.2 55 145.8 41.8 4.6-32 17.6-54.2 32-67.4-111.8-12.4-224.6-28.6-224.6-221 0-55 15.2-82.6 47.2-117.8-5.2-13-22.2-66.6 5.2-135.8 41.8-13 138 54 138 54 40-11.2 83-17 125.6-17s85.6 5.8 125.6 17c0 0 96.2-67.2 138-54 27.4 69.4 10.4 122.8 5.2 135.8 32 35.4 51.6 63 51.6 117.8 0 193-117.8 208.4-229.6 221 18.4 15.8 34 45.8 34 92.8 0 67.4-0.6 150.8-0.6 167.2 0 13 9.2 28.8 34.6 24.2C872.4 915.6 1008 725.8 1008 504 1008 226.6 783 16 505.6 16zM210.4 705.8c-2.6 2-2 6.6 1.4 10.4 3.2 3.2 7.8 4.6 10.4 2 2.6-2 2-6.6-1.4-10.4-3.2-3.2-7.8-4.6-10.4-2z m-21.6-16.2c-1.4 2.6 0.6 5.8 4.6 7.8 3.2 2 7.2 1.4 8.6-1.4 1.4-2.6-0.6-5.8-4.6-7.8-4-1.2-7.2-0.6-8.6 1.4z m64.8 71.2c-3.2 2.6-2 8.6 2.6 12.4 4.6 4.6 10.4 5.2 13 2 2.6-2.6 1.4-8.6-2.6-12.4-4.4-4.6-10.4-5.2-13-2z m-22.8-29.4c-3.2 2-3.2 7.2 0 11.8 3.2 4.6 8.6 6.6 11.2 4.6 3.2-2.6 3.2-7.8 0-12.4-2.8-4.6-8-6.6-11.2-4z" p-id="5501"></path></svg></el-button>
        <el-button type="primary" class="hide" id="2-2" style="display: none" :icon="Notebook" color="#5f962e"></el-button>
        <el-button type="primary" class="hide" id="2-3" style="display: none" :icon="Comment" color="#5f962e" @click="dialogVisible=true;whoguest=admins[1].title;guestbookfrom.user_id=admins[1].user_id"></el-button>
        <el-button type="primary" class="hide" id="2-4" style="display: none" color="#5f962e" @click="likebut(admins[1].user_id)"><svg width="24" height="25.2" viewBox="0 0 20 21" fill="none">
          <path fill-rule="evenodd" clip-rule="evenodd" d="M8.95 2.563l-1.726 3.98-1.158.008v13.748h-.648l6.462.047c3.079-.297 5.037-1.813 5.807-4.473.805-3.99 1.317-6.347 1.538-7.075.366-1.207-.032-2.178-1.235-2.178h-4.699l-.03-3.273c0-1.552-.77-2.358-2.207-2.358-.98 0-1.714.574-2.105 1.574zM4.765 20.294V6.56l-2.14.015a1.8 1.8 0 00-1.788 1.8v10.105a1.8 1.8 0 001.787 1.8l2.14.015z" fill="#FF6880"/>
        </svg></el-button>

      </div></div>
      <el-table
          v-if="table2"
          :data="admins[1].guestbooks"
          style="width: 860px; margin: 0 auto; display: flex; flex-direction: row;margin-left: -35px;
          border: 1px solid white;
      box-shadow : 6px 6px 3px rgba(0, 0, 4, 0.3);
      border-radius: 10px;
                   "
          :show-header="false"
          id="gb-1"
          class="table"
          max-height="250"
      >
        <el-table-column label="头像" prop="avatar">
          <template #default="scope">
            <div style="
        display: flex;
        margin-left: 20px;
        flex-direction: column;
        width: auto;
        height: auto;

      ">
              <div style="display: flex;width: auto;height: auto;" >
                <el-image :src="scope.row.avatar_url" fit="fill" :lazy="true" style="
          width: 50px;
          height: 50px;
          border-radius: 50%;
          box-shadow: 0 0 4px 2px;
          overflow: hidden;
          margin-top: 10px;
          margin-right:10px;
          margin-bottom: 10px;
" >
                </el-image>
                <div class="name" style="font-size: 18px; width: auto; white-space: normal;display: flex;flex-direction: row">{{ scope.row.user_name }}
                  <br>
                  {{scope.row.create_time}}
                </div>
              </div>
              <div style="margin-left: 60px; width: auto; white-space: normal;font-size: 20px;margin-top: -20px">{{ scope.row.context }}</div>
            </div>
          </template>
        </el-table-column>
      </el-table>


      <div class="card" :style="{ 'background-image': 'url('+admins[2].image_url+')' }">   <div class="xbox">
        <el-image :src="admins[2].icon_url" fit="fill" :lazy="true" style="bottom: -36px"  >
        </el-image>
        <div class="name-box">{{admins[2].user_name}}</div>
        <div class="good-box">
          <svg  width="45" height="45" viewBox="0 0 20 21" fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M8.95 2.563l-1.726 3.98-1.158.008v13.748h-.648l6.462.047c3.079-.297 5.037-1.813 5.807-4.473.805-3.99 1.317-6.347 1.538-7.075.366-1.207-.032-2.178-1.235-2.178h-4.699l-.03-3.273c0-1.552-.77-2.358-2.207-2.358-.98 0-1.714.574-2.105 1.574zM4.765 20.294V6.56l-2.14.015a1.8 1.8 0 00-1.788 1.8v10.105a1.8 1.8 0 001.787 1.8l2.14.015z" fill="#FF6880"/></svg>
          <p>114514</p>
        </div>
        <div class="good-box">
          <svg width="45" height="45" viewBox="0 0 19 18" fill="none" :style="{cursor:'pointer'}" id="2" @click="tableview(3)"><path clip-rule="evenodd" d="M.733 2.8a2 2 0 012-2h13.2a2 2 0 012 2v6.015a6 6 0 01-6 6H5.534a4 4 0 00-2.189.652L1.507 16.67a.5.5 0 01-.774-.418V2.8z" stroke="#333" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/><path d="M4.5 5.62h5" stroke="#333" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
          <p> 114514</p>
        </div>
      </div>
        <div class="xbox">
          <div class="but-box">
            <div class="desi-box">{{admins[2].title}}</div>
            <el-button type="primary"  @click="hideButtons(3)" class="show" color="#ff661a">撅一下</el-button>
          </div>
          <div class="sing-box">{{admins[2].signature}}</div>
        </div>
        <div class="xbox">
          <el-button type="primary" class="hide" id="3-1" style="display: none" color="#ff661a" @click="gogit(2)" ><svg t="1705830951101" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5500" width="32" height="32"><path d="M347.8 794.8c0 4-4.6 7.2-10.4 7.2-6.6 0.6-11.2-2.6-11.2-7.2 0-4 4.6-7.2 10.4-7.2 6-0.6 11.2 2.6 11.2 7.2z m-62.2-9c-1.4 4 2.6 8.6 8.6 9.8 5.2 2 11.2 0 12.4-4s-2.6-8.6-8.6-10.4c-5.2-1.4-11 0.6-12.4 4.6z m88.4-3.4c-5.8 1.4-9.8 5.2-9.2 9.8 0.6 4 5.8 6.6 11.8 5.2 5.8-1.4 9.8-5.2 9.2-9.2-0.6-3.8-6-6.4-11.8-5.8zM505.6 16C228.2 16 16 226.6 16 504c0 221.8 139.6 411.6 339 478.4 25.6 4.6 34.6-11.2 34.6-24.2 0-12.4-0.6-80.8-0.6-122.8 0 0-140 30-169.4-59.6 0 0-22.8-58.2-55.6-73.2 0 0-45.8-31.4 3.2-30.8 0 0 49.8 4 77.2 51.6 43.8 77.2 117.2 55 145.8 41.8 4.6-32 17.6-54.2 32-67.4-111.8-12.4-224.6-28.6-224.6-221 0-55 15.2-82.6 47.2-117.8-5.2-13-22.2-66.6 5.2-135.8 41.8-13 138 54 138 54 40-11.2 83-17 125.6-17s85.6 5.8 125.6 17c0 0 96.2-67.2 138-54 27.4 69.4 10.4 122.8 5.2 135.8 32 35.4 51.6 63 51.6 117.8 0 193-117.8 208.4-229.6 221 18.4 15.8 34 45.8 34 92.8 0 67.4-0.6 150.8-0.6 167.2 0 13 9.2 28.8 34.6 24.2C872.4 915.6 1008 725.8 1008 504 1008 226.6 783 16 505.6 16zM210.4 705.8c-2.6 2-2 6.6 1.4 10.4 3.2 3.2 7.8 4.6 10.4 2 2.6-2 2-6.6-1.4-10.4-3.2-3.2-7.8-4.6-10.4-2z m-21.6-16.2c-1.4 2.6 0.6 5.8 4.6 7.8 3.2 2 7.2 1.4 8.6-1.4 1.4-2.6-0.6-5.8-4.6-7.8-4-1.2-7.2-0.6-8.6 1.4z m64.8 71.2c-3.2 2.6-2 8.6 2.6 12.4 4.6 4.6 10.4 5.2 13 2 2.6-2.6 1.4-8.6-2.6-12.4-4.4-4.6-10.4-5.2-13-2z m-22.8-29.4c-3.2 2-3.2 7.2 0 11.8 3.2 4.6 8.6 6.6 11.2 4.6 3.2-2.6 3.2-7.8 0-12.4-2.8-4.6-8-6.6-11.2-4z" p-id="5501"></path></svg></el-button>
          <el-button type="primary" class="hide" id="3-2" style="display: none" :icon="Notebook" color="#ff661a"></el-button>
          <el-button type="primary" class="hide" id="3-3" style="display: none" :icon="Comment" color="#ff661a" @click="dialogVisible=true;whoguest=admins[2].title;guestbookfrom.user_id=admins[2].user_id"></el-button>
          <el-button type="primary" class="hide" id="3-4" style="display: none" color="#ff661a" @click="likebut(admins[2].user_id)"><svg width="24" height="25.2" viewBox="0 0 20 21" fill="none">
            <path fill-rule="evenodd" clip-rule="evenodd" d="M8.95 2.563l-1.726 3.98-1.158.008v13.748h-.648l6.462.047c3.079-.297 5.037-1.813 5.807-4.473.805-3.99 1.317-6.347 1.538-7.075.366-1.207-.032-2.178-1.235-2.178h-4.699l-.03-3.273c0-1.552-.77-2.358-2.207-2.358-.98 0-1.714.574-2.105 1.574zM4.765 20.294V6.56l-2.14.015a1.8 1.8 0 00-1.788 1.8v10.105a1.8 1.8 0 001.787 1.8l2.14.015z" fill="#FF6880"/>
          </svg></el-button>

        </div></div>
      <el-table
          v-if="table3"
          :data="admins[2].guestbooks"
          style="width: 860px; margin: 0 auto; display: flex; flex-direction: row;margin-left: -35px;
          border: 1px solid white;
      box-shadow : 6px 6px 3px rgba(0, 0, 4, 0.3);
      border-radius: 10px;
                   "
          :show-header="false"
          id="gb-1"
          class="table"
          max-height="250"
      >
        <el-table-column label="头像" prop="avatar">
          <template #default="scope">
            <div style="
        display: flex;
        margin-left: 20px;
        flex-direction: column;
        width: auto;
        height: auto;

      ">
              <div style="display: flex;width: auto;height: auto;" >
                <el-image :src="scope.row.avatar_url" fit="fill" :lazy="true" style="
          width: 50px;
          height: 50px;
          border-radius: 50%;
          box-shadow: 0 0 4px 2px;
          overflow: hidden;
          margin-top: 10px;
          margin-right:10px;
          margin-bottom: 10px;
" >
                </el-image>
                <div class="name" style="font-size: 18px; width: auto; white-space: normal;display: flex;flex-direction: row">{{ scope.row.user_name }}
                  <br>
                  {{scope.row.create_time}}
                </div>
              </div>
              <div style="margin-left: 60px; width: auto; white-space: normal;font-size: 20px;margin-top: -20px">{{ scope.row.context }}</div>
            </div>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </template>

  <style scoped lang="scss">
.header{
  height: 800px;
  width: 100%;

}
.header-text{
  position: absolute;
  top: 40%;
  left: 50%;
  transform: translate(-50%,-50%);
}
.heading-primary{
  color:#ffffff;
  text-transform: uppercase;
  font-family: Monocraft;
  backface-visibility: hidden;//解决动画异常摇晃问题
}
.heading-primary-main{
  display: block;
  font-size: 80px;
  font-weight: 400;
  letter-spacing: 25px;
  margin-left: 22.5%;
  animation-name: moveInleft;
  animation-duration: 1s;
  animation-timing-function: ease-out;
  //animation-iteration-count: 3;播放几次
  //animation-delay: 1s;延迟
}
.header-primary-sub{
  display: block;
  font-size: 60px;
  font-weight: 400;
  letter-spacing: 15px;
  animation-name: moveInright;
  animation-duration: 1s;
  animation-timing-function: ease-out;
}
@keyframes moveInleft {
  0%{
    opacity: 0;
    transform: translateX(-100px);
  }
  80%{
    transform: translateX(15px);
  }
  100%{
    opacity:1;
    transform: translate(0);
  }
}
@keyframes moveInright {
  0%{
    opacity: 0;
    transform: translateX(100px);
  }
  80%{
    transform: translateX(-15px);
  }
  100%{
    opacity:1;
    transform: translate(0);
  }
}
.header-arrow {
  position: absolute;
  height: 100px;
  width: 100px;
  top: 80%;
  left: 45%;
  clip-path: polygon(0 50%, 35% 50%, 35% 0, 66% 0, 66% 50%, 100% 51%, 50% 100%);
  background-color: white;
  animation-name: arrowblink;
  animation-duration: 3s;
  animation-timing-function: ease-out;
  animation-iteration-count: infinite; /* Set to 'infinite' for continuous loop */
}

@keyframes arrowblink {
  0%, 100% {
    opacity: 0;
    transform: translateY(-15px);
  }

  50% {
    opacity: 1;
    transform: translateY(0);
  }
}


  .overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.5); /* 半透明黑色背景 */
    z-index: 1;
    opacity: 0; /* 默认情况下不可见 */
    pointer-events: none; /* 不可交互 */
    transition: opacity 0.3s;
  }
  .overlay.active {
    opacity: 1; /* 遮罩层激活时可见 */
    pointer-events: auto; /* 可交互 */
  }
  @font-face {
    font-family: ggbond;
    src: url('../../assets/MFBoHeHaiYan_Noncommercial-Regular.otf')

  }
@font-face {
  font-family: Monocraft;
  src: url("../../assets/Monocraft.otf");
}
    .card {
      display: flex;
      flex-direction: row;
      align-items: center;
      top: 100px;
      width: 860px;
      height: auto;
      border: 1px solid white;
      box-shadow : 6px 6px 3px rgba(0, 0, 4, 0.3);
      border-radius: 10px;
      min-height: 500px;
      margin-bottom: 150px;
      bottom: 100px;
      position: relative;
     // background-color:transparent;
      background-size: cover;
      z-index: 0;
      margin-left: -50px;
      .good-box{
        width: 150px;
        margin-top: 30px;
        margin-left: 20px;
        align-items: flex-start;
        display: flex;
        flex-direction: row;
        font-size: 48px;
        word-wrap: normal;
        font-family: ggbond;
        color: #7bff91;
      }
      .xbox {
        width: auto;
        height: auto;
        flex-wrap: nowrap;
        .el-image {
          width: 130px;
          height: 130px;
          border-radius: 50%;
          box-shadow: 0 0 4px 2px;
          overflow: hidden;
          margin: 50px 30px;
          margin-top: 0px;
          flex-direction: row;

        }
        .name-box {
          font-family: ggbond;
          word-wrap: break-word;//文本自动换行
          width: 180px;
          text-align: center;
          line-height: 60px;
          font-size: 55px;
          font-weight: bold;
          color: #2fd0be;
        }
        .but-box{
          width: auto;
          height: auto;
          display: flex;
          flex-direction: row;
        }
          .el-button{
            margin-top: 50px;
            margin-left: 150px;
            width: 135px;
            height: 55px;
            border-radius: 20px;
            font-size: 30px;
            font-family: ggbond;
            box-shadow:
                inset 0 -3em 3em rgba(0, 0, 0, 0.1),
                0 0 0 2px rgb(255, 255, 255),
                0.3em 0.3em 1em rgba(0, 0, 0, 0.3);
            transition: all 0.2s;
          }
        @keyframes shake {
          0% {
            transform: translateX(0);
          }
          25% {
            transform: translateX(-5px);
          }
          50% {
            transform: translateX(5px);
          }
          75% {
            transform: translateX(-5px);
          }
          100% {
            transform: translateX(0);
          }
        }
        .el-button.hide:hover{
          transform: scale(1.1); /* 设置按钮获得焦点时的缩放效果，可以根据需要调整 */
          transition: transform 0.2s; /* 添加过渡效果以使动画更平滑 */
        }
        .el-button.show:hover{

          animation: shake 0.5s;
          background-color: #ffd04b;
          //box-shadow: 0 10px 20px ;
        }
        .el-button.show:active{
            transform: translateY(-2px);
        }
        .el-button.hide{
          width: 65px;
          height: 65px;
          border-radius: 25px;
          margin-left: 100px;

        }
        .desi-box{
          background-color: white;
          height: 65px;
          width:300px ;
          margin-top: 50px;
          border-radius: 20px;
          background-color: rgb(153, 204, 255,0.6);
          box-shadow:
              inset 0 -3em 3em rgba(0, 0, 0, 0.1),
              0 0 0 2px rgb(255, 255, 255),
              0.3em 0.3em 1em rgba(0, 0, 0, 0.3);
          text-align: center;
          font-size: 50px;
          font-family: ggbond;
        }

        .sing-box {
          font-family: ggbond;
          margin-top: 50px ;
          width: 600px;
          height: 250px;
          border-radius: 20px;
          background-color: rgba(255,255,255,0.5);
          box-shadow:
              inset 0 -3em 3em rgba(0, 0, 0, 0.1),
              0 0 0 2px rgb(255, 255, 255),
              0.3em 0.3em 1em rgba(0, 0, 0, 0.3);
          font-size: 40px;
        }

        .el-image img {
          width: 100%;
          height: 100%;
          object-fit: cover;
        }
      }
    }
  </style>