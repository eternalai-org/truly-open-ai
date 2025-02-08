#include <operations.cuh>

////////////////////// implementation ///////////////////////// 

void __conv2dFixedLongLong(
    long long* inp, long long* kernel, long long* bias, long long* out, // data io
    int kernel_size, int in_channel, int out_channel, // kernel properties
    int h, int w, // spatial size of inp,
    int padding, int stride_h, int stride_w // padding: same(0) or valid(1)
    , uint8_t* error
)
{

    // @TODO: reduce redundant iterations and support for k_w != k_h case, in_w != in_h also

    // references:
    // [1] https://www.pico.net/kb/what-is-the-difference-between-same-and-valid-padding-in-tf-nn-max-pool-of-tensorflow/

    // parameters
    // inp: [h, w, in_channel]
    // kernel: [kernel_size, kernel_size, in_channel, out_channel]
    // out: [h_out, w_out, out_channel]
    // kernel_size: size of kernel
    // in_channel: number of input channels
    // out_channel: number of output channels
    // w: width of input
    // h: height of input
    // padding: padding mode, one of 'valid': 0 or 'same': 1
    // strides: strides of kernel along width and height

    int pad_top = 0, pad_bottom = 0, pad_left = 0, pad_right = 0;

    if (padding == 1)
    {
        int out_h = (h + stride_h - 1) / stride_h;
        int out_w = (w + stride_w - 1) / stride_w;

        int pad_h = max((out_h - 1) * stride_h + kernel_size - h, 0);
        int pad_w = max((out_w - 1) * stride_w + kernel_size - w, 0);

        pad_top = pad_h / 2;
        pad_bottom = pad_h - pad_top;

        pad_left = pad_w / 2;
        pad_right = pad_w - pad_left;
    }

    long long *d_gpu;

    const int out_w = (w + pad_left + pad_right - kernel_size) / stride_w + 1;
    const int out_h = (h + pad_top + pad_bottom - kernel_size) / stride_h + 1;

    const uint64_t inpFlatSize = (w + pad_left + pad_right) * (h + pad_top + pad_bottom) * in_channel;
    const uint64_t kernelFlatSize = out_channel * kernel_size * kernel_size * in_channel;
    const uint64_t outFlatSize = out_h * out_w * out_channel;
    const uint64_t cudaMemSize = inpFlatSize + outFlatSize + kernelFlatSize + out_channel;

    if (*error = cuda_fmt_error(cudaMalloc(&d_gpu, cudaMemSize * sizeof(long long))))
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

    if (*error = cuda_fmt_error(cudaMemcpy(d_gpu + inpFlatSize, kernel, kernelFlatSize * sizeof(long long), cudaMemcpyHostToDevice)))
    {
        cudaFree(d_gpu);
        return;
    }
    
    if (*error = cuda_fmt_error(cudaMemcpy(d_gpu + inpFlatSize + kernelFlatSize, bias, out_channel * sizeof(long long), cudaMemcpyHostToDevice)))
    {
        cudaFree(d_gpu);
        return;
    }

    const int BLOCK_SIZE = 32;

    const dim3 THREAD_PER_BLOCK(BLOCK_SIZE, BLOCK_SIZE, 1);
    const dim3 BLOCK_PER_GRID((out_w + BLOCK_SIZE - 1) / BLOCK_SIZE, (out_h + BLOCK_SIZE - 1) / BLOCK_SIZE, out_channel);

    conv2dImplFixedLongLong_kernel<<<BLOCK_PER_GRID, THREAD_PER_BLOCK>>>(
        d_gpu, // inp
        d_gpu + inpFlatSize, // kernel 
        d_gpu + inpFlatSize + kernelFlatSize, // bias 
        d_gpu + inpFlatSize + kernelFlatSize + out_channel, // out 
        kernel_size, in_channel, out_channel, 
        h + pad_top + pad_bottom,  w + pad_left + pad_right, 
        out_h, out_w, padding, stride_h, stride_w
    );

    
    *error = cuda_fmt_error(cudaMemcpy(out, d_gpu + inpFlatSize + kernelFlatSize + out_channel, 
        out_h * out_w * out_channel * sizeof(long long), cudaMemcpyDeviceToHost));
    cudaFree(d_gpu);
}

void __depthwiseConv2dFixedLongLong(
    long long* inp, long long* kernel, long long* bias, long long* out, // data io
    int h, int w, int in_channel, // spatial size of inp,
    int kernel_size_h, int kernel_size_w,  // kernel properties
    int padding, int stride_h, int stride_w, // padding: same(0) or valid(1)
    uint8_t* error
)
{
    int pad_top = 0, pad_bottom = 0, pad_left = 0, pad_right = 0;

    if (padding == 1)
    {
        int out_h = (h + stride_h - 1) / stride_h;
        int out_w = (w + stride_w - 1) / stride_w;

        int pad_h = max((out_h - 1) * stride_h + kernel_size_h - h, 0);
        int pad_w = max((out_w - 1) * stride_w + kernel_size_w - w, 0);

        pad_top = pad_h / 2;
        pad_bottom = pad_h - pad_top;

        pad_left = pad_w / 2;
        pad_right = pad_w - pad_left;
    }

    long long *d_gpu;

    const int out_w = (w + pad_left + pad_right - kernel_size_w) / stride_w + 1;
    const int out_h = (h + pad_top + pad_bottom - kernel_size_h) / stride_h + 1;

    const uint64_t inpFlatSize = (w + pad_left + pad_right) * (h + pad_top + pad_bottom) * in_channel;
    const uint64_t kernelFlatSize = in_channel * kernel_size_h * kernel_size_w; // * 1
    const uint64_t outFlatSize = out_h * out_w *  in_channel;
    const uint64_t cudaMemSize = inpFlatSize + outFlatSize + kernelFlatSize + in_channel;

    if (*error = cuda_fmt_error(cudaMalloc(&d_gpu, cudaMemSize * sizeof(long long))))
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
    
    if (*error = cuda_fmt_error(cudaMemcpy(d_gpu + inpFlatSize, kernel, kernelFlatSize * sizeof(long long), cudaMemcpyHostToDevice)) 
    || cuda_fmt_error(cudaMemcpy(d_gpu + inpFlatSize + kernelFlatSize, bias, in_channel * sizeof(long long), cudaMemcpyHostToDevice)))
    {
        cudaFree(d_gpu);
        return;
    }

    const int BLOCK_SIZE = 32;
    const dim3 THREAD_PER_BLOCK(BLOCK_SIZE, BLOCK_SIZE, 1);
    const dim3 BLOCK_PER_GRID((out_w + BLOCK_SIZE - 1) / BLOCK_SIZE, (out_h + BLOCK_SIZE - 1) / BLOCK_SIZE, in_channel);

    depthwise_conv2d_kernel<<<BLOCK_PER_GRID, THREAD_PER_BLOCK>>>(
        d_gpu, // inp
        d_gpu + inpFlatSize, // kernel 
        d_gpu + inpFlatSize + kernelFlatSize, // bias 
        d_gpu + inpFlatSize + kernelFlatSize + in_channel, // out 
        h + pad_top + pad_bottom,  w + pad_left + pad_right, in_channel, 
        kernel_size_h, kernel_size_w,
        out_h, out_w, 
        padding, stride_h, stride_w
    );

    *error = cuda_fmt_error(cudaMemcpy(out, d_gpu + inpFlatSize + kernelFlatSize + in_channel,
        out_h * out_w * in_channel * sizeof(long long), cudaMemcpyDeviceToHost));

    cudaFree(d_gpu);
}

uint8_t estimateConvOutputSize(
    int kernel_size, int in_channel, int out_channel, // kernel properties
    int h, int w, // spatial size of inp,
    int padding, int stride_h, int stride_w, // padding: same(0) or valid(1)
    int* out_h, int* out_w // spatial size of out
)
{
    if (!out_h || !out_w)
    {
        return ERROR;
    }

    int pad_top = 0, pad_bottom = 0, pad_left = 0, pad_right = 0;

    if (padding == 1)
    {
        int _out_h = (h + stride_h - 1) / stride_h;
        int _out_w = (w + stride_w - 1) / stride_w;

        int pad_h = max((_out_h - 1) * stride_h + kernel_size - h, 0);
        int pad_w = max((_out_w - 1) * stride_w + kernel_size - w, 0);

        pad_top = pad_h / 2;
        pad_bottom = pad_h - pad_top;

        pad_left = pad_w / 2;
        pad_right = pad_w - pad_left;
    }

    *out_w = (w + pad_left + pad_right - kernel_size) / stride_w + 1;
    *out_h = (h + pad_top + pad_bottom - kernel_size) / stride_h + 1;

    printf("out_w %d; out_h: %d\n", *out_w, *out_h);

    return OK;
}