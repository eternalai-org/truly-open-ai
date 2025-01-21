from eth_account import Account
from web3 import Web3
from web3.middleware import construct_sign_and_send_raw_middleware
from eai.utils import publisher
from eai.utils import Logger, LayerType, handle_keras_version
import json
import os
import requests
from eai.network_config import NETWORK, COLLECTION_ADDRESS
from eai.model import Eternal
from eai.layer_config import KERAS_ACTIVATIONS
from eai.deployer import ModelDeployer
from eai.exporter import ModelExporter
import importlib
from eai.artifacts.collection.ModelCollection import CONTRACT_ARTIFACT as COLLECTION_ARTIFACT
from eai.utils import Logger as logger
from eth_abi import encode, decode


def register(address, name, owner):
    Logger.info("Registering model ...")
    register_endpoint = NETWORK[os.environ["NETWORK_MODE"]
                                ]["REGISTER_ENDPOINT"]
    try:
        headers = {"Content-Type": "application/json"}
        data = {
            "model_address": address,
            "model_name": name,
            "owner_address": owner
        }
        response = requests.post(
            register_endpoint, headers=headers, json=data)
        response.raise_for_status()  # Raises an HTTPError for bad responses (4xx or 5xx)

        if response.json().get("status") == 1:
            Logger.success("Model registered successfully.")
        else:
            Logger.error("Failed to register model. Response status not 1.")
    except Exception as e:
        Logger.error(f"An unexpected error occurred: {e}")


def transform(
        model,
        model_name: str = "Unnamed Model",
        format="keras3",
        network_mode: str = None):
    assert format in [
        "keras2", "keras3"], "Format must be either 'keras2' or 'keras3'"
    network = network_mode if network_mode is not None else os.environ["NETWORK_MODE"]
    handle_keras_version(format)
    import keras
    import time
    start = time.time()
    Logger.info(f"Transforming model on EternalAI's {network}...")
    if isinstance(model, str):
        model = keras.models.load_model(model)
    assert isinstance(model, keras.Model), "Model must be a keras model"
    try:
        model_data = ModelExporter().export_model(model)
    except Exception as e:
        Logger.error(f"Failed to export model: {e}")
        return None
    try:
        contract = ModelDeployer(network).deploy_model(model_data)
    except Exception as e:
        Logger.error(f"Failed to deploy model: {e}")
        return None
    address = contract.address
    register(address, model_name, publisher())
    try:
        eai_model = Eternal(address)
    except Exception as e:
        Logger.error(f"Failed to load model from address: {e}")
    Logger.success(
        f"Model transformed successfully on EternalAI's {network}. Time taken: {time.time() - start} seconds")
    return eai_model


def check_keras_graph(model_data, layers):
    layer_names = []
    labels = []
    for idx, layer in enumerate(model_data["config"]["layers"]):
        module = layer["module"]
        if module == "keras.layers":
            class_name = layer["class_name"]
            layer_config = layer["config"]
            if class_name == "BatchNormalization":
                input_shape = layers[idx].input.shape
                axis = layer_config["axis"]
                if isinstance(axis, list):
                    axis = axis[0]
                layer_config["input_dim"] = input_shape[axis]
            elif class_name == "Activation":
                if layer_config["activation"] in KERAS_ACTIVATIONS:
                    class_name = KERAS_ACTIVATIONS[layer_config["activation"]]
                else:
                    raise Exception(
                        f"Activation {layer_config['activation']} is not supported")
            layer_names.append(class_name)
            try:
                module = importlib.import_module("eai.layers")
                layer_class = getattr(module, class_name)(layer_config)
                labels.append(1)
            except Exception as e:
                labels.append(0)
        elif module == "keras.src.engine.functional" or module == "keras.src.models.functional":
            _layer_names, _labels = check_keras_graph(
                layer, layers[idx].layers)
            layer_names.extend(_layer_names)
            labels.extend(_labels)
        else:
            raise Exception(f"Module {module} not supported")

    return layer_names, labels


def check_keras_model(model, output_path: str = None):
    try:
        model_data = json.loads(model.to_json())
    except Exception as e:
        Logger.error(f"Failed to load model data: {e}")
        return

    Logger.info("Checking model layers ...")

    layer_names, labels = check_keras_graph(
        model_data, model.layers)
    supported_layers = 0
    unsupported_layers = 0
    error_layers = []
    for idx, layer_name in enumerate(layer_names):
        if labels[idx] == 1:
            Logger.success(f"{idx}: Layer {layer_name} is supported.")
            supported_layers += 1
        else:
            if layer_name not in error_layers:
                Logger.error(f"{idx}: Layer {layer_name} is not supported.")
                error_layers.append(layer_name)
            unsupported_layers += 1
    response = {
        "status": 1
    }
    if len(error_layers) > 0:
        response["status"] = -1
        response["error"] = []
        for error_layer in error_layers:
            response["error"].append(f"Layer {error_layer} is not supported")
        if output_path is not None:
            with open(output_path, "w") as f:
                json.dump(response, f)
    else:
        if output_path is not None:
            with open(output_path, "w") as f:
                json.dump(response, f)

    Logger.info(
        f"Summary: {supported_layers} layers supported, {unsupported_layers} layers not supported.")
    return response


def check(model, format="keras3", output_path=None):
    assert format in [
        "keras2", "keras3"], "Format must be either 'keras2' or 'keras3'"
    handle_keras_version(format)
    import keras
    Logger.info(f"Loading model from {model} ...")
    if isinstance(model, str):
        try:
            model = keras.models.load_model(model)
            Logger.success("Model loaded successfully.")
        except Exception as e:
            response = {
                "status": -1,
                "error": str(e)
            }
            with open(output_path, "w") as f:
                json.dump(response, f)
            raise Exception(f"Failed to load model: {e}")
    response = check_keras_model(model, output_path)
    if response["status"] == -1:
        return False
    return True


def layers():
    return list(LayerType.__members__.keys())


def get_model(model: str):
    eternal = Eternal(model)
    return eternal


def transfer_model(model_id: str, to_address: str, network_mode: str = None):
    network = network_mode if network_mode is not None else os.environ["NETWORK_MODE"]
    private_key = os.environ["PRIVATE_KEY"]
    node_endpoint = NETWORK[network]["NODE_ENDPOINT"]
    w3 = Web3(Web3.HTTPProvider(node_endpoint,
                                request_kwargs={'timeout': 300}))
    collection_contract_abi = COLLECTION_ARTIFACT['abi']
    collection_contract = w3.eth.contract(
        address=COLLECTION_ADDRESS, abi=collection_contract_abi)
    from_address = Account.from_key(private_key).address
    w3.middleware_onion.add(
        construct_sign_and_send_raw_middleware(private_key))

    model_id_uint = decode(['uint256'], encode(
        ['uint256'], [int(model_id)]))[0]

    logger.info("Approving token transfer...")
    approve_tx_hash = collection_contract.functions.approve(
        to_address, model_id_uint).transact({"from": from_address})
    approve_receipt = w3.eth.wait_for_transaction_receipt(approve_tx_hash)
    if approve_receipt['status'] != 1:
        raise Exception('tx failed', approve_receipt)
    logger.success(
        f"Approved transfering model with tokenId {model_id_uint} to wallet {to_address} (tx: {approve_tx_hash.hex()})")
    logger.info("Transfering token")
    transfer_tx_hash = collection_contract.functions.safeTransferFrom(
        from_address, to_address, model_id_uint).transact({"from": from_address})
    transfer_receipt = w3.eth.wait_for_transaction_receipt(transfer_tx_hash)
    if transfer_receipt['status'] != 1:
        raise Exception('tx failed', transfer_receipt)
    logger.success(
        f"Transfered model with tokenId {model_id} to wallet {to_address} (tx: {transfer_tx_hash.hex()})")
