<script>
export let maxWidth = "5xl"; // max-w-lg | max-w-xl | max-w-2xl | max-w-3xl | max-w-4xl | max-w-5xl | max-w-6xl
export let padding = "py-20"; // Custom padding classes
export let background = "bg-black"; // Background color

$: containerClass = `max-w-${maxWidth} mx-auto px-6 ${padding} font-mono`;
</script>

<section class={background}>
    <div class={containerClass}>
        <slot />
    </div>
</section>
