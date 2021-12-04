<template>
  <div>
    <v-container class="grey lighten-5">
      <v-row>
        <v-col md="5" offset-md="3" align-self="center">
          <p style="font-size: 1.9em; text-align: center; color: black">
            開催するROOMを選択
          </p>
        </v-col>
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
            @click="showSetPass(room.rid)"
          >
            {{ room.name }}
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
    <setPassword @closeModal="show=false" v-if="show" :rid="selected_rid"></setPassword>
  </div>
</template>

<script>
import post from "@/lib/post.js";
import setPassword from "@/components/EventPages/SetPassword";

export default {
  components: {
    setPassword,
  },
  data() {
    return {
      roomList: [],
      show: false,
      selected_rid: null,
    };
  },
  beforeCreate() {
    let params = new URLSearchParams();
    params.append("uid", this.$store.getters.getUser.uid);
    post("/getRoomList", params).then((res) => {
      let tmp = [];
      for (let i = 0; i < res.data.rooms.length; i++) {
        if (res.data.rooms[i].admin) {
          tmp.push(res.data.rooms[i]);
          if (tmp.length === 2) {
            this.roomList.push(tmp);
            tmp = []
          }
        }
      }
      if (tmp.length != 0) {
        this.roomList.push(tmp);
      }
    });
  },
  methods: {
    showSetPass(rid) {
      this.selected_rid = rid
      this.show = true
    }
  }
};
</script>