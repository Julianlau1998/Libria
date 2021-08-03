<template>
  <div class="container">
    <h2 class="is-size-2 mt-6">
        Welcome
    </h2>
    <h3 class="is-size-6">
      Find quick and understandable information on any topic you'd like.
    </h3>
    <Search @topicsFound="topicsFound($event)" :topics="topics" />
    <StartFeed :topics="topics" :foundTopics="foundTopics" @openEditModal="openEditModal($event)"/>
    <AddButton
      @openNewModal="openNewModal()"
     />
    <TopicModal
      :topicToEdit="topicToEdit"
      :edit="edit"
      @closeModal="closeModal()"
      v-if="openTopicModal"
    />
  </div>
</template>

<script>
import Search from '@/components/start/Search.vue'
import StartFeed from '@/components/start/StartFeed.vue'
import AddButton from '@/components/start/AddButton.vue'
import TopicModal from '@/components/modals/TopicModal.vue'
import { mapState } from 'vuex'

const chunk = require('chunk')

export default {
  name: 'Home',
  components: {
    Search,
    StartFeed,
    AddButton,
    TopicModal
  },
  data () {
    return {
      openTopicModal: false,
      edit: false,
      topicToEdit: {},
      foundTopics: []
    }
  },
  computed: {
    ...mapState(['topicModule']),
    topics () {
      let topics = (!this.topicModule.topics.loading && this.topicModule.topics.data) || []
      topics = topics.sort(function(a,b) {
        return new Date(a.created_date) - new Date(b.created_date);
      });
      return chunk(topics, 3)
    }
  },
  methods: {
    openNewModal () {
      this.openTopicModal = true
    },
    openEditModal (topic) {
      this.topicToEdit = topic
      this.edit = true
      this.openTopicModal = true
    },
    closeModal () {
      this.openTopicModal=false
      this.edit = false
    },
    topicsFound (foundTopics) {
      this.foundTopics = foundTopics
    }
  }
}
</script>
