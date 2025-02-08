import keras
import random
from tensor import Tensor
from op import Operation, execute
from utils import absolute_or_relative_error, log as wrap_log
from .test_registry import wrap_test
import time

def check_pooling2d_constraints(spatial_size, channel_in, window_size, stride, padding):
    return all([
        all([x > 0 for x in [window_size, stride, spatial_size, channel_in]]),
        window_size <= spatial_size,
        stride <= window_size,
        padding in ['valid', 'same'],
    ])

@wrap_test(
    name='maxpooling2d test',
    repeat=10,
    meta={
        'description': 'Test max pooling operation',
    },
    params={
        'spatial_size': [(1 << i) for i in range(4, 6)],
        'channel_in': [(1 << i) for i in range(3, 6)],
        'window_size': [(1 << i) for i in range(1, 6)],
        'stride': [(1 << i) for i in range(1, 6)],
        'padding': ['valid', 'same']
    },
    checker=check_pooling2d_constraints
)
def max_pooling_correct_check(spatial_size, channel_in, window_size, stride, padding):
    t = Tensor.random_tensor([spatial_size, spatial_size, channel_in])

    padding_i = 1 if padding == 'same' else 0
    maxpooling = keras.layers.MaxPooling2D(
        pool_size=(window_size, window_size), 
        strides=(stride, stride), 
        padding=padding
    )
    
    params = [window_size, window_size, stride, stride, padding_i]

    maxpooling_out , stats = execute(Operation.MAXPOOLING2D, params, t)
    
    t_start = time.time()
    expected_max_pooling = maxpooling(t.data.reshape(1, *t.shape)).numpy().flatten()
    stats['cpu_based'] = time.time() - t_start

    error = absolute_or_relative_error(maxpooling_out.data, expected_max_pooling).mean()
    stats['error'] = error
    
    return stats


@wrap_test(
    name='avgpooling2d test',
    repeat=10,
    meta={
        'description': 'Test average pooling operation',
    },
    params={
        'spatial_size': [(1 << i) for i in range(4, 6)],
        'channel_in': [(1 << i) for i in range(3, 6)],
        'window_size': [(1 << i) for i in range(1, 6)],
        'stride': [(1 << i) for i in range(1, 6)],
        'padding': ['valid', 'same']
    },
    checker=check_pooling2d_constraints
)
def avg_pooling_correct_check(spatial_size, channel_in, window_size, stride, padding):
    t = Tensor.random_tensor([spatial_size, spatial_size, channel_in])

    padding_i = 1 if padding == 'same' else 0
    maxpooling = keras.layers.AveragePooling2D(
        pool_size=(window_size, window_size), 
        strides=(stride, stride), 
        padding=padding
    )
    
    params = [window_size, window_size, stride, stride, padding_i]

    maxpooling_out , stats = execute(Operation.AVGPOOLING2D, params, t)
    
    t_start = time.time()
    expected_max_pooling = maxpooling(t.data.reshape(1, *t.shape)).numpy().flatten()
    stats['cpu_based'] = time.time() - t_start

    error = absolute_or_relative_error(maxpooling_out.data, expected_max_pooling).mean()
    stats['error'] = error
    
    return stats

@wrap_test(
    name='global avgpooling2d test',
    repeat=300,
    meta={
        'description': 'Test global average pooling operation',
    },
    params={
        'spatial_size': [(1 << i) for i in range(4, 7)],
        'channel_in': [(1 << i) for i in range(3, 7)]
    }
)
def global_avg_pooling_correct_check(spatial_size, channel_in):
    
    t = Tensor.random_tensor([spatial_size, spatial_size, channel_in])  
    global_avgpooling_2d = keras.layers.GlobalAveragePooling2D()
    global_avgpooling_2d.build(t.data.reshape(1, *t.shape).shape)
    
    output , stats = execute(Operation.GLOBAL_AVGPOOLING2D, [], t)

    t_start = time.time()
    expected_global_avg_pooling = global_avgpooling_2d(t.data.reshape(1, *t.shape)).numpy().flatten()
    stats['cpu_based'] = time.time() - t_start
    
    error = absolute_or_relative_error(output.data, expected_global_avg_pooling).mean()
    stats['error'] = error
    
    return stats
