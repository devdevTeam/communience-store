<template>
  <v-container class="grey lighten-5">
    <v-row>
      <v-col md="4" offset-md="4" align-self="center">
        <p style="font-size: 1.9em; text-align: center; color: black;">開催中イベント一覧</p>
      </v-col>
      <v-col md="2" offset-md="2">
        <v-btn class="mb-2" rounded dark color="indigo" to="/roomlist?onlyowner=true">
          イベントを開催
        </v-btn>
      </v-col>
    </v-row>
    <v-row v-for="two_events, i in eventList" :key="i">
      <v-col md="6" v-for="event, j in two_events" :key="j">
        <v-btn block color="black" elevation="2" class="pa-2" x-large outlined tile @click="joinEvent(event.eid, event.org_uid)">
          {{ event.name }}
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import post from '@/lib/post.js'

export default {
  data() {
    return {
      eventList: []
    }
  },
  beforeCreate() {
    let params = new URLSearchParams()
    params.append("uid", this.$store.getters.getUser.uid)
    post("/getEventList", params)
    .then((res) => {
      for (let i = 0; i < Math.ceil(res.data.events.length / 2); i++) {
        let multiple_cnt = i * 2
        let result = res.data.events.slice(multiple_cnt, multiple_cnt + 2) 
        this.eventList.push(result)
      }
    })
  },
  methods: {
    joinEvent(eid, org_uid) {
      let params = new URLSearchParams()
      params.append("uid", this.$store.getters.getUser.uid)
      params.append("eid", eid)
      params.append("password", "1111")
      post("/event/joinEvent", params)
      .then((res) => {
        if (res.data.error != null) {
          console.log(error)
          return
        }
        this.$store.commit("set_event_org", org_uid)
        this.$store.commit("connect_event_ws", eid)
        this.$router.push({ name: 'event-eid', params: { eid : eid } })
      })
    }
  }
}
</script>