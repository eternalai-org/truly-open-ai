import numpy as np
from concurrent.futures import ProcessPoolExecutor

from tensor import Tensor
from op import Operation, execute
import traceback
from utils import log as wrap_log, absolute_or_relative_error
from .test_registry import wrap_test
import time


@wrap_test(
    name='concatenate test',
    repeat=1000,
    meta={
        'description': 'Test concatenate operation',
    }
)
def concatenate_test(): 
    t1 = Tensor.random_tensor()
    t2 = Tensor.random_tensor(t1.shape)
    t3 = Tensor.random_tensor(t1.shape)
    
    t1_data, t2_data, t3_data = t1.data.reshape(t1.shape), t2.data.reshape(t2.shape), t3.data.reshape(t3.shape)
    axis = np.random.randint(0, len(t1.shape))
    
    params = [axis]
    
    out, stats = execute(Operation.CONCATENATE, params, [t1, t2, t3])
    
    t_start = time.time()
    expected = np.concatenate([t1_data, t2_data, t3_data], axis=axis)
    stats['cpu_based'] = time.time() - t_start
    
    error = np.allclose(out.data, expected.flatten())
    stats['error'] = error
    
    return stats
