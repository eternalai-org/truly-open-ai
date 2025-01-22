import os
os.environ['TF_CPP_MIN_LOG_LEVEL'] = '3' 
import argparse
import json
import keras
import sys
import numpy as np
from keras import layers

TEST_DIR = "scripts/tests/layers"

parser = argparse.ArgumentParser(description='Test add layer forward output')
parser.add_argument('--config-path', type=str, help='Config path')
parser.add_argument('--output-path', type=str, help='Output path')

args = parser.parse_args()
config_path = args.config_path
output_path = args.output_path

with open(config_path, "r") as f:
    config = json.loads(f.read())

input1 = np.asarray(config["input1"])
input2 = np.asarray(config["input2"])
input1 = np.expand_dims(input1, axis=0)
input2 = np.expand_dims(input2, axis=0)

input1_layer = keras.Input(shape=(input1.shape[1:]))
input2_layer = keras.Input(shape=(input2.shape[1:]))

added = layers.Add()([input1_layer, input2_layer])

model = keras.Model(inputs=[input1_layer, input2_layer], outputs=added)

output = model.predict([input1, input2])

with open(output_path, "w") as f:
    f.write(json.dumps(output[0].tolist()))