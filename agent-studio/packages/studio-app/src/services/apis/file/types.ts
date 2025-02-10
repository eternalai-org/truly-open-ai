import { AxiosProgressEvent, AxiosResponse } from "axios";

export type UploadFileProcess = {
  chunkIndex: number;
  totalChunk: number;
  chunkSize: number;
  progress: number;
  progressEvent: AxiosProgressEvent;
};

export interface IUploadFile {
  file: File;
}

export interface IInitiateMultipartUploadPayload {
  fileName: string;
  group?: string;
}

export interface IInitiateMultipartUploadResponse {
  uploadId: string;
}

export type RequestPromise = {
  request: Promise<AxiosResponse<any, any>>;
  controller: AbortController;
};

export interface ICompleteMultipartUploadResponse {
  fileUrl: string;
}
