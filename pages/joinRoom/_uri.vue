<template>
  <div>
    <faildDialog
      :text="'無効なURLです'"
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
      faild: false
    }
  },
  created() {
    let params = new URLSearchParams();
    params.append("uid", this.$store.getters.getUser.uid);
    params.append("hash", this.$route.params.uri);
    post("/checkHash", params).then((res) => {
      if (res.data.error != "") {
        this.faild = true
        return
      }
      this.$router.push("/room")
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