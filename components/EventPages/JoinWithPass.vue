<template>
  <div id="overlay">
    <div id="content">
      <h1 style="text-align: center" class="-color-black">パスワードを入力</h1>
      <v-col cols="12" md="10" offset-md="1">
        <v-text-field
          v-model="password"
          label="password"
          outlined
          light
          color="teal darken-2"
          background-color="deep-purple lighten-5"
        ></v-text-field>
      </v-col>
      <v-row justify="center" align-content="center">
        <v-btn id="submit" color="primary" light @click="joinEvent">submit</v-btn>
        <v-btn color="primary" @click="$emit('closeModal')">close</v-btn>
      </v-row>
    </div>
  </div>
</template>

<script>
import post from '@/lib/post.js';

export default {
  props: ['prop_params'],
  data() {
    return {
      password: null,
      pass_f: false,
    }
  },
  methods: {
    async joinEvent() {
      let params = new URLSearchParams();
      params.append("uid", this.prop_params.uid);
      params.append("eid", this.prop_params.eid);
      params.append("password", this.password);
      post("/event/joinEvent", params).then((res) => {
        if (res.data.error != null) {
          console.log(error);
          return;
        }
        this.$store.commit("set_event_org", this.prop_params.org_uid);
        this.$store.commit("connect_event_ws", this.prop_params.eid);
        this.$router.push({ name: "event-eid", params: { eid: this.prop_params.eid } });
      });
    }
  }
}
</script>

<style scoped lang="scss">
.-color-black {
  color: black;
}
#content {
  z-index: 10;
  width: 580px;
  height: 250px;
  padding: 1em;
  background: white;
}
#overlay {
  z-index: 4;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
}
#submit {
  margin-right: 20px;
}
</style>