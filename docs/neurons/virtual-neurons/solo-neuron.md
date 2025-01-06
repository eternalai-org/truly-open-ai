# Solo Neuron

Solo mining is the best option for securing the network. It involves operating a Neuron from your home. It's a direct way to participate in maintaining the network and earning rewards.&#x20;

Solo mining offers great control and direct rewards but requires technical know-how and a dedicated setup. Before diving in, make sure you're ready for the responsibility.

### Requirements <a href="#requirements" id="requirements"></a>

Before installing the Virtual Neuron software on your computer, you will need to meet the following requirements:

* **Operating system**: Virtual Neuron software is optimized for Linux, but Windows will be supported soon. Ensure that your operating system is compatible with the version of Eternal Node you intend to install.
* **GPU**: NVIDIA RTX 3000 series or equivalent, with at least 24 GB of VRAM. This is important because AI models tend to have larger and more complex neural networks that require more memory to operate efficiently. The NVIDIA RTX A6000 or RTX 3090 would be suitable.
* **CPU architecture**: Virtual Neuron software is compatible with 64-bit CPU architectures. Ensure that your computer has a 64-bit CPU. A modern multi-core processor (e.g., Intel i7, i9 or AMD Ryzen 7, Ryzen 9) will suffice. The CPU is less critical than the GPU but still important for overall system stability and parallel processing tasks.
* **Memory**: Virtual Neuron software requires at least 32GB of RAM to run efficiently.
* **Storage**: SSD with at least 100 GB capacity. An SSD is recommended over an HDD for faster data access and processing speeds, which is beneficial when loading large models.
* **Internet connection**: Virtual Neuron software requires a stable and high-speed internet connection to synchronize with the Eternal AI network.
* **Firewall and port forwarding**: Ensure that your firewall settings and port forwarding rules allow incoming and outgoing traffic on the ports used by Eternal AI.
* **Command-line interface**: Virtual Neuron software is primarily operated through the command line interface. Ensure that you have a basic understanding of command-line usage and are comfortable working in a terminal environment.
* **NVIDIA Driver**: Updated NVIDIA drivers to ensure compatibility and optimal performance with the latest GPU models. [Learn more.](https://ubuntu.com/server/docs/nvidia-drivers-installation)​
* **Docker**: Installation of Docker to facilitate the deployment and management of AI models within containerized environments, essential for model consistency and portability. [Learn more.](https://docs.docker.com/engine/install/ubuntu/)​

### Installation <a href="#installation" id="installation"></a>

Let's start by installing the `eternal` command line tool.

```bash
curl -L https://eternalai.org/get.sh | sh
```

### Create an Eternal AI account <a href="#create-an-eternal-ai-account" id="create-an-eternal-ai-account"></a>

[Create a self-custody wallet](../../fully-onchain-ai-models/deploy-your-first-fully-onchain-ai/create-a-self-custody-wallet.md) and send at least 25,010 EAI to the wallet. The minimum requirement to run a compute node is 25,000 EAI.

### Run a Compute Neuron <a href="#run-a-miner" id="run-a-miner"></a>

Check for updates

```bash
curl -L https://eternalai.org/get.sh | sh
```

Start a miner client

```bash
./eternal -account <private-key>
```

`-account` is optional. You can create an Eternal AI account and provide the private key or the app will generate it for you.

### Run a Verify Neuron (coming soon) <a href="#run-a-validator-coming-soon" id="run-a-validator-coming-soon"></a>

Check for updates

```bash
curl -L https://eternalai.org/get.sh | sh
```

Start a validator client

```sh
./eternal -validator -account <private-key>
```
