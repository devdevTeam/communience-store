<template>
  <v-main>
    <v-container fluid fill-height class="grey lighten-5">
      <v-row justify="center" align="center">
        <v-col cols="12" md="3" offset-md="10">
          <v-btn outlined color="pink darken-1" @click="closeEvent">
            終了
          </v-btn>
        </v-col>
      </v-row>
      <v-row
        justify="center"
        align-content="center"
        v-for="(col, i) in items"
        :key="i"
      >
        <v-col md="4" offset-md="0" align-self="center">
          <v-list-item-title v-if="!col.hidden" class="-color-black">{{
            col.colName
          }}</v-list-item-title>
          <v-list-item-title v-else class="-color-black"
            ><s>{{ col.colName }}</s></v-list-item-title
          >
        </v-col>
        <v-col md="4" offset-md="0" align-self="center">
          <v-btn color="primary" @click="changeHidden(i)">かくす</v-btn>
        </v-col>
      </v-row>
      <v-row justify="center" align-content="center">
        <v-btn color="primary" class="-top-bottom-margin" @click="updateCol"
          >更新</v-btn
        >
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
import post from "@/lib/post.js"
export default {
  props: ["items"],
  methods: {
    updateCol() {
      let hidden_list = [];
      for (let i = 0; i < this.items.length; i++) {
        hidden_list.push(this.items[i].hidden);
      }
      let msg = JSON.stringify({
        uid: this.$store.getters.getUser.uid,
        eid: this.$route.params.eid,
        hidden: hidden_list,
      });
      let ws = this.$store.getters.getEventWs;
      ws.send(msg);
    },
    changeHidden(i) {
      this.items[i].hidden = !this.items[i].hidden;
    },
    closeEvent() {
      let params = new URLSearchParams()
      params.append("eid", this.$route.params.eid)
      post("/closeEvent", params).then((res) => {
        console.log(res)
        if (res.data.error != null) {
          console.error(res.data.error)
        }
        this.$router.push("/event")
      })
    }
  },
};
</script>

<style>
.-color-black {
  color: black;
}
.-top-bottom-margin {
  margin-top: 10px;
  margin-bottom: 10px;
}
</style>

