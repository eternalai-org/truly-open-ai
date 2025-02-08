#include <operations.cuh>

uint8_t matAddLongLong(long long *A, long long *B, long long *C, int m, int n)
{
    uint8_t isOk = true;
    __matAddLongLong(A, B, C, m, n, &isOk);
    return isOk;
}

uint8_t matSubLongLong(long long *A, long long *B, long long *C, int m, int n)
{
    uint8_t isOk = true;
    __matSubLongLong(A, B, C, m, n, &isOk);
    return isOk;
}

uint8_t matMulLongLong(long long *A, long long *B, long long *C, int m, int n)
{
    uint8_t isOk = true;
    __matMulLongLong(A, B, C, m, n, &isOk);
    return isOk;
}

uint8_t matDivLongLong(long long *A, long long *B, long long *C, int m, int n )
{
    uint8_t isOk = true;
    __matDivLongLong(A, B, C, m, n, &isOk);
    return isOk;
}

uint8_t matSqrtLongLong(long long *A, long long *B, int m, int n)
{
    uint8_t isOk = true;
    __matSqrtLongLong(A, B, m, n, &isOk);
    return isOk;
}

uint8_t conv2dFixedLongLong(long long* inp, long long* kernel, long long* bias, long long* out, int kernel_size, int in_channel, int out_channel, int h, int w, int padding, int stride_h, int stride_w)
{
    uint8_t isOk = true;
    __conv2dFixedLongLong(inp, kernel, bias, out, kernel_size, in_channel, out_channel, h, w, padding, stride_h, stride_w, &isOk);
    return isOk;
}


uint8_t avgPoolingFixedLongLong(long long* inp, long long* out, int h, int w, int in_channel, int pool_size, int stride_h, int stride_w, int padding)
{
    uint8_t isOk = true;
    __avgPoolingFixedLongLong(inp, out, h, w, in_channel, pool_size, stride_h, stride_w, padding, &isOk);
    return isOk;
}

uint8_t maxPoolingFixedLongLong(long long* inp, long long* out, int h, int w, int in_channel, int pool_size, int stride_h, int stride_w, int padding
)
{
    uint8_t isOk = true;
    __maxPoolingFixedLongLong(inp, out, h, w, in_channel, pool_size, stride_h, stride_w, padding, &isOk);
    return isOk;
}

uint8_t globalAvgPoolingFixedLongLong(long long* inp, long long* out, int h, int w, int in_channel)
{
    uint8_t isOk = true;
    __globalAvgPoolingFixedLongLong(inp, out, h, w, in_channel, &isOk);
    return isOk;
}


uint8_t sumReduction(long long* inp, int n, long long* res)
{
    uint8_t isOk = true;
    *res = __sumReduction(inp, n, &isOk);
    return isOk;
}

uint8_t maxReduction(long long* inp, int n, long long* res)
{
    uint8_t isOk = true;
    *res = __maxReduction(inp, n, &isOk);
    return isOk;
}

uint8_t minReduction(long long* inp, int n, long long* res)
{
    uint8_t isOk = true;
    *res = __minReduction(inp, n, &isOk);
    return isOk;
}

uint8_t meanReduction(long long* inp, int n, long long* res)
{
    uint8_t isOk = true;
    *res = __meanReduction(inp, n, &isOk);
    return isOk;
}

uint8_t stdReduction(long long* inp, int n, long long* res)
{
    uint8_t isOk = true;
    *res = __stdReduction(inp, n, &isOk);
    return isOk;
}

uint8_t maxMinScale(long long* inp, long long* out, int n)
{
    uint8_t isOk = true;
    __maxMinScale(inp, out, n, &isOk);
    return isOk;
}

uint8_t zScore(long long* inp, long long* out, long long eps, int n)
{
    uint8_t isOk = true;
    __zScore(inp, out, eps, n, &isOk);
    return isOk;
}

uint8_t concatenate(long long* inp, long long* out, long long* shapes, long long axis, long long ndims, long long n)
{
    uint8_t isOk = true;
    __concatenate(inp, out, shapes, axis, ndims, n, &isOk);
    return isOk;
}

uint8_t softmaxFixedLongLong(long long *A, long long* B, int m)
{
    uint8_t isOk = true;
    __softmaxFixedLongLong(A, B, m, &isOk);
    return isOk;
}

uint8_t sigmoidFixedLongLong(long long *A, long long* B, int m)
{
    uint8_t isOk = true;
    __sigmoidFixedLongLong(A, B, m, &isOk);
    return isOk;
}

uint8_t tanhFixedLongLong(long long *A, long long *B, int m)
{
    uint8_t isOk = true;
    __tanhFixedLongLong(A, B, m, &isOk);
    return isOk;
}

uint8_t reluFixedLongLong(long long *A, long long *B, int m)
{
    uint8_t isOk = true;
    __reluFixedLongLong(A, B, m, &isOk);
    return isOk;
}

uint8_t relu3DFixedLongLong(long long *A, long long *B, int h, int w, int c)
{
    uint8_t isOk = true;
    __relu3DFixedLongLong(A, B, h, w, c, &isOk);
    return isOk;
}

uint8_t sigmoid3DFixedLongLong(long long *A, long long *B, int h, int w, int c)
{
    uint8_t isOk = true;
    __sigmoid3DFixedLongLong(A, B, h, w, c, &isOk);
    return isOk;
}

uint8_t tanh3DFixedLongLong(long long *A, long long *B, int h, int w, int c)
{
    uint8_t isOk = true;
    __tanh3DFixedLongLong(A, B, h, w, c, &isOk);
    return isOk;
}

uint8_t softmax2DFixedLongLong(long long* A, long long* B, int h, int w, int c)
{
    uint8_t isOk = true;
    __softmax2DFixedLongLong(A, B, h, w, c, &isOk);
    return isOk;
}

uint8_t layerNormalizeFixedLongLong(long long *X, long long *Y, long long *ma, long long *mv, long long *gamma, long long *beta, long long epsilon, int h, int w, int c)
{
    uint8_t isOk = true;
    __layerNormalizeFixedLongLong(X, Y, ma, mv, gamma, beta, epsilon, h, w, c, &isOk);
    return isOk;
}

uint8_t batchNormalizeFixedLongLong(long long *X, long long *Y, long long *ma, long long *mv, long long *gamma, long long *beta, long long epsilon, int h, int w, int c)
{
    uint8_t isOk = true;
    __batchNormalizeFixedLongLong(X, Y, ma, mv, gamma, beta, epsilon, h, w, c, &isOk);
    return isOk;
}

uint8_t maxmulFixedLongLong(long long *A, long long *B, long long *C, int m, int n, int k)
{
    uint8_t isOk = true;
    __maxmulFixedLongLong(A, B, C, m, n, k, &isOk);
    return isOk;
}

uint8_t maxmulFloat(float *A, float *B, float *C, int m, int n, int k)
{
    uint8_t isOk = true;
    __maxmulFloat(A, B, C, m, n, k, &isOk);
    return isOk;
}

uint8_t maxmulLong(long *A, long *B, long *C, long m, long n, long k)
{
    uint8_t isOk = true;
    __maxmulLong(A, B, C, m, n, k, &isOk);
    return isOk;
}

uint8_t maxmulInt(int *A, int *B, int *C, int m, int n, int k)
{
    uint8_t isOk = true;
    __maxmulInt(A, B, C, m, n, k, &isOk);
    return isOk;
}

uint8_t maxmulDouble(double *A, double *B, double *C, int m, int n, int k)
{
    uint8_t isOk = true;
    __maxmulDouble(A, B, C, m, n, k, &isOk);
    return isOk;
}