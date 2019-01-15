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

export const fetchUsersBehavior = ({ commit }) => {
  return new Promise((resolve, reject) => {
    axios.get(`${API_URL}/statistics/users-behavior`).then(response => {
      commit(types.SET_USERS_BEHAVIOR, response.data)
      resolve(response)
    }).catch(err => {
      reject(err)
    })
  })
}

export const fetchEmailStatistics = ({ commit }) => {
  return new Promise((resolve, reject) => {
    axios.get(`${API_URL}/statistics/email`).then(response => {
      commit(types.SET_EMAIL_STATISTICS, response.data)
      resolve(response)
    }).catch(err => {
      reject(err)
    })
  })
}

export const fetchSalesData = ({ commit }) => {
  return new Promise((resolve, reject) => {
    axios.get(`${API_URL}/statistics/sales`).then(response => {
      commit(types.SET_SALES_DATA, response.data)
      resolve(response)
    }).catch(err => {
      reject(err)
    })
  })
}
