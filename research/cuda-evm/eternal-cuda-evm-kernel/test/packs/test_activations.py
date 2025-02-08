from tensor import Tensor
from op import execute, Operation
from utils import sigmoid, softmax, tanh, relu, absolute_or_relative_error
from .test_registry import wrap_test
import time

@wrap_test(
    name='sigmoid activations test',
    repeat=100,
    meta={
        'description': 'Test sigmoid activation function',
    },
    params={
        'test_size': [(1 << i) for i in range(10, 16)],
        'factor': [(1 << i) for i in range(1, 7)]
    }
)
def test_sigmoid(test_size, factor):
    sigmoid_opcode = Operation.SIGMOID

    tin = Tensor.random_tensor([test_size])
    tin._data += 0.5
    tin._data *= factor

    sigmoid_out , stats = execute(sigmoid_opcode, [], tin)    
    t_start = time.time()

    expected_sigmoid = sigmoid(tin.data)
    stats['cpu_based'] = time.time() - t_start

    error = absolute_or_relative_error(sigmoid_out.data, expected_sigmoid).mean()
    stats['error'] = error

    return stats  


@wrap_test(
    name='softmax activations test',
    repeat=100,
    meta={
        'description': 'Test softmax activation function',
    },
    params={
        'test_size': [(1 << i) for i in range(10, 16)],
        'factor': [(1 << i) for i in range(1, 7)]   
    }
)
def test_softmax(test_size, factor):
    softmax_opcode = Operation.SOFTMAX

    tin = Tensor.random_tensor([test_size])
    tin._data += 0.5
    tin._data *= factor

    softmax_out , stats = execute(softmax_opcode, [], tin)    
    
    t_start = time.time()
    expected_softmax = softmax(tin.data)
    stats['cpu_based'] = time.time() - t_start
    
    error = absolute_or_relative_error(softmax_out.data, expected_softmax).mean()
    stats['error'] = error

    return stats  


@wrap_test(
    name='tanh activations test',
    repeat=100,
    meta={
        'description': 'Test tanh activation function',
    },
    params={
        'test_size': [(1 << i) for i in range(10, 16)],
        'factor': [(1 << i) for i in range(1, 7)]    
    }
)
def test_tanh(test_size, factor):
    tanh_opcode = Operation.TANH

    tin = Tensor.random_tensor([test_size])
    tin._data += 0.5
    tin._data *= factor

    tanh_out , stats = execute(tanh_opcode, [], tin)    
    
    t_start = time.time()
    expected_tanh = tanh(tin.data)
    stats['cpu_based'] = time.time() - t_start
    
    error = absolute_or_relative_error(tanh_out.data, expected_tanh).mean()
    stats['error'] = error

    return stats


@wrap_test(
    name='relu activations test',
    repeat=100,
    meta={
        'description': 'Test relu activation function',
    },
    params={
        'test_size': [2 ** i for i in range(10, 16)],
        'factor': [2 ** i for i in range(1, 7)]
    }
)
def test_relu(test_size, factor):
    relu_opcode = Operation.RELU

    tin = Tensor.random_tensor([test_size])
    tin._data += 0.5
    tin._data *= factor

    relu_out , stats = execute(relu_opcode, [], tin)    
    
    t_start = time.time()
    expected_relu = relu(tin.data)  
    stats['cpu_based'] = time.time() - t_start
    
    error = absolute_or_relative_error(relu_out.data, expected_relu).mean()
    stats['error'] = error

    return stats
