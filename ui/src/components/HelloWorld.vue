<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <p>
      <span v-if='`${info.code}` === "200"'>
      数据库大小{{ info.data }}.
      </span>
      <span v-else>
      {{ info.data }}.
      </span>
    </p>
  </div>
</template>

<script>
const  axios = require('axios')
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  data() {
    return {
      info: null
    }
  },
  created() {
    this.fetchDBTotal()
  },
  methods: {
    fetchDBTotal(){
      axios({
        method: 'GET',
        url: '/apis/example.dev/v1beta/db/total',
        headers: {'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMxNTA2NjQsInJvbGUiOiJhZG1pbiIsInVzZXJuYW1lIjoiYWRtaW4ifQ.Hd3giIhtp1pPRcjMmsRKF6HHUFqyRKTESZC2acwmeaw'},
      }).then(response => {
        this.info = response.data
      })
    }
  }

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
