<template>
  <div>
    <v-container class="grey lighten-5">
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
export default {
  // rooms = {name: hoge, rid: fuga}
  props: ["rooms"],
  data() {
    return {
      roomList: [],
    };
  },
  created() {
    for (let i = 0; i < Math.ceil(this.rooms.length / 2); i++) {
      let multiple_cnt = i * 2;
      let result = this.rooms.slice(multiple_cnt, multiple_cnt + 2);
      this.roomList.push(result);
    }
  },
  methods: {
    selectRoom(rid) {
      this.$router.push({ name: "room-rid", params: { rid: rid } });
    },
  },
};
</script>