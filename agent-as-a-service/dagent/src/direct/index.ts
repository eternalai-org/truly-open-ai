
interface IAuthenParams {
    client_id: string;
    redirect_uri: string;
}

export const getTwitterOauthUrl = ({ client_id, redirect_uri }: IAuthenParams) => {
    const rootUrl = "https://twitter.com/i/oauth2/authorize";
    const URL = `&client_id=${client_id}`;
    const options = {
        redirect_uri: redirect_uri,
        client_id: client_id,
        state: "state",
        response_type: "code",
        code_challenge: "challenge",
        code_challenge_method: "plain",
        scope: [
            "offline.access",
            "tweet.read",
            "tweet.write",
            "users.read",
            "tweet.moderate.write",
            "follows.write",
            "like.write",
            "list.write",
            "block.write",
            "bookmark.write",
            "block.read",
            "follows.read",
            "bookmark.read",
            "list.read",
            "space.read",
            "like.read",
            "users.read",
            "mute.read",
        ].join(" "),
    };
    const qs = new URLSearchParams(options).toString();
    return `${rootUrl}?${qs}`;
}