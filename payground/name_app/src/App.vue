<script>
import { defineComponent } from "vue";
import router from "./router.js";
import * as storage from "./storage.js";
import Header from "./components/Header.vue";

export default defineComponent({
  components: { Header },
  mounted() {
    let token = storage.get("token");
    if (!token) {
      router.push("/auth");
    }
  },
});
</script>

<template>
  <div class="root">
    <router-view></router-view>
  </div>
</template>

<style lang="scss">
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.root {
  width: 100%;
  min-height: 100vh;
  background: #1e1e1e;
}
</style>
