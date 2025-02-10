import AgentAPI from "../services/apis/agent";
import { EIdeaOption } from "../types/agent";

const useCollections = () => {
  const fetchCollections = async (type: EIdeaOption) => {
    try {
      const rs = await AgentAPI.getNFTCollections({
        inscription: type === EIdeaOption.ordinals,
        limit: type === EIdeaOption.ordinals ? 20 : 100,
      });
      return rs?.collections || [];
    } catch (e) {
      return [];
    }
  };

  return { fetchCollections };
};

export default useCollections;
