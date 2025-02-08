#include <operations.cuh>


// implementations - no mempory allocation or deallocation for the input
long long __sumReduction_impl(long long* d_gpu, int n, uint8_t* error)
{
    if (!n)
    {
        *error = 1;
        return 0;
    }

    long long res = 0;
    int block_sz = 256;
    int block_sz2 = block_sz * 2;
    int grid_sz = (n + block_sz2 - 1) / block_sz2;

    long long* blockSum;

    if (*error = cuda_fmt_error(cudaMalloc(&blockSum, grid_sz * sizeof(long long))))
    {
        cudaFree(blockSum);
        return 0;
    }

    if (*error = cuda_fmt_error(cudaMemset(blockSum, 0, grid_sz * sizeof(long long))))
    {
        cudaFree(blockSum);
        return 0;
    }

    sumReduction_kernel<<<grid_sz, block_sz, block_sz2 * sizeof(long long)>>>(d_gpu, blockSum, n);


    if (grid_sz > 1)
    {
        res = __sumReduction_impl(blockSum, grid_sz, error);
    }
    else
    {
        if (*error = cuda_fmt_error(cudaMemcpy(&res, blockSum, sizeof(long long), cudaMemcpyDeviceToHost)))
        {
            cudaFree(blockSum);
            return 0;
        }
    }

    cudaFree(blockSum);
    return res;
}



long long __maxReduction_impl(long long* d_gpu, int n, uint8_t* error)
{
    if (!n)
    {
        *error = 1;
        return 0;
    }

    long long res = 0;
    int block_sz = 256;
    int block_sz2 = block_sz * 2;
    int grid_sz = (n + block_sz2 - 1) / block_sz2;

    long long* blockMax;

    if (*error = cuda_fmt_error(cudaMalloc(&blockMax, grid_sz * sizeof(long long))))
    {
        cudaFree(blockMax);
        return 0;
    }

    maxReduction_kernel<<<grid_sz, block_sz, block_sz2 * sizeof(long long)>>>(d_gpu, blockMax, n);

    if (grid_sz > 1)
    {
        res = __maxReduction_impl(blockMax, grid_sz, error);
    }
    else
    {
        if (*error = cuda_fmt_error(cudaMemcpy(&res, blockMax, sizeof(long long), cudaMemcpyDeviceToHost)))
        {
            cudaFree(blockMax);
            return 0;
        }
    }

    cudaFree(blockMax);
    return res;
}



long long __minReduction_impl(long long* d_gpu, int n, uint8_t* error)
{
    if (!n)
    {
        *error = 1;
        return 0;
    }

    long long res = 0;
    int block_sz = 256;
    int block_sz2 = block_sz * 2;
    int grid_sz = (n + block_sz2 - 1) / block_sz2;

    long long* blockMin;

    if (*error = cuda_fmt_error(cudaMalloc(&blockMin, grid_sz * sizeof(long long))))
    {
        cudaFree(blockMin);
        return 0;
    }

    minReduction_kernel<<<grid_sz, block_sz, block_sz2 * sizeof(long long)>>>(d_gpu, blockMin, n);

    if (grid_sz > 1)
    {
        res = __minReduction_impl(blockMin, grid_sz, error);
    }
    else
    {
        if (*error = cuda_fmt_error(cudaMemcpy(&res, blockMin, sizeof(long long), cudaMemcpyDeviceToHost)))
        {
            cudaFree(blockMin);
            return 0;
        }
    }

    cudaFree(blockMin);
    return res;
}

void __channelWiseSumReduction_impl(long long* d_gpu, long long* d_out, int n, int c, uint8_t* error)
{
    if (!(n * c))
    {
        *error = 1;
        return;
    }

    long long* blockSum;

    const int block_size = 256;
    const int block_size_2 = block_size * 2;
    const int blocks = (n + block_size_2 - 1) / block_size_2;

    if (*error = cuda_fmt_error(cudaMalloc(&blockSum, blocks * c * sizeof(long long))))
    {
        cudaFree(blockSum);
        return;
    }

    sumReductionV2_kernel<<<
        dim3(blocks, 1, c), 
        dim3(block_size, 1, 1), 
        sizeof(long long) * block_size_2>>>(d_gpu, blockSum, n, c);

    if (blocks > 1)
    {
        __channelWiseSumReduction_impl(blockSum, d_out, blocks, c, error);
    }
    else
    {
        if (*error = cuda_fmt_error(cudaMemcpy(d_out, blockSum, c * sizeof(long long), cudaMemcpyDeviceToDevice)))
        {
            cudaFree(blockSum);
            return;
        }
    }

    cudaFree(blockSum);
}

long long __meanReduction_impl(long long* gpu_inp, int n, uint8_t* error)
{
    if (!n)
    {
        *error = 1;
        return 0;
    }

    return FixedLongLong::div(__sumReduction_impl(gpu_inp, n, error), (1LL * n) << 32);
}

long long __stdReduction_impl(long long* d_gpu, int n, uint8_t* error)
{
    if (!n)
    {
        *error = 1;
        return 0;
    }

    long long mean = __meanReduction_impl(d_gpu, n, error);
    
    if (*error)
    {
        return 0;
    }

    long long* sub = nullptr;
    
    if (*error = cuda_fmt_error(cudaMalloc(&sub, 2 * n * sizeof(long long))))
    {
        cudaFree(sub);
        return 0;
    }

    const int BLOCK_SIZE = 256;

    mat_sub_single_fixed_longlong<<<(n + BLOCK_SIZE - 1) / BLOCK_SIZE, BLOCK_SIZE>>>(d_gpu, sub, mean, n);
    mat_pow2_single_fixed_longlong<<<(n + BLOCK_SIZE - 1) / BLOCK_SIZE, BLOCK_SIZE>>>(sub, sub + n, n);
    long long res = FixedLongLong::sqrt(FixedLongLong::div(__sumReduction_impl(sub + n, n, error), (1LL * n) << 32));

    cudaFree(sub);
    return res;
}


////////////////////// wrappers //////////////////////


long long __sumReduction(long long* inp, int n, uint8_t* error)
{
    long long* gpu; 
    if (*error = cuda_fmt_error(cudaMalloc(&gpu, n * sizeof(long long))))
    {
        cudaFree(gpu);
        return 0;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu, inp, n * sizeof(long long), cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu);
        return 0;
    }
    long long res = __sumReduction_impl(gpu, n, error);
    cudaFree(gpu);
    return res;
}


long long __maxReduction(long long* inp, int n, uint8_t* error)
{
    long long* gpu; 

    if (*error = cuda_fmt_error(cudaMalloc(&gpu, n * sizeof(long long))))
    {
        cudaFree(gpu);
        return 0;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu, inp, n * sizeof(long long), cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu);
        return 0;
    }

    long long res = __maxReduction_impl(gpu, n, error);
    cudaFree(gpu);
    return res;
}

long long __minReduction(long long* inp, int n, uint8_t* error)
{
    long long* gpu; 

    if (*error = cuda_fmt_error(cudaMalloc(&gpu, n * sizeof(long long))))
    {
        cudaFree(gpu);
        return 0;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu, inp, n * sizeof(long long), cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu);
        return 0;
    }

    long long res = __minReduction_impl(gpu, n, error);
    cudaFree(gpu);
    return res;
}

long long __meanReduction(long long* inp, int n, uint8_t* error)
{
    return FixedLongLong::div(__sumReduction(inp, n, error), (1ll * n) << 32);
}

long long __stdReduction(long long* inp, int n, uint8_t* error)
{
    long long* d_gpu = nullptr;
    if (*error = cuda_fmt_error(cudaMalloc(&d_gpu, n * sizeof(long long))))
    {
        cudaFree(d_gpu);
        return 0;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(d_gpu, inp, n * sizeof(long long), cudaMemcpyHostToDevice)))
    {
        cudaFree(d_gpu);
        return 0;
    }

    long long mean = __stdReduction_impl(d_gpu, n, error);
    cudaFree(d_gpu);

    return mean;
}

void __channelWiseSumReduction(long long* inp, long long* out, int n, int c, uint8_t* error)
{
    long long* d_gpu = nullptr;
    
    if (*error = cuda_fmt_error(cudaMalloc(&d_gpu, (n * c + c) * sizeof(long long))))
    {
        cudaFree(d_gpu);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(d_gpu + c, inp, n * c * sizeof(long long), cudaMemcpyHostToDevice)))
    {
        cudaFree(d_gpu);
        return;
    }
    
    __channelWiseSumReduction_impl(d_gpu + c, d_gpu, n, c, error);

    *error = *error || cuda_fmt_error(cudaMemcpy(out, d_gpu, c * sizeof(long long), cudaMemcpyDeviceToHost));
    cudaFree(d_gpu);
}


void __channelWiseMeanReduction(long long* inp, long long* out, int n, int c, uint8_t* error)
{
    long long* d_gpu = nullptr;
    
    if (*error = cuda_fmt_error(cudaMalloc(&d_gpu, (n * c + c) * sizeof(long long))))
    {
        cudaFree(d_gpu);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(d_gpu + c, inp, n * c * sizeof(long long), cudaMemcpyHostToDevice)))
    {
        cudaFree(d_gpu);
        return;
    }
    
    __channelWiseSumReduction_impl(d_gpu + c, d_gpu, n, c, error);

    const int BLOCK_SIZE = 256;
    mat_div_single_fixed_longlong<<<(c + BLOCK_SIZE - 1) / BLOCK_SIZE, BLOCK_SIZE>>>(d_gpu, d_gpu + c, n, c);

    *error = *error || cuda_fmt_error(cudaMemcpy(out, d_gpu + c, c * sizeof(long long), cudaMemcpyDeviceToHost));
    cudaFree(d_gpu);
}