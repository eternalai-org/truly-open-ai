import os
os.environ['TF_CPP_MIN_LOG_LEVEL'] = '3' 
import argparse
import json
import keras
import sys
import numpy as np
from keras import layers

TEST_DIR = "scripts/tests/layers"

parser = argparse.ArgumentParser(description='Test maxPooling layer forward output')
parser.add_argument('--config-path', type=str, help='Config path')
parser.add_argument('--output-path', type=str, help='Output path')

args = parser.parse_args()
config_path = args.config_path
output_path = args.output_path

with open(config_path, "r") as f:
    config = json.loads(f.read())

size = config["size"]
stride = config["stride"]
padding = config["padding"]
input = np.asarray(config["inputs"])
input = np.expand_dims(input, axis=0)

model = keras.models.Sequential([
    layers.InputLayer(shape=input.shape[1:]),
    layers.MaxPooling2D(pool_size=(size[0],size[1]), strides=(stride[0],stride[1]), padding=padding),
])

output = model.predict(input)

with open(output_path, "w") as f:
    f.write(json.dumps(output[0].tolist()))