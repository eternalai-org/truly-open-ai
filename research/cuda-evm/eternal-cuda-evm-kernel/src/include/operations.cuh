#ifndef __OPERATIONS_H__
#define __OPERATIONS_H__

#include <stdint.h>
#include <helpers.cuh>
#include <fixedlonglong32x32.cuh>
#include <kernels.cuh>

void __channelWiseSumReduction_impl(long long* d_gpu, long long* d_out, int n, int c, uint8_t* error);
long long __sumReduction_impl(long long* d_gpu, int n, uint8_t* error);
long long __maxReduction_impl(long long* d_gpu, int n, uint8_t* error);
long long __minReduction_impl(long long* d_gpu, int n, uint8_t* error);
long long __meanReduction_impl(long long* gpu_inp, int n, uint8_t* error);
long long __stdReduction_impl(long long* d_gpu, int n, uint8_t* error);

// matrix mutiplications
void __maxmulFixedLongLong(long long *A, long long *B, long long *C, int m, int n, int k, uint8_t* error);
void __maxmulFloat(float *A, float *B, float *C, int m, int n, int k, uint8_t* error);
void __maxmulLong(long *A, long *B, long *C, long m, long n, long k, uint8_t* error);
void __maxmulInt(int *A, int *B, int *C, int m, int n, int k, uint8_t* error);
void __maxmulDouble(double *A, double *B, double *C, int m, int n, int k, uint8_t* error);

// element wise operations
/*
    A: input matrix
    B: input matrix
    C: output matrix
    m: number of rows
    n: number of columns
    error: error flag
*/
void __matAddLongLong(long long *A, long long *B, long long *C, int m, int n, uint8_t* error);
void __matSubLongLong(long long *A, long long *B, long long *C, int m, int n, uint8_t* error);
void __matMulLongLong(long long *A, long long *B, long long *C, int m, int n, uint8_t* error);
void __matDivLongLong(long long *A, long long *B, long long *C, int m, int n , uint8_t* error);
void __matSqrtLongLong(long long *A, long long *B, int m, int n, uint8_t* error);
void __matExpLongLong(long long *A, long long *B, int m, int n, uint8_t* error);

// conv2d operations
void __conv2dFixedLongLong(long long* inp, long long* kernel, long long* bias, long long* out, int kernel_size, int in_channel, int out_channel, int h, int w, int padding, int stride_h, int stride_w, uint8_t* error);

// pooling operations
void __avgPoolingFixedLongLong(long long* inp, long long* out, int h, int w, int in_channel, int pool_size, int stride_h, int stride_w, int padding, uint8_t* error);
void __maxPoolingFixedLongLong(long long* inp, long long* out, int h, int w, int in_channel, int pool_size, int stride_h, int stride_w, int padding, uint8_t* error);
void __globalAvgPoolingFixedLongLong(long long* inp, long long* out, int h, int w, int in_channel, uint8_t* error);


// reduction
long long __sumReduction(long long* inp, int n, uint8_t* error);
long long __maxReduction(long long* inp, int n, uint8_t* error);
long long __minReduction(long long* inp, int n, uint8_t* error);
long long __meanReduction(long long* inp, int n, uint8_t* error);
long long __stdReduction(long long* inp, int n, uint8_t* error);
void __maxMinScale(long long* inp, long long* out, int n, uint8_t* error);
void __zScore(long long* inp, long long* out, long long eps, int n, uint8_t* error);
void __channelWiseSumReduction(long long* inp, long long* out, int n, int c, uint8_t* error);

// merging
void __concatenate(long long* inp, long long* out, long long* shapes, long long axis, long long ndims, long long n, uint8_t* error);
void __concatenate_dummy(long long** inp, long long* out, long long** shapes, long long axis, long long ndims, long long n, uint8_t* error);
uint8_t estimateConcatenate_dummy(long long** shapes, long long axis, long long ndims, long long n, long long* out);


// activations
void __softmaxFixedLongLong(long long *A, long long* B, int m, uint8_t* error);
void __sigmoidFixedLongLong(long long *A, long long* B, int m, uint8_t* error);
void __tanhFixedLongLong(long long *A, long long *B, int m, uint8_t* error);
void __reluFixedLongLong(long long *A, long long *B, int m, uint8_t* error);
void __relu3DFixedLongLong(long long *A, long long *B, int h, int w, int c, uint8_t* error);
void __sigmoid3DFixedLongLong(long long *A, long long *B, int h, int w, int c, uint8_t* error);
void __tanh3DFixedLongLong(long long *A, long long *B, int h, int w, int c, uint8_t* error);
void __softmax2DFixedLongLong(long long* A, long long* B, int h, int w, int c, uint8_t* error);


// normalizations
void __layerNormalizeFixedLongLong(long long *X, long long *Y, long long *ma, long long *mv, long long *gamma, long long *beta, long long epsilon, int h, int w, int c, uint8_t* error);
void __batchNormalizeFixedLongLong(
    long long *X, 
    long long *Y, 
    long long *gamma, 
    long long *beta, 
    long long *ma, 
    long long *mv, 
    long long epsilon, 
    int h, int w, int c, 
    uint8_t* error
);

void __depthwiseConv2dFixedLongLong(
    long long* inp, long long* kernel, long long* bias, long long* out, // data io
    int h, int w, int in_channel, // spatial size of inp,
    int kernel_size_h, int kernel_size_w,  // kernel properties
    int padding, 
    int stride_h, int stride_w, // padding: same(0) or valid(1)
    uint8_t*
);

#if __cplusplus
extern "C" {
#endif

#define OK 1
#define ERROR 0

// element wise operations
uint8_t matAddLongLong(long long *A, long long *B, long long *C, int m, int n);
uint8_t matSubLongLong(long long *A, long long *B, long long *C, int m, int n);
uint8_t matMulLongLong(long long *A, long long *B, long long *C, int m, int n);
uint8_t matDivLongLong(long long *A, long long *B, long long *C, int m, int n);
uint8_t matSqrtLongLong(long long *A, long long *B, int m, int n);

// conv2d operations
uint8_t conv2dFixedLongLong(long long* inp, long long* kernel, long long* bias, long long* out, int kernel_size, int in_channel, int out_channel, int h, int w, int padding, int stride_h, int stride_w);

// @deprecated
uint8_t estimateConvOutputSize(int kernel_size, int in_channel, int out_channel, int h, int w, int padding, int stride_h, int stride_w, int* out_h, int* out_w);

// pooling operations
uint8_t avgPoolingFixedLongLong(long long* inp, long long* out, int h, int w, int in_channel, int pool_size, int stride_h, int stride_w, int padding);
uint8_t maxPoolingFixedLongLong(long long* inp, long long* out, int h, int w, int in_channel, int pool_size, int stride_h, int stride_w, int padding);

// @deprecated
uint8_t estimatePoolingOutputSize(int h, int w, int in_channel, int pool_size, int padding, int stride_h, int stride_w, int* out_h, int* out_w);

uint8_t globalAvgPoolingFixedLongLong(long long* inp, long long* out, int h, int w, int in_channel);

uint8_t sumReduction(long long* inp, int n, long long* res);
uint8_t maxReduction(long long* inp, int n, long long* res);
uint8_t minReduction(long long* inp, int n, long long* res);
uint8_t meanReduction(long long* inp, int n, long long* res);
uint8_t stdReduction(long long* inp, int n, long long* res);


uint8_t maxMinScale(long long* inp, long long* out, int n);
uint8_t zScore(long long* inp, long long* out, long long eps, int n);


// merging 
uint8_t concatenate(long long* inp, long long* out, long long* shapes, long long axis, long long ndims, long long n);

// @deprecated
uint8_t estimateConcatenate(long long* shapes, long long axis, long long ndims, long long n, long long* out);


// activations
uint8_t softmaxFixedLongLong(long long *A, long long* B, int m);
uint8_t sigmoidFixedLongLong(long long *A, long long* B, int m);
uint8_t tanhFixedLongLong(long long *A, long long *B, int m);
uint8_t reluFixedLongLong(long long *A, long long *B, int m);
uint8_t relu3DFixedLongLong(long long *A, long long *B, int h, int w, int c);
uint8_t sigmoid3DFixedLongLong(long long *A, long long *B, int h, int w, int c);
uint8_t tanh3DFixedLongLong(long long *A, long long *B, int h, int w, int c);
uint8_t softmax2DFixedLongLong(long long* A, long long* B, int h, int w, int c);

// normalizations
uint8_t layerNormalizeFixedLongLong(long long *X, long long *Y, long long *ma, long long *mv, long long *gamma, long long *beta, long long epsilon, int h, int w, int c);
uint8_t batchNormalizeFixedLongLong(long long *X, long long *Y, long long *ma, long long *mv, long long *gamma, long long *beta, long long epsilon, int h, int w, int c);


// matrix mutiplications
uint8_t maxmulFixedLongLong(long long *A, long long *B, long long *C, int m, int n, int k);
uint8_t maxmulFloat(float *A, float *B, float *C, int m, int n, int k);
uint8_t maxmulLong(long *A, long *B, long *C, long m, long n, long k);
uint8_t maxmulInt(int *A, int *B, int *C, int m, int n, int k);
uint8_t maxmulDouble(double *A, double *B, double *C, int m, int n, int k);

#if __cplusplus
}
#endif

#endif