import Vue from "vue";
import VueRouter from "vue-router";
import routes from "./routes";
import store from "../store";
Vue.use(VueRouter);

// configure router
const router = new VueRouter({
  routes, // short for routes: routes
  linkActiveClass: "active",
  mode: "history"
});

router.beforeEach((to, from, next) => {
  const vuexStore = store;
  if (!vuexStore.getters["auth/authenticated"]) {
    if (to.matched.some(m => m.meta.requiresAuth)) {
      return next({ name: "login" });
    } else {
      return next();
    }
  } else {
    return next();
  }
});

export default router;
