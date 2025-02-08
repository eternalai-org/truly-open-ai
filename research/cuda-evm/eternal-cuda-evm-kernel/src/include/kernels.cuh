#ifndef __CUDA_KERNELS_CUH__
#define __CUDA_KERNELS_CUH__

__global__ void arraySum_kernel(long long* A, long long* sum, int n);
__global__ void arrayExp_kernel(long long* A, long long* B, int n);
__global__ void softmaxImplFixedLongLong(long long *expA, long long* B, int n, long long sumExp);
__global__ void reluImplFixedLongLong(long long *A, long long* B, int m);
__global__ void sigmoidImplFixedLongLong(long long *A, long long* B, int n);
__global__ void tanhImplFixedLongLong(long long *A, long long* B, int n);
__global__ void vecmulFixedLongLong(long long *A, long long *B, long long *C, int m, int n, int k);
__global__ void vecmulInt(int *A, int *B, int *C, int m, int n, int k);
__global__ void vecmulLong(long *A, long *B, long *C, int m, int n, int k);
__global__ void vecmulFloat(float *A, float *B, float *C, int m, int n, int k);
__global__ void vecmulDouble(double *A, double *B, double *C, int m, int n, int k);
__global__ void normalizeFixedLongLong_kernel(
    long long *X, // input of shape (m, n, c) 
    long long *Y, // output of shape (m, n, c)
    long long *ma, // moving average(c)
    long long *mv, // moving variance (c)
    long long *gamma, // scale of shape (c)
    long long *beta, // offset of shape (c)
    long long epsilon, // epsilon 
    int h, int w, int c  // m, n, c 
);
__global__ void maxPoolingImplFixedLongLong_kernel(
    long long* inp, long long* out,
    int in_h, int in_w, int in_channel,
    int pool_size, int stride_h, int stride_w,
    int out_h, int out_w, 
    int padded_top, int padded_bottom, 
    int padded_left, int padded_right 
);
__global__ void avgPoolingImplFixedLongLong_kernel(
    long long* inp, long long* out,
    int in_h, int in_w, int in_channel,
    int pool_size, int stride_h, int stride_w,
    int out_h, int out_w,
    int padded_top, int padded_bottom, 
    int padded_left, int padded_right 
);
__global__ void sumReduction_kernel(long long* d_gpu, long long* blockOutput, int n);
__global__ void sumReductionV2_kernel(long long* d_gpu, long long* blockOutput, int n, int c);
__global__ void maxReduction_kernel(long long* d_gpu, long long* blockOutput, int n);
__global__ void minReduction_kernel(long long* d_gpu, long long* blockOutput, int n);
__global__ void minMaxScale_kernel(long long* d_gpu, long long* out, long long min, long long max, int n);
__global__ void zScore_kernel(long long* d_gpu, long long* out, long long mean, long long std, int n);
__global__ void maxPoolingImplFixedLongLong_kernel(
    long long* inp, long long* out,
    int in_h, int in_w, int in_channel,
    int pool_size, int stride_h, int stride_w,
    int out_h, int out_w, 
    int padded_top, int padded_bottom, 
    int padded_left, int padded_right 
);
__global__ void avgPoolingImplFixedLongLong_kernel(
    long long* inp, long long* out,
    int in_h, int in_w, int in_channel,
    int pool_size, int stride_h, int stride_w,
    int out_h, int out_w,
    int padded_top, int padded_bottom, 
    int padded_left, int padded_right 
);
__global__ void mat_add_fixed_longlong(long long *A, long long *B, long long *C, int n);
__global__ void mat_sub_fixed_longlong(long long *A, long long *B, long long *C, int n);
__global__ void mat_mul_fixed_longlong(long long *A, long long *B, long long *C, int n);
__global__ void mat_div_fixed_longlong(long long *A, long long *B, long long *C, int n);
__global__ void mat_sqrt_fixed_longlong(long long *A, long long *B, int n);
__global__ void mat_exp_fixed_longlong(long long *A, long long *B, int n);
__global__ void conv2dImplFixedLongLong_kernel(
    long long* inp, long long* kernel, long long* bias, long long* out, // data io
    int kernel_size, int in_channel, int out_channel, // kernel properties
    int in_h, int in_w, int out_h, int out_w, // spatial size of inp,
    int padding, int stride_h, int stride_w // padding mode, one of 'valid': 0 or 'same': 1
);

__global__ void mat_add_single_fixed_longlong(long long *A, long long *B, long long e, int n);
__global__ void mat_sub_single_fixed_longlong(long long *A, long long *B, long long e, int n);
__global__ void mat_mul_single_fixed_longlong(long long *A, long long *B, long long e, int n);
__global__ void mat_div_single_fixed_longlong(long long *A, long long *B, long long e, int n);

__global__ void mat_pow2_single_fixed_longlong(long long *A, long long *B, int n);

__global__ void depthwise_conv2d_kernel(
    long long* inp, long long* kernel, long long* bias, long long* out, // data io
    int in_h, int in_w, int in_channel, 
    int kernel_size_h, int kernel_size_w, // kernel properties
    int out_h, int out_w, // spatial size of inp,
    int padding, int stride_h, int stride_w // padding mode, one of 'valid': 0 or 'same': 1
);

#endif // __CUDA_KERNELS_CUH__