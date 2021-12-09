<template>
  <div>
    <faildDialog
      :text="text"
      :dialog="faild"
      @closeDialog="closeDialog"
    ></faildDialog>
    <h1>フレンドになろうとしています</h1>
  </div>
</template>

<script>
import post from "@/lib/post.js"
import faildDialog from '@/components/faildDialog.vue'

export default {
  components: {
    faildDialog,
  },
  data() {
    return {
      faild: false,
      text: "",
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
    await post("/makeFriend", params).then((res) => {
      if (res.data.error != null) {
        if (res.data.error === 'this user is your friend') {
          this.text = "ユーザーはすでにフレンドになっています"
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
      this.$router.push(`/friend`);
    }
  },
  methods: {
    closeDialog() {
      this.faild = false
      this.$router.push("/friend")
    },
  }
};
</script>