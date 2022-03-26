<template>

  <div id="app">
    <router-link to="">首页</router-link>
    <router-link to="">博客</router-link>
    <router-link to="">个人</router-link>
    <div id="top">

      <h1>上部分</h1>

    </div>

    <div id="middle">

      <h1>中部分</h1>

    </div>

    <div id="bottom">

      <h1>下部分</h1>

    </div>

  </div>

</template>

<script>
import axios from 'axios'
export default {
  data(){
    return {
    }
  },
  methods:{
    getUserInfoAll(){
      axios.post("http://localhost:8081/v1/allUserInfo").then(res=>{

           console.log(res.data)
      })
    }
  },
  created() {
    this.getUserInfoAll()
  }

}
</script>
<style>
#top{
  color: aquamarine;
  background: black;
  margin-top: 0;
}
</style>