<template>
  <div>
    <v-container class="grey lighten-5">
      <v-row>
        <v-col md="5" offset-md="3" align-self="center">
          <h1 style="text-align: center; color: black">
            フレンド一覧
          </h1>
        </v-col>
        <v-menu
          :close-on-content-click="false"
          :nudge-width="100"
          slide-y
          bottom
          left
        >
          <template v-slot:activator="{ on, attrs }">
            <v-col md="1" offset-md="3">
              <v-btn
                class="mb-2"
                fab
                dark
                color="indigo"
                v-bind="attrs"
                v-on="on"
              >
                <v-icon dark> mdi-plus </v-icon>
              </v-btn>
            </v-col>
          </template>
          <v-card>
            <v-list>
              <v-list-item link>
                <v-list-item-title @click="input_show=true">フレンドを追加する</v-list-item-title>
              </v-list-item>
              <v-divider></v-divider>
              <v-list-item link>
                <v-list-item-title @click="show=true">自分を紹介する</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-card>
        </v-menu>
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
      <v-row v-for="(two_friends, i) in friendList" :key="i">
        <v-col md="6" v-for="(friend, j) in two_friends" :key="j">
          <v-btn
            block
            color="black"
            elevation="2"
            class="pa-2"
            x-large
            outlined
            tile
            @click="selectFriend(friend.fid)"
          >
            {{ friend.displayValue }}
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
    <PresentUrl @closeModal="show=false" v-if="show" :uid="$store.getters.getUser.uid"></PresentUrl>
    <InputURL @closeInput="input_show=false" v-if="input_show"></InputURL>
  </div>
</template>


<script>
import post from "@/lib/post.js";
import PresentUrl from '@/components/Friend/presentURL.vue'
import InputURL from '@/components/Friend/inputURL.vue'

export default {
  components: {
    PresentUrl,
    InputURL
  },
  data() {
    return {
      show: false,
      input_show: false,
      friendList: [],
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
      for (let i = 0; i < Math.ceil(res.data.friends.length / 2); i++) {
        let multiple_cnt = i * 2;
        let result = res.data.friends.slice(multiple_cnt, multiple_cnt + 2);
        this.friendList.push(result);
      }
    })
  },
  methods: {
    selectFriend(fid) {
      this.$router.push(`/friend/${fid}`);
    },
    search() {
      this.friendList = []
      let params  = new URLSearchParams()
      params.append("uid", this.$store.getters.getUser.uid)
      params.append("colName", this.selected_col)
      params.append("value", this.value)
      post("/searchFriend", params).then((res) => {
        for (let i = 0; i < Math.ceil(res.data.friends.length / 2); i++) {
          let multiple_cnt = i * 2;
          let result = res.data.friends.slice(multiple_cnt, multiple_cnt + 2);
          this.friendList.push(result);
        }
      })
    },
    reset() {
      this.friendList = []
      let params  = new URLSearchParams()
      params.append("uid", this.$store.getters.getUser.uid)
      params.append("haveForm", this.haveForm);
      post("/getFriends", params).then((res) => {
        for (let i = 0; i < Math.ceil(res.data.friends.length / 2); i++) {
          let multiple_cnt = i * 2;
          let result = res.data.friends.slice(multiple_cnt, multiple_cnt + 2);
          this.friendList.push(result);
        }
      });
    }
  },
};
</script>