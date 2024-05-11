<script setup>
import Button from "./Button.vue";
import { useRouter } from "vue-router";

const router = useRouter();
const handleBackClick = () => {
  router.push("/dashboard/main");
};
</script>

<template>
  <div class="container">
    <div class="top-text">Ошибка 404</div>
    <div class="bottom-text">Страница не найдена</div>
    <Button variant="outline" @click="handleBackClick">
      Вернуться на главную</Button
    >
  </div>
</template>

<style scoped lang="scss">
.container {
  margin-top: 270px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 10px;
}

.top-text {
  font-size: 48px;
  font-weight: 700;
  color: #2e6ac3;
}

.bottom-text {
  font-size: 48px;
  font-weight: 400;
  color: white;
}
</style>
