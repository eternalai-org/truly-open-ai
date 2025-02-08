#include <operations.cuh>

void __matSqrtLongLong(long long *A, long long *B, int m, int n, uint8_t* error) {
    // Allocate device memory:
    long long *gpu;
    const int N = m * n;

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu, sizeof(long long) * N * 2)))
    {
        cudaFree(gpu);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu, A, sizeof(long long) * N, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu);
        return;
    }

    const int BLOCK_SIZE = 256;
    const int GRID_SIZE = (N + BLOCK_SIZE - 1) / BLOCK_SIZE;

    mat_sqrt_fixed_longlong<<<GRID_SIZE, BLOCK_SIZE>>>(gpu, gpu + N, N);

    if (*error = cuda_fmt_error(cudaMemcpy(B, gpu + N, sizeof(long long) * N, cudaMemcpyDeviceToHost)))
    {
        cudaFree(gpu);
        return;
    }
    cudaFree(gpu);
}


void __matExpLongLong(long long *A, long long *B, int m, int n, uint8_t* error) {
    // Allocate device memory:
    long long *gpu;
    const int N = m * n;

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu, sizeof(long long) * N * 2)))
    {
        cudaFree(gpu);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu, A, sizeof(long long) * N, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu);
        return;
    }

    const int BLOCK_SIZE = 256;
    const int GRID_SIZE = (N + BLOCK_SIZE - 1) / BLOCK_SIZE;

    mat_exp_fixed_longlong<<<GRID_SIZE, BLOCK_SIZE>>>(gpu, gpu + N, N);

    if (*error = cuda_fmt_error(cudaMemcpy(B, gpu + N, sizeof(long long) * N, cudaMemcpyDeviceToHost)))
    {
        cudaFree(gpu);
        return;
    }

    cudaFree(gpu);
}



void __zScore(long long* inp, long long* out, long long eps, int n, uint8_t* error)
{
    long long* d_gpu = nullptr;

    if (*error = cuda_fmt_error(cudaMalloc(&d_gpu, 2 * n * sizeof(long long))))
    {
        cudaFree(d_gpu);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(d_gpu, inp, n * sizeof(long long), cudaMemcpyHostToDevice)))
    {
        cudaFree(d_gpu);
        return;
    }

    long long mean = __meanReduction_impl(d_gpu, n, error);
    long long std = __stdReduction_impl(d_gpu, n, error);

    if (*error)
    {
        cudaFree(d_gpu);
        return;
    }

    const int BLOCK_SIZE = 256;
    zScore_kernel<<<(n + BLOCK_SIZE - 1) / BLOCK_SIZE, BLOCK_SIZE>>>(d_gpu, d_gpu + n, mean, std + eps, n);
    *error = cuda_fmt_error(cudaMemcpy(out, d_gpu + n, n * sizeof(long long), cudaMemcpyDeviceToHost));
    cudaFree(d_gpu);
}


void __maxMinScale(long long* inp, long long* out, int n, uint8_t* error)
{
    long long* d_gpu = nullptr;

    if (*error = cuda_fmt_error(cudaMalloc(&d_gpu, 2 * n * sizeof(long long))))
    {
        cudaFree(d_gpu);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(d_gpu, inp, n * sizeof(long long), cudaMemcpyHostToDevice)))
    {
        cudaFree(d_gpu);
        return;
    }

    long long min = __minReduction_impl(d_gpu, n, error);
    long long max = __maxReduction_impl(d_gpu, n, error);

    if (*error)
    {
        cudaFree(d_gpu);
        return;
    }

    const int BLOCK_SIZE = 256;
    minMaxScale_kernel<<<(n + BLOCK_SIZE - 1) / BLOCK_SIZE, BLOCK_SIZE>>>(d_gpu, d_gpu + n, min, max, n);
    *error = cuda_fmt_error(cudaMemcpy(out, d_gpu + n, n * sizeof(long long), cudaMemcpyDeviceToHost));
    cudaFree(d_gpu);
}