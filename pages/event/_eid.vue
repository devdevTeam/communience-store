<template>
  <v-main>
    <organizer v-if="admin" :items=eventCol :name=$store.getters.getEventInfo.name />
    <participant v-else :items=eventCol :name=$store.getters.getEventInfo.name />
  </v-main>
</template>

<script>
import organizer from "@/components/EventPages/Organizer";
import participant from "@/components/EventPages/Participant";

export default {
  name: "Organizer",

  components: {
    organizer,
    participant,
  },
  data() {
    return {
      admin: false,
      items: [],
      ws: this.$store.getters.getEventWs,
    };
  },
  created() {
    if (this.$store.getters.getUser.uid === this.$store.getters.getEventInfo.org_uid) {
      this.admin = true 
    }
    let msg = JSON.stringify({
      uid: this.$store.getters.getUser.uid,
      eid: this.$route.params.eid,
      hidden: [],
    });
    this.ws.send(msg);
  },
  computed: {
    eventCol() {
      this.ws.onmessage = (msg) => {
        this.items = JSON.parse(msg.data).colList;
      };
      return this.items
    },
  },
};
</script>