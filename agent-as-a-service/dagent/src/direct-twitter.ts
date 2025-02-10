import { Direct, getRouter } from "@eternalai-dagent/direct";
import {getTwitterOauthUrl, randomString} from "./direct";
import axios from "axios";
import dotenv from "dotenv";
import path from "path";
import {twitterStorage} from "./storage";

const getRedirectUrl = (prams: {
    install_code: string
    app_id: string,
    agent_id: string,
}) => {
    return `${process.env.TWITTER_APP_STORE_URL}/api/webhook/twitter-oauth?install_code=${prams?.install_code}`;
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
        const install_code = req.query?.install_code;

        console.log(`${TWITTER_APP_STORE_URL}/api/webhook/twitter-oauth?install_code=${install_code}`);

        const redirect_uri = `${TWITTER_APP_STORE_URL}/api/webhook/twitter-oauth?install_code=${install_code}`;
        const callbackUrl = getTwitterOauthUrl({
            client_id: TWITTER_APP_CLIENT_ID,
            redirect_uri: redirect_uri
        });

        res.redirect(callbackUrl);

    });

    router.get("/api/webhook/twitter-oauth", async (req, res) => {
        const code = req.query?.code as string;
        const install_code = req.query?.install_code;
        const redirect_uri = `${TWITTER_APP_STORE_URL}/api/webhook/twitter-oauth?install_code=${install_code}`

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
                api_key
            }))

            const return_data = Buffer.from(JSON.stringify({
                api_key
            })).toString('base64');

            res.redirect(`https://eternalai.org/agents/edit-mission?install_code=${install_code}&return_data=${return_data}`)

        } catch (e: any) {
            res.send("Error: " + e?.message ? e.message : e);
        }
    });

    return router;
}

const __dirname = path.dirname(new URL(import.meta.url).pathname);
dotenv.config({ path: path.resolve(__dirname, "../../.env") });

const service = new Direct({
    routers: [
        twitterRouters()
    ]
})

service.start(80);