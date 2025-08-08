import api from "$lib";
import { defineTransaction, Model } from "$lib/state";
import toast from "svelte-french-toast";
import type { PageProps } from "./$types";
import { goto } from "$app/navigation";

interface InboxModel {
  inbox: PageProps["data"]["inbox"] & {
    pending_delete?: boolean;
  };
  messages: NonNullable<PageProps["data"]["messages"]>;
}

let set_public_key = defineTransaction<InboxModel>()(
  async (cond: boolean, { optimisticUpdate: optimistic, update, state }) => {
    optimistic((x) => {
      x.inbox.public_key = cond ? "creating..." : "";
    });
    if (cond) {
      let result = await api().POST("/v1/inbox/{private_key}/set-public-key", {
        params: {
          path: {
            private_key: state().inbox.private_key,
          },
        },
      });
      if (result.data) {
        update((x) => {
          x.inbox.public_key = result.data.public_key;
        });
      } else {
        toast.error("Could not create a public key. Please try again.");
      }
    } else {
      let result = await api().POST(
        "/v1/inbox/{private_key}/remove-public-key",
        {
          params: {
            path: {
              private_key: state().inbox.private_key,
            },
          },
        },
      );
      if (result.error) {
        toast.error(
          "Could not remove the public key from the inbox. Please try again.",
        );
      } else {
        update((x) => {
          x.inbox.public_key = "";
        });
      }
    }
  },
);

let delete_msg = defineTransaction<InboxModel>()(
  async (
    message_id: number,
    {
      state,
      optimisticUpdate: optimistic,
      applyOptimisticUpdates: applyOptimistic,
    },
  ) => {
    optimistic((x) => {
      x.messages = x.messages.filter((m) => m.id !== message_id);
    });
    let result = await api().DELETE("/v1/inbox/{private_key}/message", {
      params: {
        path: {
          private_key: state().inbox.private_key,
        },
      },
      body: {
        message_id,
      },
    });
    if (result.error) {
      toast.error("Failed to delete message. Please try again.");
      return;
    } else {
      applyOptimistic();
    }
  },
);

let change_private_key = defineTransaction<InboxModel>()(
  async (_, { optimisticUpdate: optimistic, update, state }) => {},
);

let set_inbox_active_status = defineTransaction<InboxModel>()(
  async (cond: boolean, { optimisticUpdate: optimistic, update, state }) => {
    optimistic((x) => {
      x.inbox.active = cond;
    });
    let result = await api().POST("/v1/inbox/{private_key}/active", {
      params: {
        path: {
          private_key: state().inbox.private_key,
        },
        query: {
          status: cond,
        },
      },
    });
    if (result.error) {
      toast.error(
        "Could not set the active status on this inbox. Please try again.",
      );
    } else {
      update((x) => {
        x.inbox.active = result.data.active;
      });
    }
  },
);

let delete_inbox = defineTransaction<InboxModel>()(
  async (_, { state, optimisticUpdate: optimistic }) => {
    let promise = Promise.withResolvers<void>();
    toast.promise(promise.promise, {
      error: "error deleting inbox",
      loading: "deleting inbox...",
      success: "successfully deleted inbox",
    });

    let request = await api().DELETE("/v1/inbox/{private_key}", {
      params: {
        path: {
          private_key: state().inbox.private_key,
        },
      },
    });

    if (request.error) {
      promise.reject();
    } else {
      promise.resolve();
      goto("/");
    }
  },
);

let refresh = defineTransaction<InboxModel>()(
  async (user_triggered: boolean, { update, state }) => {
    let response = await api().GET("/v1/inbox/{private_key}/messages", {
      params: {
        path: {
          private_key: state().inbox.private_key,
        },
      },
    });

    let toastPromise = Promise.withResolvers();
    if (user_triggered) {
      toast.promise(toastPromise.promise, {
        loading: "refreshing messages",
        error: "error refreshing messages",
        success: "messsages refreshed",
      });
    }

    if (response.error) {
      toastPromise.reject();
    } else {
      toastPromise.resolve(response.error);
      update((m) => {
        m.messages = response.data.messages ?? [];
      });
    }
  },
);

export let newInboxModel = (props: PageProps["data"]) => {
  let initial_state = {
    inbox: props.inbox,
    messages: props.messages ?? [],
  };
  let state = $state.raw(initial_state);
  let model = new Model<InboxModel>(initial_state, {
    onUpdate(ns) {
      state = ns;
    },
  });
  return {
    get state() {
      return state;
    },
    set_public_key: model.addTransaction(set_public_key),
    change_private_key: model.addTransaction(change_private_key),
    delete_msg: model.addTransaction(delete_msg),
    set_inbox_active_status: model.addTransaction(set_inbox_active_status),
    delete_inbox: model.addTransaction(delete_inbox),
    refresh: model.addTransaction(refresh),
  };
};
