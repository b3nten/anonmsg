<script>
export let title = "";
export let language = "";
export let variant = "dark"; // 'dark' | 'darker'

$: bgClass = variant === "darker" ? "bg-black" : "bg-neutral-800";
$: borderClass =
	variant === "darker" ? "border-gray-700" : "border-neutral-600";
$: languageColor =
	language === "bash"
		? "text-green-400"
		: language === "json"
			? "text-blue-400"
			: language === "http"
				? "text-yellow-400"
				: "text-purple-400";
</script>

<div class="{bgClass} p-6 rounded-lg border {borderClass}">
    {#if title}
        <div class="text-gray-500 text-sm mb-2">{title}</div>
    {/if}
    {#if language}
        <div class="{languageColor} text-sm font-mono mb-2 uppercase">
            {language}
        </div>
    {/if}
    <div class="font-mono text-sm">
        <slot />
    </div>
</div>
