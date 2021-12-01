<template>
  <v-main>
    <v-container fluid fill-height>
      <v-row justify="center" align-content="center">
        <v-col cols="12" sm="7" offset-sm="0">
          <v-card>
            <v-list-item-title>col list</v-list-item-title>
            <v-list-item-title v-for="col, i in items" :key="i">{{col.colName}}{{col.hidden}}</v-list-item-title>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
export default {
  props: ['items'],
  created() {
    console.log(this.items)
  }
};
</script>