<template>
  <div>
    <v-dialog v-model="viewUpdater">
      <v-card>
        <v-card-title>
          {{title}}
        </v-card-title>
        <v-card-text>
          {{msg}}
        </v-card-text>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
export default {
  props: ["view", "title", "msg"],
  data() {
    return {
      dialogView: false
    }
  },
  watch: {
    view() {
      this.viewUpdater = this.view
    },
    dialogView() {
      this.viewUpdater = this.dialogView
    }
  },
  computed: {
    viewUpdater: {
      get() {
        return this.dialogView
      },
      set(value) {
        this.dialogView = value
        this.$emit('updateView', value)
      }
    }
  }
}
</script>