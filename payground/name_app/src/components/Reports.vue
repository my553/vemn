<script setup lang="ts">
import SectionTitle from "./SectionTitle.vue";
import { onMounted, reactive, ref } from "vue";
import Button from "./Button.vue";
import * as httpClient from "../httpClient";

const weekStatistic = ref<{
  current: { noRes: any[]; newWaf: any[] };
  last: { noRes: any[]; newWaf: any[] };
}>({
  current: { noRes: [], newWaf: [] },
  last: { noRes: [], newWaf: [] },
});

onMounted(() => {
  httpClient.Get("/api/week-statistic").then((response) => {
    let resp = JSON.parse(response.data.body);

    weekStatistic.value = {
      last: {
        noRes: resp.last.no_res_resource,
        newWaf: resp.last.new_waf_resource,
      },
      current: {
        noRes: resp.current.no_res_resource,
        newWaf: resp.current.new_waf_resource,
      },
    };
  });
});

const downloadReports = () => {
  httpClient.Get("/api/download-reports").then((response) => {
    console.log(response);
  });
};
</script>

<template>
  <SectionTitle>Отчеты</SectionTitle>
  <div class="mt-[32px] pl-[88px] w-full flex flex-col items-center">
    <div class="w-full">
      <SectionTitle size="sm">Текущая неделя</SectionTitle>
      <div class="flex w-full">
        <div
          class="pl-[97px] py-[28px] pr-[160px] flex flex-1 flex-col gap-5 border-solid border-r-[1px] border-white"
        >
          <span class="text-2xl font-bold text-white">Неактивные ресурсы</span>
          <div class="flex justify-between items-center">
            <span class="text-xl font-bold text-white">Ресурс</span>
            <span class="text-xl font-bold text-white">Дата</span>
          </div>
          <div class="flex flex-col gap-4">
            <div
              v-for="item in weekStatistic.current.noRes"
              class="flex justify-between items-center"
            >
              <span class="text-[#999999] text-lg">
                {{ (item as any).resource }}
              </span>
              <span class="text-[#999999] text-lg">
                {{ (item as any).date }}
              </span>
            </div>
          </div>
        </div>

        <div class="pl-[97px] py-[28px] pr-[160px] flex flex-1 flex-col gap-5">
          <span class="text-2xl font-bold text-white">Новые WAF-ресурсы</span>
          <div class="flex justify-between items-center">
            <span class="text-xl font-bold text-white">Ресурс</span>
            <span class="text-xl font-bold text-white">Дата</span>
          </div>
          <div class="flex flex-col gap-4">
            <div
              v-for="item in weekStatistic.current.newWaf"
              class="flex justify-between items-center"
            >
              <span class="text-[#999999] text-lg">
                {{ (item as any).resource }}
              </span>
              <span class="text-[#999999] text-lg">
                {{ (item as any).date }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="w-full">
      <SectionTitle size="sm">Прошлая неделя</SectionTitle>
      <div class="flex w-full">
        <div
          class="pl-[97px] py-[28px] pr-[160px] flex flex-1 flex-col gap-5 border-solid border-r-[1px] border-white"
        >
          <span class="text-2xl font-bold text-white">Неактивные ресурсы</span>
          <div class="flex justify-between items-center">
            <span class="text-xl font-bold text-white">Ресурс</span>
            <span class="text-xl font-bold text-white">Дата</span>
          </div>
          <div class="flex flex-col gap-4">
            <div
              v-for="item in weekStatistic.last.noRes"
              class="flex justify-between items-center"
            >
              <span class="text-[#999999] text-lg">
                {{ (item as any).resource }}
              </span>
              <span class="text-[#999999] text-lg">
                {{ (item as any).date }}
              </span>
            </div>
          </div>
        </div>

        <div class="pl-[97px] py-[28px] pr-[160px] flex flex-1 flex-col gap-5">
          <span class="text-2xl font-bold text-white">Новые WAF-ресурсы</span>
          <div class="flex justify-between items-center">
            <span class="text-xl font-bold text-white">Ресурс</span>
            <span class="text-xl font-bold text-white">Дата</span>
          </div>
          <div class="flex flex-col gap-4">
            <div
              v-for="item in weekStatistic.last.newWaf"
              class="flex justify-between items-center"
            >
              <span class="text-[#999999] text-lg">
                {{ (item as any).resource }}
              </span>
              <span class="text-[#999999] text-lg">
                {{ (item as any).date }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
    <Button @click="downloadReports">
      Выгрузить полный отчёт
      <svg
        width="17"
        height="17"
        viewBox="0 0 17 17"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M4.90795 8.54323L5.06534 8.38317L7.5393 10.8571L8.39286 11.7107V10.5036V0.5H8.60714V10.5036V11.7107L9.4607 10.8571L11.9339 8.38388L12.0916 8.54372L8.5 12.1363L4.90795 8.54323ZM0.5 15.0389V12.5967H0.714286V15.0389C0.714286 15.3812 0.862706 15.6738 1.09295 15.9052L1.09479 15.907C1.32625 16.1373 1.61876 16.2857 1.96107 16.2857H15.0389C15.3812 16.2857 15.6738 16.1373 15.9052 15.907L15.907 15.9052C16.1373 15.6738 16.2857 15.3812 16.2857 15.0389V12.5967H16.5V15.0389C16.5 15.4732 16.36 15.8097 16.0854 16.0842C15.8099 16.3598 15.4728 16.5 15.0389 16.5H1.96107C1.52685 16.5 1.19031 16.36 0.915767 16.0854C0.640188 15.8099 0.5 15.4728 0.5 15.0389Z"
          fill="white"
          stroke="white"
        />
      </svg>
    </Button>
  </div>
</template>
