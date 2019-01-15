import Vue from 'vue'
import Vuex from 'vuex'

import statistics from './modules/statistics'

Vue.use(Vuex)

export default new Vuex.Store({
  strict: true,
  modules: {
    statistics
  }
})
