<script lang="ts">
import "../app.css";
import favicon from "$lib/assets/favicon.ico";

let { children } = $props();
import toast, { Toaster } from "svelte-french-toast";
</script>

<svelte:head>
  <!-- HTML Meta Tags -->
  <title>anonmsg - anon inbox api</title>
  <meta
    name="description"
    content="anonmsg is an anonymous inbox api for developers"
  />

  <!-- Facebook Meta Tags -->
  <meta property="og:url" content="https://anonmsg.dev" />
  <meta property="og:type" content="website" />
  <meta property="og:title" content="anonmsg - anon inbox api" />
  <meta
    property="og:description"
    content="anonmsg is an anonymous inbox api for developers"
  />
  <meta property="og:image" content="" />

  <!-- Twitter Meta Tags -->
  <meta name="twitter:card" content="summary_large_image" />
  <meta property="twitter:domain" content="anonmsg.dev" />
  <meta property="twitter:url" content="https://anonmsg.dev" />
  <meta name="twitter:title" content="anonmsg - anon inbox api" />
  <meta
    name="twitter:description"
    content="anonmsg is an anonymous inbox api for developers"
  />
  <meta name="twitter:image" content="" />

  <!-- Meta Tags Generated via https://www.opengraph.xyz -->
  <link rel="icon" href={favicon} />
</svelte:head>

<Toaster
  toastOptions={{
    className: "!bg-black !text-white border border-gray-700 font-mono",
  }}
/>
{@render children?.()}
