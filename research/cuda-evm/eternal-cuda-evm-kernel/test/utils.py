import numpy as np
import ctypes
import os

def compress_uint256(a1 = 0, a2 = 0, a3 = 0, a4 = 0):
        return a4 + (a3 << 64) + (a2 << 128) + (a1 << 192)
    
def to_i64(a):
    return ctypes.c_int64(a).value
    
def unpack_uint256(a):
    
    a1 = to_i64((a >> 192) & 0xFFFFFFFFFFFFFFFF)
    a2 = to_i64((a >> 128) & 0xFFFFFFFFFFFFFFFF)
    a3 = to_i64((a >> 64) & 0xFFFFFFFFFFFFFFFF)
    a4 = to_i64(a & 0xFFFFFFFFFFFFFFFF)
    
    return a1, a2, a3, a4

def chunked(l, n):
    for i in range(0, len(l), n):
        yield l[i:i+n]
    
def relu(x):
    return np.maximum(0, x)

def sigmoid(x): 
    return 1 / (1 + np.exp(-x))

def tanh(x):
    exp_x = np.exp(x)
    exp_nx = np.exp(-x)
    return (exp_x - exp_nx) / (exp_x + exp_nx)

def softmax(x):
    e_x = np.exp(x - np.max(x))
    return e_x / e_x.sum()

def compare_tensors(t1, t2):
    return all(x == y for x, y in zip(t1, t2))

def mae(x, y):
    return np.mean(np.abs(x - y))

def compare_tensors(t1, t2, eps=1e-6):
    return all(abs(x - y) < eps for x, y in zip(t1, t2))

def absolute_or_relative_error(a, b): # absolute or relative error
    return np.abs(a - b) / np.maximum(1.0, b)

def log(*msg):
    if 'benchmark' not in os.environ:
        print(*msg)
