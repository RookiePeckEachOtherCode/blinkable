<template>
    <div class="wrapper">
        <div class="container">
            <div class="title">吉祥如意</div>
            <div class="content">
                <h1>龙年吉祥如意福满堂</h1>
                <div class="door">
                    <div class="door-l"></div>
                    <img src="https://th.bing.com/th/id/OIP.kAAGPgoe9-BfeecPUNgeUwAAAA?rs=1&pid=ImgDetMain" alt=""
                        class="pic">
                    <div class="door-r"></div>
                </div>
                <h1>春节欢庆和谐喜盈门</h1>
            </div>
        </div>
    </div>
    <div>
        <iframe frameborder="no" border="0" marginwidth="0" marginheight="0" width=298 height=52
            src="//music.163.com/outchain/player?type=2&id=396696&auto=1&height=32"></iframe>
    </div>
</template>
<script setup lang="ts">
import confetti from 'canvas-confetti'
import { onMounted } from 'vue';
onMounted(() => {
    start()
})

function randomInRange(min, max) {
    return Math.random() * (max - min) + min;
}
function start() {
    var duration = 10 * 1000;
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
</style>