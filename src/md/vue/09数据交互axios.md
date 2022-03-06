## axios
    引入：import axios from "axios"
    
```vue
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
```