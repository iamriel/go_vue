import DashboardLayout from "@/layout/dashboard/DashboardLayout.vue";

// GeneralViews
import NotFound from "@/pages/NotFoundPage.vue";

// Auth pages
import Login from "@/pages/Login.vue";

// Admin pages
import Dashboard from "@/pages/Dashboard.vue";
import UserProfile from "@/pages/UserProfile.vue";
import Notifications from "@/pages/Notifications.vue";
import Icons from "@/pages/Icons.vue";
import Maps from "@/pages/Maps.vue";
import Typography from "@/pages/Typography.vue";
import TableList from "@/pages/TableList.vue";

const routes = [
  {
    path: "/",
    component: DashboardLayout,
    redirect: "/dashboard",
    children: [
      {
        path: "dashboard",
        name: "dashboard",
        meta: { requiresAuth: true },
        component: Dashboard
      },
      {
        path: "stats",
        name: "stats",
        meta: { requiresAuth: true },
        component: UserProfile
      },
      {
        path: "notifications",
        name: "notifications",
        meta: { requiresAuth: true },
        component: Notifications
      },
      {
        path: "icons",
        name: "icons",
        meta: { requiresAuth: true },
        component: Icons
      },
      {
        path: "maps",
        name: "maps",
        meta: { requiresAuth: true },
        component: Maps
      },
      {
        path: "typography",
        name: "typography",
        meta: { requiresAuth: true },
        component: Typography
      },
      {
        path: "table-list",
        name: "table-list",
        meta: { requiresAuth: true },
        component: TableList
      }
    ]
  },
  {
    path: "/login",
    name: "login",
    meta: { requiresUnauthenticated: true },
    component: Login
  },
  {
    path: "*",
    component: NotFound
  }
];

/**
 * Asynchronously load view (Webpack Lazy loading compatible)
 * The specified component must be inside the Views folder
 * @param  {string} name  the filename (basename) of the view to load.
function view(name) {
   var res= require('../components/Dashboard/Views/' + name + '.vue');
   return res;
};**/

export default routes;
