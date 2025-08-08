export let mrgclx = (...classes: any[]) => classes.filter(Boolean).join(" ");

export let defstr = (value: any, str: string) => (value ? str : undefined);
