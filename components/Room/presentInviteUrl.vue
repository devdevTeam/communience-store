<template>
  <div id="overlay">
    <div id="content">
      <h1 style="text-align: center" class="-color-black">招待URL</h1>
      <h5 style="text-align: center" class="-color-black">
        以下のURLを招待する人に共有してください
      </h5>
      <v-row justify="center" align-content="center" class="-margin-top">
        <v-col cols="12" md="10" offset-md="0">
          <v-text-field
            v-model="url"
            label="URL"
            outlined
            light
            color="teal darken-2"
            background-color="deep-purple lighten-5"
          ></v-text-field>
        </v-col>
        <v-col cols="12" md="1">
          <v-btn class="copy-btn" large icon light @click="copy">
            <v-slide-y-transition
              mode="out-in"
            >
              <v-icon
                :key="`icon-${copied}`"
                :color="copied ? 'success' : 'primary'"
                @click="copy"
                v-text="copied ? 'mdi-check-outline' : 'mdi-checkbox-multiple-blank-outline'"
                large
              ></v-icon>
            </v-slide-y-transition>
          </v-btn>
        </v-col>
      </v-row>
      <v-row style="margin-top: -15px;" justify="center" align-content="center">
        <v-btn color="primary" @click="$emit('closeModal')">close</v-btn>
      </v-row>
    </div>
  </div>
</template>

<script>
import post from "@/lib/post.js";

export default {
  props: ["rid"],
  data() {
    return {
      copied: false,
      origin: "",
      hash: "",
      url: "",
    };
  },
  async created() {
    let origin = window.location.href.split("/room/")[0];
    let hash = null;
    let params = new URLSearchParams();
    params.append("rid", this.rid);
    await post("/getRoomInfo", params).then((res) => {
      hash = res.data.hash;
    });
    this.origin = origin;
    this.hash = hash;
    this.url = `${origin}/joinRoom/${hash}`;
  },
  watch: {
    url() {
      this.url = `${this.origin}/joinRoom/${this.hash}`;
    },
  },
  methods: {
    copy() {
      this.url = `${this.origin}/joinRoom/${this.hash}`;
      navigator.clipboard.writeText(this.url);
      this.copied = true
      setTimeout(() => {this.copied = false}, 1000)
    },
  },
};
</script>


<style scoped lang="scss">
.-margin-top {
  margin-top: 10px;
}
.copy-btn {
  margin-top: 10px;
}
.-color-black {
  color: black;
}
#content {
  z-index: 10;
  width: 580px;
  height: 250px;
  padding: 1em;
  background: white;
}
#overlay {
  z-index: 4;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
}
#submit {
  margin-right: 20px;
}
</style>