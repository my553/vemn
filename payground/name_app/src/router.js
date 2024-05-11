import { createRouter, createWebHistory } from "vue-router";
import auth from "./components/Auth/auth.vue";
import * as storage from "./storage.js";
import Main from "./components/Main.vue";
import RootLayout from "./components/RootLayout.vue";
import NotFound from "./components/NotFound.vue";
import Statistic from "./components/Statistic.vue";
import Analytics from "./components/Analytics.vue";
import Reports from "./components/Reports.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/auth", component: auth },
    {
      path: "/dashboard",
      component: RootLayout,
      children: [
        { path: "", component: NotFound },
        { path: "main", component: Main },
        { path: "statistics", component: Statistic },
        {
          path: "reports",
          component: Reports,
        },
        { path: "analytics", component: Analytics },
        { path: "/:pathMatch(.*)*", component: NotFound },
      ],
    },
    { path: "/:pathMatch(.*)*", component: NotFound },
  ],
});

router.beforeEach((to, from) => {
  if (!(!storage.getToken() && to.fullPath !== "/auth")) {
    return true;
  } else {
    return false;
  }
});

export default router;
