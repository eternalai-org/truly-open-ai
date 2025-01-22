# Guide to start Agent with Eliza framework

### 1. Install Dependencies 

- Docker

### 2. Configure Environment Variables

Config enviroment variables in `.env` file and input your agent's system prompt in `config.json` in `5.start-agent` folder.


### 3. Build docker image
*Remember to navigate to the `5.start-agent` folder*.

```bash
docker build -t eliza .
```

### 4. Run docker cointainer to start agent

```bash
docker run --env-file .env  -v ./config.json:/app/eliza/agents/config.json eliza
```