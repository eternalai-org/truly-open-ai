from enum import Enum
from tensor import Tensor
import base64
import time
from utils import log as wraplog

class Operation(int, Enum):
    CONV2D = 0, # 0
    MAXPOOLING2D = 1, # 1
    AVGPOOLING2D = 2, # 2

    # Matrix operations
    MATMUL = 3, # 3

    # Elementwise operations
    ELEMENTWISE_ADD = 4, # 4
    ELEMENTWISE_MUL = 5, # 5
    ELEMENTWISE_SUB = 6, # 6
    ELEMENTWISE_DIV = 7, # 7

    # Transforms
    TRANSFORM_EXP = 8, # 8
    TRANSFORM_SQRT = 9, # 9

    # Normalizations
    BATCH_NORM = 10, # 10
    LAYER_NORM = 11, # 11 
    ZSCORE = 12, # 12
    MIN_MAX_SCALE = 13, # 13

    # merging operations
    CONCATENATE = 14, # 14

    # Activations
    RELU = 15, # 15
    TANH = 16, # 16
    SIGMOID = 17, # 17
    SOFTMAX = 18, # 18
    LOGSOFTMAX = 19, # 19
    SOFTMAX2D = 20, # 20

    # Reductions
    REDUCTION_MAX = 21, # 21
    REDUCTION_MIN = 22, # 22
    REDUCTION_MEAN = 23, # 23
    REDUCTION_SUM = 24, # 24
    REDUCTION_ARGMAX = 25, # 25
    REDUCTION_ARGMIN = 26, # 26

    # misc
    DROPOUT = 27, # 27
    GLOBAL_AVGPOOLING2D = 28 # 28
    
    RESCALE = 29, # 29
    CHANNEL_WISE_MEAN_REDUCTION = 30, # 30    
    CHANNEL_WISE_SUM_REDUCTION = 31, # 31   
    DEPTHWISE_CONV2D = 32
    
from eth_abi import encode as abi_encode, decode as abi_decode
import ctypes
import os

dll = ctypes.CDLL(os.path.join(os.getcwd(), 'libcomputelib.so'))
dll.cuda_execute_operation.restype = ctypes.POINTER(ctypes.c_ubyte)

def encode(op: int, params: list, tensors: list[Tensor]) -> bytes:
    if not isinstance(params, (list, tuple)):
        params = [params]
        
    if not isinstance(tensors, (list, tuple)):
        tensors = [tensors]
    
    shapes = [e.shape for e in tensors]
    data = [e.compress() for e in tensors]
    return abi_encode(('uint64', 'uint64[]', 'uint64[][]', 'int256[][]'), (int(op), params, shapes, data))

def decode(b: bytes) -> Tensor:
    (data, shape) = abi_decode(('int256[]', 'uint64[]'), b)
    return Tensor.uncompress(data, shape)

def execute(op: int, params, tensor: Tensor) -> Tensor:
    time_track = {}

    t_start_0 = time.time()
    b = encode(op, params, tensor)
    time_track['encode'] = time.time() - t_start_0
    
    length_out = ctypes.c_int()
    has_error = ctypes.c_int()

    length_out_ptr = ctypes.pointer(length_out)
    has_error_ptr = ctypes.pointer(has_error)
    
    t_start = time.time()
    out = dll.cuda_execute_operation(b, len(b), length_out_ptr, has_error_ptr) 
    time_track['cuda_execute'] = time.time() - t_start

    if has_error.value != 0:
        raise ValueError('CUDA error')
    
    deref = bytearray(ctypes.cast(out, ctypes.POINTER(ctypes.c_char * length_out.value)).contents)
    
    t_start = time.time()
    (tensor_out, shape_out) = abi_decode(('uint256[]', 'uint64[]'), deref)
    tout = Tensor.uncompress(tensor_out, shape_out)
    time_track['decode'] = time.time() - t_start

    dll.deallocate_cpp_response(out)
    
    time_track['total'] = time.time() - t_start_0

    return tout, time_track