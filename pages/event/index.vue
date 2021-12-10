<template>
  <v-main>
    <v-container class="grey lighten-5">
      <v-row>
        <v-col md="5" offset-md="3" align-self="center">
          <p style="font-size: 1.9em; text-align: center; color: black">
            開催中イベント一覧
          </p>
        </v-col>
        <v-col md="2" offset-md="2">
          <v-btn class="mb-2" rounded dark color="indigo" to="/startEvent">
            イベントを開催
          </v-btn>
        </v-col>
      </v-row>
      <v-row v-for="(two_events, i) in eventList" :key="i">
        <v-col md="6" v-for="(event, j) in two_events" :key="j">
          <v-btn
            block
            color="black"
            elevation="2"
            class="pa-2"
            x-large
            outlined
            tile
            @click="selectEvent(event)"
          >
            {{ event.name }}
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
    <joinWithPass
      @closeModal="show = false"
      v-if="show"
      :prop_params="params"
    ></joinWithPass>
  </v-main>
</template>

<script>
import post from "@/lib/post.js";
import joinWithPass from "@/components/EventPages/JoinWithPass.vue";

export default {
  components: {
    joinWithPass,
  },
  data() {
    return {
      eventList: [],
      params: {
        uid: this.$store.getters.getUser.uid,
        eid: null,
        org_uid: null,
        name: null,
        rid: null,
      },
      show: false,
    };
  },
  beforeCreate() {
    let params = new URLSearchParams();
    params.append("uid", this.$store.getters.getUser.uid);
    post("/getEventList", params).then((res) => {
      for (let i = 0; i < Math.ceil(res.data.events.length / 2); i++) {
        let multiple_cnt = i * 2;
        let result = res.data.events.slice(multiple_cnt, multiple_cnt + 2);
        this.eventList.push(result);
      }
    });
  },
  methods: {
    selectEvent(event) {
      this.params.eid = event.eid
      this.params.org_uid = event.org_uid
      this.params.name = event.name
      this.params.rid = event.rid
      this.show = true
    },
  },
};
</script>