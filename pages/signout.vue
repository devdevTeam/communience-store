<template>
  <h1>ログアウトしています</h1>
</template>

<script>
export default {
  asyncData() {
    $nuxt.$fire.auth.signOut()
  }
}
</script>