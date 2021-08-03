<template>
  <div class="mb-8">
    <h1 class="is-size-3 mt-5 mb-5">
          Feed
    </h1>
    <div
        class="columns is-variable is-0-mobile is-1-tablet is-3-desktop"
        v-for="(topic, index) in topics"
        :key="index"
    >   
        <div class="column" v-if="topic.length>1">
            <card 
                :card="topic[0]"
                :type="'topic'"
                :cardNumber="Math.floor(Math.random() * 4) + 1"
                @openTopic="openTopic($event)"
                @openEditModal="openEditModal($event)"
             />
        </div>
        <div class="column" v-if="topic.length>1">
            <card
                :card="topic[1]"
                :type="'topic'"
                :cardNumber="Math.floor(Math.random() * 4) + 1"
                @openTopic="openTopic($event)"
                @openEditModal="openEditModal($event)"
            />
        </div>
        <div class="column" v-if="topic.length>2">
            <card
                :card="topic[2]"
                :type="'topic'"
                :cardNumber="Math.floor(Math.random() * 4) + 1"
                @openTopic="openTopic($event)"
                @openEditModal="openEditModal($event)"
            />
        </div>
        <div class="singleColumn" v-if="topic.length===1">
            <card 
                :card="topic[0]"
                :type="'topic'"
                :cardNumber="Math.floor(Math.random() * 4) + 1"
                @openTopic="openTopic($event)"
                @openEditModal="openEditModal($event)"
             />
        </div>
    </div>
  </div>
</template>

<script>
import card from '@/components/start/Card.vue'

const chunk = require('chunk')
export default {
    name: "startFeed",
    components: {
        card
    },
    props: {
        topics: {
            type: Array,
            required: true,
        },
        foundTopics: {
            type: Array,
            required: false
        }
    },
    data () {
        return {
        }
    },
    created () {
        this.$store.dispatch('topicModule/getAll')
},
    watch: {
        foundTopics (val) {
            this.topics = chunk(val, 3)
        },
    },
    methods: {
        openTopic (topicUuid) {
            this.$router.push(
                {
                    name: 'Topic', 
                    params: {uuid: topicUuid}
                }
            )
        },
        openEditModal (topic) {
            this.$emit('openEditModal', topic)
        }
    }
}
</script>

<style>

</style>