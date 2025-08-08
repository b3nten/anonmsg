<script lang="ts">
import type { HTMLDialogAttributes } from "svelte/elements";
import { mrgclx, defstr } from "./util.js";
import type { Snippet } from "svelte";

// #tw-safelist
// modal modal-box modal-action modal-backdrop
// modal-top modal-middle modal-bottom modal-start modal-end

interface DialogProps extends HTMLDialogAttributes {
	trigger: Snippet<[onclick: () => void]>;
	placement?: "top" | "middle" | "bottom" | "start" | "end";
	content: Snippet<[close: () => void]>;
}

let {
	trigger,
	placement,
	class: className = "",
	content,
	...restProps
}: DialogProps = $props();

let dialog: HTMLDialogElement | undefined = $state(undefined);

let onclick = () => {
	if (!dialog?.open) {
		dialog?.showModal();
	}
};

let close = () => {
	if (dialog?.open) {
		dialog.close();
	}
};

const classes = $derived(
	mrgclx("modal", defstr(placement, `modal-${placement}`), className),
);
</script>

{@render trigger(onclick)}
<dialog bind:this={dialog} class={classes} {...restProps}>
  <div class="modal-box outline outline-gray-700 rounded-2xl">
    {@render content?.(close)}
  </div>
</dialog>
