<template>
  <div id="overlay">
    <faildDialog :text="text" :dialog="faild" @closeDialog="closeDialog"></faildDialog>
    <confirm-dialog 
      :text="'MyCardをRoomに登録しますか？'" 
      :dialog="confirm"
      @closeConfirmYes="Yes" 
      @closeConfirmNo="No"
    ></confirm-dialog>
    <div id="content">
      <h1 style="text-align: center" class="-color-black">パスワードを入力</h1>
      <h5 style="text-align: center" class="-color-black">Room ID : {{rid}}</h5>
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
import ConfirmDialog from '@/components/confirmDialog.vue';


export default {
  props: ['rid'],
  components: {
    faildDialog,
    ConfirmDialog
  },
  data() {
    return {
      password: null,
      pass_f: false,
      faild: false,
      text: "",
      confirm: false,
      succeed: false,
    }
  },
  methods: {
    async joinRoom() {
      this.succeed = false
      let params = new URLSearchParams();
      params.append("uid", this.$store.getters.getUser.uid);
      params.append("rid", this.rid);
      params.append("password", this.password);
      await post("/joinRoom", params).then((res) => {
        if (res.data.error != null) {
          console.log(res);
          if (res.data.error === "isn't match password") {
            this.text = "パスワードが違います"
          }
          else if (res.data.error === 'this user exists in this room') {
            this.text = "このRoomに参加済みです"
          }
          else {
            console.error(res.data.error);
            this.text = "不具合が発生しました"
          }
          this.faild = true
          return;
        }
        else {
          this.succeed = true
        }
      });
      if (this.succeed) {
        let haveForm = true
        await post("/getRoomInfo", params).then((res) => {
          haveForm = res.data.haveForm
        })
        if (haveForm) {
          this.$router.push(`/room/${this.rid}/${this.$store.getters.getUser.uid}/updateCardValue`);
        }
        else {
          this.confirm = true
        } 
      }
    },
    closeDialog() {
      this.faild = false
    },
    Yes() {
      this.confirm = false
      this.$router.push(`/room/${this.rid}/`)
    },
    No() {
      let params = new URLSearchParams();
      params.append("uid", this.$store.getters.getUser.uid);
      params.append("rid", this.rid)
      post("/leaveRoom", params).then(() => {
        this.confirm = false
      })
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