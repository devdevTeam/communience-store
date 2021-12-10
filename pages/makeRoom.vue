<template>
  <v-main>
    <v-container fluid fill-height class="grey lighten-5">
      <v-row justify="center" align-content="center">
        <v-col md="10" offset-md="1" align-self="center">
          <h1 class="-color-black" style="text-align: center;">Room作成</h1>
        </v-col>
        <v-col md="1" />
      </v-row>
      <v-row justify="center" align-content="center">
        <v-col md="10" offset-md="1" align-self="center">
          <v-text-field
            v-model="name"
            label="Room Name"
            required
            light
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row justify="center" align-content="center">
        <v-col md="10" offset-md="1" align-self="center">
          <v-text-field
            v-model="password"
            label="Room Password"
            counter="4"
            required
            light
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row justify="center" align-content="center">
        <v-checkbox
          v-model="haveForm"
          label=""
          type="checkbox"
          light
        ></v-checkbox>
        <h2 class="-color-black" style="margin-top: 15px">Room独自のFormの作成</h2>
      </v-row>
      <template v-if="haveForm">
        <v-row justify="center" align-content="center" style="margin-bottom: -30px;">
          <v-col md="3" offset-md="1">
            <p class="-color-black" style="text-align: center; margin: 0px 25px -30px 0px;">見出し項目</p>
          </v-col>
          <v-col md="8" />
        </v-row>
        <v-row justify="center" align-content="center" v-for="i in cols.length" :key=i>
          <v-col md="1" offset-md="2">
            <v-checkbox
              v-model="display[i-1]"
              type="checkbox"
              light
              :disabled="display[i-1]"
              @click="displayChange(i-1)"
            ></v-checkbox>
          </v-col>
          <v-col md="7" offset-md="0">
            <v-text-field
              v-model="cols[i-1]"
              :label="'項目' + i"
              light
              append-outer-icon="mdi-close-box-outline"
              @click:append-outer="removeCol(i-1)"
            ></v-text-field>
          </v-col>
          <v-col md="2" />
        </v-row>
        <v-row class="-bottom-margin" justify="center" align-content="center">
          <v-btn light text icon @click="addCol">
            <v-icon>mdi-plus</v-icon>
          </v-btn>
        </v-row>
      </template>
      <v-row class="-bottom-margin" justify="center" align-content="center">
        <v-btn color="primary" :disabled="!able" light @click="makeRoom">submit</v-btn>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
import post from '@/lib/post.js'

export default {
  data() {
    return {
      name: "",
      password: "",
      haveForm: false,
      cols: [""],
      display: [true],
      able: false,
    }
  },
  watch: {
    name() {
      if (this.name.length != 0 && this.password.length > 3) {
        this.able = true
      }
      else {
        this.able = false
      }
    },
    password() {
      if (this.name.length != 0 && this.password.length > 3) {
        this.able = true
      }
      else {
        this.able = false
      }
    },
  },
  methods: {
    addCol() {
      this.cols.push("")
      this.display.push(false)
    },
    removeCol(i) {
      if (this.cols.length == i) {
        this.cols.pop()
      }
      else {
        this.cols.splice(i, 1)
      }
    },
    displayChange(i) {
      for (let n = 0; n < this.display.length; n++) {
        this.display[n] = false
      }
      this.display[i] = true
    },
    async makeRoom() {
      let rid = null
      let params = new URLSearchParams()
      params.append("roomname", this.name)
      params.append("password", this.password)
      params.append("uid", this.$store.getters.getUser.uid)
      params.append("haveForm", this.haveForm)
      await post("/makeRoom", params).then((res) => {
        if (res.data.error != null) {
          console.error(res.data.error);
        }
        else {
          rid = res.data.rid
        }
      })
      if (this.haveForm && rid != null) {
        params.append("rid", rid)
        params.append("colList", this.cols.join(","))
        params.append("displayBoolList", this.display.join(","))
        await post("/makeForm", params).then((res) => {
          if (res.data.error != null) {
            console.error(res.data.error);
          }
        })
        this.$router.push(`/room/${rid}/${this.$store.getters.getUser.uid}/updateCardValue`)
      }
      else {
        this.$router.push(`/room/${rid}`)
      }
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