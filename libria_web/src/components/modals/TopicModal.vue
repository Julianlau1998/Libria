<template>
    <div class="modal is-active is-new-topic-modal">
        <h1 class="is-size-3 mt-1 modalTitle">
            Add New Topic
        </h1>
        <validationObserver v-slot="{ invalid, handleSubmit}" class="textarea-wrapper">
            <form @submit.prevent="handleSubmit(onSubmit)">
                <validationProvider rules="required|max:255" v-slot="{errors}">
                    <input
                        rules="required"
                        class="input is-primary mt-2"
                        type="text"
                        placeholder="New Topic"
                        v-model="topic.title"
                        autofocus
                    >
                    <br>
                    <span class="is-error mb-4">
                        {{ errors[0] }} &nbsp;
                    </span>
                </validationProvider>
                <validationProvider
                    rules="max:255"
                    class="textarea-wrapper"
                    v-slot="{errors}"
                    tag="span"
                >
                    <textarea
                        rules="required"
                        class="textarea is-primary is-textarea mt-1"
                        placeholder="Optional: Your answer"
                        style="border:solid 1px orange;"
                        v-model="topic.body"
                    ></textarea>
                    <span class="is-error mb-6">
                        {{ errors[0] }}
                    </span>
                    <br>
                </validationProvider>
                <span>
                    <button type="submit" class="button is-success mx-2 mt-3" attachToDocument: true>
                        <span class="icon is-small">
                            <i class="fas fa-check"></i>
                        </span>
                        <span>
                            Save
                        </span>
                    </button>
                    <button
                        class="button is-danger is-outlined mx-2 mt-3"
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
    <!-- </div> -->
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
    name: 'newTopicModal',
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
        topicToEdit: {
            type: Object,
            required: false
        }
    },
    data () {
        return {
            topic: {
                id: '',
                title: '',
                body: '',
            }
        }
    },
    created () {
        if (this.edit) {
            this.topic = this.topicToEdit
        }
    },
    methods: {
        onSubmit () {
            if (!this.edit) {
                this.topic.username = this.$auth.user.nickname
                this.topic.user_id = this.$auth.user.sub
                this.$store.dispatch('topicModule/post', this.topic)
                this.$emit('closeModal')
            } else {
                this.$store.dispatch('topicModule/put', this.topic)
                this.$emit('closeModal')
            }
        }
    }
}
</script>

<style>

</style>