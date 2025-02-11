import { Direct, getRouter } from "@eternalai-dagent/direct";
import { getTwitterOauthUrl, randomString } from "./utils";
import axios from "axios";
import dotenv from "dotenv";
import path from "path";
import {postedStorage, twitterStorage} from "../storage";

const getRedirectUrl = (prams: {
    install_code: string
    agent_id: string,
    store_url: string
}) => {
    return `${prams.store_url}/api/webhook/twitter-oauth?install_code=${prams?.install_code}&agent_id=${prams?.agent_id}`;
}

export function twitterRouters() {
    const router = getRouter();
    const TWITTER_APP_CLIENT_ID = process.env.TWITTER_APP_CLIENT_ID as string;
    const TWITTER_APP_CLIENT_SECRET = process.env.TWITTER_APP_CLIENT_SECRET as string;
    const TWITTER_APP_STORE_URL = process.env.TWITTER_APP_STORE_URL as string;

    console.log("TWITTER_APP_CLIENT_ID: ", TWITTER_APP_CLIENT_ID);
    console.log("TWITTER_APP_CLIENT_SECRET: ", TWITTER_APP_CLIENT_SECRET);
    console.log("TWITTER_APP_STORE_URL: ", TWITTER_APP_STORE_URL);

    router.get("/", (req, res) => {
        res.send("Welcome, this is the REST API!");
    });

    router.get("/api/install", (req, res) => {
        const install_code = (req.query?.install_code || "") as string;
        const agent_id = (req.query?.agent_id || "") as string;

        if (!install_code || !agent_id) {
            res.send("Error: install_code and agent_id are required");
            return;
        }

        const redirect_uri = getRedirectUrl({
            install_code,
            agent_id,
            store_url: TWITTER_APP_STORE_URL
        });
        const callbackUrl = getTwitterOauthUrl({
            client_id: TWITTER_APP_CLIENT_ID,
            redirect_uri: redirect_uri
        });

        res.redirect(callbackUrl);

    });

    router.get("/api/webhook/twitter-oauth", async (req, res) => {
        const code = req.query?.code as string;
        const install_code = req.query?.install_code as string;
        const agent_id = req.query?.agent_id as string;
        const redirect_uri = getRedirectUrl({
            install_code,
            agent_id,
            store_url: TWITTER_APP_STORE_URL
        });

        const username = TWITTER_APP_CLIENT_ID;
        const password = TWITTER_APP_CLIENT_SECRET;
        const credentials = `${username}:${password}`;
        const encodedCredentials = Buffer.from(credentials).toString('base64');

        try {
            const response = await axios.post("https://api.twitter.com/2/oauth2/token", {
                client_id: username,
                code_verifier: "challenge",
                redirect_uri: redirect_uri,
                grant_type: "authorization_code",
                code: code
            }, {
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Basic ${encodedCredentials}`,
                }
            });
            const accessToken = response.data.access_token;
            const refreshToken = response.data.refresh_token;
            const api_key = randomString(32);

            twitterStorage.setItem(api_key, JSON.stringify({
                accessToken,
                refreshToken,
                install_code,
                api_key,
                agent_id
            }))

            const return_data = Buffer.from(JSON.stringify({
                api_key
            })).toString('base64');

            res.redirect(`https://eternalai.org/agents/edit-mission/${agent_id}?install_code=${install_code}&return_data=${return_data}`)

        } catch (e: any) {
            res.send("Error: " + e?.message ? e.message : e);
        }
    });

    router.get("/api/internal/twitter/user/tweet-by-token", async (req, res) => {
        // get api_key from the request in header
        const api_key = req.headers["api-key"] as string;
        const data = twitterStorage.getItem(api_key);
        if (!data) {
            res.status(401).send("Unauthorized");
            return;
        }
        const text = req.body?.text as string;
        res.send("Tweeted: " + text);

        const { agent_id } = JSON.parse(data);

        let posted: any = postedStorage.getItem(agent_id);
        if (!posted) {
            posted = []
        } else {
            posted = JSON.parse(posted) || [];
        }
        postedStorage.setItem(agent_id, JSON.stringify([...(posted as any), text]));

        // handle the tweet here
    });

    return router;
}

const __dirname = path.dirname(new URL(import.meta.url).pathname);
dotenv.config({ path: path.resolve(__dirname, "../../../.env") });

const service = new Direct({
    routers: [
        twitterRouters()
    ]
});

export default service;

