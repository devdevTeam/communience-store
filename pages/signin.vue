<template>
  <div>
    <v-container fluid>
      <v-row>
        <v-col cols=12 class="col-lg-4" />
        <v-col cols=12 class="col-lg-4">
          <v-card class="my-15">
            <v-tabs
            v-model="tab"
            background-color="indigo"
            dark
            left>
              <v-tab v-for="item in tabItems" :key="item">
                {{item}}
              </v-tab>
            </v-tabs>
            <v-tabs-items v-model="tab">
              <div v-if="tabItems[tab] == 'signin'">
                <signin-form />
              </div>
              <div v-else-if="tabItems[tab] == 'signup'">
                <signup-form />
              </div>
            </v-tabs-items>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import SignupForm from '@/components/signup/form.vue'
import SigninForm from '@/components/signin/form.vue'

export default {
  components: {
    'signup-form': SignupForm,
    'signin-form': SigninForm
  },
  data() {
    return {
      tab: null,
      tabItems: [
        'signin',
        'signup'
      ],
      progressCircle: false
    }
  }
}
</script>