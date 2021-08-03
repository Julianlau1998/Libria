<template>
    <div class="modal is-active is-new-topic-modal">
        <h1 class="is-size-3 mb-2">
            Add New Answer
        </h1>
        <validationObserver v-slot="{ invalid, handleSubmit}" class="textarea-wrapper">
            <form @submit.prevent="handleSubmit(onSubmit)">
                <validationProvider rules="required|max:255" v-slot="{errors}">
                    <textarea
                        class="textarea is-primary is-textarea"
                        placeholder="Your answer"
                        style="border:solid 1px orange; max-width: 5rem;"
                        cols="1"
                        v-model="answer.body"
                        autofocus
                    ></textarea>
                    <span class="is-error mb-6">
                        {{ errors[0] }}&nbsp;
                    </span>
                    <br>
                </validationProvider>
                <span>
                    <button type="submit" attachToDocument: true class="button is-success mx-2">
                        <span class="icon is-small">
                            <i class="fas fa-check"></i>
                        </span>
                        <span>
                            Save
                        </span>
                    </button>
                    <button
                        class="button is-danger is-outlined mx-2 "
                        @click="$emit('closeModal')"
                    >
                        <span>
                            Cancel
                        </span>
                        <span class="icon is-small">
                            <i class="fas fa-times"></i>
                        </span>
                    </button>
                </span>
            </form>
        </validationObserver>
    </div>
</template>

<script>
import { ValidationProvider, ValidationObserver, extend } from 'vee-validate'
import { required, max } from 'vee-validate/dist/rules'

extend('required', {
  ...required,
  message: 'This field is required'
})
extend('max', {
  ...max,
  message: 'Max 255 Chars'
})

export default {
    name: 'newAnswerModal',
    components: {
        ValidationProvider,
        ValidationObserver,
    },
    props: {
        edit: {
            type: Boolean,
            required: true,
            default: false
        },
        answerToEdit: {
            type: Object,
            required: false
        },
        topicId: {
            type: String,
            required: true,
            default: ''
        }
    },
    data () {
        return {
            answer: {
                id: '',
                topic_id: '',
                body: ''
            }
        }
    },
    created () {
        if (this.edit) {
            this.answer = this.answerToEdit
        }
        this.answer.topic_id = this.topicId
    },
    methods: {
        onSubmit () {
            if (!this.edit) {
                this.answer.username = this.$auth.user.nickname
                this.answer.user_id = this.$auth.user.sub
                this.answer.votes = [{upvote: 'true', user_id: this.$auth.user.sub}]
                this.answer.upvoted_by_me = true
                this.$store.dispatch('answerModule/post', this.answer)
                this.$emit('closeModal')
            } else {
                this.$store.dispatch('answerModule/put', this.answer)
                this.$emit('closeModal')
            }
        }
    }
}
</script>

<style>

</style>