import Vue from 'vue'
import Vuex from 'vuex'
import topicModule from './topicModule'
import answerModule from './answerModule'
import votesModule from './votesModule'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    login_modal: false,
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    topicModule,
    answerModule,
    votesModule
  }
})
