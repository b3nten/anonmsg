import { Model } from "./index";

let mock_api = <T>(value: T, time?: number) =>
  new Promise<T>((r) =>
    setTimeout(() => r(value), time ?? Math.random() * 3000),
  );

let model = new Model(
  {
    foo: "foo",
    bar: "bar",
    baz: "baz",
    c: "LOL",
  },
  {
    onUpdate: (n) => console.log(),
  },
);

let a = model.addTransaction(
  async (
    _,
    {
      applyOptimisticUpdates,
      dropOptimisticUpdates,
      optimisticUpdate,
      pendingTransactions,
      state,
      update,
    },
  ) => {
    optimisticUpdate((x) => {
      x.foo = "a update";
      x.bar = "a update";
    });
    await mock_api(undefined, 3000);
    update((x) => {
      x.foo = "a final";
      x.bar = "a final";
    });
  },
);

let b = model.addTransaction(
  async (
    _,
    {
      applyOptimisticUpdates,
      dropOptimisticUpdates,
      optimisticUpdate,
      pendingTransactions,
      state,
      update,
    },
  ) => {
    await mock_api(undefined, 500);
    console.log(state());
  },
);

let c = model.addTransaction(
  async (
    _,
    {
      applyOptimisticUpdates,
      dropOptimisticUpdates,
      optimisticUpdate,
      pendingTransactions,
      state,
      update,
    },
  ) => {
    optimisticUpdate((m) => {
      m.c = "CCCC";
    });
  },
);

a.run();
b.run();
c.run();
