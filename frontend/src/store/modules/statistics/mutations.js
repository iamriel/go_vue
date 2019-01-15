import * as types from "./mutation-types";

const mutations = {
  [types.SET_CARDS_DATA](state, payload) {
    state.cardsData = payload;
  },
  [types.SET_USERS_BEHAVIOR](state, payload) {
    state.usersBehavior = payload;
  },
  [types.SET_EMAIL_STATISTICS](state, payload) {
    state.emailStatistics = payload;
  },
  [types.SET_SALES_DATA](state, payload) {
    state.salesData = payload;
  }
};

export default mutations;
