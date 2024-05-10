import {createRouter, createWebHistory} from 'vue-router';
import auth from "./components/Auth/auth.vue";
import main from "./components/Main/main.vue";
import * as storage from "./storage.js";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {path: '/', component: auth},
        {path: '/dashboard', component: main},
    ],
})

router.beforeEach((to, from) => {
    if (!(!storage.getToken() && to.fullPath !== '/')) {
        return true
    } else {
        return false
    }
})

export default router