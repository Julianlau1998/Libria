<template>
    <button @click="openRandomTopic()" class="button is-outlined is-warning mt-4">
        {{buttonText}}
        <i class="fas fa-arrow-right ml-2"></i>
    </button>
</template>

<script>
import { mapState } from 'vuex'

export default {
    name: 'RandomTopic',
    props: {
        buttonText: {
            type: String,
            required: true
        }
    },
    created () {
        this.$store.dispatch('topicModule/getRandom')
    },
    computed: {
      ...mapState(['topicModule']),
      randomTopic () {
        return (this.topicModule.randomTopic.data) || {}
      }
    },
    methods: {
        openRandomTopic () {
            this.$store.dispatch('topicModule/getRandom')
            if (this.$route.params.uuid) {
                this.topicModule.topic.data = this.randomTopic
                this.$store.dispatch('answerModule/getAllByTopic', {topicId: this.randomTopic.id})
            } else {
                this.$router.push(
                    {
                        name: 'Topic', 
                        params: {uuid: this.randomTopic.id}
                    }
                )
            }
        }
    }
}
</script>