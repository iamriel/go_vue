import axios from "axios";

const API_URL = process.env.VUE_APP_API_URL;
const SET_AUTHENTICATED = "SET_AUTHENTICATED"

const initialState = {
  authenticated: localStorage.getItem("authenticated")
}

const getters = {
  state: state => state.authenticated
}

const mutations = {
  [SET_AUTHENTICATED](state, payload) {
    state.authenticated = payload
    localStorage.setItem("authenticated", payload)
  }
}

const actions = {
  login({commit}, {email, password}) {
    return new Promise((resolve, reject) => {
      axios
        .post(`${API_URL}/login`, {email, password})
        .then(response => {
          commit(SET_AUTHENTICATED, true);
          resolve(response);
        })
        .catch(err => {
          reject(err);
        });
    });
  }
}

export default {
  namespaced: true,
  state: initialState,
  getters,
  actions,
  mutations
};
