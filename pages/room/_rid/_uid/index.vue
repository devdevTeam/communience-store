<template>
  <v-main>
    <v-container fluid fill-height class="grey lighten-5">
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
export default {
  data() {
    return {
      items: [],
      values: [],
    }
  },
  beforeCreate() {
    let params = new URLSearchParams()
    params.append("rid", this.$route.params.rid)
    post("/getForm", params).then((res) => {
      this.items = res.data.colList
    })
    params.append("uid", this.$route.params.uid)
    post("/getCardValue", params).then((res) => {
      this.values = res.data.cardValue
    })
  }
};
</script>

<style>
.-color-black {
  color: black;
}
</style>