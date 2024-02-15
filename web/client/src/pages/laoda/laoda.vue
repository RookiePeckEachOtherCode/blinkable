<template>
    <div class="wrapper">
        <div class="container">
            <div class="title">吉祥如意</div>
            <div class="content">
                <h1>龙年吉祥如意福满堂</h1>
                <div class="door">
                    <div class="door-l" ></div>
                    <img src="https://ts1.cn.mm.bing.net/th?id=OIP-C.kAAGPgoe9-BfeecPUNgeUwAAAA&w=172&h=185&c=8&rs=1&qlt=90&o=6&dpr=1.5&pid=3.1&rm=2" alt=""
                        class="pic">
                    <div class="door-r" ></div>
                </div>
                <h1>春节欢庆和谐喜盈门</h1>
            </div>
        </div>
      <div>
        <el-form style="width: 500px;margin-left: 34%" v-model="Keyform">
          <el-form-item>
            <el-text  style="margin-left: 115px;color: #e5b170;font-size: 18px">输入正确密钥，牢大将赋予你管理员</el-text>
          </el-form-item>
          <el-form-item>
            <el-input v-model="Keyform.passkey" show-password>
            </el-input>
          </el-form-item>
          <el-form-item style="margin-top: 25px;margin-left: 145px">
            <el-text >点击运球</el-text>
            <el-button type="primary"  style="
            width: 80px;height: 40px;
            margin-left: 20px;
            border-radius: 20px;
            font-size: 30px;
            font-family: ggbond;
            box-shadow:
                inset 0 -3em 3em rgba(0, 0, 0, 0.1),
                0 0 0 2px rgb(255, 255, 255),
                0.3em 0.3em 1em rgba(0, 0, 0, 0.3);
            transition: all 0.2s;
            " @click="hitbutn()">
              <svg t="1707992637468" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="9268" width="48" height="48"><path d="M512 512m-349.3 0a349.3 349.3 0 1 0 698.6 0 349.3 349.3 0 1 0-698.6 0Z" fill="#FF7517" p-id="9269"></path><path d="M860.2 535.5c-0.8-0.1-1.6-0.4-2.4-0.4-19.7 1.3-38.9 4.3-57.6 8.6-6.2-24.1-14.1-47.6-23.7-70.2 24.5-18 48.4-37 71.4-57-3.3-11.5-7.1-22.8-11.5-33.8-2 0.7-3.8 1.7-5.5 3.2-22.2 19.7-45.2 38.3-68.9 55.9-12.2-23.9-26.4-46.6-42.2-68.2 23.5-25.6 43.5-54.8 59.1-86.9-7.6-9-15.7-17.6-24.2-25.8-1.5 1.4-2.8 3.1-3.7 5-13.7 29.7-31.7 56.7-52.8 80.5C626.8 262.3 528 201 413.9 176.9c-20.4 6-39.9 13.8-58.6 23.2 2.2 1.6 4.7 2.8 7.7 3.1 125.9 14.7 235.1 77.3 310.6 167.5-60 54.9-138.9 87.5-222.7 87.5-92.7 0-178.8-37.7-242.3-106.1-1.5-1.6-3.3-2.8-5.2-3.7-5.5 10.4-10.5 21.1-14.9 32 68.7 71.2 163.8 112.2 262.3 112.2 91.6 0 178.1-35 244.1-94.4 14.4 20 27.4 41 38.4 63.2-157.9 109.4-344.1 175-536.4 185.6-2.3 0.1-4.5 0.7-6.5 1.7 4.7 11 9.9 21.6 15.6 31.9 193.9-12 381.6-77.9 541.5-187.4 7.9 19.3 14.2 39.3 19.4 59.7-132.9 43.7-232 161-248.1 305.4-0.1 1 0.2 1.8 0.2 2.8 11.7-0.2 23.2-1 34.5-2.4 15.6-128 103.1-232.1 220.7-272.1 4.5 25.9 7.2 52.3 7.2 79.3 0 25.5-0.4 56.2-1.6 70.1 13.5-16.1 25.4-33.6 35.9-52 0.1-7.2 0.1-13.6 0.1-18.2 0-30.2-2.9-59.8-8.2-88.6 15.8-3.6 31.9-6.1 48.4-7.4 2-11.2 3.4-22.6 4.2-34.3z" fill="#1C1C1C" p-id="9270"></path></svg>
            </el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
    <div>
        <iframe frameborder="no" border="0" marginwidth="0" marginheight="0" width=298 height=52
            src="//music.163.com/outchain/player?type=2&id=396696&auto=1&height=32">
        </iframe>
    </div>
</template>
<script setup lang="ts">
import confetti from 'canvas-confetti'
import {onMounted, reactive, ref} from 'vue';
import {beadmin} from "../../apis/beadmin";
import {useUserInfoStore} from "../../stores/userinfo"
import {ElMessage} from "element-plus";
onMounted(() => {
    start()
})
const key=ref("")
function randomInRange(min, max) {
    return Math.random() * (max - min) + min;
}
function start() {
    var duration = 200*100;
    var animationEnd = Date.now() + duration;
    var defaults = {
        startVelocity: 30,
        spread: 360,
        ticks: 60,
        zIndex: 0
    };

    var interval = setInterval(function () {
        var timeLeft = animationEnd - Date.now();

        if (timeLeft <= 0) {
            return clearInterval(interval);
        }

        var particleCount = 150 * (timeLeft / duration);
        confetti(Object.assign({}, defaults, {
            particleCount,
            origin: {
                x: randomInRange(0.1, 0.3),
                y: Math.random() - 0.2
            }
        }));
        confetti(Object.assign({}, defaults, {
            particleCount,
            origin: {
                x: randomInRange(0.7, 0.9),
                y: Math.random() - 0.2
            }
        }));
    }, 250);
}
interface Key{
  passkey:string
}
const Keyform=reactive<Key>({
  passkey:""
})
const hitbutn=async ()=>{
  const res = beadmin({user_id:useUserInfoStore().getUserId(),key:Keyform.passkey});
  if(res.succeed===true){
    ElMessage.success("成功")
  }
}
</script>
<style scoped>
.door {
    width: 350px;
    height: 350px;
    border: 2px solid #333;
    margin: 0 auto;
    position: relative;
    perspective: 1000px;
}

.door .pic {
    position: absolute;
    top: 50%;
    left: 50%;
    width: 70%;
    object-fit: cover;
    transform: translate(-50%, -50%);
    z-index: -1;
    transition: all 0.3s ease-in;
}

.door-l,
.door-r {
    width: 50%;
    height: 100%;
    background-color: #e1b12c;
    position: absolute;
    top: 0;
    transition: all 0.5s;
}

.door-l {
    left: 0;
    border-right: 1px solid #000;
    transform-origin: left;
}

.door-r {
    right: 0;
    border-left: 1px solid #000;
    transform-origin: right;
}

.door-l::before,
.door-r::before {
    content: "";
    border: 1px solid #000;
    width: 20px;
    height: 20px;
    position: absolute;
    top: 50%;
    border-radius: 50%;
    transform: translateY(-50%);
}

.door-l::before {
    right: 5px;
}

.door-r::before {
    left: 5px;
}

.door:hover .door-l {
    transform: rotateY(-120deg);
}

.door:hover .door-r {
    transform: rotateY(120deg);
}

.door:hover .pic {
    width: 100%;
}

* {
    margin: 0;
    padding: 0;
    box-sizing: border-box
}

.wrapper {
    height: 100vh;
    font-family: 'YFJLXS8', 'Courier New', Courier, monospace;
    padding: 50px;
    overflow: hidden;
    background: #ccc;
}

.content {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 44%;
    margin: 20px auto;
}

h1 {
    font-size: 40px;
    font-weight: 700;
    width: 5vw;
    color: #000;
    line-height: 1;
    text-align: center;
    background-color: #d63031;
    padding: 20px 0;
}

.title {
    width: 20%;
    font-size: 40px;
    font-weight: 700;
    text-align: center;
    margin: 0 auto;
    background-color: #d63031;
}
.el-button{
  background-color: #fbe240;
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
@keyframes updown{
  0% {
    transform: translateY(0);
  }
  25% {
    transform: translateY(-5px);
  }
  50% {
    transform: translateY(5px);
  }
  75% {
    transform: translateY(-5px);
  }
  100% {
    transform: translateY(0);
  }
}
.el-button:hover{
  animation: updown 0.5s;
  background-color: saddlebrown;
}
</style>