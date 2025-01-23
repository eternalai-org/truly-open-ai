import json
import hashlib
import importlib
import numpy as np
from eai.utils import Logger, merge_float32_to_uint256
from eai.layer_config import KERAS_ACTIVATIONS


class ModelExporter:
    def __init__(self):
        pass

    def _prepare_for_export(self):
        self.layer_indices = []
        self.last_functional_layer_indices = {}
        self.layer_idx = 0

    def _export_inbound_nodes_keras2(self, layer, layer_indices):
        """
        Export inbound nodes of a given layer.

        Args:
            layer (dict): The layer configuration.
            layer_indices (list): List of layer names.

        Returns:
            list: List of inbound node data for Keras 2.
        """
        ret = []
        build_config = layer.get("build_config", {"input_shape": []})
        inbound_nodes = layer["inbound_nodes"]
        inputs = build_config["input_shape"]
        if len(inputs) > 1:
            if isinstance(inputs[0], list):
                for idx, input_shape in enumerate(inputs):
                    node = inbound_nodes[0]
                    inbound_node_data = {"args": [], "kwargs": {}}
                    args = node[idx]
                    name = args[0]
                    if name in self.last_functional_layer_indices:
                        inbound_node_data["args"].append({
                            "name": layer_indices[self.last_functional_layer_indices[name]],
                            "idx": self.last_functional_layer_indices[name],
                            "shape": input_shape
                        })
                    else:
                        inbound_node_data["args"].append({
                            "name": args[0],
                            "idx": layer_indices.index(args[0]),
                            "shape": input_shape
                        })
                    ret.append(inbound_node_data)
            else:
                node = inbound_nodes[0]
                inbound_node_data = {"kwargs": {}, "args": []}
                args = node[0]
                name = args[0]
                if name in self.last_functional_layer_indices:
                    inbound_node_data["args"].append({
                        "name": layer_indices[self.last_functional_layer_indices[name]],
                        "idx": self.last_functional_layer_indices[name],
                        "shape": inputs
                    })
                else:
                    inbound_node_data["args"].append({
                        "name": name,
                        "idx": layer_indices.index(name),
                        "shape": inputs
                    })
                ret.append(inbound_node_data)
        return ret

    def _export_inbound_nodes_keras3(self, layer, layer_indices):
        """
        Export inbound nodes of a given layer for Keras 3.

        Args:
            layer (dict): The layer configuration.
            layer_indices (list): List of layer names.

        Returns:
            list: List of inbound node data.
        """
        ret = []
        build_config = layer.get("build_config", {"input_shape": []})
        inbound_nodes = layer["inbound_nodes"]
        inputs = build_config["input_shape"]
        if len(inputs) > 1:
            if isinstance(inputs[0], list):
                for idx, input_shape in enumerate(inputs):
                    node = inbound_nodes[0]
                    inbound_node_data = {"args": [], "kwargs": node["kwargs"]}
                    args = node['args'][0][idx]
                    name = args['config']["keras_history"][0]
                    if name in self.last_functional_layer_indices:
                        inbound_node_data["args"].append({
                            "name": layer_indices[self.last_functional_layer_indices[name]],
                            "idx": self.last_functional_layer_indices[name],
                            "shape": input_shape
                        })
                    else:
                        inbound_node_data["args"].append({
                            "name": args['config']["keras_history"][0],
                            "idx": layer_indices.index(args['config']["keras_history"][0]),
                            "shape": input_shape
                        })
                    ret.append(inbound_node_data)
            else:
                node = inbound_nodes[0]
                inbound_node_data = {"args": [], "kwargs": node["kwargs"]}
                for args in node['args']:
                    name = args['config']["keras_history"][0]
                    if name in self.last_functional_layer_indices:
                        inbound_node_data["args"].append({
                            "name": layer_indices[self.last_functional_layer_indices[name]],
                            "idx": self.last_functional_layer_indices[name],
                            "shape": inputs
                        })
                    else:
                        inbound_node_data["args"].append({
                            "name": args['config']["keras_history"][0],
                            "idx": layer_indices.index(args['config']["keras_history"][0]),
                            "shape": inputs
                        })
                ret.append(inbound_node_data)

        return ret

    def _export_model_graph(self, model_data, layers):
        """
        Export the model graph.

        Args:
            model (object): The model object.
            vocabulary (str): The vocabulary.
            output_path (str): The output path.

        Returns:
            dict: The exported model graph.
        """
        graph = {
            "layers": [],
        }
        index_layer = 0
        for idx, layer in enumerate(model_data["config"]["layers"]):
            module = layer["module"]
            if module == "keras.layers":
                layer_config = layer["config"]
                layer_name = layer_config["name"]
                class_name = layer["class_name"]
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
                data = {
                    "idx": self.layer_idx,
                    "name": layer_name,
                    "class_name": class_name,
                }
                try:
                    # Get the layer class and its configuration
                    module = importlib.import_module("eai.layers")
                    layer_class = getattr(module, class_name)(layer_config)
                    layer_config = layer_class.get_layer_config()
                    Logger.success(f"Layer {class_name} exported")
                except BaseException:
                    Logger.error(f"Layer {class_name} not supported")
                    raise Exception(f"Layer {class_name} not supported")
                data["layer_config"] = layer_config
                self.layer_indices.append(layer_name)
                if model_data["class_name"] == "Functional":
                    keras_version = importlib.import_module(
                        "keras").__version__
                    if keras_version.startswith("2."):
                        data["inbound_nodes"] = self._export_inbound_nodes_keras2(
                            layer, self.layer_indices)
                    elif keras_version.startswith("3."):
                        data["inbound_nodes"] = self._export_inbound_nodes_keras3(
                            layer, self.layer_indices)
                    else:
                        raise Exception("Keras version not supported")
                elif model_data["class_name"] == "Sequential":
                    if idx == 0:
                        data["inbound_nodes"] = []
                        if data["class_name"] == "InputLayer":
                            prev_layer_out_shape = layer_config["batch_input_shape"]
                            index_layer = 0
                        else:
                            prev_layer_out_shape = list(
                                layers[index_layer].output.shape)
                    else:
                        data["inbound_nodes"] = [{
                            "args": [
                                {
                                    "name": self.layer_indices[self.layer_idx - 1],
                                    "idx": self.layer_idx - 1,
                                    "shape": prev_layer_out_shape
                                }
                            ],
                            "kwargs": {}
                        }]
                        prev_layer_out_shape = list(
                            layers[index_layer].output.shape)
                        index_layer += 1
                else:
                    raise Exception(
                        f"Model type {model_data['class_name']} not supported")
                if class_name == "InputLayer" and self.layer_idx > 0:
                    data["inbound_nodes"] = [{
                        "args": [
                            {
                                "name": self.layer_indices[self.layer_idx - 1],
                                "idx": self.layer_idx - 1,
                                "shape": list(layers[index_layer].output.shape)
                            }
                        ],
                        "kwargs": {}
                    }]
                graph["layers"].append(data)
                self.layer_idx += 1
            elif module == "keras.src.engine.functional" or module == "keras.src.models.functional":
                sub_graph = self._export_model_graph(layer, layers[idx].layers)
                graph["layers"] = graph["layers"] + sub_graph["layers"]
                self.last_functional_layer_indices[layer["name"]
                                                   ] = self.layer_idx
            else:
                raise Exception(f"Module {module} not supported")
        return graph

    def _export_weights(self, model):
        """
        Export the model weights.

        Args:
            model (object): The model object.
            output_path (str): The output path.

        Returns:
            list: List of flattened weights.
        """
        Logger.info("Exporting model weights ...")
        layer_weights = []
        for layer in model.layers:
            w = layer.get_weights()
            for param in w:
                param = param.reshape(-1).tolist()
                N = len(param)
                to_append = 0
                if N % 4 != 0:
                    to_append = 4 - N % 4
                for i in range(to_append):
                    param.append(0)
                for i in range(0, len(param), 4):
                    mergedNumber = merge_float32_to_uint256(param[i:i + 4])
                    layer_weights.append(mergedNumber)
        Logger.success("Weights exported.")
        return layer_weights

    def _export_tf_model(self, model):
        """
        Export a Tensorflow/Keras model.

        Args:
            model (tf.keras.Model): The model to be exported.
            vocabulary (list, optional): The vocabulary used by the model. Defaults to None.
            output_dir (str, optional): The directory to save the exported model. Defaults to None.

        Returns:
            dict: A dictionary containing the exported model graph and weights.
        """
        Logger.info("Exporting Tensorflow/Keras model ...")
        # Export the model graph
        Logger.info("Exporting model graph ...")
        model_data = json.loads(model.to_json())
        model_graph = self._export_model_graph(model_data, model.layers)
        Logger.success("Model graph exported successfully.")
        # Export the weights
        weights = self._export_weights(model)
        # Create a hash of the model graph
        hashed_graph = hashlib.sha256(json.dumps(
            model_graph, sort_keys=True).encode()).hexdigest()

        # Create a hash of the weights
        hashed_weights = hashlib.sha256(
            json.dumps(weights).encode()).hexdigest()

        # Concatenate the two hashes
        hashed_model = f"{hashed_graph}{hashed_weights}"
        Logger.success("Model exported successfully.")
        return {
            "model_graph": model_graph,
            "weights": weights,
            "hashed_model": hashed_model
        }

    def export_model(self, model):
        """
        Export a Tensorflow/Keras model.

        Args:
            model (tf.keras.Model): The model to be exported.
            vocabulary (list, optional): The vocabulary used by the model. Defaults to None.
            output_dir (str, optional): The directory to save the exported model. Defaults to None.

        Returns:
            dict: A dictionary containing the exported model graph and weights.
        """
        self._prepare_for_export()
        model_data = self._export_tf_model(model)
        return model_data
