# AnonMsg

AnonMsg is a simple API for sending and receiving anonymous messages without an account. Users can create an inbox tied to a private key, allowing them to receive messages anonymously. This is useful for things like contact forms on websites where you don't want to set up an email service or sign up for paid services. It is intentionally minimal, focusing on privacy and ease of use. The API does not collect any personal information.

## Features

- Create an inbox with a private key
- Send anonymous messages to any inbox
- Retrieve messages from your inbox using the private key
- Simple RESTful API
- No user accounts or personal information required

## Self Hosting

AnonMsg API is simple to self host. It's a Go server and a single postgres database. There is a docker-compose file to get you started quickly.

The frontend is a SvelteKit application that can be served separately or together with the API.
Refer to [SvelteKit's documentation](https://svelte.dev/) for [build](https://svelte.dev/docs/kit/building-your-app) and [deployment](https://svelte.dev/docs/kit/adapters) options.
