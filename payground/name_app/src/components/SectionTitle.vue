<script setup lang="ts">
import { cva, type VariantProps } from "class-variance-authority";

const title = cva(
  [
    "bg-gradient-to-r w-full from-[rgba(255,255,255,0.7)_11.69%] to-[rgba(255,255,255,0)_98.85%]",
    "[&>*]:font-bold [&>*]:text-white",
  ],
  {
    variants: {
      size: {
        sm: "py-[10px] pl-[97px] [&>*]:text-[20px]",
        md: "py-[19px] pl-[97px] [&>*]:text-[30px]",
      },
      rounded: {
        none: "",
        md: "rounded-2xl",
      },
    },
    defaultVariants: {
      size: "md",
      rounded: "none",
    },
  },
);

type ButtonProps = VariantProps<typeof title>;

defineProps<{ size?: ButtonProps["size"]; rounded?: ButtonProps["rounded"] }>();
</script>

<template>
  <div :class="title({ size, rounded })">
    <h2 class="text">
      <slot />
    </h2>
  </div>
</template>
