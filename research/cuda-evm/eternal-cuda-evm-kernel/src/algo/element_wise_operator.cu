#include <operations.cuh>


////////////////////// kernels ///////////////////////// 


////////////////////// implementations ///////////////////////// 

void __matAddLongLong(long long *A, long long *B, long long *C, int m, int n, uint8_t* error) {
    // Allocate device memory:
    long long *gpu;
    const int N = m * n;

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu, sizeof(long long) * N * 3)))
    {
        cudaFree(gpu);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu, A, sizeof(long long) * N, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu + N, B, sizeof(long long) * N, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu);
        return;
    }
    
    const int BLOCK_SIZE = 256;
    const int GRID_SIZE = (N + BLOCK_SIZE - 1) / BLOCK_SIZE;

    mat_add_fixed_longlong<<<GRID_SIZE, BLOCK_SIZE>>>(gpu, gpu + N, gpu + 2 * N, N);

    if (*error = cuda_fmt_error(cudaMemcpy(C, gpu + 2 * N, sizeof(long long) * N, cudaMemcpyDeviceToHost)))    
    {
        cudaFree(gpu);
        return;
    }
    cudaFree(gpu);
}

void __matSubLongLong(long long *A, long long *B, long long *C, int m, int n, uint8_t* error) {
    // Allocate device memory:
    long long *gpu;
    const int N = m * n;

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu, sizeof(long long) * N * 3)))
    {
        cudaFree(gpu);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu, A, sizeof(long long) * N, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu + N, B, sizeof(long long) * N, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu);
        return;
    }

    const int BLOCK_SIZE = 256;
    const int GRID_SIZE = (N + BLOCK_SIZE - 1) / BLOCK_SIZE;

    mat_sub_fixed_longlong<<<GRID_SIZE, BLOCK_SIZE>>>(gpu, gpu + N, gpu + 2 * N, N);

    if (*error = cuda_fmt_error(cudaMemcpy(C, gpu + 2 * N, sizeof(long long) * N, cudaMemcpyDeviceToHost)))    
    {
        cudaFree(gpu);
        return;
    }

    cudaFree(gpu);
    
}

void __matMulLongLong(long long *A, long long *B, long long *C, int m, int n, uint8_t* error) {
    // Allocate device memory:
    long long *gpu;
    const int N = m * n;

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu, sizeof(long long) * N * 3)))
    {
        cudaFree(gpu);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu, A, sizeof(long long) * N, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu + N, B, sizeof(long long) * N, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu);
        return;
    }

    const int BLOCK_SIZE = 256;
    const int GRID_SIZE = (N + BLOCK_SIZE - 1) / BLOCK_SIZE;

    mat_mul_fixed_longlong<<<GRID_SIZE, BLOCK_SIZE>>>(gpu, gpu + N, gpu + 2 * N, N);

    if (*error = cuda_fmt_error(cudaMemcpy(C, gpu + 2 * N, sizeof(long long) * N, cudaMemcpyDeviceToHost)))    
    {
        cudaFree(gpu);
        return;
    }
    cudaFree(gpu);
}

void __matDivLongLong(long long *A, long long *B, long long *C, int m, int n, uint8_t* error) {
    // Allocate device memory:
    long long *gpu;
    const int N = m * n;

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu, sizeof(long long) * N * 3)))
    {
        cudaFree(gpu);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu, A, sizeof(long long) * N, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu + N, B, sizeof(long long) * N, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu);
        return;
    }

    const int BLOCK_SIZE = 256;
    const int GRID_SIZE = (N + BLOCK_SIZE - 1) / BLOCK_SIZE;

    mat_div_fixed_longlong<<<GRID_SIZE, BLOCK_SIZE>>>(gpu, gpu + N, gpu + 2 * N, N);

    if (*error = cuda_fmt_error(cudaMemcpy(C, gpu + 2 * N, sizeof(long long) * N, cudaMemcpyDeviceToHost)))    
    {
        cudaFree(gpu);
        return;
    }
    cudaFree(gpu);
}

