import os
import time
import json
import requests
import numpy as np
from typing import List
from web3 import Web3, HTTPProvider
from eai.utils import merge_float32_to_uint256, parse_uint256_to_float32, Logger
from eai.network_config import NETWORK, INFER_TIMEOUT
from eai.artifacts.models.FunctionalModel import CONTRACT_ARTIFACT


class TensorData:
    def __init__(self, data: bytes, shapes: List[int]):
        self.data = data
        self.shapes = shapes

    def toContractParams(self):
        return (self.data, self.shapes)

    @staticmethod
    def from_numpy(arr):
        original_shape = arr.shape
        arr_32 = arr.reshape(-1).tolist()
        merged_arr = []
        to_append = 0
        N = len(arr_32)
        if N % 4 != 0:
            to_append = 4 - N % 4
        for i in range(to_append):
            arr_32.append(0)
        for i in range(0, N, 4):
            merged_arr.append(merge_float32_to_uint256(arr_32[i:i + 4]))
        return TensorData(merged_arr, original_shape)

    def to_numpy(self):
        arr = []
        for i in range(len(self.data)):
            float32_arr = parse_uint256_to_float32(self.data[i])
            arr.extend(float32_arr)
        return np.array(arr[:np.prod(self.shapes)]).reshape(self.shapes)


class Eternal:
    def __init__(self, id_or_address: str = None):
        self.address = None
        self.model_id = None
        self.price = None
        self.name = None
        self.owner = None
        self.status = None
        if id_or_address is not None:
            try:
                self.load(id_or_address)
            except Exception as e:
                Logger.error(f"Failed to load model from {id_or_address}: {e}")

    def __str__(self):
        return f"id: {self.model_id}, address: {self.address}, name: {self.name}, price: {self.price}, owner: {self.owner}, status: {self.status}."

    def set_price(self, price: float):
        self.price = price

    def get_price(self):
        return self.price

    def set_name(self, name: str):
        self.name = name

    def get_name(self):
        return self.name

    def get_publisher(self):
        return self.owner

    def load(self, model: str):
        if model.startswith("0x"):
            checksum_address = Web3.to_checksum_address(model)
            self._load_metadata_from_address(checksum_address)
        else:
            self._load_metadata_from_id(model)

    def _load_metadata_from_id(self, id):
        self.id = id
        Logger.info(f"Loading metadata from id {self.id}")
        endpoint = NETWORK[os.environ["NETWORK_MODE"]]["MODEL_INFO_BY_ID"]
        try:
            url = f"{endpoint}/{self.id}"
            response = requests.get(url)
            response = response.json()
            if response["status"] == 1:
                self.status = response["data"]["status"]
                self.address = response["data"]["model_address"]
                self.name = response["data"]["model_name"]
                self.price = 0
                self.owner = response["data"]["owner"]["address"]
                Logger.success("Metadata loaded successfully.")
            else:
                Logger.error("Failed to load metadata. Response status not 1.")
        except Exception as e:
            Logger.error(f"An unexpected error occurred: {e}")
            return False
        return True

    def _load_metadata_from_address(self, address: str):
        self.address = address
        Logger.info(f"Loading metadata from address {self.address}")
        endpoint = NETWORK[os.environ["NETWORK_MODE"]]["MODEL_INFO_BY_ADDRESS"]
        try:
            url = f"{endpoint}/{self.address}"
            response = requests.get(url)
            response = response.json()
            if response["status"] == 1:
                self.status = response["data"]["status"]
                self.id = response["data"]["model_id"]
                self.name = response["data"]["model_name"]
                self.price = 0
                self.owner = response["data"]["owner"]["address"]
                Logger.success("Metadata loaded successfully.")
            else:
                Logger.error("Failed to load metadata. Response status not 1.")
        except Exception as e:
            Logger.error(f"An unexpected error occurred: {e}")

    def get_address(self):
        return self.address

    def get_id(self):
        return self.id

    def to_json(self, output_path=None):
        metadata = {
            "address": self.address,
            "id": self.id,
            "status": self.status,
            "price": self.price,
            "name": self.name,
            "owner": self.owner
        }
        if output_path is not None:
            with open(output_path, "w") as f:
                json.dump(metadata, f)
            Logger.success(
                f"Transformed model metadata saved to {output_path}.")
        return metadata

    def predict(self,
                inputs: List[np.ndarray],
                output_path: str = None) -> np.ndarray:
        address = self.address
        Logger.info(
            "Making prediction on EternalAI's {} at {} ...".format(
                os.environ["NETWORK_MODE"], address))
        w3 = Web3(HTTPProvider(NETWORK[os.environ["NETWORK_MODE"]]
                  ["NODE_ENDPOINT"], request_kwargs={'timeout': INFER_TIMEOUT}))
        model_contract = w3.eth.contract(
            address=self.address, abi=CONTRACT_ARTIFACT['abi'])
        input_tensors = list(map(lambda x: TensorData.from_numpy(x), inputs))
        input_params = list(map(lambda x: x.toContractParams(), input_tensors))
        start = time.time()
        result = model_contract.functions.predict(input_params).call()
        output_tensor = TensorData(result[0], result[1])
        output_numpy = output_tensor.to_numpy()
        Logger.success(
            "Prediction made successfully in {} seconds. Output: {}".format(
                time.time() - start,
                output_numpy.tolist()))
        if output_path is not None:
            np.save(output_path, output_numpy)
            Logger.success(
                f"Prediction saved to {output_path}.")
        return output_numpy
