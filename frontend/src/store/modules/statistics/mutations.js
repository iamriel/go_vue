import * as types from './mutation-types'

const mutations = {
  [types.SET_CARDS_DATA] (state, payload) {
    state.cardsData = payload
  },
  [types.SET_USERS_BEHAVIOR] (state, payload) {
    state.usersBehavior = payload
  }
}

export default mutations
