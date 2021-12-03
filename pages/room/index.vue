<template>
  <div>
    <v-container class="grey lighten-5">
      <v-row>
        <v-col md="5" offset-md="3" align-self="center">
          <p style="font-size: 1.9em; text-align: center; color: black">
            参加しているroom一覧
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
              >
                <v-icon dark> mdi-plus </v-icon>
              </v-btn>
            </v-col>
          </template>
          <v-card>
            <v-list>
              <v-list-item link>
                <v-list-item-title>roomを作成</v-list-item-title>
              </v-list-item>
              <v-divider></v-divider>
              <v-list-item link>
                <v-list-item-title @click="toJoin">roomに参加</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-card>
        </v-menu>
      </v-row>
      <v-row v-for="(two_rooms, i) in roomList" :key="i">
        <v-col md="6" v-for="(room, j) in two_rooms" :key="j">
          <v-btn
            block
            color="black"
            elevation="2"
            class="pa-2"
            x-large
            outlined
            tile
            @click="selectRoom(room.rid)"
          >
            {{ room.name }}
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
      roomList: [],
    };
  },
  beforeCreate() {
    let params = new URLSearchParams();
    params.append("uid", this.$store.getters.getUser.uid);
    post("/getRoomList", params).then((res) => {
      for (let i = 0; i < Math.ceil(res.data.rooms.length / 2); i++) {
        let multiple_cnt = i * 2;
        let result = res.data.rooms.slice(multiple_cnt, multiple_cnt + 2);
        this.roomList.push(result);
      }
    });
  },
  methods: {
    selectRoom(rid) {
      this.$router.push({ name: "room-rid", params: { rid: rid } });
    },
    toJoin() {
      this.$router.push("/joinRoom")
    }
  },
};
</script>