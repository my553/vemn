import { createApp } from "vue";
import App from "./App.vue";
import router from "./router.js";
import PrimeVue from "primevue/config";
import "primevue/resources/themes/lara-light-teal/theme.css";

import "./style.css";

createApp(App).use(router).use(PrimeVue).mount("#app");
