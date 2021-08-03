<template>
  <div>
      <br><br><br><br><br>
      <img :src="$auth.user.picture" class="mb-6">
      <div class="columns is-variable is-0-mobile is-1-tablet is-3-desktop">   
        <div class="column" @click="test()">
            <h1 class="heading is-size-4 is-pill">
                Username: <br>{{this.$auth.user.nickname}}
            </h1>
        </div>
        <div class="column">
            <h1 class="heading is-size-4 is-pill">
                Email verified: <br>{{this.$auth.user.email_verified}}
            </h1>
        </div>
    </div>
    <div class="column ">
            <h1 class="heading is-size-4 is-pill is-extra-pill">
                Email: <br>{{this.$auth.user.email}}
            </h1>
        </div>
  </div>
</template>

<script>
export default {
    methods: {
        test () {
            console.log(this.$auth.user)
        }
    }
}
</script>
