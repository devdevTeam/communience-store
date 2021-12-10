<template>
  <v-main>
    <v-row>
      <v-col cols="12" md="3" offset-md="3">
        <v-btn text style="margin: -40px 0 0px 0" to="/friend">＜ フレンド一覧</v-btn>
      </v-col>
      <v-col cols="12" md="6" />
    </v-row>
    <my-card :info=info :hobby=hobby :friend=friend :canEdit=false></my-card>
  </v-main>
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