import Vue from "vue";
import App from "./App";
import router from "./router/index";

import VueMoment from "vue-moment";
import PaperDashboard from "./plugins/paperDashboard";
import moment from "moment-timezone";
import "vue-notifyjs/themes/default.css";

import store from "./store";

Vue.use(PaperDashboard);
Vue.use(VueMoment, {
  moment
});

/* eslint-disable no-new */
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
