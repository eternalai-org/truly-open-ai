import os
os.environ['TF_CPP_MIN_LOG_LEVEL'] = '3'

from .import test_registry

from .test_abi import *
from .test_activations import *
from .test_conv2d import *
from .test_elementwise import *
from .test_matmul import *
from .test_merging import *
from .test_norm import *
from .test_pooling import *
from .test_reduction import *
from .test_transformation import *

from .test_registry import run_tests, list_names