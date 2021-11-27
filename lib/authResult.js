const authResult = {
  errors: [
    {
      code: 'auth/user-not-found',
      title: 'ユーザが存在しません．',
      msg: 'メールアドレスが正しいか確認してください'
    },
    {
      code: 'auth/wrong-password',
      title: 'メールアドレスとパスワードの組み合わせが存在しません',
      msg: '認証情報を確認してください'
    }
  ]
}

export default function errorResult(code) {
  for(let i = 0; i < authResult.errors.length; i++) {
    if (authResult.errors[i].code === code) {
      return authResult.errors[i]
    }
  }
  return null
}

