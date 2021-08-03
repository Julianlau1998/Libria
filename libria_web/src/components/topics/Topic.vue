<template>
  <div>
    <div class="card my-2 is-wide-card mt-6">
            <div class="card-content is-wide-card-content ml-5">
                <span class="is-vote-icons">
                    <br>
                    <br>
                    <br>
                </span>
                <h2 class="is-size-4 is-card-title">
                    <i class="fas fa-crown is-crown"></i>
                    <br>
                    <b>
                        {{topic.title}}
                    </b>
                </h2>
                <h2 class="is-size-5">
                    {{bestAnswer.body}}
                </h2>
            </div>
        </div>
        <RandomTopic
            @openRandomTopic="openRandomTopic($event)"
            :buttonText="'Get Random Topic'"
            class="mb-6"
        />
        <div
            class="columns is-variable is-0-mobile is-1-tablet is-3-desktop"
            v-for="(answer, index) in answers"
            :key="index"
        >   
            <div class="column" v-if="answer.length>1">
                <card
                    :type="'answer'"
                    :card="answer[0]"
                    @openEditModal="openEditModal($event)"
                    :cardNumber="Math.floor(Math.random() * 4) + 1"
                />
            </div>
            <div class="column" v-if="answer.length>1">
                <card
                    :type="'answer'"
                    :card="answer[1]"
                    @openEditModal="openEditModal($event)"
                    :cardNumber="Math.floor(Math.random() * 4) + 1"
                />
            </div>
            <div class="column" v-if="answer.length>2">
                <card
                    :type="'answer'"
                    :card="answer[2]"
                    @openEditModal="openEditModal($event)"
                    :cardNumber="Math.floor(Math.random() * 4) + 1"
                />
            </div>

            <div class="singleColumn" v-if="answer.length===1">
                <card
                    :type="'answer'"
                    :card="answer[0]"
                    @openEditModal="openEditModal($event)"
                    :cardNumber="Math.floor(Math.random() * 4) + 1"
                />
            </div>
        </div>
        <AnswerModal
            v-if="openAnswerModal"
            :answerToEdit="answerToEdit"
            :edit="edit"
            :topicId="topicId"
            @closeModal="closeModal()"
        />
        <AddButton
            @openNewModal="openNewModal()"
        />
  </div>
</template>

<script>
import card from '@/components/start/Card.vue'
import RandomTopic from '@/components/start/RandomTopic.vue'
import AnswerModal from '@/components/modals/AnswerModal.vue'
import AddButton from '@/components/start/AddButton.vue'
import { mapState } from 'vuex'

const chunk = require('chunk')
export default {
    components: {
        card,
        RandomTopic,
        AnswerModal,
        AddButton
    },
    data () {
        return {
            topicId: '',
            openAnswerModal: false,
            answerToEdit: {}
        }
    },
    created () {
        this.topicId = this.$route.params.uuid
        if (this.$auth.user === undefined) {
            this.$auth.user = {}
        } 
        this.$store.dispatch('answerModule/getAllByTopic', {topicId: this.topicId, userId: this.$auth.user.sub})
        this.$store.dispatch('topicModule/getOne', this.topicId)
    },
    computed: {
        ...mapState(['answerModule']),
        ...mapState(['topicModule']),
        answers () {
            let answers = (!this.answerModule.answers.loading && this.answerModule.answers.data) || []
            answers = answers.sort(function(a,b) {
                return new Date(a.created_date) - new Date(b.created_date);
            });
            return chunk(answers, 3)
        },
        bestAnswer () {
            if (this.answers === []) {
                return {text: '', votes: 0}
            }
            let bestAnswer = {votes: -5}
            this.answers.forEach(el => {
                el.forEach(answer => {
                    let votes = 0
                    answer.votes.forEach(vote => {
                        if (vote.upvote === 'true') {
                            votes++
                        } else if (vote.upvote === 'false') {
                            votes--
                        }
                    })
                    if (votes > bestAnswer.votes) {
                        bestAnswer.body = answer.body
                        bestAnswer.votes = votes
                    }
                })
            });
            return bestAnswer
        },
        topic () {
            return (!this.topicModule.topic.loading && this.topicModule.topic.data) || {}
        }
    },
    methods: {
        openNewModal () {
            this.openAnswerModal = true
        },
        openEditModal (answer) {
            this.openAnswerModal = true
            this.edit = true
            this.answerToEdit = answer
        },
        closeModal () {
            this.openAnswerModal=false
            this.edit = false
        },
    }
}
</script>