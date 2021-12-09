<template>
  <div>
    <v-container class="grey lighten-5">
      <v-row>
        <v-col md="5" offset-md="3" align-self="center">
          <h1 style="text-align: center; color: black">
            フレンド一覧
          </h1>
        </v-col>
      </v-row>
      <v-row>
        <v-col cols="12" md="3" offset-md="1">
          <v-select
            v-model="selected_col"
            :items="cols"
            label="項目"
            light
            filled
            dense
          ></v-select>
        </v-col>
        <v-col cols="12" md="7">
          <v-text-field
            v-model="value"
            append-icon="mdi-magnify"
            append-outer-icon="mdi-close-box-outline"
            label="検索"
            light
            filled
            dense
            @click:append="search"
            @click:append-outer="reset"
          ></v-text-field>
        </v-col>
      </v-row>
      <v-row v-for="(two_users, i) in userList" :key="i">
        <v-col md="6" v-for="(user, j) in two_users" :key="j">
          <v-btn
            block
            color="black"
            elevation="2"
            class="pa-2"
            x-large
            outlined
            tile
            @click="selectUser(user.uid)"
          >
            {{ user.displayValue }}
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>


<script>
import post from "@/lib/post.js";

export default {
  data() {
    return {
      userList: [],
      selected_col: [],
      cols: [
        {text: "名前", value: "name"},
        {text: "生年月日", value: "birthday"},
        {text: "趣味", value: "hobby"},
        {text: "【自己紹介】", value: "free"},
      ],
      value: "",
    };
  },
  async beforeCreate() {
    let params = new URLSearchParams();
    params.append("uid", this.$store.getters.getUser.uid)
    post("/getFriends", params).then((res) => {
      for (let i = 0; i < Math.ceil(res.data.users.length / 2); i++) {
        let multiple_cnt = i * 2;
        let result = res.data.users.slice(multiple_cnt, multiple_cnt + 2);
        this.userList.push(result);
      }
    })
  },
  methods: {
    selectUser(uid) {
      this.$router.push(`/friend/${uid}`);
    },
    search() {
      this.userList = []
      let params  = new URLSearchParams()
      params.append("uid", this.$store.getters.getUser.uid)
      params.append("colName", this.selected_col)
      params.append("value", this.value)
      post("/searchFriend", params).then((res) => {
        for (let i = 0; i < Math.ceil(res.data.users.length / 2); i++) {
          let multiple_cnt = i * 2;
          let result = res.data.users.slice(multiple_cnt, multiple_cnt + 2);
          this.userList.push(result);
        }
      })
    },
    reset() {
      this.userList = []
      let params  = new URLSearchParams()
      params.append("uid", this.$store.getters.getUser.uid)
      params.append("haveForm", this.haveForm);
      post("/getFriends", params).then((res) => {
        for (let i = 0; i < Math.ceil(res.data.users.length / 2); i++) {
          let multiple_cnt = i * 2;
          let result = res.data.users.slice(multiple_cnt, multiple_cnt + 2);
          this.userList.push(result);
        }
      });
    }
  },
};
</script>