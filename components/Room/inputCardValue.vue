<template>
  <v-main>
    <v-container fluid fill-height class="grey lighten-5">
      <v-row justify="center" align-content="center" v-for="col, i in colList" :key="i">
        <v-col md="4" offset-md="0" align-self="center">
          <v-list-item-title class="-color-black">{{col}}</v-list-item-title>
        </v-col>
        <v-col md="4" offset-md="0" align-self="center">
          <v-text-field
            v-model="value"
            required
            light
          ></v-text-field>
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
export default {
  props: ['colList'],
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