import service from "./service";
import { Direct } from "@eternalai-dagent/direct";
import { Wallet } from "ethers";
import path from "path";
import dotenv from "dotenv";
import { dagentLogger, StoreAgent } from "@eternalai-dagent/core";
import { AppAgent } from "./constant";

const __dirname = path.dirname(new URL(import.meta.url).pathname);
dotenv.config({ path: path.resolve(__dirname, "../../../.env") });

class ServiceManager {
    protected service: Direct;
    protected signer: Wallet;
    protected storeApi: StoreAgent

    constructor() {
        this.service = service;
        this.signer = new Wallet(process.env.STORE_OWNER_PRIVATE_KEY || "");
        this.storeApi = new StoreAgent({
            endpoint: process.env.ETERNAL_AI_URL || "",
        });

        dagentLogger.info("Signer Address", this.signer.address);
    }

    async init() {
        const signerAddress = this.signer.address;
        const message = `Sign this message to get access token: ${signerAddress}`;
        const signature = await this.signer.signMessage(message);

        const accessToken = await this.storeApi.getAccessToken({
            address: signerAddress,
            signature: signature,
            message: message,
        });

        this.storeApi.setAuthToken(accessToken);
    }

    start() {
        this.service.start(80);
    }

    stop() {
        this.service.stop();
    }

    async createApp() {
        await this.storeApi.createApp({
            info: {
                ...AppAgent.info,
                owner_address: this.signer.address,
            },
            missions: AppAgent.missions,
        })
    }
}

const serviceManager = new ServiceManager();

// Init environment, authen, and get access token interact with eternal ai api
await serviceManager.init();

// Create your app if you don't have one
await serviceManager.createApp();

// Start the service if the app has been created
serviceManager.start();