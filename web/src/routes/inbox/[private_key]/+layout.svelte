<script lang="ts">
	let { children, data } = $props();
	import { newInboxModel, setInboxModelContext } from "./model.svelte";
	let model = newInboxModel(data);
	setInboxModelContext(model);
	setInterval(() => model.refresh.run(false), 10000);
</script>

{@render children()}
