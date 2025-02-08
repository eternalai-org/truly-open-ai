import numpy as np
from utils import compress_uint256, unpack_uint256, chunked
import random

class Tensor(object):
    def __init__(self, data, shape) -> None:
        self._data = data
        self._shape = shape

        if not isinstance(data, np.ndarray):
            self._data = np.array(data).astype(np.float64)

        self._data = self._data.flatten()

    @property
    def data(self): return self._data

    @property
    def shape(self): return self._shape

    def compress(self):
        return [compress_uint256(*i) for i in chunked((self._data * (2 ** 32)).astype(np.int64).tolist(), 4)]
    
    @staticmethod
    def uncompress(b: list, shape = None) -> 'Tensor':
        data = [unpack_uint256(i) for i in b]
        data = np.array(data, dtype=np.float64)

        if shape is not None:
            data = data.flatten()[:np.prod(shape)] # remove padding
            data = data.reshape(shape)
        else:
            data = data.flatten()

        return Tensor(data / (2 ** 32), shape)

    def __str__(self) -> str:
        return f'Tensor(shape={self._shape}, data={self._data})'
    
    @staticmethod
    def random_tensor(s=None) -> 'Tensor':
        if s is not None:
            shapes = s
        else:
            shapes = [random.randint(1, 10) for _ in range(4)]

        flatten = np.prod(shapes)
        tensor = np.random.rand(flatten).astype(np.float64) - 0.5
        return Tensor(tensor, shapes)
    
    @staticmethod
    def zeros_tensor(s=None) -> 'Tensor':
        if s is not None:
            shapes = s
        else:
            shapes = [random.randint(1, 10) for _ in range(4)]

        flatten = np.prod(shapes)
        tensor = np.zeros(flatten).astype(np.float64)
        return Tensor(tensor, shapes)