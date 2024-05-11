<script setup lang="ts">
import Table, { Statistics } from "./Table.vue";
import SectionTitle from "./SectionTitle.vue";
import Button from "./Button.vue";
import { computed, onMounted, ref, watch, watchEffect } from "vue";
import * as httpClient from "../httpClient";

const data = ref<Statistics[]>([]);
const rowsCount = ref("5");
const withSsl = ref(false);
const withWaf = ref(false);
const withActive = ref(false);

const filteredData = computed(() => {
  const filters = [
    withSsl.value ? (item: Statistics) => item.ssl : () => true,
    withWaf.value ? (item: Statistics) => item.waf === "За WAF" : () => true,
    withActive.value
      ? (item: Statistics) => item.resolve === "Активен"
      : () => true,
  ];

  return data.value.filter((stat) => filters.every((filter) => filter(stat)));
});

onMounted(() => {
  httpClient.Get("/api/statistic").then((response) => {
    let resp = JSON.parse(response.data.body) as { stats: any[] };

    data.value = resp.stats.map(({ resource }) => ({
      resource: resource.nameUrl,
      ip: resource.ip,
      ssl: resource.dateCert,
      date: resource.nameUrl,
      waf: resource.wafStatus ? "За WAF" : "Не за WA",
      vulnerability: resource.vulnerabilites,
      resolve: resource.errStatus ? "Не активен" : "Активен",
      lastScanDate: resource.vulnerabilites.date,
    }));
  });
});
</script>

<template>
  <div class="flex flex-col gap-4 w-full pb-6">
    <SectionTitle>Статистика</SectionTitle>
    <div class="flex px-[88px] w-full gap-[40px] justify-between">
      <div class="flex-1 flex relative items-center">
        <SectionTitle size="sm" rounded="md">
          Получить информацию о ресурсе
        </SectionTitle>
        <input
          placeholder="Ресурс"
          class="right-0 rounded-2xl h-[48px] w-[310px] placeholder:text-main-blue text-main-blue px-3 text-md"
        />
      </div>
      <Button>Проверить ресурс</Button>
    </div>
    <div class="px-[88px] w-full">
      <div
        class="w-full p-[22px] bg-white/10 rounded-xl flex items-end flex-col gap-[40px]"
      >
        <div class="flex gap-4 justify-between w-full">
          <div class="flex flex-1 gap-[40px]">
            <span class="text-white">Найдено: {{ filteredData.length }}</span>
            <div class="flex gap-4">
              <label class="text-white flex gap-3 cursor-pointer select-none">
                <span class="text-white">С SSL</span>
                <input v-model="withSsl" type="checkbox" />
              </label>

              <label class="text-white flex gap-3 cursor-pointer select-none">
                <span class="text-white">С WAF</span>
                <input v-model="withWaf" type="checkbox" />
              </label>

              <label class="text-white flex gap-3 cursor-pointer select-none">
                <span class="text-white">Активные</span>
                <input v-model="withActive" type="checkbox" />
              </label>
            </div>
          </div>
          <div class="flex gap-2">
            <label for="pagination-select" class="text-white">
              Количество строк:
            </label>
            <select
              v-model="rowsCount"
              id="pagination-select"
              class="bg-transparent text-white cursor-pointer"
            >
              <option value="5">5</option>
              <option value="10">10</option>
              <option value="25">25</option>
            </select>
          </div>
        </div>
        <Table :rows-count="+rowsCount" :data="filteredData" />
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss"></style>
