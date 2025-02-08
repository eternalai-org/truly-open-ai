import numpy as np
from concurrent.futures import ProcessPoolExecutor

from tensor import Tensor
from utils  import absolute_or_relative_error, log as wrap_log
from op import Operation, execute
from .test_registry import wrap_test
import time

@wrap_test(
    name='add element wise test',
    repeat=100,
    meta={
        'description': 'Test add element wise operation',
    },
    params={'test_size': [(1 << i) for i in range(10, 16)]}
)
def add_element_wise(test_size):
    t1 = Tensor.random_tensor([test_size])
    t2 = Tensor.random_tensor(t1.shape)

    add_out , stats = execute(Operation.ELEMENTWISE_ADD, [], [t1, t2])
    
    t_start = time.time()
    expected_add = t1.data + t2.data
    stats['cpu_based'] = time.time() - t_start
    
    
    error = absolute_or_relative_error(add_out.data, expected_add).mean()
    stats['error'] = error

    return stats

@wrap_test(
    name='sub element wise test',
    repeat=100,
    meta={
        'description': 'Test sub element wise operation',
    },
    params={'test_size': [(1 << i) for i in range(10, 16)]}
)
def sub_element_wise(test_size):
    t1 = Tensor.random_tensor([test_size])
    t2 = Tensor.random_tensor(t1.shape)

    sub_out , stats = execute(Operation.ELEMENTWISE_SUB, [], [t1, t2])
    
    t_start = time.time()
    expected_sub = t1.data - t2.data
    stats['cpu_based'] = time.time() - t_start
    
    error = absolute_or_relative_error(sub_out.data, expected_sub).mean()
    stats['error'] = error

    return stats

@wrap_test(
    name='mul element wise test',
    repeat=100,
    meta={
        'description': 'Test mul element wise operation',
    },
    params={'test_size': [(1 << i) for i in range(10, 16)]}
)
def mul_element_wise(test_size):
    t1 = Tensor.random_tensor([test_size])
    t2 = Tensor.random_tensor(t1.shape)

    mul_out , stats = execute(Operation.ELEMENTWISE_MUL, [], [t1, t2])

    t_start = time.time()
    expected_mul = t1.data * t2.data
    stats['cpu_based'] = time.time() - t_start

    error = absolute_or_relative_error(mul_out.data, expected_mul).mean()
    stats['error'] = error

    return stats

@wrap_test(
    name='div element wise test',
    repeat=100,
    meta={
        'description': 'Test div element wise operation',
    },
    params={'test_size': [(1 << i) for i in range(10, 16)]}
)
def div_element_wise(test_size):
    t1 = Tensor.random_tensor([test_size])
    t2 = Tensor.random_tensor(t1.shape)

    t2._data += 0.5 + 1e-4
    div_out , stats = execute(Operation.ELEMENTWISE_DIV, [], [t1, t2])

    t_start = time.time()
    expected_div = t1.data / t2.data
    stats['cpu_based'] = time.time() - t_start
    
    error = absolute_or_relative_error(div_out.data, expected_div).mean()
    stats['error'] = error

    return stats
