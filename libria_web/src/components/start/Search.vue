<template>
  <div class="control has-icons-right">
      <input
        v-model="seachText"
        class="input is-info mt-4"
        type="text"
        placeholder="Search Topic"
        size="1"
      >
      <span class="icon is-small is-right mt-4 is-search-icon">
          <i class="fas fa-search"></i>
      </span>
      <br>
      <RandomTopic :buttonText="'Or Get Random Topic'" />
  </div>
</template>

<script>
import RandomTopic from '@/components/start/RandomTopic.vue'

export default {
    name: 'search',
    components: { RandomTopic },
    props: {
      topics: {
          type: Array,
          required: true,
      }
    },
    data () {
      return {
        seachText: '',
        topicsList: []
      }
    },
    watch: {
      seachText (val) { 
        let foundTopics = []
        if (val !== '') {
          foundTopics = this.topicsList.filter(el => el.title.toLowerCase().includes(val.toLowerCase()))
        } else {
          foundTopics = this.topicsList
        }
        this.$emit('topicsFound', foundTopics)
      },
      topics () {
        this.topicsList = []
        this.topics.forEach(el => {
          el.forEach(topic => {
            this.topicsList.push(topic)
          })
        });
      }
    }
}
</script>