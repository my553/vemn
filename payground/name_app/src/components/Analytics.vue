<script setup lang="ts">
import SectionTitle from "./SectionTitle.vue";
import { Bar, Pie } from "vue-chartjs";
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale,
  ChartData,
  ChartOptions,
  ArcElement,
} from "chart.js";
import { computed, onMounted, ref, watchEffect } from "vue";
import * as httpClient from "../httpClient";

ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
  ArcElement,
);

const bars = ref({
  resource: [],
  waf: [],
});

const wafPieData = ref([0, 0]);
const allResourcePieData = ref([0, 0]);

const barChartData = computed<ChartData>(() => ({
  labels: [
    "Январь",
    "Февраль",
    "Март",
    "Апрель",
    "Май",
    "Июнь",
    "Июль",
    "Август",
    "Сентябрь",
    "Октябрь",
    "Ноябрь",
    "Декабрь",
  ],
  datasets: [
    {
      label: "Ресурс",
      backgroundColor: "#FFFFFF4D",
      data: bars.value.resource,
      barPercentage: 0.55,
      categoryPercentage: 0.7,
    },
    {
      label: "За WAF",
      backgroundColor: "#2E6AC3",
      data: bars.value.waf,
      barPercentage: 0.55,
      categoryPercentage: 0.7,
    },
  ],
}));

const barChartOptions: ChartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  backgroundColor: "#FFFFFF1A",
  color: "#FFFFFF",
  scales: {
    x: {
      ticks: {
        color: "#FFFFFF",
      },
      grid: {
        display: false,
      },
    },
    y: {
      ticks: {
        color: "#FFFFFF",
      },
      grid: {
        display: false,
      },
    },
  },
};

const wafChartData = computed<ChartData>(() => ({
  labels: ["Все ресурсы", "За WAF"],
  datasets: [
    {
      backgroundColor: ["#2E6AC3", "#FFFFFF4D"],
      data: wafPieData.value,
    },
  ],
}));

const allResourcesChartData = computed<ChartData>(() => ({
  labels: ["Все ресурсы", "За WAF"],
  datasets: [
    {
      backgroundColor: ["#2E6AC3", "#FFFFFF4D"],
      data: allResourcePieData.value,
    },
  ],
}));

const pieChartOptions: ChartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: {
      labels: {
        color: "white",
        boxHeight: 20,
        boxWidth: 20,
        useBorderRadius: true,
        usePointStyle: true,
        borderRadius: 9999,
        pointStyleWidth: 0,
        pointStyle: "transparent",
      },
    },
  },
};

onMounted(() => {
  httpClient.Get("/api/chart-statistic").then((response) => {
    let resp = JSON.parse(response.data.body) as Record<any, any>;

    // Change when second object will be ready
    allResourcePieData.value = [resp.wafStat.withWAF, resp.wafStat.noWAF];
    wafPieData.value = [resp.wafStat.withWAF, resp.wafStat.noWAF];

    bars.value = Object.keys(resp.year).reduce(
      (acc, key) => {
        const withWaf = Number(resp.year[key].chart.withWAF);
        const erResource = Number(resp.year[key].chart.erServers);
        const allResource = Number(resp.year[key].chart.allServers);

        const resource = allResource - erResource;

        acc.waf.push(isNaN(withWaf) ? 0 : withWaf);
        acc.resource.push(isNaN(resource) ? 0 : resource);

        return acc;
      },
      {
        resource: [],
        waf: [],
      },
    );
  });
});
</script>

<template>
  <div class="w-full">
    <SectionTitle> Аналитика</SectionTitle>
    <div class="px-[88px] py-[40px] flex gap-[64px]">
      <div
        class="h-[680px] w-1/2 rounded-2xl overflow-hidden px-[40px] py-[60px] bg-white/10 flex flex-col gap-6"
      >
        <span class="text-lg font-bold text-white"
          >Динамика изменения активности ресурсов</span
        >
        <Bar :options="barChartOptions as any" :data="barChartData as any" />
      </div>
      <div class="flex flex-col gap-[40px] h-[680px] w-1/2">
        <div
          class="h-1/2 bg-white/10 flex flex-col rounded-2xl w-full p-[26px]"
        >
          <span class="text-lg font-bold text-white">WAF</span>
          <div class="h-[256px]">
            <Pie
              :options="pieChartOptions as any"
              :data="wafChartData as any"
            />
          </div>
        </div>
        <div
          class="h-1/2 bg-white/10 flex flex-col rounded-2xl w-full p-[26px]"
        >
          <span class="text-lg font-bold text-white">Все ресурсы</span>
          <div class="h-[256px]">
            <Pie
              :options="pieChartOptions as any"
              :data="allResourcesChartData as any"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss"></style>
