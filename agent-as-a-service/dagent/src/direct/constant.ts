import { Store } from "@eternalai-dagent/core";

export const AppAgent: Store = {
    info: {
        owner_address: "", // Default get from the environment
        authen_url: "https://composed-rarely-feline.ngrok-free.app/api/install", // Config the authen_url for the store here
        name: "Twitter",
        description: "The price fluctuations of coins in the past 24 hours",
        type: "store",
        icon: "https://s2.coinmarketcap.com/static/img/coins/64x64/1.png"
    },
    missions: [
        {
            "name": "Get coin price fluctuations 24h",
            "prompt": "API to get the price fluctuations of coins in the past 24 hours",
            "price": 1,
            "tool_list": [
                {
                    "headers": {},
                    "method": "GET",
                    "label": "query",
                    "executor": "https://agent.api.eternalai.org/api/bubble/list",
                    "name": "get_coin_price_fluctuations_24h",
                    "description": "API to get the price fluctuations of coins in the past 24 hours.",
                    "params": []
                },
                {
                    "headers": {
                        "Content-Type": "application/json",
                        "api-key": "<api-key>",
                    },
                    "label": "action",
                    "method": "POST",
                    "executor": "https://composed-rarely-feline.ngrok-free.app/api/internal/twitter/user/tweet-by-token",
                    "name": "post",
                    "description": "Post something to twitter",
                    "params": [
                        {
                            "name": "text",
                            "dtype": "string"
                        }
                    ]
                }
            ]
        }
    ]
}