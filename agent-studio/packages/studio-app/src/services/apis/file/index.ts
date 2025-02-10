import { sleep } from "../../../utils/time";
import { agentAPIClient } from "../clients";
import {
  ICompleteMultipartUploadResponse,
  IInitiateMultipartUploadPayload,
  IInitiateMultipartUploadResponse,
  IUploadFile,
  RequestPromise,
  UploadFileProcess,
} from "./types";

import { AxiosProgressEvent } from "axios";

const API_PATH = `/api/files/multipart`;

const FileAPI = {
  initiateMultipartUpload: async (payload: IInitiateMultipartUploadPayload) => {
    const res: IInitiateMultipartUploadResponse = await agentAPIClient.post(
      `${API_PATH}`,
      payload
    );
    return res;
  },
  putUploadHandler: (id: string, requestPromises: RequestPromise[]) => {
    if (!(window as any).uploadMultipart) {
      (window as any).uploadMultipart = {};
    }
    (window as any).uploadMultipart = {
      ...(window as any).uploadMultipart,
      [id]: requestPromises,
    };
  },
  completeMultipartUpload: async (payload: { uploadId: string }) => {
    const res: ICompleteMultipartUploadResponse = await agentAPIClient.post(
      `${API_PATH}/${payload.uploadId}`,
      {}
    );
    return res;
  },
  deleteUploadHandler: (id: string) => {
    if (!(window as any).uploadMultipart) {
      (window as any).uploadMultipart = {};
    }
    delete (window as any).uploadMultipart[id];
  },
  uploadFile: async (
    payload: IUploadFile,
    chunkSize = 800_000,
    onCallback?: (id: string) => void,
    onProcess?: (process: UploadFileProcess) => void
  ): Promise<string> => {
    const { uploadId } = await FileAPI.initiateMultipartUpload({
      fileName: payload.file.name,
      group: "eternalai",
    });
    onCallback && onCallback(uploadId);
    const totalChunks = Math.ceil(payload.file.size / chunkSize);

    const fileName = payload.file.name;
    const requestPromises = Array.from(Array(totalChunks)).map((_, index) => {
      const start = index * chunkSize;
      const end = Math.min(start + chunkSize, payload.file.size);

      const chunk = payload.file.slice(start, end);
      const formData = new FormData();
      formData.append("file", chunk, fileName);

      const partNumber = index + 1;

      const controller = new AbortController();

      const request = agentAPIClient.put(
        `${API_PATH}/${uploadId}?partNumber=${partNumber}`,
        formData,
        {
          signal: controller.signal,
          headers: {
            "Content-Type": "application/octet-stream",
          },
          onUploadProgress: (progressEvent: AxiosProgressEvent) => {
            onProcess &&
              onProcess({
                chunkIndex: index,
                totalChunk: totalChunks,
                chunkSize: chunkSize,
                progress: progressEvent.progress || 0,
                progressEvent,
              });
          },
        }
      );
      return {
        request,
        controller,
      };
    });

    try {
      FileAPI.putUploadHandler(uploadId, requestPromises);
      const requests = requestPromises.map((item) => item.request);
      await Promise.all(requests);
      await sleep(1);

      const completeUploadRes = await FileAPI.completeMultipartUpload({
        uploadId: uploadId,
      });

      FileAPI.deleteUploadHandler(uploadId);
      return completeUploadRes.fileUrl;
    } catch (error: unknown) {
      throw Error("Upload chunk file error");
    }
  },
};

export default FileAPI;
