import mutations from './mutations'
import * as actions from './actions'
import * as getters from './getters'

const initialState = {
  cardsData: [{}],
  usersBehavior: {
    labels: [],
    series: []
  }
}

export default {
  namespaced: true,
  state: initialState,
  getters,
  actions,
  mutations
}
