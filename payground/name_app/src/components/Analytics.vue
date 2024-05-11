<script setup lang="ts">
import SectionTitle from "./SectionTitle.vue";
import { Bar, Doughnut } from "vue-chartjs";
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

// const plugin = {
//   id: "increase-legend-spacing",
//   beforeInit(chart: any) {
//     // Get reference to the original fit function
//     // const originalFit = chart.legend.fit;
//     //
//     // // Override the fit function
//     // chart.legend.fit = function fit() {
//     //   // Call original function and bind scope in order to use `this` correctly inside it
//     //   originalFit.bind(chart.legend)();
//     //   // Change the height as suggested in another answers
//     //   this.height += 20;
//     // };
//   },
// };

ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
  ArcElement,
  // plugin,
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
      barPercentage: 0.7,
      categoryPercentage: 0.8,
      // barThickness: 11,
      borderRadius: 4,
    },
    {
      label: "За WAF",
      backgroundColor: "#2E6AC3",
      data: bars.value.waf,
      barPercentage: 0.7,
      categoryPercentage: 0.8,
      // barThickness: 11,
      borderRadius: 4,
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
        minRotation: 0,
        maxRotation: 90,
        count: 12,
        font: {
          size: 9,
        },
      },
      grid: {
        display: false,
      },
      border: {
        display: false,
      },
    },
    y: {
      border: {
        display: false,
      },
      ticks: {
        display: false,
        minRotation: 0,
        maxRotation: 0,
      },
      angleLines: {
        display: false,
      },
      grid: {
        display: false,
      },
    },
  },

  plugins: {
    legend: {
      align: "start",
      title: {
        font: {
          size: 14,
        },
      },
      labels: {
        color: "white",
        boxHeight: 20,
        boxWidth: 20,
        useBorderRadius: true,
        usePointStyle: true,
        textAlign: "right",
        borderRadius: 9999,
        pointStyleWidth: 0,
        pointStyle: "transparent",
      },
    },
  },
};

const wafChartData = computed<ChartData>(() => ({
  labels: ["ресурсы", "За WAF"],
  datasets: [
    {
      backgroundColor: ["#35D073", "#FFFFFF4D"],
      data: wafPieData.value,
    },
  ],
}));

const allResourcesChartData = computed<ChartData>(() => ({
  labels: ["с SSL-сертификатами", "без SSL-сертификатов"],
  datasets: [
    {
      backgroundColor: ["#2E6AC3", "#D03535"],
      data: allResourcePieData.value,
    },
  ],
}));

const wafPieChartOptions: ChartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  borderColor: "transparent",
  plugins: {
    legend: {
      position: "bottom",
      labels: {
        font: {
          size: 14,
        },
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

const sslPieChartOptions: ChartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  borderColor: "transparent",
  plugins: {
    legend: {
      position: "bottom",
      labels: {
        font: {
          size: 14,
        },
        color: "white",
        padding: 30,
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
        class="h-[644px] w-[800px] rounded-2xl overflow-hidden px-[40px] py-[60px] bg-white/10 flex flex-col gap-6"
      >
        <span class="text-lg font-bold text-white"
          >Динамика изменения активности ресурсов</span
        >
        <div class="h-[calc(100%-28px)]">
          <Bar :options="barChartOptions as any" :data="barChartData as any" />
        </div>
      </div>
      <div class="flex flex-col gap-[40px] h-[684px] w-1/2">
        <div
          class="h-[calc(50%-40px)] bg-white/10 flex flex-col rounded-2xl w-full p-[26px]"
        >
          <span class="text-lg font-bold text-white">WAF</span>
          <div class="h-[calc(100%-28px)]">
            <Doughnut
              :options="wafPieChartOptions as any"
              :data="wafChartData as any"
            />
          </div>
        </div>
        <div
          class="h-[calc(50%-40px)] bg-white/10 flex flex-col rounded-2xl w-full p-[26px]"
        >
          <span class="text-lg font-bold text-white">Все ресурсы</span>
          <div class="h-[calc(100%-28px)]">
            <Doughnut
              :options="sslPieChartOptions as any"
              :data="allResourcesChartData as any"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss"></style>
