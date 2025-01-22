import os
os.environ['TF_CPP_MIN_LOG_LEVEL'] = '3' 
import argparse
import json
import keras
import sys
import numpy as np
from keras import layers, activations

TEST_DIR = "scripts/tests/layers"

def splice(params, shape):
    cnt = 1
    for x in shape:
        cnt *= x
    return np.asarray(params[:cnt]).reshape(shape), params[cnt:]

parser = argparse.ArgumentParser(description='Test batchnorm layer forward output')
parser.add_argument('--config-path', type=str, help='Config path')
parser.add_argument('--output-path', type=str, help='Output path')

args = parser.parse_args()
config_path = args.config_path
output_path = args.output_path

with open(config_path, "r") as f:
    config = json.loads(f.read())

inputUnits = config["inputUnits"]
momentum = config["momentum"]
epsilon = config["epsilon"]
params = config["params"]
input = np.asarray(config["inputs"])
input = np.expand_dims(input, axis=0)

model = keras.models.Sequential([
    layers.InputLayer(shape=input.shape[1:]),
    layers.BatchNormalization(momentum=momentum,epsilon=epsilon),
])

weights_gamma, params = splice(params, [inputUnits])
weights_beta, params = splice(params,[inputUnits])
weights_moving_mean, params = splice(params, [inputUnits])
weights_moving_variance, params = splice(params, [inputUnits])

all_weights = [weights_gamma]
all_weights.append(weights_beta)
all_weights.append(weights_moving_mean)
all_weights.append(weights_moving_variance)

model.layers[0].set_weights(all_weights)

output = model.predict(input)

with open(output_path, "w") as f:
    f.write(json.dumps(output[0].tolist()))