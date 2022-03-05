## 项目创建
    vue create <项目名>

## 属性赋值
    {{ msg }}


## 绑定属性 v-bind:
    <img v-bind:src="url" alt="图片">
    简写  <img :src="url" alt="图片">
   

## 绑定事件 v—on:
    <button v-on:click="sayHi">测试按钮</button>
    简写 <button @click="sayHi">测试按钮</button>


## this获取vue实例属性
```javascript
export default {
    name: 'HelloWorld',
    // props: {
    //   msg: String
    // },
    el:".hello",
    data: function() {
        return {
            msg: "哈哈",
            url: "../assets/logo.png"
        };
    },
    //绑定事件
    methods: {
        //方法
        sayHi(){
            alert("hi")
        }
    }
}
```