import * as types from './mutation-types'

const mutations = {
  [types.SET_CARDS_DATA] (state, payload) {
    state.cardsData = payload
  }
}

export default mutations
