<script setup lang="ts">
import {
  FlexRender,
  getCoreRowModel,
  useVueTable,
  createColumnHelper,
  getPaginationRowModel,
} from "@tanstack/vue-table";
import {
  computed,
  h,
  onMounted,
  onUpdated,
  reactive,
  ref,
  watch,
  watchEffect,
} from "vue";
import Pagination from "./Pagination.vue";
import Vulnerability from "./Vulnerability.vue";

export interface Statistics {
  resource: string;
  ip: string;
  resolve: string;
  waf: string;
  ssl: string;
  lastScanDate: string;
  vulnerability: {
    critical: number;
    high: number;
    medium: number;
    light: number;
    inform: number;
  };
}

const props = defineProps<{ data: Statistics[]; rowsCount: number }>();

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
  columnHelper.accessor("vulnerability", {
    header: "Уязвимости",
    cell: (info) => h(Vulnerability, { vulnerability: info.getValue() }),
  }),
  columnHelper.accessor("lastScanDate", {
    header: "Дата сканирования",
    cell: (info) => info.getValue(),
  }),
];

const table = useVueTable({
  get data() {
    return props.data;
  },
  columns,
  initialState: {
    pagination: {
      pageSize: 5,
      pageIndex: 0,
    },
  },
  getCoreRowModel: getCoreRowModel(),
  getPaginationRowModel: getPaginationRowModel(),
});

watch(
  () => props.rowsCount,
  () => {
    table.setPageSize(props.rowsCount);
  },
);

const handlePageClick = (page: number) => {
  table.setPageIndex(page - 1);
};
</script>

<template>
  <table class="w-full flex flex-col gap-3">
    <thead class="w-full">
      <tr
        class="grid grid-cols-7"
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
        class="grid grid-cols-7 border-b-[1px] border-solid border-white pb-4"
        v-for="row in table.getRowModel().rows"
        :key="row.id"
      >
        <td class="w-full" v-for="cell in row.getVisibleCells()" :key="cell.id">
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
      :rows-per-page="props.rowsCount"
      :current-page="table.getState().pagination.pageIndex + 1"
      :on-page-click="handlePageClick"
    />
  </div>
</template>

<style scoped lang="scss"></style>
