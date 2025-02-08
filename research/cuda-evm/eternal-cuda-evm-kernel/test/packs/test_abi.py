from tensor import Tensor
from op import execute
from utils import compare_tensors, log as wrap_log, absolute_or_relative_error
from .test_registry import wrap_test

@wrap_test(
    name='abi random test', 
    repeat=1000, 
    meta={
        'description': 'Test ABI encode/decode with random tensor'
    },
    params={}
)
def run_abi_test_random_case(**_):
    sample = Tensor.random_tensor()
    tout, stats = execute(27, [], [sample])
    error = absolute_or_relative_error(tout.data, sample.data).mean()
    stats['error'] = error
    return stats


@wrap_test(
    name='abi zero test', 
    repeat=1000, 
    meta={
        'description': 'Test ABI encode/decode with zero tensor'
    },
    params={}
)
def run_abi_test_zero_case(**_):
    sample = Tensor.random_tensor()
    tout , stats = execute(27, [], [sample])
    error = absolute_or_relative_error(sample.data, tout.data).mean()
    stats['error'] = error
    return stats