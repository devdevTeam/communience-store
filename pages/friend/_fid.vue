<template>
  <my-card :info=info :hobby=hobby :friend=friend></my-card>
</template>

<script>
import MyCard from '@/components/My-card.vue'
import post from '@/lib/post.js'
export default {
  components: { MyCard },
  async asyncData({route}) {
    let info = null, hobby = [], friend = []
    let params = new URLSearchParams()
    params.append('uid',route.params.fid)
    await post("/getDefaultCard",params).then((res) => { 
      console.log(res);
      info = res.data;
    })
    return {
      info: info,
      hobby: hobby,
      friend: friend,
    }
  },
}
</script>