<template>
  <div class="container mb-8">
        <Topic v-if="!$auth.loading || $auth === undefined" />
  </div>
</template>

<script>
import Topic from '@/components/topics/Topic.vue'

export default {
    components: {
        Topic
    }
}
</script>

<style>

</style>