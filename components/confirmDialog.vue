<template>
  <div class="text-center">
    <v-dialog v-model="dialog" width="400">
      <v-container style="background-color: #455A64;">
        <v-row justify="center" align-content="center">
          <h2 class="-margin-top" style="color: white;">{{ text }}</h2>
        </v-row>
        <v-row justify="center" align-content="center">
          <v-btn class="yes-btn -margin-top -margin-bottom" color="blue lighten-1" outlined @click="closeYes"> Yes </v-btn>
          <v-btn class="-margin-top -margin-bottom" color="pink accent-3" outlined @click="closeNo"> Cnacel </v-btn>
        </v-row>
      </v-container>
    </v-dialog>
  </div>
</template>

<script>
export default {
  props: ["text", "dialog"],
  methods: {
    closeYes() {
      this.$emit("closeConfirmYes");
    },
    closeNo() {
      this.$emit("closeConfirmNo");
    },
  },
};
</script>

<style>
.-margin-top {
  margin-top: 30px;
}
.-margin-bottom {
  margin-bottom: 30px;
}
.yes-btn {
  margin-right: 20px;
}
</style>>
