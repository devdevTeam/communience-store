<template>
  <div>
    <v-container fluid>
      <v-row>
        <v-col cols=12>
          <v-text-field
          label="user name"
          v-model="signupInfo.username"
          :rules="[requireRule]" />
        </v-col>

        <v-col cols=12>
          <v-text-field
          label="mail"
          v-model="signupInfo.mail"
          :rules="[requireRule, mailRule]" />
        </v-col>

        <v-col cols=12>
          <v-text-field
          label="password"
          v-model="signupInfo.password"
          type="password"
          :rules="[requireRule, passwordRule]" />
        </v-col>
      </v-row>
      <v-row>
        <v-col cols=12 class="text-right">
          <v-btn color="primary" @click="send">送信</v-btn>
        </v-col>
      </v-row>
      <div v-if="!progressCircle" />
      <div v-else align="center" style="vertical-align: center;">
        <v-overlay>
          <v-progress-circular
            :size="100"
            :width="10"
            absolute
            color="white"
            indeterminate
          ></v-progress-circular>
        </v-overlay>
      </div>
      <err-dialog :view="view" :title="error.title" :msg="error.msg" @updateView="updateView" />
    </v-container>
  </div>
</template>

<script>
import dialog from '@/components/dialog.vue'
import post from '@/lib/post.js'
import errorResult from '@/lib/authResult.js'

export default {
  components: {
    'err-dialog': dialog
  },
  data() {
    return {
      signupInfo: {
        username: '',
        mail: '',
        password: '',
      },
      error: {
        title: '',
        msg: ''
      },
      view: false,
      progressCircle: false,
      requireRule: value => !!value || "入力してください",
      passwordRule: value => value.length >= 8 || "パスワードは8文字以上必要です．",
      mailRule: value => {
        const pattern = /^(([^<>()[\]\\.,;:\s@]+(\.[^<>()[\]\\.,;:\s@]+)*)|(.+))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        return pattern.test(value) || '不正なメールアドレスです．'
      },
    }
  },
  methods: {
    async send() {
      let params = new URLSearchParams()
      params.append("mail", this.signupInfo.mail)
      params.append("password", this.signupInfo.password)
      params.append("username", this.signupInfo.username,)
      this.progressCircle = true
      let res = await post("/signup", params)
      if (res.data.error !== null) {
        console.error(res.data.error)
        this.progressCircle = false
        this.view = true
        this.error.title = "ユーザ登録に失敗しました"
        this.error.msg = "既に登録されたメールです。"
      }
      else {
        let error = null
        await this.$fire.auth.signInWithEmailAndPassword(this.signupInfo.mail, this.signupInfo.password)
        .catch((err) => error = err.code)
        this.progressCircle = false
        if (error !== null) {
          error = errorResult(error)
          this.view = true
          this.error.title = error.title
          this.error.msg = error.msg
        }
        else {
          if (this.$nuxt.context.from.path.match("/joinRoom")) {
            setTimeout(() => {
              this.$router.go(-1)
            }, 500)
          }
          this.$router.push({ path: '/room' })
        }
      }
    },
    updateView(nv) {
      this.view = nv
    }
  }
}
</script>

<style>

</style>