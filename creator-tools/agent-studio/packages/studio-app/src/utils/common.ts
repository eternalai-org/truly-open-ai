export const getImageIPFS = (hash: string): string => {
  if (!hash) {
    return "";
  }
  const _hash = hash.replace("ipfs://", "");

  return `https://gateway.lighthouse.storage/ipfs/${_hash}`;
};

export const getImageIPFSCreateAgent = (hash: string): string => {
  if (!hash) {
    return "";
  }

  if (hash.startsWith("ipfs://")) {
    const _hash = hash.replace("ipfs://", "");

    return `https://gateway.pinata.cloud/ipfs/${_hash}`;
  } else {
    return hash;
  }
};
