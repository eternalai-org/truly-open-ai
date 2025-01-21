// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { Float32x32 } from "./../Float32x32/Lib32x32.sol";
import "./Tensors.sol";
import { CUDA } from "./../libCudaTensor.sol";



library TensorMethods {
    enum CUDA_OPCODE {
        CONV2D, // 0
        MAXPOOLING2D, // 1
        AVGPOOLING2D, // 2

        // Matrix operations
        MATMUL, // 3

        // Elementwise operations
        ELEMENTWISE_ADD, // 4
        ELEMENTWISE_MUL, // 5
        ELEMENTWISE_SUB, // 6
        ELEMENTWISE_DIV, // 7

        // Transforms
        TRANSFORM_EXP, // 8
        TRANSFORM_SQRT, // 9

        // Normalizations
        BATCH_NORM, // 10
        LAYER_NORM, // 11 
        ZSCORE, // 12
        MIN_MAX_SCALE, // 13

        // merging operations
        CONCATENATE, // 14

        // Activations
        RELU, // 15
        TANH, // 16
        SIGMOID, // 17
        SOFTMAX, // 18
        LOGSOFTMAX, // 19
        SOFTMAX2D, // 20

        // Reductions
        REDUCTION_MAX, // 21
        REDUCTION_MIN, // 22
        REDUCTION_MEAN, // 23
        REDUCTION_SUM, // 24
        REDUCTION_ARGMAX, // 25
        REDUCTION_ARGMIN, // 26

        DROPOUT, // 27
        GLOBAL_AVGPOOLING2D, //28

        // batch 2 operations
        RESCALE, // 29
        CHANNEL_WISE_MEAN_REDUCTION, // 30    
        CHANNEL_WISE_SUM_REDUCTION, // 31
        DEPTHWISE_CONV2D // 31
    }

    // Assuming x is untouched
    function initStorageEmptyTensor(Tensors.Tensor storage w, uint64[] memory shape) internal {
        w.shapes = shape;
    }

    // Assuming x is untouched
    function initStorageZerosTensor(Tensors.Tensor storage w, uint64[] memory shape) internal {
        uint len = Tensors.getWeightCount(shape);
        uint256[] storage data = w.data;
        assembly {
            sstore(data.slot, len)
        }
        w.shapes = shape;
    }

    function appendStorageTensor(Tensors.Tensor storage w, uint[] calldata x, uint ptrLayer, uint ptr, uint idx) internal returns (uint, uint, uint) {
        uint n = Tensors.getWeightCount(w.shapes);
        unchecked {
            while (idx < x.length && ptr < n) {
                w.data[ptr] = x[idx];
                ptr++; idx++;
            }            
        }
        if (ptr == n) { ++ptrLayer; ptr = 0; }
        return (ptrLayer, ptr, idx);
    }

    function execute_operation(
        uint64 opcode, 
        int64[] memory params, 
        uint64[][] memory shapes, 
        uint256[][] memory tensors
    ) internal view returns (Tensors.Tensor memory)
    {   
        bytes memory data = abi.encode(opcode, params, shapes, tensors);
        bytes memory output = CUDA.processBytesArray(data);

        (uint256[] memory output_data, uint64[] memory output_shapes) = abi.decode(output, (uint256[], uint64[]));

        return toTensor(output_data, output_shapes);
    }

    function matmul(
        Tensors.Tensor memory a, 
        Tensors.Tensor memory b
    ) internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.MATMUL);

        uint64[][] memory shapes = new uint64[][](2);
        shapes[0] = a.shapes;
        shapes[1] = b.shapes;

        uint256[][] memory tensors = new uint256[][](2);
        tensors[0] = a.data;
        tensors[1] = b.data;

        Tensors.Tensor memory result = execute_operation(opcode, new int64[](0), shapes, tensors);
        return squeeze(result, 0);
    }

    function add(Tensors.Tensor memory a, Tensors.Tensor memory b) 
    internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.ELEMENTWISE_ADD);
        
        uint64[][] memory shapes = new uint64[][](2);
        shapes[0] = a.shapes;
        shapes[1] = b.shapes;

        uint256[][] memory tensors = new uint256[][](2);
        tensors[0] = a.data;
        tensors[1] = b.data;

        return execute_operation(opcode, new int64[](0), shapes, tensors);
    }

    function mul(Tensors.Tensor memory a, Tensors.Tensor memory b) 
    internal view returns (Tensors.Tensor memory) 
    {   
        uint64 opcode = uint64(CUDA_OPCODE.ELEMENTWISE_MUL);
        
        uint64[][] memory shapes = new uint64[][](2);
        shapes[0] = a.shapes;
        shapes[1] = b.shapes;

        uint256[][] memory tensors = new uint256[][](2);
        tensors[0] = a.data;
        tensors[1] = b.data;

        return execute_operation(opcode, new int64[](0), shapes, tensors);
    }

    function sub(Tensors.Tensor memory a, Tensors.Tensor memory b) 
    internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.ELEMENTWISE_SUB);
        
        uint64[][] memory shapes = new uint64[][](2);
        shapes[0] = a.shapes;
        shapes[1] = b.shapes;

        uint256[][] memory tensors = new uint256[][](2);
        tensors[0] = a.data;
        tensors[1] = b.data;

        return execute_operation(opcode, new int64[](0), shapes, tensors);
    }

    function div(Tensors.Tensor memory a, Tensors.Tensor memory b) 
    internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.ELEMENTWISE_DIV);
        
        uint64[][] memory shapes = new uint64[][](2);
        shapes[0] = a.shapes;
        shapes[1] = b.shapes;

        uint256[][] memory tensors = new uint256[][](2);
        tensors[0] = a.data;
        tensors[1] = b.data;

        return execute_operation(opcode, new int64[](0), shapes, tensors);
    }

    function exp(Tensors.Tensor memory a) 
    internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.TRANSFORM_EXP);

        uint64[][] memory shapes = new uint64[][](1);
        shapes[0] = a.shapes;

        uint256[][] memory tensors = new uint256[][](1);
        tensors[0] = a.data;

        return execute_operation(opcode, new int64[](0), shapes, tensors);
    }

    function sqrt(Tensors.Tensor memory a) 
    internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.TRANSFORM_SQRT);

        uint64[][] memory shapes = new uint64[][](1);
        shapes[0] = a.shapes;

        uint256[][] memory tensors = new uint256[][](1);
        tensors[0] = a.data;

        return execute_operation(opcode, new int64[](0), shapes, tensors);
    }

    function relu(Tensors.Tensor memory a, int64 negative_slope) 
    internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.RELU);

        int64[] memory params = new int64[](1);
        params[0] = negative_slope;

        uint64[][] memory shapes = new uint64[][](1);
        shapes[0] = a.shapes;

        uint256[][] memory tensors = new uint256[][](1);
        tensors[0] = a.data;

        return execute_operation(opcode, params, shapes, tensors);
    }

    function tanh(Tensors.Tensor memory a) 
    internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.TANH);

        uint64[][] memory shapes = new uint64[][](1);
        shapes[0] = a.shapes;

        uint256[][] memory tensors = new uint256[][](1);
        tensors[0] = a.data;

        return execute_operation(opcode, new int64[](0), shapes, tensors);
    }

    function sigmoid(Tensors.Tensor memory a) 
    internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.SIGMOID);

        uint64[][] memory shapes = new uint64[][](1);
        shapes[0] = a.shapes;

        uint256[][] memory tensors = new uint256[][](1);
        tensors[0] = a.data;

        return execute_operation(opcode, new int64[](0), shapes, tensors);
    }

    function softmax(Tensors.Tensor memory a, uint axis) 
    internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.SOFTMAX);

        int64[] memory params = new int64[](1);
        params[0] = int64(uint64(axis));

        uint64[][] memory shapes = new uint64[][](1);
        shapes[0] = a.shapes;

        uint256[][] memory tensors = new uint256[][](1);
        tensors[0] = a.data;

        return execute_operation(opcode, params, shapes, tensors);
    }

    // this is a wrapper
    function activation(Tensors.Tensor memory a, Tensors.ActivationFunc actv) 
    internal view returns (Tensors.Tensor memory) 
    {
        // leaky relu? negative slope?
        if (actv == Tensors.ActivationFunc.Linear) {
            return a;
        } else if (actv == Tensors.ActivationFunc.ReLU) {
            return relu(a, 0);
        } else if (actv == Tensors.ActivationFunc.Sigmoid) {
            return sigmoid(a);
        } else if (actv == Tensors.ActivationFunc.Tanh) {
            return tanh(a);
        } else if (actv == Tensors.ActivationFunc.Softmax) {
            return softmax(a, dims(a) - uint(1));
        } else {
            revert InvalidActivationFunction();
        }
    }

    function squeeze(Tensors.Tensor memory a, uint dim)
    internal pure returns (Tensors.Tensor memory) 
    {
        if (dims(a) == dim)
        {
            // remove all dims with size 1
            uint64 new_dims = 0;

            for (uint64 i = 0; i < dims(a); i++)
            {
                if (a.shapes[i] != 1)
                {
                    new_dims++;
                }
            }

            uint64[] memory new_shapes = new uint64[](new_dims);
            uint it = 0;

            for (uint64 i = 0; i < dims(a); i++)
            {
                if (a.shapes[i] != 1)
                {
                    new_shapes[it] = a.shapes[i];
                    it++;
                }
            }

            return toTensor(a.data, new_shapes);
        }

        else if (0 <= dim && dim < dims(a) && a.shapes[dim] == 1)
        {
            uint64[] memory new_shapes = new uint64[](dims(a) - 1);
            uint64 j = 0;

            for (uint64 i = 0; i < dims(a); i++)
            {
                if (i != dim)
                {
                    new_shapes[j] = a.shapes[i];
                    j++;
                }
            }

            return toTensor(a.data, new_shapes);
        }

        return toTensor(a.data, a.shapes);
    }

    function unsqueeze(Tensors.Tensor memory a, uint dim)
    internal pure returns (Tensors.Tensor memory) 
    {
        if (dim >= 0 && dim < dims(a))
        {
            uint64[] memory new_shapes = new uint64[](dims(a) + 1);
            uint64 j = 0;

            for (uint64 i = 0; i < dims(a); i++)
            {
                if (i != dim)
                {
                    new_shapes[j] = a.shapes[i];
                    j++;
                }
                else
                {
                    new_shapes[j] = 1;
                    new_shapes[j + 1] = a.shapes[i];
                    j += 2;
                }
            }

            return toTensor(a.data, new_shapes);
        }

        return toTensor(a.data, a.shapes);
    }

    function conv2d(
        Tensors.Tensor memory a, 
        Tensors.Tensor memory kernel, 
        Tensors.Tensor memory bias, 
        uint stride_y, 
        uint stride_x, 
        uint padding
    ) 
    internal view returns (Tensors.Tensor memory) 
    {
        int64[] memory params = new int64[](3);
        params[0] = int64(uint64(stride_y));
        params[1] = int64(uint64(stride_x));
        params[2] = int64(uint64(padding));

        uint64[][] memory shapes = new uint64[][](3);
        shapes[0] = a.shapes;
        shapes[1] = kernel.shapes;
        shapes[2] = bias.shapes;

        uint256[][] memory tensors = new uint256[][](3);
        tensors[0] = a.data;
        tensors[1] = kernel.data;
        tensors[2] = bias.data;

        uint64 opcode = uint64(CUDA_OPCODE.CONV2D);

        return execute_operation(opcode, params, shapes, tensors);   
    }

    function maxpooling2d(
        Tensors.Tensor memory a, 
        uint window_size_y,
        uint window_size_x,
        uint stride_y, 
        uint stride_x, 
        uint padding
    )
    internal view returns (Tensors.Tensor memory) 
    {
        int64[] memory params = new int64[](5);
        params[0] = int64(uint64(window_size_y));
        params[1] = int64(uint64(window_size_x));
        params[2] = int64(uint64(stride_y));
        params[3] = int64(uint64(stride_x));
        params[4] = int64(uint64(padding));

        uint64[][] memory shapes = new uint64[][](1);
        shapes[0] = a.shapes;

        uint256[][] memory tensors = new uint256[][](1);
        tensors[0] = a.data;

        uint64 opcode = uint64(CUDA_OPCODE.MAXPOOLING2D);

        return execute_operation(opcode, params, shapes, tensors);
    }

    function avgpooling2d(
        Tensors.Tensor memory a, 
        uint window_size_y,
        uint window_size_x,
        uint stride_y, 
        uint stride_x, 
        uint padding
    ) 
    internal view returns (Tensors.Tensor memory) 
    {
        int64[] memory params = new int64[](5);
        params[0] = int64(uint64(window_size_y));
        params[1] = int64(uint64(window_size_x));
        params[2] = int64(uint64(stride_y));
        params[3] = int64(uint64(stride_x));
        params[4] = int64(uint64(padding));


        uint64[][] memory shapes = new uint64[][](1);
        shapes[0] = a.shapes;

        uint256[][] memory tensors = new uint256[][](1);
        tensors[0] = a.data;

        uint64 opcode = uint64(CUDA_OPCODE.AVGPOOLING2D);

        return execute_operation(opcode, params, shapes, tensors);
    }

    function global_avgpooling2d(Tensors.Tensor memory a)
    internal view returns (Tensors.Tensor memory){

        uint64 opcode = uint64(CUDA_OPCODE.GLOBAL_AVGPOOLING2D);

        uint64[][] memory shapes = new uint64[][](1);
        shapes[0] = a.shapes;

        uint256[][] memory tensors = new uint256[][](1);
        tensors[0] = a.data;

        return execute_operation(opcode, new int64[](0), shapes, tensors);
    }

    function batchnorm(
        Tensors.Tensor memory a,
        uint epsilon,
        uint momentum,
        Tensors.Tensor memory gamma,
        Tensors.Tensor memory beta,
        Tensors.Tensor memory movingMean,
        Tensors.Tensor memory movingVariance
    )
    internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.BATCH_NORM);

        int64[] memory params = new int64[](2);
        params[0] = int64(uint64(epsilon));
        params[1] = int64(uint64(momentum));

        uint64[][] memory shapes = new uint64[][](5);
        shapes[0] = a.shapes;
        shapes[1] = gamma.shapes;
        shapes[2] = beta.shapes;
        shapes[3] = movingMean.shapes;
        shapes[4] = movingVariance.shapes;

        uint256[][] memory tensors = new uint256[][](5);
        tensors[0] = a.data;
        tensors[1] = gamma.data;
        tensors[2] = beta.data;
        tensors[3] = movingMean.data;
        tensors[4] = movingVariance.data;

        return execute_operation(opcode, params, shapes, tensors);
    }

    function sum(Tensors.Tensor memory a, uint axis)
    internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.REDUCTION_SUM);

        int64[] memory params = new int64[](1);
        params[0] = int64(uint64(axis));

        uint64[][] memory shapes = new uint64[][](1);
        shapes[0] = a.shapes;

        uint256[][] memory tensors = new uint256[][](1);
        tensors[0] = a.data;

        return execute_operation(opcode, params, shapes, tensors);
    }

    function mean(Tensors.Tensor memory a, uint axis)
    internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.REDUCTION_MEAN);

        int64[] memory params = new int64[](1);
        params[0] = int64(uint64((axis)));

        uint64[][] memory shapes = new uint64[][](1);
        shapes[0] = a.shapes;

        uint256[][] memory tensors = new uint256[][](1);
        tensors[0] = a.data;

        return execute_operation(opcode, params, shapes, tensors);
    }

    function argmax(Tensors.Tensor memory a, uint axis)
    internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.REDUCTION_ARGMAX);

        int64[] memory params = new int64[](1);
        params[0] = int64(uint64(axis));

        uint64[][] memory shapes = new uint64[][](1);
        shapes[0] = a.shapes;

        uint256[][] memory tensors = new uint256[][](1);
        tensors[0] = a.data;

        return execute_operation(opcode, params, shapes, tensors);
    }

    function argmin(Tensors.Tensor memory a, uint axis)
    internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.REDUCTION_ARGMIN);

        int64[] memory params = new int64[](1);
        params[0] = int64(uint64(axis));

        uint64[][] memory shapes = new uint64[][](1);
        shapes[0] = a.shapes;

        uint256[][] memory tensors = new uint256[][](1);
        tensors[0] = a.data;

        return execute_operation(opcode, params, shapes, tensors);
    }

    function max(Tensors.Tensor memory a, uint axis)
    internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.REDUCTION_MAX); 

        int64[] memory params = new int64[](1);
        params[0] = int64(uint64(axis));

        uint64[][] memory shapes = new uint64[][](1);
        shapes[0] = a.shapes;

        uint256[][] memory tensors = new uint256[][](1);
        tensors[0] = a.data;
        
        return execute_operation(opcode, params, shapes, tensors);
    }

    function min(Tensors.Tensor memory a, uint axis)
    internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.REDUCTION_MIN); 

        int64[] memory params = new int64[](1);
        params[0] = int64(uint64(axis));

        uint64[][] memory shapes = new uint64[][](1);
        shapes[0] = a.shapes;

        uint256[][] memory tensors = new uint256[][](1);
        tensors[0] = a.data;
        
        return execute_operation(opcode, params, shapes, tensors);
    }

    function depthwiseConv2D(
        Tensors.Tensor memory a, // h, w, c
        Tensors.Tensor memory kernel, // kh, kw, c (or kh, kw, c, 1)
        Tensors.Tensor memory bias, // c
        uint stride_y, uint stride_x, 
        uint padding
    )
    internal view returns (Tensors.Tensor memory) 
    {
        int64[] memory params = new int64[](3);
        params[0] = int64(uint64(stride_y));
        params[1] = int64(uint64(stride_x));
        params[2] = int64(uint64(padding));

        uint64[][] memory shapes = new uint64[][](3);
        shapes[0] = a.shapes;
        shapes[1] = kernel.shapes;
        shapes[2] = bias.shapes;

        uint256[][] memory tensors = new uint256[][](3);
        tensors[0] = a.data;
        tensors[1] = kernel.data;
        tensors[2] = bias.data;

        uint64 opcode = uint64(CUDA_OPCODE.DEPTHWISE_CONV2D);
        return execute_operation(opcode, params, shapes, tensors);
    }

    function toScalar(int64 value)
    internal pure returns (Tensors.Tensor memory) 
    {
        uint256[] memory data = new uint256[](1);
        data[0] = uint256(uint64(value)) << 192;
        
        uint64[] memory shapes = new uint64[](1);
        shapes[0] = 1;

        return toTensor(data, shapes);
    }

    function scalar(Tensors.Tensor memory a)   
    internal pure returns (int64)
    {
        if (isScalar(a))
        {
            return int64(uint64(a.data[0] >> 192));
        }
        else
        {
            revert NotAScalar();
        }
    }

    function dims(Tensors.Tensor memory a)
    internal pure returns (uint) 
    {
        return a.shapes.length;   
    }

    function toTensor(uint256[] memory data, uint64[] memory shapes)
    internal pure returns (Tensors.Tensor memory)
    {
        return Tensors.Tensor(data, shapes);
    }

    function isScalar(Tensors.Tensor memory a)
    internal pure returns (bool) 
    {
        return a.shapes.length == 1 && a.shapes[0] == 1;
    }

    function concat(Tensors.Tensor[] memory a, uint axis)
    internal view returns (Tensors.Tensor memory) 
    {
        uint64 opcode = uint64(CUDA_OPCODE.CONCATENATE); 

        int64[] memory params = new int64[](1);
        params[0] = int64(uint64(axis));

        uint64[][] memory shapes = new uint64[][](a.length);
        for (uint i=0; i<a.length; i++)
            shapes[i] = a[i].shapes;

        uint256[][] memory tensors = new uint256[][](a.length);
        for (uint i=0; i<a.length; i++)
            tensors[i] = a[i].data;
        
        return execute_operation(opcode, params, shapes, tensors);
    }
}
