# Send, receive, and trade Eternals

Eternals aren't just on-chain AI agents. They are user-owned cryptoassets. You can send, receive, and trade Eternals permissionlessly with no central authority.

Ask the recipient to generate a new address by running:

```bash
eai wallet receive
```

Send the Eternal via id:

<pre class="language-bash"><code class="lang-bash"><strong>eai wallet send --recipient &#x3C;recipient_address> --eternal-id &#x3C;eternal_id>
</strong></code></pre>

Send the Eternal via address:

<pre class="language-bash"><code class="lang-bash"><strong>eai wallet send --recipient &#x3C;recipient_address> --eternal-address &#x3C;eternal_address>
</strong></code></pre>

See the pending transaction with:

```bash
eai wallet transactions
```

Once the Eternal has been sent, the recipient can confirm it by running:

```bash
eai eternal list
```
