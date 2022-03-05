<template>
  <div id="app">
      <h1>Vue实例周期</h1>
  </div>
</template>

<script>
export default {
  beforeCreate() {
    console.log("beforeCreate")
  },
  created() {
    console.log("created")
  },
  beforeMount() {
    console.log("beforeMount")
  },
  mounted() {
    console.log("mounted")
  }
}

</script>

<style>

</style>
