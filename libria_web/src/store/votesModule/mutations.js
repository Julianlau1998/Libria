export function GET_VOTES (state) {
  state.votes.loading = true
  state.err = null
}

export function RECEIVE_VOTES (state, votes) {
  state.votes.data = votes
  state.votes.loading = false
}

export function GET_VOTE (state) {
  state.vote.loading = true
  state.err = null
}

export function RECEIVE_VOTE (state, vote) {
  state.vote.data = vote
  state.vote.loading = false
}
export function POST_VOTE (state) {
  state.vote.submitting = true
  state.err = null
}
export function VOTE_POSTED (state, vote) {
  state.votes.data.push(vote)
  state.votes.submitting = false
}
export function DELETE_VOTE (state, vote_id) {
  state.votes.data = state.votes.data.filter(el => el.id !== vote_id)
  state.vote.loading = true
  state.err = null
}
export function VOTE_DELETED (state) {
  state.vote.loading = false
}
export function PUT_VOTE (state) {
  state.vote.submitting = true
  state.err = null
}
export function VOTE_PUT (state) {
  state.vote.submitting = false
}
