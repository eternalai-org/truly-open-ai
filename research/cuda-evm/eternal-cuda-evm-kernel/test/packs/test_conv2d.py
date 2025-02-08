from tensor import Tensor
from op import execute, Operation
import tensorflow as tf
import time
from utils import absolute_or_relative_error
from .test_registry import wrap_test

def check_conv2d_constraint(
    spatial_size, 
    channel_in, 
    ksize, 
    stride, 
    padding,
    channel_out = None 
):
    channel_out = channel_out or channel_in
    return all([
        all([x > 0 for x in [spatial_size, channel_in, channel_out, ksize, stride]]),
        ksize <= spatial_size,
        stride <= ksize,
        padding in ['valid', 'same'],
    ])

@wrap_test(
    name='conv2d correct random test', 
    repeat=5, 
    meta={
        'description': 'Test conv2d operation, execution time includes tensorflow operation on cpu',
    },
    params={
        'spatial_size': [(1 << i) for i in range(3, 5)],
        'channel_in': [(1 << i) for i in range(3, 6)],
        'channel_out': [(1 << i) for i in range(3, 6)],
        'ksize': [(1 << i) for i in range(1, 5)],
        'stride': [(1 << i) for i in range(1, 5)],
        'padding': ['valid', 'same']
    },
    checker=check_conv2d_constraint
)
def test_correct_conv2d(
    spatial_size, 
    channel_in, 
    channel_out, 
    ksize, 
    stride, 
    padding
):
    t = Tensor.random_tensor([spatial_size, spatial_size, channel_in])
    padding_i = 1 if padding == 'same' else 0
    
    random_kernel = Tensor.random_tensor((ksize, ksize, channel_in, channel_out))
    random_bias = Tensor.random_tensor((channel_out,))

    params = [stride, stride, padding_i]
    
    conv2d = tf.keras.layers.Conv2D(
        filters=channel_out, 
        kernel_size=(ksize, ksize), 
        strides=(stride, stride), 
        padding=padding
    )
    conv2d.build(t.data.reshape(1, *t.shape).shape)    
    conv2d.set_weights([random_kernel.data.reshape(random_kernel.shape), random_bias.data.reshape(random_bias.shape)])

    conv2d_out , stats = execute(Operation.CONV2D, params, [t, random_kernel, random_bias])

    t_start = time.time()
    expected_conv2d = conv2d(t.data.reshape(1, *t.shape)).numpy().flatten()
    stats['cpu_based'] = time.time() - t_start
    
    error = absolute_or_relative_error(conv2d_out.data, expected_conv2d).mean()
    stats['error'] = error
    
    return stats


@wrap_test(
    name='depthwise conv2d correct random test', 
    repeat=5, 
    meta={
        'description': 'Test deothwise conv2d operation, execution time includes tensorflow operation on cpu',
    },
    params={
        'spatial_size': [(1 << i) for i in range(3, 5)],
        'channel_in': [(1 << i) for i in range(3, 6)],
        'ksize': [(1 << i) for i in range(1, 5)],
        'stride': [(1 << i) for i in range(1, 5)],
        'padding': ['valid', 'same']
    },
    checker=check_conv2d_constraint
)
def test_deothwise_correct_conv2d(
    spatial_size, 
    channel_in, 
    ksize, 
    stride, 
    padding
):
    t = Tensor.random_tensor([spatial_size, spatial_size, channel_in])
    padding_i = 1 if padding == 'same' else 0
    
    random_kernel = Tensor.random_tensor((ksize, ksize, channel_in))
    random_bias = Tensor.zeros_tensor((channel_in,))

    params = [stride, stride, padding_i]
    
    depthwise_conv2d = tf.keras.layers.DepthwiseConv2D(
        kernel_size=(ksize, ksize), 
        strides=(stride, stride), 
        padding=padding
    )

    depthwise_conv2d.build(t.data.reshape(1, *t.shape).shape)    
    depthwise_conv2d.set_weights([random_kernel.data.reshape([*random_kernel.shape, 1]), random_bias.data.reshape(random_bias.shape)])

    conv2d_out , stats = execute(Operation.DEPTHWISE_CONV2D, params, [t, random_kernel, random_bias])

    t_start = time.time()
    expected_conv2d = depthwise_conv2d(t.data.reshape(1, *t.shape)).numpy().flatten()
    stats['cpu_based'] = time.time() - t_start
    
    error = absolute_or_relative_error(conv2d_out.data, expected_conv2d).mean()
    stats['error'] = error
    
    return stats