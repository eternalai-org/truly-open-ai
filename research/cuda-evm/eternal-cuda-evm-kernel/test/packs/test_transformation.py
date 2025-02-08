import numpy as np
from tensor import Tensor
from utils  import absolute_or_relative_error
from op import Operation, execute
from .test_registry import wrap_test
import time

@wrap_test(
    name='array sqrt test',
    repeat=100,
    meta={
        'description': 'Test array sqrt transformation operation',
    },
    params={
        'test_size': [(1 << i) for i in range(16)]
    }
)
def array_sqrt(test_size):
    t1 = Tensor.random_tensor([test_size])
    t1._data += 0.5
    add_out , stats = execute(Operation.TRANSFORM_SQRT, [], [t1])

    t_start = time.time()
    expected_sqrt = np.sqrt(t1.data)
    stats['cpu_based'] = time.time() - t_start
    
    error = absolute_or_relative_error(add_out.data, expected_sqrt).mean()
    stats['error'] = error
    
    return stats


@wrap_test(
    name='array exp test',
    repeat=100,
    meta={
        'description': 'Test array exp transformation operation',
        'accepted_error': 1e-4
    },
    params={
        'test_size': [(1 << i) for i in range(16)]
    }
)
def array_exp(test_size):
    t1 = Tensor.random_tensor([test_size])
    add_out , stats = execute(Operation.TRANSFORM_EXP, [], [t1])
    
    t_start = time.time()
    expected_exp = np.exp(t1.data)
    stats['cpu_based'] = time.time() - t_start
    
    error = absolute_or_relative_error(add_out.data, expected_exp).mean()
    stats['error'] = error
    
    return stats


@wrap_test(
    name='array max/min scale test',
    repeat=100,
    meta={
        'description': 'Test array max/min scale transformation operation',
        'accepted_error': 1e-4
    },
    params={
        'test_size': [(1 << i) for i in range(16)]
    }
)
def array_max_min_scale(test_size):
    t1 = Tensor.random_tensor([test_size])
    _min, _max = t1.data.min(), t1.data.max()
    add_out , stats = execute(Operation.MIN_MAX_SCALE, [], [t1])
    
    t_start = time.time()
    expected_max_min_scale = (t1.data - _min) / (_max - _min)
    stats['cpu_based'] = time.time() - t_start

    error = absolute_or_relative_error(add_out.data, expected_max_min_scale).mean()
    stats['error'] = error
    
    return stats


@wrap_test(
    name='array zscore test',
    repeat=100,
    meta={
        'description': 'Test array zscore transformation operation',
    },
    params={
        'test_size': [(1 << i) for i in range(16)]
    }
)
def array_zscore(test_size):
    t1 = Tensor.random_tensor([test_size])
    _mean, _std = t1.data.mean(), t1.data.std()

    eps = 1e-4
    add_out , stats = execute(Operation.ZSCORE, [int(eps * (1 << 32))], [t1])
    
    t_start = time.time()
    expected_zscore = (t1.data - _mean) / (_std + eps)
    stats['cpu_based'] = time.time() - t_start
    
    error = absolute_or_relative_error(add_out.data, expected_zscore).mean()
    stats['error'] = error
    
    return stats