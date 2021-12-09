<template>
  <div id="overlay">
    <div id="content">
      <v-row justify="center" align="center">
        <v-col cols="12" md="10" offset-md="1">
          <v-form ref="form" v-model="valid" lazy-validation>
            <v-text-field
              v-model="URL"
              label="フレンドのURLを入力"
              required
              light
              class="-margin-top"
            ></v-text-field>
            <div class="text-center">
              <v-btn
                id="submit"
                center
                color="primary"
                light
                class="-margin-bottom"
                @click="submitUrl"
              >
                submit
              </v-btn>
              <v-btn
                center
                color="orange"
                light
                class="-margin-bottom"
                @click="$emit('closeInput')"
              >
                close
              </v-btn>
            </div>
          </v-form>
        </v-col>
        <v-col cols='12' md='1'></v-col>
      </v-row>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      URL: null,
    };
  },
  methods: {
    submitUrl() {
      let uri = this.URL.split("/makeFriend/")[1]
      if (uri == undefined) {
        this.$router.push({ name: "makeFriend-uri", params: { uri: "faild" }})
      }
      else {
        this.$router.push(`/makeFriend/${uri}`)
      }
    }
  }
};
</script>

<style>
.-margin-top {
  margin-top: 40px;
}
.-margin-bottom {
  margin-bottom: 40px;
}
.-color-black {
  color: black;
}
#content {
  z-index: 10;
  width: 580px;
  height: 200px;
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