<template>
    <span>
        <div
            class="card is-one-third is-smaller-card-width"
        >
            <div class="card-content" :id="`card-content-${type}`" ref="cardTitle">
                <p class="is-username">
                    @{{card.username}}
                </p>
                <span class="is-vote-icons" v-if="this.type === 'answer'">
                    <i 
                        class="fas fa-chevron-up"
                        :class="(card.upvoted_by_me) ? 'upvote' : 'no-upvote'"
                        @click="upvoteAnswer()"
                        ref="upvote"
                    ></i>
                    <br>
                    {{votes}}
                    <br>
                    <i
                        class="fas fa-chevron-down"
                        :class="(card.downvoted_by_me) ? 'downvote' : 'no-downvote'"
                        @click="downvoteAnswer()"
                        ref="downvote"
                    ></i>
                </span>
                <span class="is-vote-icons" v-else>
                    <br><br>
                </span>
                <span @click="$emit('openTopic', card.id)">
                    <h2 class="is-size-4 is-card-title" v-if="card.title">
                        <b>
                            {{card.title}}
                        </b>
                    </h2>
                    <h2 class="is-size-6 is-card-title" v-else>
                        <b>
                            &nbsp;
                        </b>
                    </h2>
                    <h2 class="is-size-5 card-body" :id="`card-body-${cardNumber}`">
                        {{card.body}}
                    </h2>
                </span>
                <i class="fas fa-comment-alt is-comment-icon"></i>
            </div>
        </div>
        <div class="dropdown is-active">
            <div class="dropdown-trigger" v-if="!$auth.loading && $auth.isAuthenticated">
                <i 
                    v-if="card.user_id === $auth.user.sub"
                    class="fas fa-ellipsis-h is-option-icon"
                    @click="dropdownActive=!dropdownActive"
                ></i>
            </div>
            <div
                v-if="dropdownActive"
                class="dropdown-menu mt-2 is-dropdown-item"
                id="dropdown-menu"
                role="menu"
            >
                <div class="dropdown-content">
                    <a @click="deleteCard()" class="dropdown-item">
                        Delete
                    </a>
                    <a @click="editCard()" class="dropdown-item">
                        Edit
                    </a>
                </div>
            </div>
        </div>
    </span>
</template>

<script>
import { mapState } from 'vuex'

export default {
    name: 'card',
    props: {
        card: {
            required: true,
            type: Object
        },
        type: {
            type: String,
            required: true,
            default: 'topic'
        },
        cardNumber: {
            type: Number,
            required: false
        }
    },
    data () {
        return {
            dropdownActive: false,
            upvote: false,
            downvote: false,
            user_id: ''
        }
    },
    watch: {
        is_loading (val) {
            if (val === false) {
                this.user_id = this.$auth.user.sub
            }
        }
    },
    computed: {
        is_loading () {
            return this.$auth.loading
        },
        votes () {
            let votes = 0
            this.card.votes.forEach(vote => {
                if (vote.upvote === 'true') {
                    votes++
                } else if (vote.upvote === 'false') {
                    votes--
                }
            })
            return votes
        }
    },
    methods: {
        ...mapState(['answerModule']),
        deleteCard () {
            if (this.type === 'topic') {
                this.$store.dispatch('topicModule/deleteOne', this.card)
                this.dropdownActive = false
            } else if (this.type === 'answer') {
                this.$store.dispatch('answerModule/deleteOne', this.card)
                this.dropdownActive = false
            }
    
        }, 
        editCard () {
            this.$emit('openEditModal', this.card)
            this.dropdownActive = false
        },
        upvoteAnswer () {
            if (this.$auth.isAuthenticated) {
                if (this.type === 'answer' && !this.card.upvoted_by_me && !this.card.downvoted_by_me) {
                    this.upvote = true
                    this.$store.state.answerModule.answers.data.forEach(answer => {
                        if (answer.id === this.card.id) {
                            answer.upvoted_by_me = true
                            answer.votes.push({ upvote: 'true', user_id: this.$auth.user.sub})
                        }
                    })
                    this.$store.dispatch('votesModule/post', {upvote: 'true', answer_id: this.card.id, user_id: this.$auth.user.sub})
                }
            } else {
                this.$store.state.login_modal = true
            }
        },
        downvoteAnswer () {
            if (this.$auth.isAuthenticated) {
                if (this.type === 'answer' && !this.card.upvoted_by_me && !this.card.downvoted_by_me) {
                    this.downvote = true
                    this.$store.state.answerModule.answers.data.forEach(answer => {
                        if (answer.id === this.card.id) {
                            answer.downvoted_by_me = true
                            answer.votes.push({ upvote: 'false', user_id: this.$auth.user.sub})
                        }
                    })
                    this.$store.dispatch('votesModule/post', {upvote: 'false', answer_id: this.card.id, user_id: this.$auth.user.sub})
                }
            } else {
                this.$store.state.login_modal = true
            }
        }
    }
}
</script>

<style>

</style>