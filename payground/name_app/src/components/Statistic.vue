<script setup lang="ts">
import Table, { Statistics } from "./Table.vue";
import SectionTitle from "./SectionTitle.vue";
import Button from "./Button.vue";
import { onMounted, ref } from "vue";
import * as httpClient from "../httpClient";

const data = ref<Statistics[]>([]);

onMounted(() => {
  httpClient.Get("/api/statistic").then((response) => {
    let resp = JSON.parse(response.data.body) as { stats: any[] };

    console.log(
      resp.stats.map(({ resource }) => ({
        resource: resource.nameUrl,
        ip: resource.ip,
        ssl: resource.dateCert,
        date: resource.nameUrl,
        waf: resource.wafStatus,
        vulnerability: resource.vulnerabilites,
        resolve: resource.errStatus,
        lastScanDate: resource.vulnerabilites.date,
      })),
    );

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
    <div class="px-[88px]">
      <Table :data="data" />
    </div>
  </div>
</template>

<style scoped lang="scss"></style>
