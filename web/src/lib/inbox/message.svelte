<script lang="ts">
import TerminalButton from "$lib/landing/TerminalButton.svelte";
import type { MouseEventHandler } from "svelte/elements";

let props: {
	timestamp: number;
	content: string;
	delete_msg: MouseEventHandler<HTMLButtonElement>;
} = $props();

let date_string = $derived.by(() => {
	let date = new Date(props.timestamp * 1000);
	return `${date.getDay()}-${date.getMonth()}-${date.getFullYear()} @ ${date.getHours()}:${date.getMinutes()}`;
});
</script>

<div class="bg-black p-2 rounded-lg border border-gray-700">
  <div class="flex items-center justify-between mb-2">
    <div class="text-purple-400">{date_string}</div>
    <TerminalButton variant="critical" size="sm" onclick={props.delete_msg}>
      delete
    </TerminalButton>
  </div>
  <div class="bg-neutral-800 p-4 rounded mb-1 overflow-scroll">
    <div class="text-gray-500 text-sm mb-2"># content</div>
    <div class="text-green-400">
      {props.content}
    </div>
  </div>
</div>
