<template>
  <div>
    <faildDialog
      :text="text"
      :dialog="faild"
      @closeDialog="closeDialog"
    ></faildDialog>
    <confirm-dialog 
      :text="'MyCardをRoomに登録しますか？'" 
      :dialog="confirm"
      @closeConfirmYes="Yes" 
      @closeConfirmNo="No"
    ></confirm-dialog>
    <h1>Roomに参加しようとしています</h1>
  </div>
</template>

<script>
import post from "@/lib/post.js"
import faildDialog from '@/components/faildDialog.vue'
import ConfirmDialog from '../../components/confirmDialog.vue';

export default {
  components: {
    faildDialog,
    ConfirmDialog
  },
  data() {
    return {
      faild: false,
      text: "",
      confirm: false,
      rid: null,
      succeed: false,
    }
  },
  async created() {
    this.succeed = false
    let haveForm = true
    let params = new URLSearchParams();
    params.append("uid", this.$store.getters.getUser.uid);
    params.append("hash", this.$route.params.uri);
    await post("/checkHash", params).then((res) => {
      if (res.data.error != null) {
        if (res.data.error === 'this user exists in this room') {
          this.text = "このRoomに参加済みです"
        }
        else if (res.data.error === "invalid hash") {
          this.text = "無効なURLです"
        }
        else {
          console.error(res.data.error);
          this.text = "不具合が発生しました"
        }
        this.faild = true
        return;
      }
      else {
        this.rid = res.data.rid
        this.succeed = true
      }
    });
    console.log(this.succeed)
    if (this.succeed) {
      console.log(this.succeed)
      params.append("rid", this.rid)
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
  methods: {
    closeDialog() {
      this.faild = false
      this.$router.push("/joinRoom")
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
        this.$router.push("/room")
      })
    }
  }
};
</script>