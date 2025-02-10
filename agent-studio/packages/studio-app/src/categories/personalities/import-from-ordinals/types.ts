import { Collection, INFTInfo } from "../../../types/collection";
import { IMessageVerifier } from "../../../types/message";

export type ImportFromOrdinalsFormData = {
  stepper: number;

  selectedOption: Collection | undefined;

  selectedNFT: INFTInfo | undefined;

  tokenIdErrorMessage: string | undefined;

  tokenId: string;

  personality: string;

  signData: IMessageVerifier | undefined;
};

export type ImportFromOrdinalsState = {
  collections: Collection[];
};
