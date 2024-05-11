<script setup>
import { onMounted, reactive, ref } from "vue";
import * as storage from "../../storage.js";
import * as httpClient from "../../httpClient.js";
import SectionTitle from "../SectionTitle.vue";

const accFIO = ref("");
const generalStatistic = reactive({});
const weekStatistic = reactive({});
const certs = reactive({});

onMounted(() => {
  // Берем данные для главного меню
  httpClient.Get("/api/general-statistic").then((response) => {
    storage.set("cache", response.data.cache);
    let resp = JSON.parse(response.data.body);

    generalStatistic.value = {
      date: resp.date,
      resources: resp.resources,
      deactivateResource: resp.deactivateResource,
      owners: resp.owners,
      waf: resp.waf,
    };
  });

  httpClient.Get("/api/week-statistic").then((response) => {
    let resp = JSON.parse(response.data.body);
    weekStatistic.value = {
      last: {
        noResCount: resp.last.no_resolve,
        newWafCount: resp.last.new_waf,
        noRes: resp.last.no_res_resource,
        newWaf: resp.last.new_waf_resource,
      },
      current: {
        noResCount: resp.current.no_resolve,
        newWafCount: resp.current.new_waf,
        noRes: resp.current.no_res_resource,
        newWaf: resp.current.new_waf_resource,
      },
    };
  });

  //берем данные для сертификатов
  httpClient.Get("/api/certificates").then((response) => {
    let resp = JSON.parse(response.data.body);
    console.log(resp.current, resp.next);
    certs.value = {
      certs: [...resp.current, ...resp.next],
      num: resp.current.length + resp.next.length,
    };
  });

  // создаем строки для списка ресурсов
  // const rowsContainer = document.getElementById("rowsContainer");
  // const dataArray = [
  //   {
  //     resource: "zni.pag.ase.com",
  //     ip: "255.255.255.255",
  //     resolve: "активен",
  //     waf: "за WAF",
  //     ssl: "*.soc.mosreg.ru",
  //     date: "09.10.2023",
  //   },
  //   {
  //     resource: "zni.pag.ase.com",
  //     ip: "255.255.255.255",
  //     resolve: "активен",
  //     waf: "за WAF",
  //     ssl: "*.soc.mosreg.ru",
  //     date: "09.10.2023",
  //   },
  //   {
  //     resource: "zni.pag.ase.com",
  //     ip: "255.255.255.255",
  //     resolve: "активен",
  //     waf: "за WAF",
  //     ssl: "*.soc.mosreg.ru",
  //     date: "09.10.2023",
  //   },
  // ];
  // dataArray.forEach((item) => {
  //   const rowDiv = document.createElement("div");
  //   rowDiv.classList.add("rows");
  //   Object.keys(item).forEach((key) => {
  //     const span = document.createElement("span");
  //     span.classList.add(key);
  //     span.textContent = item[key];
  //     rowDiv.appendChild(span);
  //   });
  //   const lineSpan = document.createElement("span");
  //   lineSpan.classList.add("line");
  //   rowDiv.appendChild(lineSpan);
  //   rowsContainer.appendChild(rowDiv);
  // });
});
</script>

<template>
  <SectionTitle><span>Главная</span></SectionTitle>
  <div class="flex flex-col pl-[88px] w-full">
    <div
      class="py-[40px] pl-[97px] grid grid-cols-2 col-gap gap-x-[87px] gap-y-[48px]"
    >
      <div class="flex flex-col gap-2">
        <div class="text-2xl font-bold text-white">
          Общее количество неактивных ресурсов -
          <span class="text-main-blue">{{
            generalStatistic.value?.deactivateResource
          }}</span>
        </div>
        <div class="text-md font-normal text-main-gray">
          Данные актуальны на {{ generalStatistic.value?.date }} 14:20
        </div>
      </div>

      <div class="flex flex-col gap-2">
        <div class="text-2xl font-bold text-white">
          Клиентов -
          <span class="text-main-blue">{{
            generalStatistic.value?.owners
          }}</span>
        </div>
        <div class="text-md font-normal text-main-gray">
          Данные актуальны на {{ generalStatistic.value?.date }} 14:20
        </div>
      </div>

      <div class="flex flex-col gap-2">
        <div class="text-2xl font-bold text-white">
          Общее количество ресурсов -
          <span class="text-main-blue">{{
            generalStatistic.value?.resources
          }}</span>
        </div>
        <div class="text-md font-normal text-main-gray">
          Данные актуальны на {{ generalStatistic.value?.date }} 14:20
        </div>
      </div>

      <div class="flex flex-col gap-2">
        <div class="text-2xl font-bold text-white">
          Ресурсов за WAF -
          <span class="text-main-blue">{{ generalStatistic.value?.waf }}</span>
        </div>
        <div class="text-md font-normal text-main-gray">
          Данные актуальны на {{ generalStatistic.value?.date }} 14:20
        </div>
      </div>
    </div>
    <SectionTitle small="true"><span>Главная</span></SectionTitle>
    <div
      class="py-[40px] pl-[97px] grid grid-cols-2 col-gap gap-x-[87px] gap-y-[48px]"
    >
      <div class="relative">
        <div class="text-2xl font-bold text-white">Прошлая неделя</div>
        <span
          class="absolute w-[1px] h-[84px] bg-white left-[-8px] top-[16px]"
        />
        <div class="flex items-center mt-4 text-main-gray">
          <span class="w-[101px] h-[1px] bg-white translate-x-[-8px]" />
          <span class="text-nowrap">
            Неактивные ресурсы {{ weekStatistic.value?.last.noResCount }}
          </span>
        </div>
        <div class="flex items-center mt-4 text-main-gray">
          <span class="w-[200px] h-[1px] bg-white translate-x-[-8px]" />
          <span class="text-nowrap">
            Новые WAF-ресурсы {{ weekStatistic.value?.last.newWafCount }}
          </span>
        </div>
      </div>
      <div class="relative">
        <div class="text-2xl font-bold text-white">Текущая неделя</div>
        <span
          class="absolute w-[1px] h-[84px] bg-white left-[-8px] top-[16px]"
        />
        <div class="flex items-center mt-4 text-main-gray">
          <span class="w-[101px] h-[1px] bg-white translate-x-[-8px]" />
          <span class="text-nowrap">
            Неактивные ресурсы {{ weekStatistic.value?.last.noResCount }}
          </span>
        </div>
        <div class="flex items-center mt-4 text-main-gray">
          <span class="w-[200px] h-[1px] bg-white translate-x-[-8px]" />
          <span class="text-nowrap">
            Новые WAF-ресурсы {{ weekStatistic.value?.last.newWafCount }}
          </span>
        </div>
      </div>
    </div>

    <SectionTitle small="true">
      <div>
        Сертификаты, срок действия которых закончится в ближайшие 2 месяца
        <span class="text-main-blue ml-2">{{ certs.value?.num }}</span>
      </div>
    </SectionTitle>
    <div class="py-[40px] pl-[97px] flex flex-col pr-[100px] gap-[20px]">
      <div class="flex justify-between">
        <div class="text-2xl font-bold text-main-blue">Ресурс</div>
        <div class="text-2xl font-bold text-main-blue">Дата</div>
      </div>
      <div
        class="grid grid-cols-[fit-content(90px)_1fr_fit-content(90px)] gap-4"
      >
        <div class="flex flex-col gap-4">
          <div
            class="text-md font-normal text-main-gray"
            v-for="cert in certs.value?.certs"
          >
            {{ cert.resource }}
          </div>
        </div>
        <div class="flex flex-col gap-4">
          <span
            class="flex items-center h-[24px]"
            v-for="_ in certs.value?.certs"
          >
            <span class="w-full h-[1px] bg-white" />
          </span>
        </div>
        <div class="flex flex-col gap-4">
          <div
            class="text-md font-normal text-main-gray"
            v-for="cert in certs.value?.certs"
          >
            {{ cert.date }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
