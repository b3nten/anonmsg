<script lang="ts">
  import { env } from "$env/dynamic/public";
  import { goto } from "$app/navigation";
  import api from "$lib";
  import { TerminalButton } from "$lib/landing";
  import { onMount } from "svelte";
  import Dialog from "$lib/elements/Dialog.svelte";

  let mobileMenuOpen = $state(false);
  let mobileMenuRef: HTMLDivElement | undefined = $state(undefined);
  let hamburgerButtonRef: HTMLButtonElement | undefined = $state(undefined);

  let create_inbox = async () => {
    let response = await api().POST("/v1/inbox/");
    if (response.error) {
      alert(response.error.title);
    } else {
      goto("/inbox/" + response.data.private_key);
    }
  };

  let toggle_mobile_menu = () => {
    mobileMenuOpen = !mobileMenuOpen;
    if (mobileMenuOpen) {
      document.body.style.overflow = "hidden";
      setTimeout(() => {
        const firstFocusable = mobileMenuRef?.querySelector("button, a");
        firstFocusable?.focus();
      }, 100);
    } else {
      document.body.style.overflow = "";
      hamburgerButtonRef?.focus();
    }
  };

  let close_mobile_menu = () => {
    mobileMenuOpen = false;
    document.body.style.overflow = "";
    hamburgerButtonRef?.focus();
  };

  let handle_keydown = (event: KeyboardEvent) => {
    if (event.key === "Escape" && mobileMenuOpen) {
      close_mobile_menu();
    }

    if (mobileMenuOpen && event.key === "Tab") {
      const focusableElements = mobileMenuRef?.querySelectorAll(
        'button, a, [tabindex]:not([tabindex="-1"])',
      );
      if (focusableElements && focusableElements.length > 0) {
        const firstElement = focusableElements[0];
        const lastElement = focusableElements[focusableElements.length - 1];

        if (event.shiftKey && document.activeElement === firstElement) {
          event.preventDefault();
          lastElement.focus();
        } else if (!event.shiftKey && document.activeElement === lastElement) {
          event.preventDefault();
          firstElement.focus();
        }
      }
    }
  };

  onMount(() => {
    document.addEventListener("keydown", handle_keydown);
    return () => {
      document.removeEventListener("keydown", handle_keydown);
      document.body.style.overflow = "";
    };
  });
</script>

<nav class="bg-black border-b border-gray-800 sticky top-0 z-50">
  <div class="max-w-7xl mx-auto px-4 sm:px-6">
    <div class="flex items-center justify-between h-16 font-mono">
      <a
        href="/"
        class="text-green-400 font-bold text-lg"
        onclick={close_mobile_menu}
      >
        <span class="text-gray-500 px-1">#</span>anonmsg
      </a>

      <!-- Desktop Navigation -->
      <div class="hidden md:flex items-center space-x-8">
        <!-- <TerminalButton href="/self_host.md" variant="ghost" size="sm">
     self_hosting.md
    </TerminalButton> -->
        <TerminalButton
          href={`${env.PUBLIC_API_URL}/docs`}
          variant="ghost"
          size="sm"
          newTab={true}
        >
          api_docs
        </TerminalButton>
        <TerminalButton
          href="https://github.com/b3nten/anonmsg"
          variant="ghost"
          size="sm"
          newTab={true}
        >
          github
        </TerminalButton>
        <TerminalButton variant="primary" size="sm" onclick={create_inbox}>
          ./create_inbox.sh
        </TerminalButton>
      </div>

      <!-- Mobile Hamburger Button -->
      <button
        bind:this={hamburgerButtonRef}
        onclick={toggle_mobile_menu}
        class="md:hidden inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-green-400 hover:bg-gray-800 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-green-400"
        aria-expanded={mobileMenuOpen}
        aria-controls="mobile-menu"
        aria-label={mobileMenuOpen ? "Close main menu" : "Open main menu"}
      >
        <span class="sr-only"
          >{mobileMenuOpen ? "Close main menu" : "Open main menu"}</span
        >
        {#if !mobileMenuOpen}
          <svg
            class="block h-6 w-6"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            aria-hidden="true"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4 6h16M4 12h16M4 18h16"
            />
          </svg>
        {:else}
          <svg
            class="block h-6 w-6"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            aria-hidden="true"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        {/if}
      </button>
    </div>
  </div>
</nav>

<!-- Fullscreen Mobile Menu Overlay -->
{#if mobileMenuOpen}
  <div
    class="fixed inset-0 z-50 md:hidden"
    role="dialog"
    aria-modal="true"
    aria-labelledby="mobile-menu-title"
  >
    <!-- Background overlay -->
    <div
      class="fixed inset-0 bg-black bg-opacity-75 transition-opacity"
      onclick={close_mobile_menu}
      onkeydown={(e) => e.key === "Enter" && close_mobile_menu()}
      tabindex="-1"
    ></div>

    <!-- Menu panel -->
    <div
      bind:this={mobileMenuRef}
      id="mobile-menu"
      class="fixed inset-0 flex flex-col bg-black font-mono"
    >
      <!-- Header with close button -->
      <div
        class="flex items-center justify-between h-16 px-4 border-b border-gray-800"
      >
        <span id="mobile-menu-title" class="text-green-400 font-bold text-lg">
          <span class="text-gray-500 px-1">#</span>anonmsg
        </span>
        <button
          onclick={close_mobile_menu}
          class="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-green-400 hover:bg-gray-800 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-green-400"
          aria-label="Close menu"
        >
          <svg
            class="h-6 w-6"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            aria-hidden="true"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M6 18L18 6M6 6l12 12"
            />
          </svg>
        </button>
      </div>

      <!-- Menu items -->
      <div class="flex-1 px-4 py-8 space-y-6">
        <div class="space-y-4">
          <TerminalButton
            href={`${env.PUBLIC_API_URL}/docs`}
            variant="ghost"
            size="lg"
            newTab={true}
            onclick={close_mobile_menu}
            class="w-full justify-start text-lg"
          >
            api_docs
          </TerminalButton>

          <TerminalButton
            href="https://github.com"
            variant="ghost"
            size="lg"
            newTab={true}
            onclick={close_mobile_menu}
            class="w-full justify-start text-lg"
          >
            github
          </TerminalButton>
        </div>

        <div class="pt-8">
          <TerminalButton
            variant="primary"
            size="lg"
            onclick={() => {
              create_inbox();
              close_mobile_menu();
            }}
            class="w-full justify-center text-lg"
          >
            ./create_inbox.sh
          </TerminalButton>
        </div>
      </div>
    </div>
  </div>
{/if}
