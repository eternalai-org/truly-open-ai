import numpy as np
from tensor import Tensor
from utils  import absolute_or_relative_error
from op import Operation, execute
from .test_registry import wrap_test
import time

@wrap_test(
    name='sum reduction test',
    repeat=100,
    meta={
        'description': 'Test sum reduction operation',
    },
    params={
        'test_size': [(1 << i) for i in range(16)]
    }
)
def sum_reduction(test_size):
    t1 = Tensor.random_tensor([test_size])
    add_out , stats = execute(Operation.REDUCTION_SUM, [], [t1])
    
    t_start = time.time()
    expected_sum = np.sum(t1.data)
    stats['cpu_based'] = time.time() - t_start

    error = absolute_or_relative_error(add_out.data, expected_sum).mean()
    stats['error'] = error
    
    return stats


@wrap_test(
    name='mean reduction test',
    repeat=100,
    meta={
        'description': 'Test mean reduction operation',
    },
    params={
        'test_size': [(1 << i) for i in range(16)]
    }
)
def mean_reduction(test_size):
    t1 = Tensor.random_tensor([test_size])
    add_out , stats = execute(Operation.REDUCTION_MEAN, [], [t1])
    
    t_start = time.time()
    expected_mean = np.mean(t1.data)
    stats['cpu_based'] = time.time() - t_start
    
    error = absolute_or_relative_error(add_out.data, expected_mean).mean()
    stats['error'] = error
    
    return stats


@wrap_test(
    name='max reduction test',
    repeat=100,
    meta={
        'description': 'Test max reduction operation',
    },
    params={
        'test_size': [(1 << i) for i in range(16)]
    }
)
def max_reduction(test_size):
    t1 = Tensor.random_tensor([test_size])
    
    add_out , stats = execute(Operation.REDUCTION_MAX, [], [t1])
    
    t_start = time.time()
    expected_max = np.max(t1.data)
    stats['cpu_based'] = time.time() - t_start
    
    error = absolute_or_relative_error(add_out.data, expected_max).mean()
    stats['error'] = error
    
    return stats

@wrap_test(
    name='min reduction test',
    repeat=100,
    meta={
        'description': 'Test min reduction operation',
    },
    params={
        'test_size': [(1 << i) for i in range(16)]
    }
)
def min_reduction(test_size):
    t1 = Tensor.random_tensor([test_size])
    add_out , stats = execute(Operation.REDUCTION_MIN, [], [t1])
    
    t_start = time.time()
    expected_min = np.min(t1.data)
    stats['cpu_based'] = time.time() - t_start    

    error = absolute_or_relative_error(add_out.data, expected_min).mean()
    stats['error'] = error

    return stats
    