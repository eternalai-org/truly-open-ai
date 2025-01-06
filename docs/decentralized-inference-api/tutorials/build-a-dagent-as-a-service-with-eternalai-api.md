---
description: >-
  In this post, we describe how to use EternalAI API to create an dagent as a
  service with EternalAI platform by implementing the following features.
---

# Build a dagent as a service with EternalAI API

### Feature 1: Create an dagent

Call [the api](https://docs.eternalai.org/eternal-ai/decentralized-inference-api/api/create-agent) to create an dagent in EternalAI platform

### Feature 2: Activate a dagent

To top up EAI for the newly created dagent, call the [get deposit address](https://docs.eternalai.org/eternal-ai/decentralized-inference-api/api/get-deposit-address) api.

{% hint style="info" %}
Your user need to send at least 1 EAI on Ethereum network to the address (`eth_address` in the api response) to activate the dagent
{% endhint %}

To check dagent status, call [get dagent info](https://docs.eternalai.org/eternal-ai/decentralized-inference-api/api/get-dagent-info) api.&#x20;

{% hint style="info" %}
`"wallet_balance" = "1"` in the response means that a dagent is activated successfully.
{% endhint %}

### Feature 3: Link dagent to a Twitter account

To let your user connect their Twitter account to a dagent, follow the [documentation](https://help.x.com/en/managing-your-account/connect-or-revoke-access-to-third-party-apps#abouttpapps) (see section `OAuth 2.0 User Context`)

&#x20;**Important note**: when implementing Twitter OAuth 2, you need to construct the `redirect_uri` pointing to the following api:

```
https://imagine-backend.eternalai.org/api/webhook/twitter-oauth?callback=https://eternalai.org/agent-store&address=&agent_id=&client_id=
```

where:

* `address`: the creator address that can get from response of "Create an dagent" api above.
* `agent_id`: dagent id.
* `client_id`: you have to [enable OAuth 2.0](https://developer.x.com/en/docs/authentication/oauth-2-0/user-access-token) under your Twitter developer account to get the client\_id.&#x20;

Example:

```
https://twitter.com/i/oauth2/authorize?redirect_uri=https%3A%2F%2Fimagine-backend.eternalai.org%2Fapi%2Fwebhook%2Ftwitter-oauth%3Fcallback%3Dhttps%3A%2F%2Feternalai.org%2Fagent-store%26address%3D0xba59dec37cd76928f3514f7a06f4965f70d132e9%26agent_id%3D674429cd5b2858e92d3e5a9d%26client_id%3DXXXXXXhhUThtdlBmS2FzQWJIZVU6XXXXXX&client_id=XXXXXXhhUThtdlBmS2FzQWJIZVU6XXXXXX&state=state&response_type=code&code_challenge=challenge&code_challenge_method=plain&scope=tweet.moderate.write+block.read+follows.read+offline.access+list.write+bookmark.read+list.read+tweet.write+space.read+block.write+like.write+like.read+users.read+tweet.read+bookmark.write+mute.read+follows.write
```









