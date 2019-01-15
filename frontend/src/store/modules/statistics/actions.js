import axios from 'axios'

import * as types from './mutation-types'

const API_URL = process.env.VUE_APP_API_URL

export const fetchCardsData = ({ commit }) => {
  return new Promise((resolve, reject) => {
    axios.get(`${API_URL}/statistics/cards`).then(response => {
      commit(types.SET_CARDS_DATA, response.data)
      resolve(response)
    }).catch(err => {
      reject(err)
    })
  })
}
