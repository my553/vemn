<script setup lang="ts">
import SectionTitle from "./SectionTitle.vue";
import { Bar } from "vue-chartjs";

const data = {
  labels: [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
  ],
  datasets: [
    {
      label: "Data One",
      backgroundColor: "#f87979",
      data: [40, 20, 12, 39, 10, 40, 39, 80, 40, 20, 12, 11],
    },
  ],
};

const options = {
  responsive: true,
  maintainAspectRatio: false,
};
</script>

<template>
  <div>
    <SectionTitle> Аналитика</SectionTitle>
    <Bar :options="options" :data="data" />
  </div>
</template>

<style scoped lang="scss"></style>
