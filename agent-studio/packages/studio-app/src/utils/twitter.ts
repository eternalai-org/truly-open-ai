import last from "lodash/last";

enum ETwitterImageProfileSize {
  normal = "normal",
  medium = "200x200",
  high = "400x400",
}

export const getUrlAvatarTwitter = (
  url: string,
  size: "normal" | "medium" | "high" = "normal"
) => {
  if (url) {
    if (!url.includes("pbs.twimg.com") && !url.includes("abs.twimg.com")) {
      return url;
    }

    if (url?.includes("default_profile_normal.png")) {
      return undefined;
    }

    const urls = url?.split("/");

    let finalUrl = urls.splice(0, urls.length - 1).join("/");

    const lastPartUrl = last(urls)?.split("_");

    if (lastPartUrl?.[0] === "default") {
      return url;
    }

    finalUrl += `/${lastPartUrl
      ?.splice(0, lastPartUrl.length - 1)
      ?.join("_")}_${ETwitterImageProfileSize[size]}.${last(
      last(lastPartUrl)?.split(".")
    )}`;

    return finalUrl;
  }
  return undefined;
};
