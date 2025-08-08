<script lang="ts">
    import type { HTMLButtonAttributes } from "svelte/elements";
    import { mrgclx, defstr } from "./util.js";

    // #tw-safelist
    // btn-neutral btn-primary btn-secondary btn-accent btn-info btn-success btn-warning btn-error
    // btn-outline btn-soft btn-dash btn-ghost btn-link
    // btn-xs btn-sm btn-md btn-lg btn-xl
    // btn-wide btn-block btn-square btn-circle btn-active btn-disabled

    interface ButtonProps extends HTMLButtonAttributes {
        variant?:
            | "neutral"
            | "primary"
            | "secondary"
            | "accent"
            | "info"
            | "success"
            | "warning"
            | "error";
        style?: "outline" | "soft" | "dash" | "ghost" | "link";
        size?: "xs" | "sm" | "md" | "lg" | "xl";
        wide?: boolean;
        block?: boolean;
        square?: boolean;
        circle?: boolean;
        active?: boolean;
    }

    let {
        variant,
        style,
        size,
        wide,
        block,
        square,
        circle,
        active,
        disabled,
        class: className = "",
        children,
        ...restProps
    }: ButtonProps = $props();

    const classes = $derived(
        mrgclx(
            "btn",
            defstr(variant, `btn-${variant}`),
            defstr(style, `btn-${style}`),
            defstr(size, `btn-${size}`),
            defstr(wide, "btn-wide"),
            defstr(block, "btn-block"),
            defstr(square, "btn-square"),
            defstr(circle, "btn-circle"),
            defstr(active, "btn-active"),
            defstr(disabled, "btn-disabled"),
            className,
        ),
    );
</script>

<button class={classes} {disabled} {...restProps}>
    {@render children?.()}
</button>
