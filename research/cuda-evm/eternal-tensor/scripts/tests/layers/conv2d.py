import os
os.environ['TF_CPP_MIN_LOG_LEVEL'] = '3' 
import argparse
import json
import keras
import sys
import numpy as np
from keras import layers, activations

TEST_DIR = "scripts/tests/layers"

def leaky_relu(x):
  return activations.relu(x, negative_slope=0.2)


def get_activation_func(name):
    if name == "relu":
        return activations.relu
    if name == "tanh":
        return activations.tanh
    if name == "leakyrelu":
        return leaky_relu
    if name == "sigmoid":
        return activations.sigmoid
    if name == "softmax":
        return activations.softmax
    if name == "linear":
        return activations.linear
    sys.exit("Invalid activation function name")


def splice(params, shape):
    cnt = 1
    for x in shape:
        cnt *= x
    return np.asarray(params[:cnt]).reshape(shape), params[cnt:]

parser = argparse.ArgumentParser(description='Test conv2d layer forward output')
parser.add_argument('--config-path', type=str, help='Config path')
parser.add_argument('--output-path', type=str, help='Output path')

args = parser.parse_args()
config_path = args.config_path
output_path = args.output_path

with open(config_path, "r") as f:
    config = json.loads(f.read())

inputFilters = config["inputFilters"]
outputFilters = config["outputFilters"]
size = config["size"]
stride = config["stride"]
useBias = config["useBias"]
activation = get_activation_func(config["activation"])
padding = config["padding"]
params = config["params"]
input = np.asarray(config["inputs"])
input = np.expand_dims(input, axis=0)

model = keras.models.Sequential([
    layers.InputLayer(shape=input.shape[1:]),
    layers.Conv2D(kernel_size=size, activation=activation, strides=stride, padding=padding),
])

weights, params = splice(params, [size[0], size[1], inputFilters, outputFilters])
if useBias:
    bias, params = splice(params, [outputFilters])
else:
    bias = []

all_weights = [weights]
if useBias:
    all_weights.append(bias)

model.layers[0].set_weights(all_weights)

output = model.predict(input)

with open(output_path, "w") as f:
    f.write(json.dumps(output[0].tolist()))