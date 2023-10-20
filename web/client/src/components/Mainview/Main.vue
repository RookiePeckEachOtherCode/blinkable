  <script setup lang="ts">

  </script>

  <template>

    <div class="card"> <el-button type="primary">Primary</el-button></div>
    <div class="card"> <el-button type="primary">Primary</el-button></div>
    <div class="card"> <el-button type="primary">Primary</el-button></div>
  </template>

  <style scoped lang="scss">
  .card{
    top:100px;
    width: 720px;
    height: 450px;
    border: 1px solid white;
    box-shadow: 4px 4px 3px rgba(0, 0, 4, 0.2);
    border-radius: 10px;
    display: flex;
    flex-direction: column;
    justify-content: center; // 水平居中
    margin-bottom: 150px;
    align-items: center;
    bottom: 100px;
    position: relative;
    background-color: white;
    z-index: 0;
  }

  </style>