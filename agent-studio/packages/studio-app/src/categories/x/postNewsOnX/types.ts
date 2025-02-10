export type PostNewsOnXFormData = {
  id: string;
  frequency: string;
  details: string;
  topics: PostNewsTopics;
  model: string;
  modelName: string;
};

export type PostNewsTopics = {
  values: string;
  bingSearch: boolean;
  twitterSearch: boolean;
};
