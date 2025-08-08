<script lang="ts">
  import api from "$lib";
  import { goto } from "$app/navigation";
  import {
    TerminalBlock,
    TerminalButton,
    TerminalPrompt,
    TerminalSection,
    ApiEndpoint,
    Nav,
  } from "$lib/landing";

  let create_inbox = async () => {
    let response = await api().POST("/v1/inbox/");
    if (response.error) {
      alert(response.error.title);
    } else {
      goto("/inbox/" + response.data.private_key);
    }
  };
</script>

<Nav />

<!-- Hero Section -->
<section class="bg-black py-48 flex items-center">
  <div class="max-w-4xl mx-auto px-6">
    <div class="font-mono">
      <TerminalPrompt command="cat README.md" />
      <h1
        class="text-4xl md:text-6xl font-bold text-white mb-8 leading-tight mt-4"
      >
        # anonmsg
      </h1>
      <div class="text-lg md:text-xl text-gray-300 mb-8 space-y-2">
        <div>Anonymous inbox API. No signup, no bs.</div>
        <div class="text-gray-500">
          // perfect for contact forms, feedback, whatever
        </div>
      </div>

      <TerminalBlock title="Quick start:">
        <div class="text-green-400">
          curl -X POST https://api.anonmsg.dev/v1/inbox
        </div>
        <div class="text-gray-400 mt-2">
          # returns: {`{ "private_key": "xyz789" }`}
        </div>
      </TerminalBlock>

      <div class="flex flex-col sm:flex-row gap-4 mt-8">
        <TerminalButton variant="primary" onclick={create_inbox}>
          ./create_inbox.sh
        </TerminalButton>
        <TerminalButton
          href="https://api.anonmsg.dev/docs"
          variant="secondary"
          newTab={true}
        >
          api_docs.html
        </TerminalButton>
      </div>
    </div>
  </div>
</section>

<!-- API Documentation -->
<TerminalSection>
  <div class="mb-12">
    <TerminalPrompt command="cat api_docs.md" />
    <h2 class="text-2xl text-white mb-4 mt-2"># API Reference</h2>
    <div class="text-gray-400">// three (main) endpoints. That's it.</div>
  </div>

  <div class="space-y-8">
    <ApiEndpoint
      method="POST"
      endpoint="/api/inbox"
      description="Creates a new inbox"
      requestExample={`curl -X POST https://api.anonmsg.dev/v1/inbox \\
   -H "Content-Type: application/json"`}
      responseExample={`{
  "private_key": "0KKLJKJDKJCLDJCDKLCJKLDJLCD"
}`}
    />

    <ApiEndpoint
      method="POST"
      endpoint={`/api/send/{inboxID}`}
      description="Send message to inbox"
      requestExample={`curl -X POST https://api.anonmsg.dev/v1/send/{inbox_key} \\
  -H "Content-Type: application/json" \\
  -d { "greeting": "Hello from anonymous!" }`}
      responseExample={`{
  "message": {
    "id": 3467273,
    "content": "{ \\"greeting\\": \\"Hello from anonymous!\\" }",
    "created_at": 1754633348
  }
}`}
    />

    <ApiEndpoint
      method="GET"
      endpoint="/api/inbox/:inbox_id"
      description="Read messages (requires private key)"
      requestExample={`curl https://api.anonmsg.dev/inbox/{private_key}/messages`}
      responseExample={`{
  "messages": [
    {
      "id": 3467273,
      "content": "{ \\"greeting\\": \\"Hello from anonymous!\\" }",
      "created_at": 1754633348
    }
  ]
}`}
    />
  </div>
</TerminalSection>

<!-- Use Cases -->
<TerminalSection>
  <div class="mb-12">
    <TerminalPrompt command="grep -r &quot;use.*case&quot; examples/" />
    <h2 class="text-2xl text-white mb-4 mt-2"># Common Use Cases</h2>
  </div>

  <div class="grid md:grid-cols-2 gap-8">
    <TerminalBlock variant="darker">
      <div class="text-yellow-400 mb-3">// contact forms</div>
      <div class="text-gray-300 mb-4">websites, portfolios, landing pages</div>
      <TerminalBlock variant="darker">
        <pre class="text-green-400">
{`<form
  onClick={(e) => postToAnonMsg(e)}
/>`}</pre>
      </TerminalBlock>
    </TerminalBlock>

    <TerminalBlock variant="darker">
      <div class="text-yellow-400 mb-3">// feedback collection</div>
      <div class="text-gray-300 mb-4">bug reports, suggestions, feedback</div>
      <TerminalBlock variant="darker">
        <div class="text-blue-400">
          fetch(anonmsg_api_url, {`{`}
        </div>
        <div class="text-blue-400 ml-4">method: 'POST', body: feedback</div>
        <div class="text-blue-400">{`})`}</div>
      </TerminalBlock>
    </TerminalBlock>

    <TerminalBlock variant="darker">
      <div class="text-yellow-400 mb-3">// webhooks</div>
      <div class="text-gray-300 mb-4">receive notifications, alerts</div>
      <TerminalBlock variant="darker">
        <div class="text-purple-400">curl -X POST webhook_inbox \</div>
        <div class="text-purple-400 ml-4">-d "Server down: $hostname"</div>
      </TerminalBlock>
    </TerminalBlock>

    <TerminalBlock variant="darker">
      <div class="text-yellow-400 mb-3">// anything else</div>
      <div class="text-gray-300 mb-4">and anything else you want to save</div>
      <TerminalBlock variant="darker">
        <div class="text-green-400">#no account</div>
        <div class="text-red-400">#no tracking</div>
      </TerminalBlock>
    </TerminalBlock> -->
  </div>

  <div class="mt-16 text-center">
    <TerminalBlock>
      <div class="text-center">
        <TerminalPrompt command="./start_hacking.sh" />
        <div class="text-white text-xl mb-6 mt-4">Ready to build?</div>
        <TerminalButton variant="primary" size="lg" onclick={create_inbox}
          >Create inbox</TerminalButton
        >
        <div class="text-gray-500 text-sm mt-4"># Takes ~0.1 seconds</div>
      </div>
    </TerminalBlock>
  </div>
</TerminalSection>

<!-- Footer -->
<footer class="bg-black border-t border-gray-800">
  <TerminalSection maxWidth="5xl" padding="py-12">
    <div class="flex flex-col md:flex-row justify-between items-start">
      <div class="mb-6 md:mb-0">
        <TerminalPrompt command="whoami" />
        <div class="text-white text-lg font-bold mt-2">anonmsg</div>
        <div class="text-gray-500"># anonymous inbox api</div>
        <div class="text-gray-500"># no auth, no bs</div>
      </div>
      <div class="space-y-2">
        <TerminalPrompt command="ls docs/" variant="gray" />
        <div class="space-y-1 ml-4">
          <TerminalButton href="#" variant="ghost" size="sm"
            >api-docs</TerminalButton
          >
          <TerminalButton
            newTab
            href="https://github.com/b3nten/anonmsg"
            variant="ghost"
            size="sm">github_repo</TerminalButton
          >
        </div>
      </div>
    </div>
    <div class="mt-8 pt-8 border-t border-gray-800 text-center">
      <div class="text-gray-500 text-sm">
        Built for developers who hate unnecessary complexity
      </div>
    </div>
  </TerminalSection>
</footer>
