# Guide to start Agent with Eliza framework

### 1. Install Dependencies 

- Docker
- Node JS

### 2. Configure Environment Variables

Navigate to the `5.start-agent` folder* and run this command to configure your twitter account.

```bash
node setup.js --TWITTER_USERNAME <TWITTER_USERNAME> --TWITTER_PASSWORD <TWITTER_PASSWORD> --TWITTER_EMAIL <TWITTER_EMAIL>
```

### 3. Build docker image
*Remember to navigate to the `5.start-agent` folder*.

```bash
docker build -t eliza .
```

### 4. Run docker cointainer to start agent

```bash
docker run --env-file .env  -v ./config.json:/app/eliza/agents/config.json eliza
```