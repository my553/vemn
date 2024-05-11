<script setup lang="ts">
import {
  FlexRender,
  getCoreRowModel,
  useVueTable,
  createColumnHelper,
  getPaginationRowModel,
} from "@tanstack/vue-table";
import { computed, h, onMounted, reactive, ref, watchEffect } from "vue";
import Pagination from "./Pagination.vue";

export interface Statistics {
  resource: string;
  ip: string;
  resolve: string;
  waf: string;
  ssl: string;
  date: string;
}

const props = defineProps<{ data: Statistics[] }>();

const columnHelper = createColumnHelper<Statistics>();

const columns = [
  columnHelper.accessor("resource", {
    header: () => "Ресурс",
    cell: (info) => info.getValue(),
  }),
  columnHelper.accessor("ip", {
    cell: (info) => info.getValue(),
    header: () => "IP",
  }),
  columnHelper.accessor("resolve", {
    header: () => "Статус активности",
    cell: (info) => info.getValue(),
  }),
  columnHelper.accessor("waf", {
    header: () => "Активен",
    cell: (info) => info.getValue(),
  }),
  columnHelper.accessor("ssl", {
    header: "SSL",
    cell: (info) => info.getValue(),
  }),
  columnHelper.accessor("date", {
    header: "Дата",
    cell: (info) => info.getValue(),
  }),
];

const rowsPerPage = 10;
const currentPage = ref(1);

const table = useVueTable({
  data: props.data,
  columns,
  pageCount: Math.ceil(props.data.length / rowsPerPage),
  initialState: {
    pagination: {
      pageIndex: currentPage.value,
      pageSize: rowsPerPage,
    },
  },
  getCoreRowModel: getCoreRowModel(),
  getPaginationRowModel: getPaginationRowModel(),
});

onMounted(() => {});

const handlePageClick = (page: number) => {
  currentPage.value = page;
  table.setPageIndex(page - 1);
};
</script>

<template>
  <div class="w-full p-[22px] bg-white/10 rounded-xl flex items-end flex-col">
    <table class="w-full flex flex-col gap-3">
      <thead class="w-full">
        <tr
          class="grid grid-cols-6"
          v-for="headerGroup in table.getHeaderGroups()"
          :key="headerGroup.id"
        >
          <th
            class="w-full overflow-hidden"
            v-for="header in headerGroup.headers"
            :key="header.id"
            :colSpan="header.colSpan"
          >
            <span
              class="block text-start line-clamp-2 overflow-ellipsis text-main-blue"
            >
              <FlexRender
                v-if="!header.isPlaceholder"
                :render="header.column.columnDef.header"
                :props="header.getContext()"
              />
            </span>
          </th>
        </tr>
      </thead>
      <tbody class="flex flex-col gap-4">
        <tr
          class="grid grid-cols-6 border-b-[1px] border-solid border-white pb-4"
          v-for="row in table.getRowModel().rows"
          :key="row.id"
        >
          <td
            class="w-full"
            v-for="cell in row.getVisibleCells()"
            :key="cell.id"
          >
            <span
              class="block text-start line-clamp-2 overflow-ellipsis text-white"
            >
              <FlexRender
                :render="cell.column.columnDef.cell"
                :props="cell.getContext()"
              />
            </span>
          </td>
        </tr>
      </tbody>
    </table>
    <div class="mt-2.5">
      <Pagination
        :total-rows="data.length"
        :rows-per-page="rowsPerPage"
        :current-page="currentPage"
        :on-page-click="handlePageClick"
      />
    </div>
  </div>
</template>

<style scoped lang="scss"></style>
