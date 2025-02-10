import { eternalAPIClient } from "../clients";

const ModelAPI = {
  getModelDescription: async () => {
    const res = await eternalAPIClient.get(`/api/v1/models/description`);
    return res;
  },
};

export default ModelAPI;
