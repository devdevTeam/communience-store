<template>
  <div id="overlay">
    <div id="content">
      <h1 style="text-align: center" class="-color-black">パスワードを設定</h1>
      <h5 style="text-align: center" class="-color-black">Room name: {{name}}</h5>
      <h5 style="text-align: center" class="-color-black">Room ID: {{rid}}</h5>
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
        <v-btn id="submit" color="primary" light @click="startEvent">submit</v-btn>
        <v-btn color="primary" @click="$emit('closeModal')">close</v-btn>
      </v-row>
    </div>
  </div>
</template>

<script>
import post from '@/lib/post.js';

export default {
  props: ['rid', 'name'],
  data() {
    return {
      password: null,
      pass_f: false,
    }
  },
  methods: {
    async startEvent() {
      let params = new URLSearchParams()
      params.append('org_uid', this.$store.getters.getUser.uid)
      params.append('rid', this.rid)
      params.append('password', this.password)
      let res = await post('/startEvent', params)
      if (res.data.error != null) {
        console.error(res.data.error)
      }
      else {
        this.$router.push("/event")
      }
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