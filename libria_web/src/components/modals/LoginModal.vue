<template>
  <div class="modal is-active">
  <div class="modal-background"></div>
  <div class="modal-content is-login-modal">
      <h1 class="heading is-size-4 mb-4 mt-5">
          You need to be logged in for that.
      </h1>
      <span v-if="!$auth.loading">
        <!-- show login when not authenticated -->
        <button
          v-if="!$auth.isAuthenticated"
          @click="login"
          class="button is-log-in-button is-outlined is-dark mr-3 mb-4"
        >
          Log in
        </button>
        <button
          v-if="!$auth.isAuthenticated"
          @click="register"
          class="button is-log-in-button is-danger"
        >
          Register
        </button>
      </span>
      <br><br>
      <button
        class="button is-danger is-outlined mx-2 "
        @click="closeModal()"
      >
        <span>
            Cancel
        </span>
        <span class="icon is-small">
            <i class="fas fa-times"></i>
        </span>
    </button>
    <br><br>
  </div>
  <button class="modal-close is-large" aria-label="close" @click="closeModal()"></button>
</div>
</template>

<script>
export default {
    methods: {
        closeModal () {
            this.$store.state.login_modal = false;
        },
        login() {
          this.$auth.loginWithRedirect();
        },
        register() {
          this.$auth.loginWithRedirect({
            screen_hint: 'signup'
          });
        },
    },
}
</script>

<style>
    .modal-content {
        background-color: white;
        color: black;
    }
</style>