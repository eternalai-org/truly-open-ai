export enum EIdeaOption {
  generic = "generic",
  nft = "nft",
  link_token = "link_token",
  ordinals = "ordinals",
  twitter = "twitter",
  clone = "clone",
}

export interface Entities {
  annotations: any;
  cashtags: any;
  hashtags: any;
  mentions: any;
  urls: any;
}

export interface PublicMetrics {
  followers_count: number;
  following_count: number;
  listed_count: number;
  tweet_count: number;
}

export interface Withheld {
  copyright: boolean;
  country_codes: any;
}

export interface ISearchTwitterInfo {
  created_at: string;
  description: string;
  entities: Entities;
  id: string;
  location: string;
  name: string;
  pinned_tweet_id: string;
  profile_image_url: string;
  protected: boolean;
  public_metrics: PublicMetrics;
  url: string;
  username: string;
  verified: boolean;
  withheld: Withheld;
}
