import os
import pickle
import tqdm
import time
from web3 import Web3
from web3 import Account
from eth_abi import encode
from typing import List, Tuple
from web3.contract import Contract
from eai.utils import Logger as logger, ETHER_PER_WEI, MAX_32_BITS
from eai.layer_config import LayerType, InputType
from eai.layer_config import LayerData, LayerConfig
from web3.middleware import construct_sign_and_send_raw_middleware
import importlib
from colorama import Fore
from eai.network_config import GAS_LIMIT, CHUNK_LEN, NETWORK, MAX_PRIORITY_FEE_PER_GAS, DEPLOY_TIMEOUT
from eai.utils import index_last, getLayerType, getActivationType, getPaddingType, get_script_path, getZeroPadding2DType


class ModelDeployer():
    def __init__(self, network: str = None):
        network_mode = network if network is not None else os.environ["NETWORK_MODE"]
        self.node_endpoint = NETWORK[network_mode]["NODE_ENDPOINT"]
        self.w3 = Web3(
            Web3.HTTPProvider(
                self.node_endpoint,
                request_kwargs={
                    'timeout': DEPLOY_TIMEOUT}))
        self.private_key = os.environ["PRIVATE_KEY"]
        self.chunk_len = CHUNK_LEN
        self.address = Account.from_key(self.private_key).address
        self.w3.middleware_onion.add(
            construct_sign_and_send_raw_middleware(self.private_key))
        self.cache_data = {}
        self.cache_file = os.path.join(get_script_path(), ".cache")
        if os.path.exists(self.cache_file):
            with open(self.cache_file, "rb") as f:
                self.cache_data = pickle.load(f)

    def _update_cache(self):
        with open(self.cache_file, "wb") as f:
            pickle.dump(self.cache_data, f)

    def _deploy_from_artifact(self, contract_artifact) -> type[Contract]:
        start = time.time()
        tx_hash = self.w3.eth.contract(
            abi=contract_artifact['abi'],
            bytecode=contract_artifact['bytecode']).constructor().transact(
            {
                "from": self.address,
                "maxPriorityFeePerGas": MAX_PRIORITY_FEE_PER_GAS})
        receipt = self.w3.eth.wait_for_transaction_receipt(tx_hash)
        if receipt['status'] != 1:
            raise Exception('tx failed', receipt)
        total_time = time.time() - start
        contract_address = receipt['contractAddress']
        total_cost = (receipt.gasUsed *
                      receipt.effectiveGasPrice) / ETHER_PER_WEI
        logger.success(f'Contract has been deployed to address {contract_address}, tx={tx_hash.hex()}, gas used: {receipt.gasUsed}, transaction cost: {total_cost} EAI, total time: {total_time} seconds.')
        return {
            "address": contract_address,
            "tx": tx_hash.hex(),
            "total_cost": total_cost,
            "gas_used": receipt.gasUsed,
            "total_time": total_time}

    def _get_model_config(self, layers) -> Tuple[List[LayerData], int]:
        layersData = []
        totalWeights = 0
        for i in range(len(layers)):
            layer = layers[i]
            configData = ""
            layerType = getLayerType(layer['class_name'])
            logger.info(f"Layer {i}: {layer['class_name']}, type: {layerType}")
            inputIndices = list(
                map(lambda node: node['args'][0]['idx'], layer['inbound_nodes']))

            if layerType == LayerType.Dense:
                inputNode = layer['inbound_nodes'][0]['args'][0]
                inputUnits = inputNode['shape'][1]
                units = layer['layer_config']['units']
                useBias = layer['layer_config']['use_bias']
                activationFn = getActivationType(
                    layer['layer_config']['activation'])
                configData = encode(["uint8", "uint256", "uint256", "bool"], [
                                    activationFn.value, units, inputUnits, useBias])
                if useBias:
                    totalWeights += inputUnits * units + units
                else:
                    totalWeights += inputUnits * units + units
            elif layerType == LayerType.Flatten:
                configData = encode([], [])
            elif layerType == LayerType.Rescaling:
                n1 = int(layer['layer_config']['scale'] * MAX_32_BITS)
                n2 = int(layer['layer_config']['offset'] * MAX_32_BITS)
                configData = encode(["int64", "int64"], [n1, n2])
            elif layerType == LayerType.Softmax:
                axis = layer['layer_config']['axis']
                configData = encode(["int"], [axis])
            elif layerType == LayerType.Concatenate:
                axis = layer['layer_config']['axis']
                configData = encode(["int"], [axis])
            elif layerType == LayerType.ReLU:
                configData = encode([], [])
            elif layerType == LayerType.Sigmoid:
                configData = encode([], [])
            elif layerType == LayerType.Linear:
                configData = encode([], [])
            elif layerType == LayerType.InputLayer:
                dim = layer['layer_config']['batch_input_shape']
                pos = index_last(dim, None)
                dim = dim[pos + 1:]
                if len(dim) == 0:
                    configData = encode(["uint8"], [InputType.Scalar.value])
                elif len(dim) == 1:
                    n = dim[0]
                    configData = encode(["uint8", "uint[]"], [
                                        InputType.Tensor1D.value, [n]])
                elif len(dim) == 2:
                    n = dim[0]
                    m = dim[1]
                    configData = encode(["uint8", "uint[]"], [
                                        InputType.Tensor2D.value, [n, m]])
                elif len(dim) == 3:
                    n = dim[0]
                    m = dim[1]
                    p = dim[2]
                    configData = encode(["uint8", "uint[]"], [
                                        InputType.Tensor3D.value, [n, m, p]])
            elif layerType == LayerType.MaxPooling2D:
                f_w = layer['layer_config']['pool_size'][0]
                f_h = layer['layer_config']['pool_size'][1]
                s_w = layer['layer_config']['strides'][0]
                s_h = layer['layer_config']['strides'][1]
                padding = layer['layer_config']['padding']

                configData = encode(["uint[2]", "uint[2]", "uint8"], [
                    [f_w, f_h],
                    [s_w, s_h],
                    getPaddingType(padding).value,
                ])
            elif layerType == LayerType.Conv2D:
                inputNode = layer['inbound_nodes'][0]['args'][0]
                inputFilters = inputNode['shape'][3]
                outputFilters = layer['layer_config']['filters']
                f_w = layer['layer_config']['kernel_size'][0]
                f_h = layer['layer_config']['kernel_size'][1]
                s_w = layer['layer_config']['strides'][0]
                s_h = layer['layer_config']['strides'][1]
                padding = layer['layer_config']['padding']
                activationFn = getActivationType(
                    layer['layer_config']['activation'])
                useBias = layer['layer_config']['use_bias']

                configData = encode(["uint8", "uint", "uint", "uint[2]", "uint[2]", "uint8", "bool"], [
                    activationFn.value,
                    inputFilters,
                    outputFilters,
                    [f_w, f_h],
                    [s_w, s_h],
                    getPaddingType(padding).value,
                    useBias,
                ])
                if useBias:
                    totalWeights += f_w * f_h * inputFilters * outputFilters + outputFilters
                else:
                    totalWeights += f_w * f_h * inputFilters * outputFilters

            elif layerType == LayerType.BatchNormalization:
                inputDim = layer['layer_config']['input_dim']
                momentum = int(layer['layer_config']['momentum'] * MAX_32_BITS)
                epsilon = int(layer['layer_config']['epsilon'] * MAX_32_BITS)
                configData = encode(["uint256", "int64", "int64"], [
                                    inputDim, momentum, epsilon])
                totalWeights += inputDim * 4
            elif layerType == LayerType.Embedding:
                inputDim = layer['layer_config']['input_dim']
                outputDim = layer['layer_config']['output_dim']
                configData = encode(["uint256", "uint256"], [
                                    inputDim, outputDim])
                totalWeights += inputDim * outputDim
            elif layerType == LayerType.SimpleRNN:
                inputNode = layer['inbound_nodes'][0]['args'][0]
                inputUnits = inputNode['shape'][2]
                units = layer['layer_config']['units']
                activationFn = getActivationType(
                    layer['layer_config']['activation'])
                configData = encode(["uint8", "uint256"], [
                                    activationFn.value, units])
                totalWeights += inputUnits * units + units * units + units
            elif layerType == LayerType.LSTM:
                inputNode = layer['inbound_nodes'][0]['args'][0]
                inputUnits = inputNode['shape'][2]
                units = layer['layer_config']['units']
                activationFn = getActivationType(
                    layer['layer_config']['activation'])
                recActivationFn = getActivationType(
                    layer['layer_config']['recurrent_activation'])
                configData = encode(["uint8", "uint8", "uint256", "uint256"], [
                                    activationFn.value, recActivationFn.value, units, inputUnits])
                totalWeights += inputUnits * units * 4 + units * units * 4 + units * 4
            elif layerType == LayerType.Dropout:
                configData = encode([], [])
            elif layerType == LayerType.ZeroPadding2D:
                dataFormat = getZeroPadding2DType(
                    layer['layer_config']['data_format'])
                padding = layer['layer_config']['padding']
                configData = encode(["uint[4]", "uint8"], [
                                    padding, dataFormat.value])
            elif layerType == LayerType.GlobalAveragePooling2D:
                configData = encode([], [])
            elif layerType == LayerType.AveragePooling2D:
                pool_size = layer['layer_config']['pool_size']
                strides = layer['layer_config']['strides']
                paddingType = getPaddingType(padding)
                configData = encode(["uint[2]", "uint[2]", "uint8"], [
                                    pool_size, strides, paddingType.value])
            layersData.append(LayerData(
                layerType,
                layer['class_name'],
                configData,
                inputIndices,
            ))

        return layersData, totalWeights

    def _uploadChunk(
            self,
            model: type[Contract],
            weights: List[float],
            start_idx: int,
            end_idx: int):
        start = time.time()
        weightsToUpload = weights[start_idx: end_idx]
        appendWeightTxHash = model.functions.appendWeights(
            weightsToUpload,
            start_idx +
            1).transact(
            {
                "from": self.address,
                "gas": GAS_LIMIT,
                "maxPriorityFeePerGas": MAX_PRIORITY_FEE_PER_GAS})
        receipt = self.w3.eth.wait_for_transaction_receipt(
            appendWeightTxHash)
        if receipt['status'] != 1:
            raise Exception('tx failed', receipt)
        total_time = time.time() - start
        total_cost = (receipt.gasUsed *
                      receipt.effectiveGasPrice) / ETHER_PER_WEI
        logger.success(
            f"tx: {appendWeightTxHash.hex()}, gas used: {receipt.gasUsed}, transaction cost: {total_cost} EAI, total time: {total_time} seconds.")
        return {
            "tx": appendWeightTxHash.hex(),
            "gas_used": receipt.gasUsed,
            "total_cost": total_cost,
            "total_time": total_time}

    def _uploadWeights(self, contract, weights):
        logger.info(
            f"Uploading model weights with chunk length: {self.chunk_len} ...")
        data = contract.functions.model().call()
        required_weights = data[1]
        appended_weights = data[2]
        assert len(weights) == required_weights, f"Expected {required_weights} weights, got {len(weights)}"
        retry = True
        total_cost = 0.0
        total_time = 0.0
        logger.info(
            f"Weights size: {len(weights)}, total length: {len(weights) *32} bytes")
        while self.chunk_len > 0 and retry:
            try:
                logger.info(
                    f"Uploading weights from index {appended_weights} ...")
                for l in tqdm.tqdm(
                    range(
                        appended_weights,
                        len(weights),
                        self.chunk_len),
                    desc="Uploading weights",
                    bar_format="{l_bar}%s{bar}%s{r_bar}" %
                    (Fore.WHITE,
                     Fore.RESET)):
                    result = self._uploadChunk(
                        contract, weights, l, l + self.chunk_len)
                    self.deployed_model_contract["weights"] = result
                    total_cost += result["total_cost"]
                    total_time += result["total_time"]
                retry = False
            except Exception as e:
                logger.error(f"Error uploading weights: {e}")
                time.sleep(5)
                data = contract.functions.model().call()
                appended_weights = data[2]
                self.chunk_len //= 2
                logger.warning(
                    f"Retrying with chunk length: {self.chunk_len} ...")
        return {"total_cost": total_cost, "total_time": total_time}

    def _deploy_layer(self, layer_data: LayerData) -> LayerConfig:
        start = time.time()
        artifact_name = layer_data.layerName
        if not artifact_name.endswith("Layer"):
            artifact_name += "Layer"
        try:
            submodule = importlib.import_module(
                f"eai.artifacts.layers.{artifact_name}")
        except Exception as e:
            raise Exception(
                f"Layer {layer_data.layerName} is not supported: {e}")
        artifact = getattr(submodule, "CONTRACT_ARTIFACT")
        tx_hash = self.w3.eth.contract(
            abi=artifact['abi'],
            bytecode=artifact['bytecode']).constructor(
            layer_data.configData).transact(
            {
                "from": self.address,
                "maxPriorityFeePerGas": MAX_PRIORITY_FEE_PER_GAS})
        receipt = self.w3.eth.wait_for_transaction_receipt(tx_hash)
        if receipt['status'] != 1:
            raise Exception('tx failed', receipt)
        total_time = time.time() - start
        contract_address = receipt['contractAddress']
        tx = tx_hash.hex()
        total_cost = (receipt.gasUsed *
                      receipt.effectiveGasPrice) / ETHER_PER_WEI
        logger.success(
            f"Layer {layer_data.layerName} has been deployed to address {contract_address}, tx={tx}, gas used: {receipt.gasUsed}, transaction cost: {total_cost} EAI, total time: {total_time} seconds.")
        return {
            "layerTypeIndex": layer_data.layerType.value,
            "address": contract_address,
            "inputIndices": layer_data.inputIndices,
            "tx": tx,
            "total_cost": total_cost,
            "gas_used": receipt.gasUsed,
            "total_time": total_time}

    def _get_contract(self, address):
        submodule = importlib.import_module(
            f"eai.artifacts.models.FunctionalModel")
        contract_artifact = getattr(submodule, "CONTRACT_ARTIFACT")

        return self.w3.eth.contract(
            address=address, abi=contract_artifact['abi'])

    def _deploy_model_contract(self):
        logger.info("Deploying FunctionalModel contract...")
        submodule = importlib.import_module(
            f"eai.artifacts.models.FunctionalModel")
        contract_artifact = getattr(submodule, "CONTRACT_ARTIFACT")

        deployed_contract = self._deploy_from_artifact(contract_artifact)
        self.deployed_model_contract = {
            "address": deployed_contract["address"],
            "layer_configs": [],
            "model_construct_data": {
                "is_constructed": False,
            },
            "tx": deployed_contract["tx"],
            "gas_used": deployed_contract["gas_used"],
            "total_cost": deployed_contract["total_cost"],
            "total_time": deployed_contract["total_time"]
        }
        self.cache_data[self.hash_value] = self.deployed_model_contract
        self._update_cache()
        return self.w3.eth.contract(
            address=deployed_contract["address"],
            abi=contract_artifact['abi'])

    def _deploy_layers(self, layers_data):
        logger.info("Deploying layer contracts...")
        layer_configs = []
        total_cost = 0.0
        total_time = 0.0
        for idx, deployed_layer_config in enumerate(
                self.deployed_model_contract["layer_configs"]):
            layer_configs.append(LayerConfig(
                deployed_layer_config["layerTypeIndex"],
                deployed_layer_config["address"],
                deployed_layer_config["inputIndices"]
            ))
            total_cost += deployed_layer_config["total_cost"]
            total_time += deployed_layer_config["total_time"]
            logger.warning(
                f'Layer {layers_data[idx].layerName} has already been deployed to address {deployed_layer_config["address"]}, tx={deployed_layer_config["tx"]}, gas used: {deployed_layer_config["gas_used"]}, transaction cost: {deployed_layer_config["total_cost"]} EAI, total time: {deployed_layer_config["total_time"]} seconds.')
        
        for idx in range(len(layer_configs), len(layers_data)):
            config = self._deploy_layer(layers_data[idx])
            layer_configs.append(LayerConfig(
                config["layerTypeIndex"],
                config["address"],
                config["inputIndices"]
            ))
            self.deployed_model_contract["layer_configs"].append(config)
            total_cost += config["total_cost"]
            total_time += config["total_time"]
            self.cache_data[self.hash_value] = self.deployed_model_contract
            self._update_cache()
        return {
            "layer_configs": layer_configs,
            "total_cost": total_cost,
            "total_time": total_time}

    def _construct_model_if_needed(self, contract, layerConfigs):
        if not self.deployed_model_contract["model_construct_data"]["is_constructed"]:
            logger.info("Constructing model...")
            start = time.time()
            layer_config_params = [lc.toContractParams()
                                   for lc in layerConfigs]
            construct_model_tx_hash = contract.functions.constructModel(layer_config_params).transact({
                "from": self.address,
                "gas": GAS_LIMIT,
                "maxPriorityFeePerGas": MAX_PRIORITY_FEE_PER_GAS
            })
            receipt = self.w3.eth.wait_for_transaction_receipt(
                construct_model_tx_hash)
            if receipt['status'] != 1:
                raise Exception(
                    'Model construction transaction failed', receipt)
            total_time = time.time() - start
            tx = construct_model_tx_hash.hex()
            total_cost = (receipt.gasUsed *
                          receipt.effectiveGasPrice) / ETHER_PER_WEI
            self.deployed_model_contract["model_construct_data"] = {
                "is_constructed": True,
                "tx": construct_model_tx_hash.hex(),
                "gas_used": receipt.gasUsed,
                "total_cost": total_cost,
                "total_time": total_time
            }
            self.cache_data[self.hash_value] = self.deployed_model_contract
            self._update_cache()
            logger.success(
                f'Model constructed successfully, tx: {tx}, gas used: {receipt.gasUsed}, transaction cost: {total_cost} EAI, total time: {total_time} seconds.')
        else:
            logger.warning(
                f'Model has already been constructed with tx: {self.deployed_model_contract["model_construct_data"]["tx"]}, gas used: {self.deployed_model_contract["model_construct_data"]["gas_used"]}, transaction cost: {self.deployed_model_contract["model_construct_data"]["total_cost"]} EAI, total time: {self.deployed_model_contract["model_construct_data"]["total_time"]} seconds.')

    def deploy_model(self, model_data):
        assert self.private_key is not None, "Private key is required to deploy contract, please run command 'eai set-private-key' to set private key"
        total_time = 0.0
        logger.info("Deploying model to EternalAI chain ...")
        layers = model_data["model_graph"]["layers"]
        weights = model_data["weights"]
        hashed_model = model_data["hashed_model"]
        self.hash_value = hashed_model + self.node_endpoint
        total_cost = 0.0
        self.deployed_model_contract = self.cache_data.get(self.hash_value, {})
        if len(self.deployed_model_contract) > 0:
            logger.warning(
                f"Model has already been deployed at address: {self.deployed_model_contract['address']}, tx: {self.deployed_model_contract['tx']}, gas used: {self.deployed_model_contract['gas_used']}, transaction cost: {self.deployed_model_contract['total_cost']} EAI.")
            contract = self._get_contract(
                self.deployed_model_contract["address"])
        else:
            contract = self._deploy_model_contract()
        total_cost += self.deployed_model_contract["total_cost"]
        total_time += self.deployed_model_contract["total_time"]
        layers_data, total_weights = self._get_model_config(layers)
        ret = self._deploy_layers(layers_data)
        total_cost += ret["total_cost"]
        total_time += ret["total_time"]
        self._construct_model_if_needed(contract, ret["layer_configs"])
        if len(weights) > 0:
            ret = self._uploadWeights(contract, weights)
            total_cost += ret["total_cost"]
            total_time += ret["total_time"]
        logger.success(
            f'Model deployed at address: {contract.address}, total transaction cost: {total_cost} EAI, total time: {total_time} seconds.')
        self.cache_data.pop(self.hash_value)
        del self.hash_value
        del self.deployed_model_contract
        self._update_cache()
        return contract
