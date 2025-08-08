<script lang="ts">
 import type { Snippet } from "svelte";
 import type { HTMLButtonAttributes } from "svelte/elements";

 let variants = {
  primary: "bg-green-400 text-black hover:bg-green-300",
  secondary: "border border-neutral-600 text-gray-300 hover:bg-neutral-800",
  ghost: "text-gray-400 hover:text-green-400",
  critical: "bg-red-400 text-black hover:bg-red-500",
 };

 let sizes = {
  sm: "px-4 py-2 text-sm",
  md: "px-6 py-3",
  lg: "px-8 py-4 text-lg",
 };

 let {
  variant = "primary",
  size = "md",
  href,
  disabled,
  newTab,
  onclick,
  class: className,
  children,
 }: {
  variant?: keyof typeof variants;
  size?: keyof typeof sizes;
  href?: string;
  disabled?: boolean;
  newTab?: boolean;
  onclick?: HTMLButtonAttributes["onclick"];
  class?: string;
  children: Snippet;
 } = $props();

 let baseClasses = "font-mono font-bold rounded transition-colors inline-block";
 let classes = `${className} ${baseClasses} ${variants[variant]} ${sizes[size]} ${disabled ? "opacity-50 cursor-not-allowed" : ""}`;
</script>

{#if href}
 <a
  {href}
  class={classes}
  class:pointer-events-none={disabled}
  target={newTab ? "_blank" : undefined}
 >
  {@render children()}
 </a>
{:else}
 <button class={classes} {disabled} {onclick}>
  {@render children()}
 </button>
{/if}
