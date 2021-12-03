<template>
  <v-container fluid>
    <v-row>
      <v-col cols="12" class="col-md-8" offset-md="2">
        <v-card light>
          <v-toolbar flat>
            <template v-slot:extension>
              <v-tabs v-model="tabs" fixed-tabs light>
                <v-tabs-slider></v-tabs-slider>
                <v-tab class="primary--text"><h3>URLで参加</h3></v-tab>

                <v-tab class="primary--text"><h3>Roomを検索</h3></v-tab>
              </v-tabs>
            </template>
          </v-toolbar>

          <v-tabs-items v-model="tabs">
            <v-tab-item>
              <v-card light>
                <InputURL></InputURL>
              </v-card>
            </v-tab-item>
            <v-tab-item>
              <v-card light>
                <SearchRoom></SearchRoom>
              </v-card>
            </v-tab-item>
          </v-tabs-items>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import InputURL from '../../components/Room/inputURL.vue';
import SearchRoom from '../../components/Room/searchRoom.vue'

export default {
  components: {
    InputURL,
    SearchRoom
  },
  data() {
    return {
      tabs: null,
    };
  },
};
</script>