<template>
  <v-main>
    <v-container fluid fill-height class="grey lighten-5">
      <v-row justify="center" align-content="center" v-for="col, i in items" :key="i">
        <v-col md="4" offset-md="0" align-self="center">
          <v-list-item-title class="-color-black">{{col.colName}}</v-list-item-title>
        </v-col>
        <v-col md="4" offset-md="0" align-self="center">
          <v-list-item-title v-if="!col.hidden" class="-color-black">内容</v-list-item-title>
          <v-btn v-else color="primary" @click="submit(i)">解答</v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
export default {
  props: ['items'],
  methods: {
    submit(i) {
      this.items[i].hidden = !this.items[i].hidden
    }
  }
};
</script>

<style>
.-color-black {
  color: black;
}
</style>