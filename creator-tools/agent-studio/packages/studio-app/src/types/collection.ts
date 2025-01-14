export interface Collection {
  collection: string;
  name: string;
  description: string;
  image_url: string;
  banner_image_url: string;
  owner: string;
  safelist_status: string;
  category: string;
  is_disabled: boolean;
  is_nsfw: boolean;
  trait_offers_enabled: boolean;
  collection_offers_enabled: boolean;
  opensea_url: string;
  project_url: string;
  wiki_url: string;
  discord_url: string;
  telegram_url: string;
  twitter_username: string;
  instagram_username: string;
  contracts: Contract[];
}

export interface Contract {
  address: string;
  chain: string;
}

export interface INFTInfo {
  token_address: string;
  token_id: string;
  amount: string;
  owner_of: string;
  token_hash: string;
  contract_type: string;
  name: string;
  symbol: string;
  token_uri: string;
  metadata: string;
  block_number_minted: string;
  normalized_metadata: NormalizedMetadata;
}

export interface NormalizedMetadata {
  image: string;
  name: string;
  description: string;
  external_link: string;
  animation_url: string;
  attributes: Attribute[];
}

export interface Attribute {
  trait_type: string;
  value: string;
  display_type: any;
  max_value: any;
  trait_count: number;
  order: any;
  rarity_label: string;
  count: number;
  percentage: number;
}
