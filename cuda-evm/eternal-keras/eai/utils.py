import os
import sys
import ctypes
import subprocess
from web3 import Account
from loguru import logger
from typing import List
from eai.layer_config import LayerType, Activation, Padding, ZeroPaddingFormat

TENSORFLOW_KERAS2 = "2.15.1"
TENSORFLOW_KERAS3 = "2.16.1"
ETHER_PER_WEI = 10**18
MAX_32_BITS = 1 << 32
MAX_63_BITS = 1 << 63
MAX_64_BITS = 1 << 64


def to_i64(a):
    return ctypes.c_int64(a).value


Logger = logger
Logger.remove()
Logger.add(
    lambda msg: print(msg, end=""),
    format="<green>{time:YYYY-MM-DD HH:mm:ss}</green> | "
    "<level>{level: <8}</level> | "
    "<cyan>{name}</cyan>:<cyan>{function}</cyan>:<cyan>{line}</cyan> - "
    "<level>{message}</level>",
    level="INFO",
    colorize=True
)


def create_web3_account():
    account = Account.create()
    return {"address": account.address,
            "private_key": account._private_key.hex()}


def publisher():
    private_key = os.environ.get("PRIVATE_KEY", None)
    if private_key is None:
        return None
    return Account.from_key(private_key).address


def getLayerType(name: str) -> int:
    try:
        return LayerType[name]
    except BaseException:
        raise Exception("Layer type not found")


def getActivationType(name: str) -> int:
    try:
        return Activation[name]
    except BaseException:
        raise Exception("Activation function type not found")


def getPaddingType(name: str) -> int:
    try:
        return Padding[name]
    except BaseException:
        raise Exception("Padding type not found")


def getZeroPadding2DType(name: str) -> int:
    try:
        return ZeroPaddingFormat[name]
    except BaseException:
        raise Exception("ZeroPadding2D type not found")


def getConvSize(
        dim: List[int],
        size: List[int],
        stride: List[int],
        padding: str):
    out = []
    pad = []
    for i in range(dim):
        if padding == "same":
            out.append((dim[i] + stride[i] - 1) // stride[i])
            total_pad = max(
                size[i] -
                stride[i],
                0) if (
                dim[i] %
                stride[i] == 0) else max(
                size[i] -
                dim[i] %
                stride[i],
                0)
            pad.append(total_pad // 2)
        elif padding == "valid":
            out.append((dim[i] - size[i]) // stride[i] + 1)
            pad.append(0)
    return {out, pad}


def convert_float32_to_uint64(arr: List[int]):
    assert len(arr) == 4
    a1 = int(arr[0] * MAX_32_BITS)
    if (a1 < 0):
        a1 = MAX_64_BITS + a1
    a2 = int(arr[1] * MAX_32_BITS)
    if (a2 < 0):
        a2 = MAX_64_BITS + a2
    a3 = int(arr[2] * MAX_32_BITS)
    if (a3 < 0):
        a3 = MAX_64_BITS + a3
    a4 = int(arr[3] * MAX_32_BITS)
    if (a4 < 0):
        a4 = MAX_64_BITS + a4
    return [a1, a2, a3, a4]


def convert_uint64_to_float32(arr: List[int]):
    assert len(arr) == 4
    if (arr[0] > MAX_63_BITS):
        arr[0] = arr[0] - MAX_64_BITS
    if (arr[1] > MAX_63_BITS):
        arr[1] = arr[1] - MAX_64_BITS
    if (arr[2] > MAX_63_BITS):
        arr[2] = arr[2] - MAX_64_BITS
    if (arr[3] > MAX_63_BITS):
        arr[3] = arr[3] - MAX_64_BITS
    return [
        arr[0] /
        MAX_32_BITS,
        arr[1] /
        MAX_32_BITS,
        arr[2] /
        MAX_32_BITS,
        arr[3] /
        MAX_32_BITS]


def merge_float32_to_uint256(arr: List[int]):
    a1, a2, a3, a4 = convert_float32_to_uint64(arr)
    return (a1 << 192) + (a2 << 128) + (a3 << 64) + a4


def parse_uint256_to_float32(bigNum: int):
    a1 = ((bigNum >> 192) & 0xFFFFFFFFFFFFFFFF)
    a2 = ((bigNum >> 128) & 0xFFFFFFFFFFFFFFFF)
    a3 = ((bigNum >> 64) & 0xFFFFFFFFFFFFFFFF)
    a4 = (bigNum & 0xFFFFFFFF)
    return convert_uint64_to_float32([a1, a2, a3, a4])


def index_last(arr, item):
    for r_idx, elt in enumerate(reversed(arr)):
        if elt == item:
            return len(arr) - 1 - r_idx


def get_script_path():
    if hasattr(sys, 'frozen'):
        return os.path.dirname(sys.executable)
    return os.path.dirname(os.path.realpath(__file__))


ENV_PATH = os.path.join(get_script_path(), ".env")


def get_keras_version():
    import subprocess
    version = "3.4.1"
    try:
        result = subprocess.run(
            ['pip', 'show', "keras"], capture_output=True, text=True, check=True)
        for line in result.stdout.splitlines():
            if line.startswith('Version:'):
                version = line.split()[1]
    except subprocess.CalledProcessError:
        return version
    return version


def handle_keras_version(format_version):
    keras_version = get_keras_version()
    if format_version == "keras2":
        if keras_version.startswith("3."):
            Logger.warning(
                f"Your Keras version is now {keras_version} not compatible with Keras2. Downgrading to Keras2 ...")
            subprocess.run(
                ["pip", "install", "tensorflow=={}".format(TENSORFLOW_KERAS2)])
    else:
        if keras_version.startswith("2."):
            Logger.warning(
                f"Your Keras version is now {keras_version} not compatible with Keras3. Upgrading to Keras3 ...")
            subprocess.run(
                ["pip", "install", "tensorflow=={}".format(TENSORFLOW_KERAS3)])
    import keras
    Logger.success(f"Keras version is now {keras.__version__}")
