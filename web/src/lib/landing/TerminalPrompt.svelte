<script>
export let command = "";
export let user = "$";
export let variant = "green"; // 'green' | 'blue' | 'yellow' | 'purple' | 'red' | 'gray'

$: colorClass = variant === "gray" ? "text-gray-400" : `text-${variant}-400`;
</script>

<div class="{colorClass} font-mono">
	{user} {command}
</div>