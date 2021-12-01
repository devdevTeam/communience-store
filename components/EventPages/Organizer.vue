<template>
  <v-main>
    <v-container fluid fill-height>
      <v-row justify="center" align-content="center">
        <v-col cols="12" sm="7" offset-sm="0">
          <v-card>
            <v-list-item-title>col list</v-list-item-title>
            <v-list-item-title v-for="(col, i) in items" :key="i">{{col.colName}}{{col.hidden}}</v-list-item-title>
            <v-btn color="indigo" @click="updateCol">更新!</v-btn>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
export default {
  props: ["items"],
  methods: {
    updateCol() {
      let msg = JSON.stringify({
        uid: this.$store.getters.getUser.uid,
        eid: this.$route.params.eid,
        hidden: [true, false, true, false, true],
      });
      let ws = this.$store.getters.getEventWs
      ws.send(msg);
    },
  },
};
</script>

