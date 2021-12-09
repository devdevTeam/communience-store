
<template>
  <div class="d-flex justify-center align-center column fill-height mt-3 mb-3">
    <v-card
      class="mx-auto mt-20"
      width="500px"
      max-width="600"
    >
        <v-fab-transition
          >
            <v-btn
              v-if="!edit"
              color="blue"
              fab
              dark
              small
              bottom
              left
              class="mt-3 ml-3"
              @click="edit=!edit"
            >
              <v-icon>mdi-pencil</v-icon>
            </v-btn>
          </v-fab-transition>

      <div class="details">
        <v-card-title class = "d-flex flex-column align-center justify-space-between">
          <div class="d-flex justify-between">
            <div class="text-h4 mb-2" v-if="!edit">
              {{info.name}}
            </div>
            <v-text-field
              class="text-h4 mb-2"
              v-else
              v-model="name"
              label="name"
              outlined
            ></v-text-field>
          </div>
          <div class="d-flex justify-space-between">
              <div class="text-h6 font-weight-regular grey--text" v-if="!edit">
                {{info.hurigana}}
              </div>
              <v-text-field
                class="text-h6"
                v-else
                v-model="hurigana"
                label="hurigana"
                dense
                outlined
              ></v-text-field>
          </div>
        </v-card-title>
        <div
          class="px-4 grey--text"

        >
          <div class="introduction">
            <h3 style="color: white">【自己紹介】</h3>
            <p style="margin-left: 15px; margin-right: 15px" v-if="!edit">{{info.free}}</p>
            <v-textarea
              v-else
              v-model="free"
              :counter="120"
              outlined
              label="free area"
            ></v-textarea>
          </div>
        </div>
        <v-divider class="mt-6 mx-4"></v-divider>
        <v-card-text>
          <v-list>
            <v-list-group
              :prepend-icon="'mdi-cake'"
              no-action
            >
              <template v-slot:activator>
                <v-list-item-content>
                  <v-list-item-title v-text="'生年月日'"></v-list-item-title>
                </v-list-item-content>
              </template>
              <v-list-item>
                <v-list-item-title v-if="!edit" v-text="info.birthday"></v-list-item-title>
                <v-list-item-title v-else>
                  <v-text-field
                    v-model="birthday"
                    label="birthday"
                  ></v-text-field>
                </v-list-item-title>
              </v-list-item>
            </v-list-group>
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
                <v-list-item-title v-if="!edit" v-text="item"></v-list-item-title>
                <v-list-item-title v-else>
                  <v-text-field
                    v-model="hobby[i]"
                    label="hobby"
                    append-outer-icon="mdi-close-box-outline"
                    @click:append-outer="removeHobby(i)"
                  ></v-text-field>
                </v-list-item-title>
              </v-list-item>
              <v-list-item v-if="edit">
                <v-row justify="center" align="center">
                  <v-btn text icon @click="addHobby">
                    <v-icon>mdi-plus</v-icon>
                  </v-btn>
                </v-row>
              </v-list-item>
            </v-list-group>
          </v-list>
          <div v-if="!edit" style="text-align: center;">
            <v-chip class="mr-2 mb-2" @click="toLinck(info.twitter)">
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
            <v-chip class="mr-2 mb-2" @click="toLinck(info.facebook)">
              <v-icon left>
                mdi-facebook
              </v-icon>
              Facebook
            </v-chip>
          </div>
          <div v-else>
            <v-text-field
              v-model="twitter"
              prepend-inner-icon=mdi-twitter
              label="Twitter"
            ></v-text-field>
            <v-text-field
              v-model="instagram"
              prepend-inner-icon=mdi-instagram
              label="instagram"
            ></v-text-field>
            <v-text-field
              v-model="facebook"
              prepend-inner-icon=mdi-facebook
              label="Facebook"
            ></v-text-field>
          </div>
          <div v-if="edit" style="margin: 20px 0px 20px 0px;">
            <v-row justify="center" align="center">
              <v-btn 
                color="primary"
                style="margin-right: 20px" 
                @click="update"
              >
                update
              </v-btn>
              <v-btn color="pink lighten-1" @click="cancel">cancel</v-btn>
            </v-row>
          </div>
        </v-card-text>
      </div>
    </v-card>
  </div>
</template>

<script>
import post from "@/lib/post.js"
  export default {
    props: ["info", "hobby", "friend"],
    data() {
      return {
        edit: false,
        name: this.info.name,
        hurigana: this.info.hurigana,
        free: this.info.free,
        birthday: this.info.birthday,
        twitter: this.info.twitter,
        instagram: this.info.instagram,
        facebook: this.info.facebook,
      }
    },
    created() {
      this.hobby = ["サッカー","ゲーム"]
    },
    methods: {  
      toLinck(url){
        window.open(url, '_blank')
      },
      addHobby() {
        this.hobby.push("")
      },
      removeHobby(i) {
        this.hobby.splice(i, 1)
      },
      update() {
        const Shobby = this.hobby.join(",")
        let params = new URLSearchParams()
        params.append("uid", this.$store.getters.getUser.uid)
        params.append("name", this.name)
        params.append("hurigana", this.hurigana)
        params.append("birthday", this.birthday)
        params.append("instagram", this.instagram)
        params.append("twitter", this.twitter)
        params.append("facebook", this.facebook)
        params.append("free", this.free)
        params.append("hobby", Shobby)
        post("/updateDefaultCard", params).then((res) => {
          if (res.data.error != null) {
            console.error(res.data.error);
          }
          this.edit = false
        })
      },
      cancel() {
        this.edit = false
        this.name = this.info.name
        this.hurigana = this.info.hurigana
        this.free = this.info.free
        this.birthday = this.info.birthday
        this.twitter = this.info.twitter
        this.instagram = this.info.instagram
        this.facebook = this.info.facebook
      }
    },
  }
</script>

<style>
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