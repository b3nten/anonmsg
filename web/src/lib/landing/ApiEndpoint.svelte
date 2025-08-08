<script>
export let method = "GET"; // 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH'
export let endpoint = "";
export let description = "";
export let requestExample = "";
export let responseExample = "";

const methodColors = {
	GET: "text-purple-400",
	POST: "text-green-400",
	PUT: "text-blue-400",
	DELETE: "text-red-400",
	PATCH: "text-yellow-400",
};

$: methodColorClass = methodColors[method] || "text-purple-400";
</script>

<div class="bg-black p-6 rounded-lg border border-gray-700">
  <div class="{methodColorClass} mb-2">## {method} {endpoint}</div>
  <div class="text-gray-300 mb-4">{description}</div>

  {#if requestExample}
    <div class="bg-neutral-800 p-4 rounded mb-4 overflow-scroll">
      <div class="text-gray-500 text-sm mb-2"># Request</div>
      <pre class="text-green-400">
{requestExample.trim()}
      </pre>
    </div>
  {/if}

  {#if responseExample}
    <div class="bg-neutral-800 p-4 rounded overflow-scroll">
      <div class="text-gray-500 text-sm mb-2"># Response</div>
      <pre class="text-blue-400">
{responseExample.trim()}
      </pre>
    </div>
  {/if}
</div>
