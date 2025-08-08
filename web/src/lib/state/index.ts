import { produce, type Objectish, enablePatches, type Patch } from "immer";

enablePatches();

class InvariantError extends Error {
  constructor(message: unknown, cause?: unknown) {
    super(
      `Oops! Something went wrong internally. Please report this bug.\n${String(message)}`,
      { cause },
    );
  }
}

function INVARIANT(cond: unknown, msg?: string): asserts cond {
  if (!cond) throw new InvariantError(msg);
}

class CallbackQueue {
  static enqueue = (callback: Function) => {
    if (this.#queued) return;
    this.#queued = true;
    queueMicrotask(() => {
      this.#queued = false;
      callback();
    });
  };
  static #queued = false;
}

type UpdateInstance = ((draft: unknown) => void) & {
  patch_type: "local" | "remote" | "optimistic";
};

let new_patch_instance = (
  update_fn: Function,
  type: "local" | "remote" | "optimistic",
): UpdateInstance => {
  (<UpdateInstance>update_fn).patch_type = type;
  return update_fn as UpdateInstance;
};

type TransactionFunction<Args = void, Store = Objectish> = (
  /**
   * Arguments passed to the operation.
   */
  args: Args,
  fn_args: {
    /**
     * @returns the state of the model, not including future state updates
     */
    state: () => Store;
    /**
     * Apply an optimistic update to the store. This will be rolled back when the operation is completed.
     * Must be a pure function.
     * @param fn
     * @returns void
     */
    optimisticUpdate: (fn: (draft: Store) => void) => void;
    /**
     * Update the store. This will be applied in order of operations.
     * Must be a pure function.
     * @param fn
     * @returns void
     */
    update: (fn: (draft: Store) => void) => void;
    /**
     * Converts previous optimistic updates tied to this transaction to regular updates which will be applied to the store.
     * It's normally best to reapply data from the server as that guarantees it's synced with the server but if you are sure
     * that it will match your optimistic update, this can save some work.
     * @returns void
     */
    applyOptimisticUpdates: () => void;
    /**
     * Drop all optimistic updates tied to this transaction.
     * @returns void
     */
    dropOptimisticUpdates: () => void;
    /**
     * A promise that resolves when there are no previous pending transactions.
     */
    pendingTransactions: Promise<void>;
  },
) => Promise<void>;

class TransactionInstance<Store = Objectish> {
  complete = false;
  current_state?: Store;
  prev_transaction_promise = Promise.withResolvers<void>();
  update_fns: Array<UpdateInstance> = [];
}

class Transaction<Args, Store extends Objectish> {
  constructor(
    private mutation_fn: TransactionFunction<Args, Store>,
    private model: Model<Store>,
  ) {}
  run = (args: Args) => {
    let op = new TransactionInstance<Store>();

    // add transaction to model
    let model = model_internals.get(this.model);
    INVARIANT(model, "MODEL NOT SET??");
    model.transaction_queue.push(op);

    // if our transaction is first in the queue, immediately resolve the previous transaction promise
    if (model.transaction_queue.indexOf(op) === 0) {
      op.prev_transaction_promise.resolve();
    }

    // set the current state to our previous state initially
    op.current_state = model.get_previous_state(op);

    // run mutation function
    this.mutation_fn(args, {
      state: () => model.get_previous_state(op),
      optimisticUpdate: (fn) => {
        if (op.complete) return;
        op.update_fns.push(
          new_patch_instance((draft: any) => void fn(draft), "optimistic"),
        );
        model.queue_update(op);
      },
      update: (fn) => {
        if (op.complete) return;
        op.update_fns.push(
          new_patch_instance((draft: any) => void fn(draft), "remote"),
        );
        model.queue_update(op);
      },
      applyOptimisticUpdates: () => {
        for (let patch of op.update_fns) {
          if (patch.patch_type === "optimistic") {
            patch.patch_type = "local";
          }
        }
      },
      dropOptimisticUpdates: () => {
        for (let i = op.update_fns.length - 1; i >= 0; i--) {
          if (op.update_fns[i].patch_type === "optimistic") {
            op.update_fns.splice(i, 1);
          }
        }
        model.queue_update(op);
      },
      pendingTransactions: op.prev_transaction_promise.promise,
    }).finally(() => {
      op.complete = true;
      model.queue_update(op);
    });
  };
}

class InternalModel<Store extends Objectish = Objectish> {
  constructor(args: {
    base_state: Store;
    config: {
      onUpdate?: (newState: Store, oldState: Store, patches: Patch[]) => void;
    };
  }) {
    this.base_state = args.base_state;
    this.current_state = args.base_state;
    this.config = args.config;
  }

  base_state: Store;
  current_state: Store;
  transaction_queue: TransactionInstance<Store>[] = [];
  config: {
    onUpdate?: (newState: Store, oldState: Store, patches: Patch[]) => void;
  };

  get_previous_state = (op: TransactionInstance<Store>): Store => {
    let op_index = this.transaction_queue.findIndex((x) => op === x);
    switch (op_index) {
      case -1: {
        throw new InvariantError(
          "InternalModel.get_previous_state(op): op not found.",
        );
      }
      case 0:
        return this.base_state;
      default:
        let result = this.transaction_queue.at(op_index - 1)?.current_state;
        INVARIANT(
          result,
          "InternalModel.get_previous_state(op): prev state not found",
        );
        return result;
    }
  };

  queue_update = (t?: TransactionInstance<Store>) => {
    CallbackQueue.enqueue(() => {
      // todo: pass in earliest provided state
      // correctness more important for now
      this.update_state();
    });
  };

  private update_state = (t?: TransactionInstance<Store>) => {
    let og_state = this.current_state;
    let patches: Array<Patch> = [];

    // find the transaction that queued an update,
    // or else start at the beginning
    let i = t ? this.transaction_queue.findIndex((x) => (x = t)) : 0;

    INVARIANT(i === 0, "NON-ZERO INDEX NOT SUPPORTED YET");

    for (; i < this.transaction_queue.length; ) {
      let op = this.transaction_queue[i];

      // handle any completed transactions
      // removing them until first reaching first uncompleted transaction
      if (op.complete && i === 0) {
        this.transaction_queue.shift(); // remove
        for (let patch_fn of op.update_fns) {
          if (patch_fn.patch_type !== "optimistic") {
            this.base_state = produce(this.base_state, patch_fn);
          }
        }
        this.current_state = this.base_state;
        // resolve the next transaction
        this.transaction_queue[0]?.prev_transaction_promise.resolve();
        continue;
      }

      let prev_state = this.get_previous_state(op);

      // we use the prev_state as a base to patch updates onto
      // each node has it's current state attached
      for (let patch_fn of op.update_fns) {
        // skip optimistics if complete
        if (op.complete && patch_fn.patch_type === "optimistic") {
          continue;
        }
        op.current_state = prev_state = produce(prev_state, patch_fn);
      }
      // move on to next transaction
      this.current_state = prev_state;
      i++;
    }
    this.config.onUpdate?.(this.current_state, og_state, patches);
  };
}

let model_internals = new WeakMap<Model<any>, InternalModel<any>>();

export class Model<T extends Objectish> {
  constructor(
    state: T,
    config: {
      onUpdate?: (newState: T, oldState: T, patches: Patch[]) => void;
    } = {},
  ) {
    model_internals.set(
      this,
      new InternalModel({
        base_state: state,
        config: config,
      }),
    );
  }

  /**
   * Get the current state of the model.
   */
  state = (): T => {
    return model_internals.get(this)!.current_state;
  };

  /**
   * Create an operation for the model
   * @param mutation_fn
   * @returns
   */
  addTransaction = <Args = void>(mutation_fn: TransactionFunction<Args, T>) =>
    new Transaction(mutation_fn, this);
}

export let defineTransaction =
  <Store>() =>
  <Args>(fn: TransactionFunction<Args, Store>) =>
    fn;
