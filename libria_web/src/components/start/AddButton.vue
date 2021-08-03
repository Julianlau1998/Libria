<template>
  <div>
      <button
        @click="test()"
        class="button is-circle is-outlined is-add-button"
      >
          <i class="fas fa-plus"></i>
      </button>
  </div>
</template>

<script>
export default {
    name: 'AddButton',
    methods: {
        test () {
            if (this.$auth.isAuthenticated) { 
                this.$emit('openNewModal')
            } else {
                this.$store.state.login_modal = true
            }
        }
    }
}
</script>

<style>
</style>