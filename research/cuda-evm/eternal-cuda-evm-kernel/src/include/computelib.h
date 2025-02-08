#ifndef __COMPUTELIB_H__
#define __COMPUTELIB_H__

#include <stdint.h>
#include <tensor.h>
#include <helpers.cuh>
#include <operations.cuh>



struct operation_pack
{
    int64_t op;
    std::vector<int64_t> params;
    std::vector<TensorWrapper> tensors;  
};

enum opcode
{
    CONV2D = 0, // 0
    MAXPOOLING2D = 1, // 1
    AVGPOOLING2D = 2, // 2

    // Matrix operations
    MATMUL = 3, // 3

    // Elementwise operations
    ELEMENTWISE_ADD = 4, // 4
    ELEMENTWISE_MUL = 5, // 5
    ELEMENTWISE_SUB = 6, // 6
    ELEMENTWISE_DIV = 7, // 7

    // Transforms
    TRANSFORM_EXP = 8, // 8
    TRANSFORM_SQRT = 9, // 9

    // Normalizations
    BATCH_NORM = 10, // 10
    LAYER_NORM = 11, // 11 
    ZSCORE = 12, // 12
    MIN_MAX_SCALE = 13, // 13

    // merging operations
    CONCATENATE = 14, // 14

    // Activations
    RELU = 15, // 15
    TANH = 16, // 16
    SIGMOID = 17, // 17
    SOFTMAX = 18, // 18
    LOGSOFTMAX = 19, // 19
    SOFTMAX2D = 20, // 20

    // Reductions
    REDUCTION_MAX = 21, // 21
    REDUCTION_MIN = 22, // 22
    REDUCTION_MEAN = 23, // 23
    REDUCTION_SUM = 24, // 24
    REDUCTION_ARGMAX = 25, // 25
    REDUCTION_ARGMIN = 26, // 26

    // misc
    DROPOUT = 27, // 27
    GLOBAL_AVGPOOLING2D = 28, // 28

    // batch 2 operations
    RESCALE = 29, // 29
    CHANNEL_WISE_MEAN_REDUCTION = 30, // 30    
    CHANNEL_WISE_SUM_REDUCTION = 31, // 31
    DEPTHWISE_CONV2D = 32 // 31
};

// abi operations
uint8_t* abi_encode_tensor(const TensorWrapper& tensor, int32_t* length);
operation_pack abi_decode_op(const int64_t* inp, uint8_t *__error);

uint8_t* conv2d_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* maxpooling2d_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* avgpooling2d_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* matmul_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* elementwise_add_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* elementwise_mul_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* elementwise_sub_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* elementwise_div_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* transform_exp_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* transform_sqrt_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* batch_norm_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* layer_norm_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* zscore_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* min_max_scale_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* concatenate_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* relu_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* tanh_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* sigmoid_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* softmax_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* logsoftmax_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* softmax2d_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* reduction_max_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* reduction_min_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* reduction_mean_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* reduction_sum_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* reduction_argmax_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* reduction_argmin_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* dropout_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* globalavgpooling_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* rescale_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);
uint8_t* depthwise_conv2d_call(const operation_pack& pack, uint32_t* length_out, uint8_t* eerror);

uint8_t* _execute(
    uint8_t* payload_in, // bytes: opcode, params, shapes, tensors
    int32_t length_in, 
    int32_t* length_out,
    uint8_t* has_eerror
);

#if __cplusplus
extern "C" {
#endif

const uint8_t* cuda_execute_operation(
    uint8_t* payload_in, // bytes: opcode, params, shapes, tensors
    int32_t length_in, 
    int32_t* length_out,
    uint8_t* has_eerror
);

void deallocate_cpp_response(const uint8_t* payload);

#if __cplusplus
}
#endif

#endif // __COMPUTELIB_H__