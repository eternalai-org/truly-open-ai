import { Direct, getRouter } from "@eternalai-dagent/direct";

export function twitterRouters() {
    const router = getRouter();

    router.get("/", (req, res) => {
        res.send("Welcome, this is the REST API!");
    });

    return router;
}

const service = new Direct({
    routers: [
        twitterRouters()
    ]
})

service.start(80);