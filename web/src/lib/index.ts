import createClient from "openapi-fetch";
import type { paths } from "../../schema";
import { env } from "$env/dynamic/public";

let api = (args: { fetch?: any } = {}) =>
	createClient<paths>({
		baseUrl: env.PUBLIC_API_URL,
		fetch: args.fetch,
	});

export default api;
