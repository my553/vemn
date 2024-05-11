<script setup lang="ts">
import { twMerge } from "tailwind-merge";
import { computed } from "vue";

const props = defineProps<{
  rowsPerPage: number;
  totalRows: number;
  currentPage: number;
  onPageClick: (page: number) => void;
}>();

const maxPages = 4;

const totalPages = computed(() =>
  Math.ceil(props.totalRows / props.rowsPerPage),
);

const isOverflow = computed(() => maxPages < totalPages.value);

const pages = computed(() => {
  if (isOverflow.value) {
    let start = Math.max(1, props.currentPage - maxPages + 2);

    start =
      start > totalPages.value - maxPages
        ? totalPages.value - maxPages + 1
        : start;

    let end = Math.min(start + maxPages, totalPages.value);

    end = start > totalPages.value - maxPages ? totalPages.value + 1 : end;

    return Array.from({ length: end - start }, (_, i) => i + start);
  }

  return Array.from(
    { length: Math.ceil(props.totalRows / props.rowsPerPage) },
    (_, i) => i + 1,
  );
});
</script>

<template>
  <div class="flex justify-center items-center gap-2">
    <button
      @click="onPageClick(Math.max(currentPage - 1, 1))"
      class="flex justify-center items-center w-[24px] h-[24px] text-white rounded-md"
    >
      <svg
        class="w-[16px] h-[16px]"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M15 19l-7-7 7-7"
        ></path>
      </svg>
    </button>
    <div class="flex justify-center items-center gap-2">
      <div
        class="flex items-center gap-2"
        v-if="isOverflow && currentPage >= maxPages"
      >
        <button
          @click="onPageClick(1)"
          :class="twMerge('w-[24px] h-[24px] text-sm text-white rounded-md')"
        >
          1
        </button>
        <span class="text-white">...</span>
      </div>
      <button
        @click="onPageClick(index)"
        :class="
          twMerge(
            'w-[24px] h-[24px] text-sm text-white rounded-md',
            currentPage === index ? 'bg-main-blue' : '',
          )
        "
        v-for="index in pages"
        :key="index"
      >
        {{ index }}
      </button>
      <div
        class="flex items-center gap-2"
        v-if="isOverflow && currentPage < totalPages - 1"
      >
        <span class="text-white">...</span>
        <button
          @click="onPageClick(totalPages)"
          :class="twMerge('w-[24px] h-[24px] text-sm text-white rounded-md')"
        >
          {{ totalPages }}
        </button>
      </div>
    </div>
    <button
      @click="onPageClick(Math.min(currentPage + 1, totalPages))"
      class="flex justify-center items-center w-[24px] h-[24px] text-white rounded-md"
    >
      <svg
        class="w-[16px] h-[16px]"
        fill="none"
        stroke="currentColor"
        viewBox="0 0 24 24"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M9 5l7 7-7 7"
        ></path>
      </svg>
    </button>
  </div>
</template>

<style scoped lang="scss"></style>
