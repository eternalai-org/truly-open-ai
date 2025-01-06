import { NeynarAPIClient } from "@neynar/nodejs-sdk";
import { FeedTrendingProvider, FetchRepliesAndRecastsForUserFilterEnum, FetchTrendingFeedTimeWindowEnum, ForYouProvider } from "@neynar/nodejs-sdk/build/api";
import * as core from "express-serve-static-core";

export const neynarAPIRouterInit = (app: core.Express) => {
    const neynarAPI = new NeynarAPIClient({
        apiKey: process.env.NEYNAR_API_KEY
    });
    app.get('/neynar/lookupUserByUsername', async (req, res) => {
        try {
            const user = await neynarAPI.lookupUserByUsername(
                req.query as unknown as {
                    username: string;
                    viewerFid?: number;
                }
            );
            res
                .status(200)
                .send({
                    result: user,
                })
        } catch (error) {
            res
                .status(400)
                .send({
                    result: null,
                    error: {
                        message: error['message']
                    }
                })
        }
    });
    app.get('/neynar/fetchPopularCastsByUser', async (req, res) => {
        try {
            const casts = await neynarAPI.fetchPopularCastsByUser(
                req.query as unknown as {
                    fid: number;
                    viewerFid?: number;
                }
            )
            res
                .status(200)
                .send({
                    result: casts,
                })
        } catch (error) {
            res
                .status(400)
                .send({
                    result: null,
                    error: {
                        message: error['message']
                    }
                })
        }
    });
    app.get('/neynar/fetchRepliesAndRecastsForUser', async (req, res) => {
        try {
            const casts = await neynarAPI.fetchRepliesAndRecastsForUser(
                req.query as unknown as {
                    fid: number;
                    filter?: FetchRepliesAndRecastsForUserFilterEnum;
                    limit?: number;
                    cursor?: string;
                    viewerFid?: number;
                }
            )
            res
                .status(200)
                .send({
                    result: casts,
                })
        } catch (error) {
            res
                .status(400)
                .send({
                    result: null,
                    error: {
                        message: error['message']
                    }
                })
        }
    });
    app.get('/neynar/fetchRepliesAndRecastsForUser', async (req, res) => {
        try {
            const casts = await neynarAPI.fetchRepliesAndRecastsForUser(
                req.query as unknown as {
                    fid: number;
                    filter?: FetchRepliesAndRecastsForUserFilterEnum;
                    limit?: number;
                    cursor?: string;
                    viewerFid?: number;
                }
            )
            res
                .status(200)
                .send({
                    result: casts,
                })
        } catch (error) {
            res
                .status(400)
                .send({
                    result: null,
                    error: {
                        message: error['message']
                    }
                })
        }
    });
    app.get('/neynar/searchCasts', async (req, res) => {
        try {
            const casts = await neynarAPI.searchCasts(
                req.query as unknown as {
                    q: string;
                    authorFid?: number;
                    viewerFid?: number;
                    parentUrl?: string;
                    channelId?: string;
                    priorityMode?: boolean;
                    limit?: number;
                    cursor?: string;
                }
            )
            res
                .status(200)
                .send({
                    result: casts,
                })
        } catch (error) {
            res
                .status(400)
                .send({
                    result: null,
                    error: {
                        message: error['message']
                    }
                })
        }
    });
    app.get('/neynar/fetchUserFollowingFeed', async (req, res) => {
        try {
            const feeds = await neynarAPI.fetchUserFollowingFeed(
                req.query as unknown as {
                    fid: number;
                    viewerFid?: number;
                    withRecasts?: boolean;
                    limit?: number;
                    cursor?: string;
                }
            )
            res
                .status(200)
                .send({
                    result: feeds,
                })
        } catch (error) {
            res
                .status(400)
                .send({
                    result: null,
                    error: {
                        message: error['message']
                    }
                })
        }
    });
    app.get('/neynar/fetchFeedForYou', async (req, res) => {
        try {
            const feeds = await neynarAPI.fetchFeedForYou(
                req.query as unknown as {
                    fid: number;
                    viewerFid?: number;
                    provider?: ForYouProvider;
                    limit?: number;
                    cursor?: string;
                    providerMetadata?: string;
                }
            )
            res
                .status(200)
                .send({
                    result: feeds,
                })
        } catch (error) {
            res
                .status(400)
                .send({
                    result: null,
                    error: {
                        message: error['message']
                    }
                })
        }
    });
    app.get('/neynar/fetchTrendingFeed', async (req, res) => {
        try {
            const feeds = await neynarAPI.fetchTrendingFeed(
                req.query as unknown as {
                    limit?: number;
                    cursor?: string;
                    viewerFid?: number;
                    timeWindow?: FetchTrendingFeedTimeWindowEnum;
                    channelId?: string;
                    provider?: FeedTrendingProvider;
                    providerMetadata?: string;
                }
            )
            res
                .status(200)
                .send({
                    result: feeds,
                })
        } catch (error) {
            res
                .status(400)
                .send({
                    result: null,
                    error: {
                        message: error['message']
                    }
                })
        }
    });
}