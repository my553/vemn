<script setup>
import { cva } from "class-variance-authority";

const button = cva(
  ["flex rounded-2xl h-fit items-center gap-2 text-md text-white px-4 py-2.5"],
  {
    variants: {
      variant: {
        outline: "border border-main-blue bg-transparent",
        main: "bg-main-blue",
      },
    },
    defaultVariants: { variant: "main" },
  },
);

defineProps(["variant"]);
</script>

<template>
  <button v-bind="$props" :class="button({ variant })">
    <slot />
  </button>
</template>
