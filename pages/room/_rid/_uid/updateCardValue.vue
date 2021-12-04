<template>
  <v-main>
    <faildDialog
      :text="'無効な操作です'"
      :dialog="faild"
      @closeDialog="closeDialog"
    ></faildDialog>
    <v-container fluid fill-height class="grey lighten-5">
      <v-row justify="center" align-content="center" v-for="col, i in colList" :key="i">
        <v-col md="4" offset-md="0" align-self="center">
          <v-list-item-title class="-color-black">{{col}}</v-list-item-title>
        </v-col>
        <v-col md="4" offset-md="0" align-self="center">
          <v-text-field
            v-model="values[i]"
            :label="col"
            required
            light
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row justify="center" align-content="center">
        <v-btn class="-bottom-margin" x-large color="primary" @click="updateCardValue">submit</v-btn>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
import post from '@/lib/post.js'
import faildDialog from '@/components/faildDialog.vue'

export default {
  name: "updateCardValue",
  components: {
    faildDialog
  },
  async asyncData(ctx) {
    let values = [], colList = []
    let params = new URLSearchParams()
    params.append("rid", ctx.params.rid)
    await post("/getForm", params).then((res) => {
      colList = res.data.colList
    })
    params.append("uid", ctx.params.uid)
    await post("/getCardValue", params).then((res) => {
      values = res.data.cardValue
    })
    return {
      colList: colList,
      values: values,
      faild: false
    }
  },
  created() {
    if (this.$route.params.uid != this.$store.getters.getUser.uid) {
      this.faild = true
    }
  },
  methods: {
    updateCardValue() {
      let params = new URLSearchParams()
      params.append('uid', this.$route.params.uid)
      params.append('rid', this.$route.params.rid)
      params.append('valueList', this.values.join(','))
      post('/updateCardValue', params).then((res) => {
        if (res.data.error != null) {
          console.error(res.data.error);
          return
        }
        this.$router.push(`/room/${params.rid}/${params.uid}`)
      })
    },
    closeDialog() {
      this.faild = false
      this.$router.push(`/room/${this.$route.params.rid}/${this.$route.params.uid}`)
    }
  },
};
</script>

<style>
.-color-black {
  color: black;
}
.-bottom-margin {
  margin-bottom: 10px;
}
</style>