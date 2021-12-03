<template>
  <v-main>
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
        <v-btn x-large color="primaly" @click="updateCardValue">submit</v-btn>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
import post from '@/lib/post.js'

export default {
  props: ['colList'],
  asyncData() {
    let values = []
    for (let i = 0; i < this.colList.length; i++) {
      values.push(null)
    }
    return {
      values: values
    }
  },
  methods: {
    updateCardValue() {
      let params = new URLSearchParams()
      params.append('uid', this.$store.getters.getUser.uid)
      params.append('rid', this.$route.params.rid)
      params.append('valueList', this.values.join(','))
      post('/updateCardValue', params).then((res) => {
        if (res.data.error != null) {
          console.error(res.data.error);
        }
      })
    },
  },
};
</script>

<style>
.-color-black {
  color: black;
}
.-top-bottom-margin {
  margin-top: 10px;
  margin-bottom: 10px;
}
</style>