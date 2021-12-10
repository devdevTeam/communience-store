<template>
  <v-main>
    <v-container fluid fill-height class="grey lighten-5">
      <v-row justify="center" align-content="center">
        <v-col cols="12" md="6" offset-md="3">
          <h1 class="-color-black" style="text-align: center;">{{name}}</h1>
        </v-col>
        <v-col cols="12" md="3" />
      </v-row>
      <v-row justify="center" align-content="center" v-for="col, i in items" :key="i">
        <v-col md="4" offset-md="0" align-self="center">
          <v-list-item-title class="-color-black"><h2>{{col.colName}}</h2></v-list-item-title>
        </v-col>
        <v-col md="4" offset-md="0" align-self="center">
          <v-list-item-title v-if="!col.hidden" class="-color-black"><h2>{{values[i]}}</h2></v-list-item-title>
          <v-btn v-else color="primary" @click="submit(i)">解答</v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
import post from "@/lib/post.js"
export default {
  props: ['items', "name"],
  data() {
    return {
      values: [],
    }
  },
  async created() {
    let values = []
    for (let i = 0; i < this.items.length; i++) {
      values.push(null)
    }
    let params = new URLSearchParams()
    params.append("uid", this.$store.getters.getUser.uid)
    params.append("rid", this.$store.getters.getEventInfo.rid)
    await post("/getCardValue", params).then((res) => {
      values = res.data.cardValue
    })
    this.values = values
  },
  methods: {
    submit(i) {
      this.items[i].hidden = !this.items[i].hidden
    }
  }
};
</script>

<style>
.-color-black {
  color: black;
}
</style>