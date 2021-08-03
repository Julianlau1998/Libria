export function GET_TOPICS (state) {
  state.topics.loading = true
  state.err = null
}

export function RECEIVE_TOPICS (state, topics) {
  state.topics.data = topics
  state.topics.loading = false
}

export function GET_TOPIC (state) {
  state.topic.loading = true
  state.err = null
}

export function RECEIVE_TOPIC (state, topic) {
  state.topic.data = topic
  state.topic.loading = false
}

export function GET_RANDOM_TOPIC (state) {
  state.randomTopic.loading = true
  state.err = null
}

export function RECEIVE_RANDOM_TOPIC (state, topic) {
  state.randomTopic.data = topic
  state.randomTopic.loading = false
}

export function POST_TOPIC (state) {
  state.topic.submitting = true
  state.err = null
}
export function TOPIC_POSTED (state, topic) {
  state.topics.data.push(topic)
  state.topic.submitting = false
}
export function DELETE_TOPIC (state, topicId) {
  state.topics.data = state.topics.data.filter(el => el.id !== topicId)
  state.topic.loading = true
  state.err = null
}
export function TOPIC_DELETED (state) {
  state.topic.loading = false
}
export function PUT_TOPIC (state) {
  state.topic.submitting = true
  state.err = null
}
export function TOPIC_PUT (state) {
  state.topic.submitting = false
}
