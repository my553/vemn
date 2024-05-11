<template>
  <header
    class="w-full flex flex-wrap justify-between items-center px-[88px] py-[40px] gap-2"
  >
    <div class="flex items-center gap-2 justify-between">
      <img class="w-[70px] h-[70px]" alt="" src="../images/logo-main.svg" />
      <div class="lg:flex items-center gap-20 hidden text-white">
        VULNERABILITY EXPLORER <br />
        & MANAGEMENT
      </div>
    </div>

    <div class="flex gap-[48px]">
      <RouterLink
        exact-active-class="!text-main-blue"
        to="/dashboard/main"
        class="flex flex-wrap items-center text-white"
      >
        <img class="" src="../images/main.svg" alt="" />
        <div class="">Главная</div>
      </RouterLink>
      <RouterLink
        exact-active-class="!text-main-blue"
        to="/dashboard/statistics"
        class="flex flex-wrap items-center text-white"
      >
        <img class="" src="../images/statistic.svg" alt="" />
        <div class="">Статистика</div>
      </RouterLink>
      <RouterLink
        exact-active-class="!text-main-blue"
        to="/dashboard/reports"
        class="flex flex-wrap items-center text-white"
      >
        <img class="" src="../images/reports.svg" alt="" />
        <div class="">Отчеты</div>
      </RouterLink>
      <RouterLink
        exact-active-class="!text-main-blue"
        to="/dashboard/analytics"
        class="flex flex-wrap items-center text-white"
      >
        <img class="" src="../images/analytics.svg" alt="" />
        <div class="">Аналитика</div>
      </RouterLink>
    </div>

    <Button @click="handleLogoutClick">
      <img class="w-[30px] h-[30px]" src="../images/nav-profile.svg" alt="" />
      {{ accFIO || "Нет имени" }}
      <img src="../images/nav-exit.svg" alt="" />
    </Button>
  </header>
</template>

<script setup>
import { onMounted, ref } from "vue";
import * as storage from "../storage.js";
import { useRouter } from "vue-router";

import Button from "./Button.vue";

const router = useRouter();

const accFIO = ref(storage.get("accFIO"));

const handleLogoutClick = () => {
  storage.remove("token");
  storage.remove("accFIO");
  router.push("/auth");
};
</script>
