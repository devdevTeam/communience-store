<template>
  <v-main>
    <my-card v-if="!form" :info="info" :hobby="hobby" :friend="friend"></my-card>
    <v-container v-else fluid fill-height class="grey lighten-5">
      <v-row justify="center" align-content="center" v-for="col, i in items" :key="i">
        <v-col md="4" offset-md="0" align-self="center">
          <v-list-item-title class="-color-black">{{col}}</v-list-item-title>
        </v-col>
        <v-col md="4" offset-md="0" align-self="center">
          <v-list-item-title class="-color-black">{{values[i]}}</v-list-item-title>
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
import post from '@/lib/post.js'
import MyCard from '@/components/My-card.vue';
export default {
  components: { MyCard },
  async asyncData({route}) {
    let form = true
    let info = null, hobby = [], friend = []
    let items = [], values = []
    let params = new URLSearchParams()
    params.append("rid", route.params.rid)
    params.append("uid", route.params.uid)
    await post("/getRoomInfo", params).then((res) => {
      if (res.data.error != null) {
        console.error(res.data.error)
      }
      form = res.data.haveForm
    })
    if (form !== true) {
      form = false
      await post("/getDefaultCard",params).then((res) => { 
        console.log(res);
        info = res.data;
      })  
    }
    else {
      await post("/getForm", params).then((res) => {
        items = res.data.colList
      })
      await post("/getCardValue", params).then((res) => {
        values = res.data.cardValue
      })
    }
    return {
      info: info,
      hobby: hobby,
      friend: friend,
      items: items,
      values: values,
      form: form,
    }
  },
};
</script>

<style>
.-color-black {
  color: black;
}
</style>