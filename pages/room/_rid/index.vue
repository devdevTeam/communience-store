<template>
  <div>
    <v-container class="grey lighten-5">
      <v-row>
        <v-col md="5" offset-md="3" align-self="center">
          <p style="font-size: 1.9em; text-align: center; color: black">
            参加しているuser一覧
          </p>
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
                v-if="admin"
              >
                <v-icon dark> mdi-plus </v-icon>
              </v-btn>
            </v-col>
          </template>
          <v-card>
            <v-list>
              <v-list-item link>
                <v-list-item-title @click="toStartEvent">イベントを開催</v-list-item-title>
              </v-list-item>
              <v-divider></v-divider>
              <v-list-item link>
                <v-list-item-title>Roomに招待</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-card>
        </v-menu>
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
            {{ user.name }}
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
      admin: false,
    };
  },
  beforeCreate() {
    let params = new URLSearchParams();
    params.append("rid", this.$route.params.rid);
    post("/getRoomUsers", params).then((res) => {
      for (let i = 0; i < Math.ceil(res.data.users.length / 2); i++) {
        let multiple_cnt = i * 2;
        let result = res.data.users.slice(multiple_cnt, multiple_cnt + 2);
        this.userList.push(result);
      }
    });
    params.append("uid", this.$store.getters.getUser.uid)
    post("/getRoomAdmin", params).then((res) => {
      console.log(res);
      if (res.data.error != null) {
        console.error(res.data.error)
      }
      if (res.data.admin) {
        this.admin = true
      }
    })
  },
  methods: {
    selectUser(uid) {
      let rid = this.$route.params.rid
      this.$router.push({ name: "room-rid-uid", params: { rid: rid, uid } });
    },
    toStartEvent() {
      this.$router.push("/startEvent")
    },
  },
};
</script>