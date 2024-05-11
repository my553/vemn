<script setup>
import Header from "./Header.vue";
</script>

<template>
  <div class="rootLayout">
    <Header />
    <router-view />
  </div>
</template>

<style scoped lang="scss">
.rootLayout {
  width: 100%;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  background: #1e1e1e;
}
</style>
