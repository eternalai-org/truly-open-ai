import { Direct, getRouter } from "@eternalai-dagent/direct";
import { getTwitterOauthUrl } from "./direct";

export function twitterRouters() {
    const router = getRouter();

    router.get("/", (req, res) => {
        res.send("Welcome, this is the REST API!");
    });

    router.get("/api/install", (req, res) => {
        const install_code = req.query?.install_code;
        const client_id = "";

        const callbackUrl = getTwitterOauthUrl({
            client_id,
            redirect_uri: "https://composed-rarely-feline.ngrok-free.app/api/webhook/twitter-oauth?install_code=" + install_code
        });

        res.redirect(callbackUrl);

    });

    router.get("/api/webhook/twitter-oauth", (req, res) => {
        const code = req.query?.code;
        const install_code = req.query?.install_code;
        res.send(`Install code: ${install_code}, Twitter OAuth code: ${code}`);
    });

    return router;
}

const service = new Direct({
    routers: [
        twitterRouters()
    ]
})

service.start(80);