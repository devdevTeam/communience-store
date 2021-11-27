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
      <err-dialog :view="view" :title="error.title" :msg="error.msg" @updateView="updateView" />
    </v-container>
  </div>
</template>

<script>
import dialog from '@/components/dialog.vue'

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
        phoneNumber: ''
      },
      error: {
        title: '',
        msg: ''
      },
      view: false,
      requireRule: value => !!value || "入力してください",
      passwordRule: value => value.length >= 8 || "パスワードは8文字以上必要です．",
      mailRule: value => {
        const pattern = /^(([^<>()[\]\\.,;:\s@]+(\.[^<>()[\]\\.,;:\s@]+)*)|(.+))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        return pattern.test(value) || '不正なメールアドレスです．'
      },
    }
  },
  methods: {
    send() {
      // let auth = new Authentication()
      // auth.setSignupInfo(this.signupInfo.username, this.signupInfo.mail, this.signupInfo.password, this.signupInfo.phoneNumber)
      // auth.sendSignup().then(() => {
      //   location.reload()
      // }).catch(() => {
      //   this.view = true
      //   this.error.title = "ユーザ登録に失敗しました"
      //   this.error.msg = "既に登録されたメール、電話番号です。"
      // })
    },
    updateView(nv) {
      this.view = nv
    }
  }
}
</script>

<style>

</style>