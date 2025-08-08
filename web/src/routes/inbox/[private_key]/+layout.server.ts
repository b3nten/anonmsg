import api from "$lib";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({ params, fetch }) => {
	let inboxPromise = api({ fetch }).GET("/v1/inbox/{private_key}", {
		params: {
			path: {
				private_key: params.private_key,
			},
		},
	});

	let messagesPromise = api({ fetch }).GET("/v1/inbox/{private_key}/messages", {
		params: {
			path: {
				private_key: params.private_key,
			},
		},
	});

	let [inbox, messages] = await Promise.all([inboxPromise, messagesPromise]);

	if (inbox.error || messages.error) {
		// todo: make work properly
		throw new Error("Whoops!");
	}

	return {
		inbox: inbox.data,
		messages: messages.data.messages,
	};
};
