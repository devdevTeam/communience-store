<template>
  <div>
    <v-row justify="center" align="center">
      <v-col cols="12" md="10" offset-md="1">
        <v-form ref="form" v-model="valid" lazy-validation>
          <v-text-field
            v-model="name"
            label="Roomの名前"
            required
            class="-margin-top"
          ></v-text-field>
          <v-text-field v-model="rid" label="Room ID" required></v-text-field>
          <div class="text-center">
            <v-btn
              center
              :disabled="valid"
              color="primary"
              class="-margin-bottom"
              @click="search"
            >
              search
            </v-btn>
          </div>
        </v-form>
      </v-col>
      <v-col cols="12" md="1"></v-col>
    </v-row>
    <v-row v-if="search_res" justify="center" align="center">
      <v-col cols="12" md="12">
        <roomList @selectRoom="selectRoom" :rooms="rooms"></roomList>
      </v-col>
    </v-row>
    <joinWithPass @closeModal="show=false" v-if="show" :rid="selected_rid"></joinWithPass>
  </div>
</template>

<script>
import roomList from "./roomList.vue";
import post from "@/lib/post.js";
import joinWithPass from "./joinWithPass.vue"

export default {
  components: {
    roomList,
    joinWithPass,
  },
  data() {
    return {
      show: false,
      valid: true,
      name: "",
      rid: "",
      rooms: [],
      search_res: false,
      selected_rid: null,
    };
  },
  watch: {
    name() {
      if (this.name.length != 0 && this.rid.length == 0) {
        this.valid = false;
      } else {
        this.valid = true;
      }
    },
    rid() {
      if (this.rid.length != 0 && this.name.length == 0) {
        this.valid = false;
      } else {
        this.valid = true;
      }
    },
    rooms() {
      if (this.rooms.length != 0) {
        this.search_res = true;
      } else {
        this.search_res = false;
      }
    },
  },
  methods: {
    search() {
      let params = new URLSearchParams();
      params.append("name", this.name);
      params.append("rid", this.rid);
      post("/searchRoom", params).then((res) => {
        if (res.data.rooms != null) {
          this.rooms = res.data.rooms;
        }
        else {
          this.rooms = []
        }
      });
    },
    selectRoom(rid) {
      this.selected_rid = rid
      this.show = true
    }
  },
};
</script>

<style>
.-margin-top {
  margin-top: 40px;
}
.-margin-bottom {
  margin-bottom: 40px;
}
</style>