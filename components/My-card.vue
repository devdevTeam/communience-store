
<template>
  <v-app>
      <div class="d-flex justify-center align-center column fill-height mt-3 mb-3">
        <v-card
          class="mx-auto mt-20"
          max-width="450"
        >
          <v-img class="visual"
            src="https://cdn.vuetifyjs.com/images/cards/house.jpg"
            :aspect-ratio="16/9"
          >
            <v-fab-transition
             >
                <v-btn
                  color="blue"
                  fab
                  dark
                  small
                  bottom
                  left
                  class="mt-3 ml-3"
                  to = "/edit"
                >
                  <v-icon>mdi-pencil</v-icon>
                </v-btn>
              </v-fab-transition>
          </v-img>

          <div class="details">
              <v-card-title class = "d-flex flex-column align-center justify-space-between">
                <div class="d-flex justify-between">
                <div class="text-h4 mb-2">
                  {{info.name}}
                </div>
                </div>
                <div class="d-flex justify-space-between">
                    <div class="text-h6 font-weight-regular grey--text">
                      {{info.hurigana}}
                    </div> 
                </div>
              </v-card-title>
              <div
                class="px-4 grey--text"

              >
                <div class="introduction">
                  <h3 style="color: white">【自己紹介】</h3>
                  <p style="margin-left: 15px; margin-right: 15px">{{info.free}}</p>
                </div>
              </div>
              <v-divider class="mt-6 mx-4"></v-divider>
              <v-card-text>
                <v-chip
                  class="mr-2 mb-2" @click="toLinck(info.twitter)"
                >
                  <v-icon left>
                    mdi-twitter
                  </v-icon>
                  Twitter
                </v-chip>
                <v-chip class="mr-2 mb-2" @click="toLinck(info.instagram)">
                  <v-icon left>
                    mdi-instagram
                  </v-icon>
                  Instagram
                </v-chip>
                <v-chip class="mr-2 mb-2" @click="toLinck(info.facebook)"
                >
                  <v-icon left>
                    mdi-facebook
                  </v-icon>
                  Facebook
                </v-chip>
                <v-list>
                    <v-list-group
                      :prepend-icon="'mdi-account-heart-outline'"
                      no-action
                    >
                      <template v-slot:activator>
                        <v-list-item-content>
                          <v-list-item-title v-text="'趣味'"></v-list-item-title>
                        </v-list-item-content>
                      </template>
                      <v-list-item
                        v-for="item, i in hobby"
                        :key="i"
                      >
                        <v-list-item-title v-text="item"></v-list-item-title>
                      </v-list-item>
                    </v-list-group>
                    <v-list-group
                      :prepend-icon="'mdi-account-supervisor-outline'"
                      no-action
                    >
                      <template v-slot:activator>
                        <v-list-item-content>
                          <v-list-item-title v-text="'親しい友だち'"></v-list-item-title>
                        </v-list-item-content>
                      </template>
                      <v-list-item
                        v-for="item, i in friend"
                        :key="i"
                      >
                        <v-list-item-title v-text="item"></v-list-item-title>
                      </v-list-item>
                    </v-list-group>
                </v-list>
              </v-card-text>
          </div>
        </v-card>
      </div>
  </v-app>
</template>

<script>
  export default {
    props: ["info", "hobby", "friend"],
    methods: {  
      toLinck(url){
      window.open(url, '_blank')
      },
    },
  }
</script>

<style>
  .footer{
    font-size: 0.8rem;
    font-weight: light;
  }
  .introduction{
    width: 100%;
  }
  .mdi-twitter::before{
    color: #2BA4F2;
  }
  .mdi-instagram::before{
    color: #FE3972;
  }
  .mdi-facebook::before{
    color: #237CF3;
  }

</style>