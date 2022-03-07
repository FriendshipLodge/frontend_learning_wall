<template>
  <div id="app">
      <h1>插槽</h1>
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

</style>
