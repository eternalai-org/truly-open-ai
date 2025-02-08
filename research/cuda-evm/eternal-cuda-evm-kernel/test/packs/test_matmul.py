import numpy as np
from concurrent.futures import ProcessPoolExecutor

import random
from tensor import Tensor
from utils import absolute_or_relative_error, log as wrap_log
from op import Operation, execute 
from .test_registry import wrap_test
import time 


@wrap_test(
    name='matmul test',
    repeat=10,
    meta={
        'description': 'Test matmul operation',
    },
    params={
        'h1': [(1 << i) for i in range(10)],
        'w1': [(1 << i) for i in range(10)],
        'w2': [(1 << i) for i in range(10)]
    }
)
def matul_test(h1, w1, w2): 

    h1, w1, w2 = random.randint(1, h1), random.randint(1, w1), random.randint(1, w2)
    t1 = Tensor.random_tensor([h1, w1])
    t2 = Tensor.random_tensor([w1, w2])

    t1_data = t1.data.reshape(t1.shape)
    t2_data = t2.data.reshape(t2.shape)
    
    out , stats = execute(Operation.MATMUL, [], [t1, t2])
    
    t_start = time.time()
    np_result = np.matmul(t1_data, t2_data)
    stats['cpu_based'] = time.time() - t_start
    
    error = absolute_or_relative_error(out.data, np_result.flatten()).mean()
    stats['error'] = error

    return stats
