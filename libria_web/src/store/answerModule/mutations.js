export function GET_ANSWERS (state) {
  state.answers.loading = true
  state.err = null
}

export function RECEIVE_ANSWERS (state, answers) {
  state.answers.data = answers
  state.answers.loading = false
}

export function GET_ANSWER (state) {
  state.answer.loading = true
  state.err = null
}

export function RECEIVE_ANSWER (state, answer) {
  state.answer.data = answer
  state.answer.loading = false
}
export function POST_ANSWER (state) {
  state.answer.submitting = true
  state.err = null
}
export function ANSWER_POSTED (state, answer) {
  console.log(answer)
  state.answers.data.push(answer)
  state.answer.submitting = false
}
export function DELETE_ANSWER (state, answerId) {
  state.answers.data = state.answers.data.filter(el => el.id !== answerId)
  state.answer.loading = true
  state.err = null
}
export function ANSWER_DELETED (state) {
  state.answer.loading = false
}
export function PUT_ANSWER (state) {
  state.answer.submitting = true
  state.err = null
}
export function ANSWER_PUT (state) {
  state.answer.submitting = false
}
