<template>
  <div id="overlay">
    <faildDialog :text="'パスワードが違います'" :dialog="faild" @closeDialog="closeDialog"></faildDialog>
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
        <v-btn id="submit" color="primary" light @click="joinRoom">submit</v-btn>
        <v-btn color="primary" @click="$emit('closeModal')">close</v-btn>
      </v-row>
    </div>
  </div>
</template>

<script>
import post from '@/lib/post.js';
import faildDialog from '@/components/faildDialog.vue'

export default {
  props: ['rid'],
  components: {
    faildDialog
  },
  data() {
    return {
      password: null,
      pass_f: false,
      faild: false,
    }
  },
  methods: {
    async joinRoom() {
      let params = new URLSearchParams();
      params.append("uid", this.$store.getters.getUser.uid);
      params.append("rid", this.rid);
      params.append("password", this.password);
      post("/joinRoom", params).then((res) => {
        if (res.data.error != null) {
          console.error(res.data.error);
          this.faild = true
          return;
        }
        this.$router.push("/room");
      });
    },
    closeDialog() {
      this.faild = false
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