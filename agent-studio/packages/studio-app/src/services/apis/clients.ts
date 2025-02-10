import axios from "axios";
import useAgentServiceStore from "../../stores/useAgentServiceStore";
import { AGENT_AI_URL, ETERNAL_AI_URL } from "../../configs";

export const TIMEOUT = 5 * 60000;
export const HEADERS = { "Content-Type": "application/json" };

const createAxiosInstance = ({ baseURL = "" }: { baseURL: string }) => {
  const instance = axios.create({
    baseURL,
    timeout: TIMEOUT,
    headers: {
      ...HEADERS,
    },
  });

  instance.interceptors.request.use(
    (config) => {
      const authToken = useAgentServiceStore.getState().accessToken;
      config.headers.Authorization = `${authToken}`;
      return config;
    },
    (error) => {
      Promise.reject(error);
    }
  );

  instance.interceptors.response.use(
    (res) => {
      const result = res?.data?.data || res?.data?.result;
      if (res?.data?.count !== undefined) {
        result.count = res.data.count;
      }
      const error = res?.data?.error;
      if (error && Object.keys(error).length) {
        return Promise.reject(error);
      }

      if (!result) {
        return Promise.resolve(result);
      }
      if (typeof result === "object") {
        // return Promise.resolve(camelCaseKeys(result));
        return result;
      }
      return Promise.resolve(result);
    },
    (error: any) => {
      if (!error.response) {
        return Promise.reject(error);
      }
      const response = error?.response?.data || error;
      const errorMessage =
        response?.error || error?.Message || JSON.stringify(error);
      return Promise.reject(errorMessage);
    }
  );

  return instance;
};

export const agentAPIClient = createAxiosInstance({
  baseURL: AGENT_AI_URL,
});

export const eternalAPIClient = createAxiosInstance({
  baseURL: ETERNAL_AI_URL,
});
