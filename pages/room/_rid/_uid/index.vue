<template>
  <v-main>
    <my-card v-if="!form" :info="info" :hobby="hobby" :friend="friend"></my-card>
    <v-container v-else fluid fill-height class="grey lighten-5">
      <v-row />
      <v-row justify="center" align-content="center" v-for="col, i in items" :key="i">
        <v-col md="4" offset-md="0" align-self="center">
          <h2 class="-color-black"><b>{{col}}</b></h2>
        </v-col>
        <v-col md="4" offset-md="0" align-self="center">
          <h2 class="-color-gray">{{values[i]}}</h2>
        </v-col>
      </v-row>
      <v-row style="margin: 0px 0 30px 0" />
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
.-color-gray {
  color: rgb(170, 54, 0);
}
</style>