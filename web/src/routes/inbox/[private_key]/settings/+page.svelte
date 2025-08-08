<script lang="ts">
  import Dialog from "$lib/elements/Dialog.svelte";
  import Nav from "$lib/landing/Nav.svelte";
  import TerminalButton from "$lib/landing/TerminalButton.svelte";
  import TerminalPrompt from "$lib/landing/TerminalPrompt.svelte";
  import { newInboxModel } from "../model.svelte";
  import type { PageProps } from "./$types";
  let { data }: PageProps = $props();

  let model = newInboxModel(data);

  let newPrivateKey = $state("");
  let newPublicKey = $state("");
</script>

<svelte:head>
  <title>anonmsg: inbox settings</title>
</svelte:head>

<Nav />

<main class="flex flex-col items-center font-mono mt-16 p-4 md:p-8">
  <div class="w-full max-w-5xl space-y-12">
    <div class="mb-12">
      <div class="text-gray-500">
        // private_key {model.state.inbox.private_key}
      </div>
      <div class="text-gray-500">
        // public_key {model.state.inbox.public_key.length
          ? model.state.inbox.public_key
          : "not set"}
      </div>
      <div class="text-gray-500">
        // active {model.state.inbox.active ? "true" : "false"}
      </div>
      <div class="flex w-full items-center justify-between">
        <h1 class="text-4xl md:text-6xl font-bold text-white leading-tight">
          # settings
        </h1>
        <TerminalButton
          href={`/inbox/${model.state.inbox.private_key}`}
          size="sm"
          variant="secondary">messages</TerminalButton
        >
      </div>
    </div>

    <div class="space-y-4">
      <p class="text-gray-500">// inactive inboxes cannot receive messages.</p>
      <div class="flex items-center space-x-8">
        {#if model.state.inbox.active}
          <p class="text-green-400">inbox is active</p>
          <TerminalButton
            variant="critical"
            size="sm"
            onclick={() => model.set_inbox_active_status.run(false)}
            >deactivate_inbox</TerminalButton
          >
        {:else}
          <p class="text-red-400">inbox is inactive</p>
          <TerminalButton
            variant="primary"
            size="sm"
            onclick={() => model.set_inbox_active_status.run(true)}
            >activate_inbox</TerminalButton
          >
        {/if}
      </div>
    </div>

    <div class="space-y-4">
      <p class="text-gray-500">
        // a public key provides send-only privilages. use it as a public api
        key.
      </p>
      <div class="flex flex-col justify-start items-start space-y-4">
        {#if model.state.inbox.public_key}
          <p class="text-green-400">
            public key enabled ({model.state.inbox.public_key})
          </p>
          <div>
            <TerminalButton
              variant="critical"
              size="sm"
              onclick={() => model.set_public_key.run(false)}
              >remove_public_key</TerminalButton
            >
            <TerminalButton
              size="sm"
              variant="secondary"
              onclick={() => model.set_public_key.run(true)}
              >change_key</TerminalButton
            >
          </div>
        {:else}
          <p class="text-red-400">public key disabled</p>
          <TerminalButton
            variant="primary"
            size="sm"
            onclick={() => model.set_public_key.run(true)}
            >add_public_key</TerminalButton
          >
        {/if}
      </div>
    </div>

    <div>
      <div class="mb-2">
        <TerminalPrompt command="./delete_inbox.sh" variant="red" />
      </div>
      <Dialog>
        {#snippet trigger(onclick)}
          <TerminalButton class="w-full" {onclick} size="md" variant="critical">
            delete inbox
          </TerminalButton>
        {/snippet}
        {#snippet content(close)}
          <div class="flex flex-col space-y-4">
            <p>
              Are you sure? This is a destructive action which will delete all
              messages.
            </p>
            <div class="flex space-x-2">
              <TerminalButton variant="secondary" onclick={close}>
                don't_delete
              </TerminalButton>
              <TerminalButton
                onclick={() => {
                  model.delete_inbox.run(undefined);
                  close();
                }}
                variant="critical"
              >
                delete_it
              </TerminalButton>
            </div>
          </div>
        {/snippet}
      </Dialog>
    </div>
  </div>
</main>
