<template>
  <v-main>
    <my-card :info=info :hobby=hobby :friend=friend :canEdit=true></my-card>
  </v-main>
</template>

<script>
import MyCard from '../../components/My-card.vue'
import post from '../../lib/post.js'
export default {
  components: { MyCard },
  async asyncData({store}) {
    let info = null, hobby = [], friend = []
    let params = new URLSearchParams()
    params.append('uid',store.getters.getUser.uid)
    await post("/getDefaultCard",params).then((res) => { 
      console.log(res);
      info = res.data;
      hobby = res.data.hobby.split(",")
    })
    return {
      info: info,
      hobby: hobby,
      friend: friend,
    }
  },
}
</script>