<template>
  <div>
    <faildDialog
      :text="text"
      :dialog="faild"
      @closeDialog="closeDialog"
    ></faildDialog>
    <h1>Roomに参加しようとしています</h1>
  </div>
</template>

<script>
import post from "@/lib/post.js"
import faildDialog from '@/components/faildDialog.vue'

export default {
  components: {
    faildDialog
  },
  data() {
    return {
      faild: false,
      text: "",
    }
  },
  created() {
    let params = new URLSearchParams();
    params.append("uid", this.$store.getters.getUser.uid);
    params.append("hash", this.$route.params.uri);
    post("/checkHash", params).then((res) => {
      if (res.data.error != "") {
        if (res.data.error != null) {
          console.log(res);
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
      }
      this.$router.push(`/room/${res.data.rid}/${this.$store.getters.getUser.uid}/updateCardValue`);
    });
  },
  methods: {
    closeDialog() {
      this.faild = false
      this.$router.push("/joinRoom")
    }
  }
};
</script>