
import axios, { Axios } from "axios";
import {IAuthenticParams, IGetAccessTokenParams} from "../types";

export const TIMEOUT = 5 * 60000;
export const HEADERS = { "Content-Type": "application/json" };

const createAxiosInstance = ({
  baseURL,
  authToken,
}: {
  baseURL: string;
  authToken: string;
}) => {
  const instance = axios.create({
    baseURL: `${baseURL}/api`,
    timeout: TIMEOUT,
    headers: {
      ...HEADERS,
    },
  });

  instance.interceptors.request.use(
    (config: { headers: any }) => {
      if (authToken) {
        config.headers.Authorization = `${authToken}`;
      }
      return config;
    },
    (error: any) => {
      Promise.reject(error);
    }
  );

  instance.interceptors.response.use(
    (res: { data: { data: any; result: any; count: any; error: any } }) => {
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


/** Get access token */
// getAccessToken: (params: IGetAccessTokenParams) => Promise<string>;

export interface IBaseAPI {
    setAuthToken: (authToken: string) => void;
    getAccessToken: (params: IGetAccessTokenParams) => Promise<string>;
}

export default class BaseAPI implements IBaseAPI {
    protected api: Axios;

    constructor(params: IAuthenticParams) {
        this.api = createAxiosInstance({
            baseURL: params.endpoint,
            authToken: params?.accessToken || "",
        });
    }

    setAuthToken = (authToken: string) => {
        (this.api.defaults.headers as any).Authorization = authToken;
    };

    getAccessToken = async (params: IGetAccessTokenParams): Promise<string> => {
        try {
            const signature = params.signature.startsWith("0x")
                ? params.signature.replace("0x", "")
                : params.signature;
            const authenCode: string = await this.api.post('auth/verify', { ...params, signature });
            return authenCode;
        } catch (e) {
            throw e;
        }
    };
}
