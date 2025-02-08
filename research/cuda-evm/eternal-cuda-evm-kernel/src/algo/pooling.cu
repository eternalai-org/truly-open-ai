#include <operations.cuh>

////////////////////// implementation ///////////////////////// 


void globalAvgPoolingFixedLongLong_impl(
    long long* inp, long long* out,
    int h, int w, int in_channel
    , uint8_t* error
)
{
    const int BLOCK_SZ = 256;
    const int BLOCK_SZ2 = BLOCK_SZ * 2;
    const int grid_sz_x = (h * w + BLOCK_SZ2 - 1) / BLOCK_SZ2;

    long long* blockSum = nullptr;
    
    if (*error = cuda_fmt_error(cudaMalloc(&blockSum, sizeof(long long) * grid_sz_x * in_channel)))
    {
        cudaFree(blockSum);
        return;
    }

    sumReductionV2_kernel<<<
        dim3(grid_sz_x, 1, in_channel), 
        dim3(BLOCK_SZ, 1, 1), 
        sizeof(long long) * BLOCK_SZ2
    >>>(inp, blockSum, h * w, in_channel);

    if (grid_sz_x > 1)
    {
        globalAvgPoolingFixedLongLong_impl(blockSum, out, grid_sz_x, 1, in_channel, error);
    }
    else
    {
        *error = cuda_fmt_error(cudaMemcpy(out, blockSum, in_channel * sizeof(long long), cudaMemcpyDeviceToDevice));
    }

    cudaFree(blockSum);    
}

void __maxPoolingFixedLongLong(
    long long* inp, long long* out, // data io
    int h, int w, int in_channel, // in spatial size, in_channel
    int pool_size, int stride_h, int stride_w, // pooling size, stride 
    int padding // padding mode, one of 'valid': 0 or 'same': 1
    , uint8_t* error
)
{
    // only support for squared pool size, squared input

    if (w != h)
    {
        // not sure if it works
        return;
    }

    if (stride_h <= 0)
    {
        stride_h = pool_size;
    }

    if (stride_w <= 0)
    {
        stride_w = pool_size;
    }

    int pad_top = 0, pad_bottom = 0, pad_left = 0, pad_right = 0;

    if (padding == 1)
    {
        int out_h = (h + stride_h - 1) / stride_h;
        int out_w = (w + stride_w - 1) / stride_w;

        int pad_h = max((out_h - 1) * stride_h + pool_size - h, 0);
        int pad_w = max((out_w - 1) * stride_w + pool_size - w, 0);
        
        pad_top = pad_h / 2;
        pad_bottom = pad_h - pad_top;
        
        pad_left = pad_w / 2;
        pad_right = pad_w - pad_left;
    }

    long long *d_gpu;

    int out_w = (w + pad_left + pad_right - pool_size) / stride_w + 1;
    int out_h = (h + pad_top + pad_bottom - pool_size) / stride_h + 1;

    uint64_t inpFlatSize = (w + pad_left + pad_right) * (h + pad_top + pad_bottom) * in_channel;
    uint64_t outFlatSize = out_h * out_w * in_channel;
    uint64_t flatSize = inpFlatSize + outFlatSize;

    if (*error = cuda_fmt_error(cudaMalloc(&d_gpu, flatSize * sizeof(long long))))
    {
        cudaFree(d_gpu);
        return;
    }
    
    if (padding == 1)
    {
        long long* padded_inp = new long long[inpFlatSize];
        memset(padded_inp, 0x00, inpFlatSize * sizeof(long long));

        for (int i = 0; i < h; ++i)
        {
            for (int j = 0; j < w; ++j)
            {
                memcpy(padded_inp + ((i + pad_top) * (w + pad_left + pad_right) + j + pad_left) * in_channel, inp + (i * w + j) * in_channel, in_channel << 3);
            }
        }

        if (*error = cuda_fmt_error(cudaMemcpy(d_gpu, padded_inp, inpFlatSize * sizeof(long long), cudaMemcpyHostToDevice)))
        {
            cudaFree(d_gpu);
            delete[] padded_inp;
            return;
        }

        delete[] padded_inp;
    }
    else
    {
        if (*error = cuda_fmt_error(cudaMemcpy(d_gpu, inp, inpFlatSize * sizeof(long long), cudaMemcpyHostToDevice)))
        {
            cudaFree(d_gpu);
            return;
        }
    }
    
    const int thread_x = 32, thread_y=32, thread_z=1; 
    dim3 threads_per_block(thread_x, thread_y, thread_z);

    const dim3 block_per_grid(
        (out_w + threads_per_block.x - 1) / threads_per_block.x,
        (out_h + threads_per_block.y - 1) / threads_per_block.y,
        (in_channel + threads_per_block.z - 1) / threads_per_block.z
    );

    maxPoolingImplFixedLongLong_kernel<<<block_per_grid, threads_per_block>>>(
        d_gpu, d_gpu + inpFlatSize, 
        h + pad_bottom + pad_top, w + pad_left + pad_right, in_channel, 
        pool_size, stride_h, stride_w, 
        out_h, out_w, 
        pad_top, pad_bottom, pad_left, pad_right
    );

    if (*error = cuda_fmt_error(cudaMemcpy(out, d_gpu + inpFlatSize, outFlatSize * sizeof(long long), cudaMemcpyDeviceToHost)))
    {
        cudaFree(d_gpu);
        return;
    }

    cudaFree(d_gpu);
}

void __avgPoolingFixedLongLong(
    long long* inp, long long* out,
    int h, int w, int in_channel,
    int pool_size, int stride_h, int stride_w,
    int padding
    , uint8_t* error
)
{
    // only support for squared pool size, squared input

    if (w != h)
    {
        // not sure if it works
        return;
    }

    if (stride_h <= 0)
    {
        stride_h = pool_size;
    }

    if (stride_w <= 0)
    {
        stride_w = pool_size;
    }

    int pad_top = 0, pad_bottom = 0, pad_left = 0, pad_right = 0;

    if (padding == 1)
    {
        int out_h = (h + stride_h - 1) / stride_h;
        int out_w = (w + stride_w - 1) / stride_w;

        int pad_h = max((out_h - 1) * stride_h + pool_size - h, 0);
        int pad_w = max((out_w - 1) * stride_w + pool_size - w, 0);
        
        pad_top = pad_h / 2;
        pad_bottom = pad_h - pad_top;
        
        pad_left = pad_w / 2;
        pad_right = pad_w - pad_left;
    }

    long long *d_gpu;

    int out_w = (w + pad_left + pad_right - pool_size) / stride_w + 1;
    int out_h = (h + pad_top + pad_bottom - pool_size) / stride_h + 1;

    uint64_t inpFlatSize = (w + pad_left + pad_right) * (h + pad_top + pad_bottom) * in_channel;
    uint64_t outFlatSize = out_h * out_w * in_channel;
    uint64_t flatSize = inpFlatSize + outFlatSize;

    if (*error = cuda_fmt_error(cudaMalloc(&d_gpu, flatSize * sizeof(long long))))
    {
        cudaFree(d_gpu);
        return;
    }
    
    if (padding == 1)
    {
        long long* padded_inp = new long long[inpFlatSize];
        memset(padded_inp, 0x00, inpFlatSize * sizeof(long long));

        for (int i = 0; i < h; ++i)
        {
            for (int j = 0; j < w; ++j)
            {
                memcpy(
                    padded_inp + ((i + pad_top) * (w + pad_left + pad_right) + j + pad_left) * in_channel, 
                    inp + (i * w + j) * in_channel, 
                    in_channel << 3
                );
            }
        }
        if (*error = cuda_fmt_error(cudaMemcpy(d_gpu, padded_inp, inpFlatSize * sizeof(long long), cudaMemcpyHostToDevice)))
        {
            cudaFree(d_gpu);
            delete[] padded_inp;
            return;
        }
        delete[] padded_inp;
    }
    else
    {
        if (*error = cuda_fmt_error(cudaMemcpy(d_gpu, inp, inpFlatSize * sizeof(long long), cudaMemcpyHostToDevice)))
        {
            cudaFree(d_gpu);
            return;
        }
    }

    const int thread_x = 32, thread_y = 32, thread_z = 1;
    dim3 threads_per_block(thread_x, thread_y, thread_z);

    const dim3 block_per_grid(
        (out_w + threads_per_block.x - 1) / threads_per_block.x,
        (out_h + threads_per_block.y - 1) / threads_per_block.y,
        (in_channel + threads_per_block.z - 1) / threads_per_block.z
    );

    avgPoolingImplFixedLongLong_kernel<<<block_per_grid, threads_per_block>>>(
        d_gpu, d_gpu + inpFlatSize, 
        h + pad_top + pad_bottom, w + pad_left + pad_right, in_channel, 
        pool_size, stride_h, stride_w, 
        out_h, out_w,
        pad_top, pad_bottom, pad_left, pad_right
    );

    if (*error = cuda_fmt_error(cudaMemcpy(out, d_gpu + inpFlatSize, outFlatSize * sizeof(long long), cudaMemcpyDeviceToHost)))
    {
        cudaFree(d_gpu);
        return;
    }
    cudaFree(d_gpu);
}

void __globalAvgPoolingFixedLongLong(
    long long* inp, long long* out,
    int h, int w, int in_channel,
    uint8_t* error
)
{
    long long* gpu = nullptr;

    if (*error = cuda_fmt_error(cudaMalloc(&gpu, sizeof(long long) * (h * w * in_channel + in_channel))))
    {
        cudaFree(gpu);
        return;
    }

    if (*error = cuda_fmt_error(cudaMemcpy(gpu + in_channel, inp, h * w * in_channel * sizeof(long long), cudaMemcpyHostToDevice)))
    {
        cudaFree(gpu);
        return;
    }

    globalAvgPoolingFixedLongLong_impl(gpu + in_channel, gpu, h, w, in_channel, error);

    if (!*error)
    {
        // div and copy
        const int BLOCK_SIZE = 256;
        mat_div_single_fixed_longlong<<<(in_channel + BLOCK_SIZE -1) / BLOCK_SIZE, BLOCK_SIZE>>>(gpu, gpu + in_channel, (1LL * h * w) << 32, in_channel);
        *error = cuda_fmt_error(cudaMemcpy(out, gpu + in_channel, in_channel * sizeof(long long), cudaMemcpyDeviceToHost));
    }

    cudaFree(gpu);
}

// @deprecated
uint8_t estimatePoolingOutputSize(
    int h, int w, int in_channel,
    int pool_size, int padding, 
    int stride_h, int stride_w,
    int* out_h, int* out_w
)
{
    if (!out_h || !out_w)
    {
        return ERROR;
    }

    if (stride_h <= 0)
    {
        stride_h = pool_size;
    }

    if (stride_w <= 0)
    {
        stride_w = pool_size;
    }

    int pad_top = 0, pad_bottom = 0, pad_left = 0, pad_right = 0;

    if (padding == 1)
    {
        int out_h = (h + stride_h - 1) / stride_h;
        int out_w = (w + stride_w - 1) / stride_w;

        int pad_h = max((out_h - 1) * stride_h + pool_size - h, 0);
        int pad_w = max((out_w - 1) * stride_w + pool_size - w, 0);
        
        pad_top = pad_h / 2;
        pad_bottom = pad_h - pad_top;
        
        pad_left = pad_w / 2;
        pad_right = pad_w - pad_left;
    }

    *out_w = (w + pad_left + pad_right - pool_size) / stride_w + 1;
    *out_h = (h + pad_top + pad_bottom - pool_size) / stride_h + 1;
    return OK;
}
