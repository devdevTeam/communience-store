<template>
  <v-row justify="center" align="center">
    <v-col cols="12" md="10" offset-md="1">
      <v-form ref="form" v-model="valid" lazy-validation>
        <v-text-field
          v-model="URL"
          label="参加用URLを入力"
          required
          class="-margin-top"
        ></v-text-field>
        <div class="text-center">
          <v-btn
            center
            :disabled="valid"
            color="primary"
            class="-margin-bottom"
            @click="submitUrl"
          >
            submit
          </v-btn>
        </div>
      </v-form>
    </v-col>
    <v-col cols='12' md='1'></v-col>
  </v-row>
</template>

<script>
export default {
  data() {
    return {
      valid: true,
      URL: null,
    };
  },
  watch: {
    URL() {
      if (this.URL.length != 0) {
        this.valid = false
      }
      else {
        this.valid = true
      }
    },
  },
  methods: {
    submitUrl() {
      let uri = this.URL.split("/joinRoom/")[1]
      if (uri == undefined) {
        this.$router.push({ name: "joinRoom-uri", params: { uri: "faild" }})
      }
      else {
        this.$router.push({ name: "joinRoom-uri", params: { uri: uri }})
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
</style>