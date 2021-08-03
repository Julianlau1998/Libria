<template>
  <div id="app" class="app">
    <navigation class="navbar" />
    <transition name="fade" mode="out-in">
      <router-view/>
    </transition>
    <LoginModal v-if="loginModal" />
  </div>
</template>

<script>
import navigation from '@/components/navigation/Navbar.vue'
import LoginModal from '@/components/modals/LoginModal.vue'
export default {
  components: {
    navigation,
    LoginModal
  },
  computed: {
    loginModal () {
      return this.$store.state.login_modal
    }
  }
  // async created () {
  //   setInterval(async () => {
  //     if (this.$auth.isAuthenticated) {
  //       await this.$auth.getTokenSilently()
  //       window.localStorage.setItem('userId', this.$auth.user.sub)
  //     }
  //   }, 36000)
  // }
}
</script>

<style lang="scss">
@import './assets/styles/variables.scss';

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: white;
}
body {
  background-color: $mainBackground;
}
</style>
