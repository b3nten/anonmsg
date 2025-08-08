<script lang="ts">
  import Message from "$lib/inbox/message.svelte";
  import ApiEndpoint from "$lib/landing/ApiEndpoint.svelte";
  import Nav from "$lib/landing/Nav.svelte";
  import TerminalButton from "$lib/landing/TerminalButton.svelte";
  import TerminalPrompt from "$lib/landing/TerminalPrompt.svelte";
  import type { PageProps } from "./$types";
  import { newInboxModel } from "./model.svelte";
  import { refreshAll } from "$app/navigation";
  let { data }: PageProps = $props();
  let model = newInboxModel(data);

  setInterval(model.refresh.run, 10000);
</script>

<svelte:head>
  <title>anonmsg: inbox</title>
</svelte:head>

<Nav />

<main class="flex flex-col items-center font-mono mt-16 p-4 md:p-8">
  <div class="w-full max-w-5xl">
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
          # inbox
        </h1>
        <div class="flex space-x-2 items-center">
          <TerminalButton
            onclick={() => model.refresh.run(true)}
            size="sm"
            variant="secondary">refresh</TerminalButton
          >
          <TerminalButton
            href={`${model.state.inbox.private_key}/settings`}
            size="sm"
            variant="ghost"
          >
            settings
          </TerminalButton>
        </div>
      </div>
    </div>

    <div class="py-4">
      <TerminalPrompt command="./get_messages.sh" />
    </div>
    <div class="mx-auto space-y-4 w-full min-w-4/5">
      {#each model.state.messages as msg}
        <Message
          content={msg.content}
          timestamp={msg.created_at}
          delete_msg={() => model.delete_msg.run(msg.id)}
        />
      {/each}
      {#if model.state.messages.length < 1}
        <TerminalPrompt variant="red" command="no messages found" />
        <ApiEndpoint
          method="POST"
          endpoint={`https://api.anonmsg.dev/v1/send/{inbox_key}`}
          description="Get started"
          requestExample={`curl 'https://api.anonmsg.dev/v1/send/${model.state.inbox.private_key}' \\
  --request POST \\
  --header 'Content-Type: text/plain' \\
  --data 'Hello world!'`}
        />
      {/if}
    </div>
  </div>
</main>
