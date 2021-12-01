<template>
  <div>
    <organizer v-if="admin" :items=eventCol />
    <participant v-else :items=eventCol />
  </div>
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
    if (this.$store.getters.getUser.uid === this.$store.getters.getEventOrg) {
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