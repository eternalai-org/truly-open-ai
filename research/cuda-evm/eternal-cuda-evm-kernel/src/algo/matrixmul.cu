#include <operations.cuh>

void __maxmulFloat(float *A, float *B, float *C, int m, int n, int k, uint8_t* error) {

    // Allocate device memory:
    float *gpu_A;
    float *gpu_B;
    float *gpu_C;

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu_A, sizeof(float) * m * n)))
    {
        cudaFree(gpu_A);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu_A, A, sizeof(float) * m * n, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu_A);
        return;
    }

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu_B, sizeof(float) * n * k)))
    {
        cudaFree(gpu_A), cudaFree(gpu_B);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu_B, B, sizeof(float) * n * k, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu_A), cudaFree(gpu_B);
        return;
    }

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu_C, sizeof(float) * m * k)))
    {
        cudaFree(gpu_A), cudaFree(gpu_B), cudaFree(gpu_C);
        return;
    }

    int BLOCK_SIZE = 16;
    // Blocks & grids:
    unsigned int grid_rows = (m + BLOCK_SIZE - 1) / BLOCK_SIZE;
    unsigned int grid_cols = (k + BLOCK_SIZE - 1) / BLOCK_SIZE;
    dim3 blocks(BLOCK_SIZE, BLOCK_SIZE);
    dim3 grid(grid_cols, grid_rows);

    // Call the kernel:
    vecmulFloat<<<grid, blocks>>>(gpu_A, gpu_B, gpu_C, m, n, k);

    // Get the result Matrix:
    if (*error = cuda_fmt_error(cudaMemcpy(C, gpu_C, sizeof(float) * m * k, cudaMemcpyDeviceToHost)))
    {
        cudaFree(gpu_A), cudaFree(gpu_B), cudaFree(gpu_C);
        return;
    }

    //Free device matrices
    cudaFree(gpu_A);
    cudaFree(gpu_B);
    cudaFree(gpu_C);
}

void __maxmulFixedLongLong(long long *A, long long *B, long long *C, int m, int n, int k, uint8_t* error) {
    // cout << "Start maxmulFixedLongLong\n";

    // Allocate device memory:
    long long *gpu;

    if (*error = cuda_fmt_error(cudaMallocManaged(&gpu, sizeof(long long) * (m * n + n * k + m * k))))
    {
        cudaFree(gpu);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu, A, sizeof(long long) * m * n, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu + n * m, B, sizeof(long long) * n * k, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu);
        return;
    }

    int BLOCK_SIZE = 32;
    unsigned int grid_rows = (m + BLOCK_SIZE - 1) / BLOCK_SIZE;
    unsigned int grid_cols = (k + BLOCK_SIZE - 1) / BLOCK_SIZE;
    dim3 blocks(BLOCK_SIZE, BLOCK_SIZE);
    dim3 grid(grid_cols, grid_rows);

    // Call the kernel:
    vecmulFixedLongLong<<<grid, blocks>>>(
        gpu, 
        gpu + n * m, 
        gpu + m * n + n * k, 
        m, n, k
    );

    if (*error = cuda_fmt_error(cudaMemcpy(C, gpu + m * n + n * k, sizeof(long long) * m * k, cudaMemcpyDeviceToHost)))
    {
        cudaFree(gpu);
        return;
    }
    cudaFree(gpu);
}

void __maxmulLong(long *A, long *B, long *C, long m, long n, long k, uint8_t* error) {

    // Allocate device memory:
    long *gpu_A;
    long *gpu_B;
    long *gpu_C;

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu_A, sizeof(long) * m * n)))
    {
        cudaFree(gpu_A);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu_A, A, sizeof(long) * m * n, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu_A);
        return;
    }

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu_B, sizeof(long) * n * k)))
    {
        cudaFree(gpu_A), cudaFree(gpu_B);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu_B, B, sizeof(long) * n * k, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu_A), cudaFree(gpu_B);
        return;
    }

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu_C, sizeof(long) * m * k)))
    {
        cudaFree(gpu_A), cudaFree(gpu_B), cudaFree(gpu_C);
        return;
    }

    int BLOCK_SIZE = 16;
    // Blocks & grids:
    unsigned int grid_rows = (m + BLOCK_SIZE - 1) / BLOCK_SIZE;
    unsigned int grid_cols = (k + BLOCK_SIZE - 1) / BLOCK_SIZE;
    dim3 blocks(BLOCK_SIZE, BLOCK_SIZE);
    dim3 grid(grid_cols, grid_rows);

    // Call the kernel:
    vecmulLong<<<grid, blocks>>>(gpu_A, gpu_B, gpu_C, m, n, k);

    // Get the result Matrix:
    if (*error = cuda_fmt_error(cudaMemcpy(C, gpu_C, sizeof(long) * m * k, cudaMemcpyDeviceToHost)))
    {
        cudaFree(gpu_A), cudaFree(gpu_B), cudaFree(gpu_C);
        return;
    }

    //Free device matrices
    cudaFree(gpu_A);
    cudaFree(gpu_B);
    cudaFree(gpu_C);
}

void __maxmulInt(int *A, int *B, int *C, int m, int n, int k, uint8_t* error) {

    // Allocate device memory:
    int *gpu_A;
    int *gpu_B;
    int *gpu_C;

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu_A, sizeof(int) * m * n)))
    {
        cudaFree(gpu_A);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu_A, A, sizeof(int) * m * n, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu_A);
        return;
    }

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu_B, sizeof(int) * n * k)))
    {
        cudaFree(gpu_A), cudaFree(gpu_B);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu_B, B, sizeof(int) * n * k, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu_A), cudaFree(gpu_B);
        return;
    }

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu_C, sizeof(int) * m * k)))
    {
        cudaFree(gpu_A), cudaFree(gpu_B), cudaFree(gpu_C);
        return;
    }

    int BLOCK_SIZE = 16;
    // Blocks & grids:
    unsigned int grid_rows = (m + BLOCK_SIZE - 1) / BLOCK_SIZE;
    unsigned int grid_cols = (k + BLOCK_SIZE - 1) / BLOCK_SIZE;
    dim3 blocks(BLOCK_SIZE, BLOCK_SIZE);
    dim3 grid(grid_cols, grid_rows);

    // Call the kernel:
    vecmulInt<<<grid, blocks>>>(gpu_A, gpu_B, gpu_C, m, n, k);

    // Get the result Matrix:
    if (*error = cuda_fmt_error(cudaMemcpy(C, gpu_C, sizeof(int) * m * k, cudaMemcpyDeviceToHost)))
    {
        cudaFree(gpu_A), cudaFree(gpu_B), cudaFree(gpu_C);
        return;
    }

    //Free device matrices
    cudaFree(gpu_A);
    cudaFree(gpu_B);
    cudaFree(gpu_C);
}

void __maxmulDouble(double *A, double *B, double *C, int m, int n, int k, uint8_t* error) {

    // Allocate device memory:
    double *gpu_A;
    double *gpu_B;
    double *gpu_C;

    // cout << "------------------\n";
    // cout << "A = {";
    //for (int i = 0; i < m * n; ++i) cout << A[i] << (i != m * n - 1 ? ", " : "}\n");

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu_A, sizeof(double) * m * n)))
    {
        cudaFree(gpu_A);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu_A, A, sizeof(double) * m * n, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu_A);
        return;
    }

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu_B, sizeof(double) * n * k)))
    {
        cudaFree(gpu_A), cudaFree(gpu_B);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu_B, B, sizeof(double) * n * k, cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu_A), cudaFree(gpu_B);
        return;
    }

    if (*error = cuda_fmt_error(cudaMalloc((void **) &gpu_C, sizeof(double) * m * k)))
    {
        cudaFree(gpu_A), cudaFree(gpu_B), cudaFree(gpu_C);
        return;
    }

    int BLOCK_SIZE = 32;
    // Blocks & grids:
    unsigned int grid_rows = (m + BLOCK_SIZE - 1) / BLOCK_SIZE;
    unsigned int grid_cols = (k + BLOCK_SIZE - 1) / BLOCK_SIZE;
    dim3 blocks(BLOCK_SIZE, BLOCK_SIZE);
    dim3 grid(grid_cols, grid_rows);

    // Call the kernel:
    vecmulDouble<<<grid, blocks>>>(gpu_A, gpu_B, gpu_C, m, n, k);

    // Get the result Matrix:
    if (*error = cuda_fmt_error(cudaMemcpy(C, gpu_C, sizeof(double) * m * k, cudaMemcpyDeviceToHost)))
    {
        cudaFree(gpu_A), cudaFree(gpu_B), cudaFree(gpu_C);
        return;
    }

    // cout << "C = {";
    // for (int i = 0; i < m * k; ++i) cout << C[i] << (i != m * k - 1 ? ", " : "}\n");
    //Free device matrices
    cudaFree(gpu_A);
    cudaFree(gpu_B);
    cudaFree(gpu_C);
}