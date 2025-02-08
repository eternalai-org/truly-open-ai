from tensor import Tensor
from op import Operation, execute
import random
from keras import layers
from utils import absolute_or_relative_error, log as wrap_log
import time
from .test_registry import wrap_test

@wrap_test(
    name='batch norm test',
    repeat=100,
    meta={
        'description': 'Test batch norm operation',
    },
    params={
        'spatial_size': [(1 << i) for i in range(3, 6)],
        'channel_in': [(1 << i) for i in range(3, 9)]
    }
)
def batch_norm_test(spatial_size, channel_in): 
    t1 = Tensor.random_tensor([spatial_size, spatial_size, channel_in])
    t1_data = t1.data.reshape(t1.shape)

    ma, mv, beta, gama = Tensor.random_tensor([channel_in]), Tensor.random_tensor([channel_in]), \
        Tensor.random_tensor([channel_in]), Tensor.random_tensor([channel_in])
        
    mv._data += 0.5

    batch_norm_layer = layers.BatchNormalization()
    batch_norm_layer.build(t1_data.shape)
    batch_norm_layer.set_weights([gama.data, beta.data, ma.data, mv.data])

    actual , stats = execute(Operation.BATCH_NORM, [int(1e-4 * 2 ** 32)], [t1, gama, beta, ma, mv])

    t_start = time.time()
    expected = batch_norm_layer(t1_data).numpy().flatten()
    stats['cpu_based'] = time.time() - t_start
    
    error = absolute_or_relative_error(actual.data, expected).mean()
    stats['error'] = error

    return stats