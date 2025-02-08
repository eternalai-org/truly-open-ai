#include <fixedlonglong32x32.cuh>

__global__ void arraySum_kernel(long long* A, long long* sum, int n)
{
    int block_size = sqrt(1.0f * n);
    int offset = blockIdx.x * block_size;
    int r_bound = min(offset + block_size, n);

    long long s = 0;
    
    for (int i = offset; i < r_bound; ++i)
    {
        s = A[i] + s;
    }

    sum[blockIdx.x] = s;
}

__global__ void arrayExp_kernel(long long* A, long long* B, int n)
{
    int i = blockIdx.x * blockDim.x + threadIdx.x;
    if (i < n)
    {
        B[i] = FixedLongLong::exp(A[i]);
    }
}

// softmax
__global__ void softmaxImplFixedLongLong(long long *expA, long long* B, int n, long long sumExp)
{
    int tid = threadIdx.x;
    int i = blockIdx.x * blockDim.x + tid;

    if (i < n)
    {
        B[i] = FixedLongLong::div(expA[i], sumExp);
    }
}

// relu activation
__global__ void reluImplFixedLongLong(long long *A, long long* B, int m)
{
    int tid = threadIdx.x;
    int i = blockIdx.x * blockDim.x + tid;
    
    if (i < m)
    {
        B[i] = max(0ll, A[i]);
    }
}


// sigmoid activation
__global__ void sigmoidImplFixedLongLong(long long *A, long long* B, int n)
{
    int tid = threadIdx.x;
    int i = blockIdx.x * blockDim.x + tid;

    if (i < n)
    {
        if (A[i] < -20LL << 32)
        {
            B[i] = 0;
        }
        else if (A[i] > 20LL << 32)
        {
            B[i] = 1LL << 32;
        }
        else {
            long long expNegA = FixedLongLong::exp(-A[i]);
           B[i] = FixedLongLong::reciprocal(expNegA + (1ll << 32)); 
        }
    }
}


// tanh activation 
__global__ void tanhImplFixedLongLong(long long *A, long long* B, int n)
{
    int tid = threadIdx.x;
    int i = blockIdx.x * blockDim.x + tid;


    if (i < n)
    {
        if (A[i] < -20LL << 32)
        {
            B[i] = (1LL << 32) * -1;
        }
        else if (A[i] > 20LL << 32)
        {
            B[i] = 1LL << 32;
        }
        else {
            long long expAi = FixedLongLong::exp(A[i]);
            long long expNegAi = FixedLongLong::exp(-A[i]);
            B[i] = FixedLongLong::div(expAi - expNegAi, expAi + expNegAi);    
        }
    }
}

__global__ void vecmulFixedLongLong(long long *A, long long *B, long long *C, int m, int n, int k) {
    // Row and Column indexes:
    int row = blockIdx.y * blockDim.y + threadIdx.y;
    int col = blockIdx.x * blockDim.x + threadIdx.x;

    // Are they bellow the maximum?
    if (col < k && row < m) {
        long long result = 0;
        for (int ix = 0; ix < n; ix++) {
            result = FixedLongLong::add(result, FixedLongLong::mul(A[row * n + ix], B[ix * k + col]));
        }
        C[row * k + col] = result;
    }
}

__global__ void vecmulInt(int *A, int *B, int *C, int m, int n, int k) {
    // Row and Column indexes:
    int row = blockIdx.y * blockDim.y + threadIdx.y;
    int col = blockIdx.x * blockDim.x + threadIdx.x;

    // Are they bellow the maximum?
    if (col < k && row < m) {
        int result = 0;
        for (int ix = 0; ix < n; ix++) {
            result += A[row * n + ix] * B[ix * k + col];
        }
        C[row * k + col] = result;
    }
}

__global__ void vecmulLong(long *A, long *B, long *C, int m, int n, int k) {
    // Row and Column indexes:
    int row = blockIdx.y * blockDim.y + threadIdx.y;
    int col = blockIdx.x * blockDim.x + threadIdx.x;

    // Are they bellow the maximum?
    if (col < k && row < m) {
        long result = 0;
        for (int ix = 0; ix < n; ix++) {
            result += A[row * n + ix] * B[ix * k + col];
        }
        C[row * k + col] = result;
    }
}

__global__ void vecmulFloat(float *A, float *B, float *C, int m, int n, int k) {
    // Row and Column indexes:
    int row = blockIdx.y * blockDim.y + threadIdx.y;
    int col = blockIdx.x * blockDim.x + threadIdx.x;

    // Are they bellow the maximum?
    if (col < k && row < m) {
        float result = 0;
        for (int ix = 0; ix < n; ix++) {
            result += A[row * n + ix] * B[ix * k + col];
        }
        C[row * k + col] = result;
    }
}

__global__ void vecmulDouble(double *A, double *B, double *C, int m, int n, int k) {
    // Row and Column indexes:
    int row = blockIdx.y * blockDim.y + threadIdx.y;
    int col = blockIdx.x * blockDim.x + threadIdx.x;

    // Are they bellow the maximum?
    if (col < k && row < m) {
        double result = 0;
        for (int ix = 0; ix < n; ix++) {
            result += A[row * n + ix] * B[ix * k + col];
        }
        C[row * k + col] = result;
    }
}

__global__ void normalizeFixedLongLong_kernel(
    long long *X, // input of shape (m, n, c) 
    long long *Y, // output of shape (m, n, c)
    long long *ma, // moving average(c)
    long long *mv, // moving variance (c)
    long long *gamma, // scale of shape (c)
    long long *beta, // offset of shape (c)
    long long epsilon, // epsilon 
    int h, int w, int c  // m, n, c 
) 
{
    int out_row = blockIdx.y * blockDim.y + threadIdx.y; 
    int out_col = blockIdx.x * blockDim.x + threadIdx.x;
    int out_c = blockIdx.z * blockDim.z + threadIdx.z;

    if (out_row < h && out_col < w) {
        int idx = out_row * w * c + out_col * c + out_c;
        Y[idx] = FixedLongLong::mul(    
            FixedLongLong::div(X[idx] - ma[out_c], FixedLongLong::sqrt(mv[out_c] + epsilon)), 
            gamma[out_c]
        ) + beta[out_c];
    }
}

__global__ void sumReduction_kernel(long long* d_gpu, long long* blockOutput, int n)
{
    const int tid = threadIdx.x;
    const int glbl_tid = 2 * blockDim.x * blockIdx.x + tid;

    extern __shared__ long long s_out[];

    s_out[tid] = 0;
    s_out[tid + blockDim.x] = 0;

    if (glbl_tid < n)
    {
        s_out[tid] = d_gpu[glbl_tid];
        if (glbl_tid + blockDim.x < n)
            s_out[tid + blockDim.x] = d_gpu[glbl_tid + blockDim.x];
    }
    __syncthreads();

    for (unsigned int s = blockDim.x; s > 0; s >>= 1) {
        if (tid < s) {
            s_out[tid] += s_out[tid + s];
        }
        __syncthreads();
    }

    if (tid == 0)
        blockOutput[blockIdx.x] = s_out[0];
}

__global__ void maxReduction_kernel(long long* d_gpu, long long* blockOutput, int n)
{
    const int tid = threadIdx.x;
    const int glbl_tid = 2 * blockDim.x * blockIdx.x + tid;

    extern __shared__ long long s_out[];

    s_out[tid] = d_gpu[0];
    s_out[tid + blockDim.x] = d_gpu[0];

    if (glbl_tid < n)
    {
        s_out[tid] = d_gpu[glbl_tid];
        if (glbl_tid + blockDim.x < n)
            s_out[tid + blockDim.x] = d_gpu[glbl_tid + blockDim.x];
    }
    __syncthreads();

    for (unsigned int s = blockDim.x; s > 0; s >>= 1) {
        if (tid < s) {
            s_out[tid] = max(s_out[tid + s], s_out[tid]);
        }
        __syncthreads();
    }

    if (tid == 0)
        blockOutput[blockIdx.x] = s_out[0];
}

__global__ void minReduction_kernel(long long* d_gpu, long long* blockOutput, int n)
{
    const int tid = threadIdx.x;
    const int glbl_tid = 2 * blockDim.x * blockIdx.x + tid;

    extern __shared__ long long s_out[];

    s_out[tid] = d_gpu[0];
    s_out[tid + blockDim.x] = d_gpu[0];

    if (glbl_tid < n)
    {
        s_out[tid] = d_gpu[glbl_tid];
        if (glbl_tid + blockDim.x < n)
            s_out[tid + blockDim.x] = d_gpu[glbl_tid + blockDim.x];
    }
    __syncthreads();

    for (unsigned int s = blockDim.x; s > 0; s >>= 1) {
        if (tid < s) {
            s_out[tid] = min(s_out[tid + s], s_out[tid]);
        }
        __syncthreads();
    }

    if (tid == 0)
        blockOutput[blockIdx.x] = s_out[0];
}

__global__ void minMaxScale_kernel(long long* d_gpu, long long* out, long long min, long long max, int n)
{
    const int tid = threadIdx.x;
    const int glbl_tid = blockDim.x * blockIdx.x + tid;

    if (glbl_tid < n)
    {
        out[glbl_tid] = FixedLongLong::div(d_gpu[glbl_tid] - min, max - min);
    }
}

__global__ void zScore_kernel(long long* d_gpu, long long* out, long long mean, long long std, int n)
{
    const int tid = threadIdx.x;
    const int glbl_tid = blockDim.x * blockIdx.x + tid;

    if (glbl_tid < n)
    {
        out[glbl_tid] = FixedLongLong::div(d_gpu[glbl_tid] - mean, std);
    }
}


__global__ void maxPoolingImplFixedLongLong_kernel(
    long long* inp, long long* out,
    int in_h, int in_w, int in_channel,
    int pool_size, int stride_h, int stride_w,
    int out_h, int out_w, 
    int padded_top, int padded_bottom, 
    int padded_left, int padded_right 
)
{
    const int out_row = blockIdx.y * blockDim.y + threadIdx.y; 
    const int out_col = blockIdx.x * blockDim.x + threadIdx.x;
    const int c       = blockIdx.z * blockDim.z + threadIdx.z;

    if (out_row < out_h && out_col < out_w && c < in_channel)
    {
        const int 
            idx = (out_row * out_w + out_col) * in_channel + c;
        
        int in_row = out_row * stride_h,
            in_col = out_col * stride_w;

        in_row = max(in_row, padded_top);
        in_col = max(in_col, padded_left);

        int pool_y = min(pool_size, in_h - in_row - padded_bottom),
            pool_x = min(pool_size, in_w - in_col - padded_right);

        out[idx] = inp[(in_row * in_w + in_col) * in_channel + c];

        for (int row = 0; row < pool_y; ++row)
        {
            const int rr = (in_row + row) * in_w;
            for (int col = 0; col < pool_x; ++col)
            {
                out[idx] = max(out[idx], inp[(rr + in_col + col) * in_channel + c]);
            }
        }
    }
}

__global__ void avgPoolingImplFixedLongLong_kernel(
    long long* inp, long long* out,
    int in_h, int in_w, int in_channel,
    int pool_size, int stride_h, int stride_w,
    int out_h, int out_w,
    int padded_top, int padded_bottom, 
    int padded_left, int padded_right 
)
{
    const int out_row = blockIdx.y * blockDim.y + threadIdx.y; 
    const int out_col = blockIdx.x * blockDim.x + threadIdx.x;
    const int c       = blockIdx.z * blockDim.z + threadIdx.z;

    if (out_row < out_h && out_col < out_w && c < in_channel)
    {
        long long sum = 0;
        const int 
            idx = (out_row * out_w + out_col) * in_channel + c;

        int in_row = out_row * stride_h,
            in_col = out_col * stride_w;

        in_row = max(in_row, padded_top);
        in_col = max(in_col, padded_left);

        int pool_y = min(pool_size, in_h - in_row - padded_bottom),
            pool_x = min(pool_size, in_w - in_col - padded_right);

        for (int row = 0; row < pool_size; ++row)
        {
            const int rr = (in_row + row) * in_w;
            for (int col = 0; col < pool_size; ++col)
            {
                sum = FixedLongLong::add(sum, inp[(rr + in_col + col) * in_channel + c]);
            }
        }

        out[idx] = FixedLongLong::div(sum, (1LL << 32) * (pool_x * pool_y));
    }
}


__global__ void sumReductionV2_kernel(long long* d_gpu, long long* blockOutput, int n, int c)
{
    const int tid = threadIdx.x;
    const int z = blockIdx.z * blockDim.z + threadIdx.z;
    const int glbl_tid = (2 * blockDim.x * blockIdx.x + tid) * c + z;

    extern __shared__ long long s_out[];

    s_out[tid] = 0;
    s_out[tid + blockDim.x] = 0;

    if (glbl_tid < n * c)
    {
        s_out[tid] = d_gpu[glbl_tid];
        if (glbl_tid + blockDim.x * c < n * c)
            s_out[tid + blockDim.x] = d_gpu[glbl_tid + blockDim.x * c];
    }
    __syncthreads();

    for (unsigned int s = blockDim.x; s > 0; s >>= 1) {
        if (tid < s) {
            s_out[tid] += s_out[tid + s];
        }
        __syncthreads();
    }

    if (tid == 0)
        blockOutput[blockIdx.x * c + z] = s_out[0];
}

__global__ void mat_add_fixed_longlong(long long *A, long long *B, long long *C, int n) {
    int x = threadIdx.x + blockIdx.x * blockDim.x; 

    if(x < n) {
        C[x] = A[x] + B[x];
    }
}

__global__ void mat_sub_fixed_longlong(long long *A, long long *B, long long *C, int n) {
    int x = threadIdx.x + blockIdx.x * blockDim.x; 	// Row address

    if(x < n) {
        C[x] = A[x] - B[x];
    }
}

__global__ void mat_mul_fixed_longlong(long long *A, long long *B, long long *C, int n) {
    int x = threadIdx.x + blockIdx.x * blockDim.x;	// Column Address

    if(x < n) {
        C[x] = FixedLongLong::mul(A[x], B[x]);
    }
}


__global__ void mat_div_fixed_longlong(long long *A, long long *B, long long *C, int n) {
    int x = threadIdx.x + blockIdx.x * blockDim.x;	// Column Address

    if(x < n && B[x] != 0) {
        C[x] = FixedLongLong::div(A[x], B[x]);
    }
}



__global__ void mat_add_single_fixed_longlong(long long *A, long long *B, long long e, int n) {
    int x = threadIdx.x + blockIdx.x * blockDim.x; 

    if(x < n) {
        B[x] = A[x] + e;
    }
}

__global__ void mat_sub_single_fixed_longlong(long long *A, long long *B, long long e, int n) {
    int x = threadIdx.x + blockIdx.x * blockDim.x; 	// Row address

    if(x < n) {
        B[x] = A[x] - e;
    }
}

__global__ void mat_mul_single_fixed_longlong(long long *A, long long *B, long long e, int n) {
    int x = threadIdx.x + blockIdx.x * blockDim.x;	// Column Address

    if(x < n) {
        B[x] = FixedLongLong::mul(A[x], e);
    }
}

__global__ void mat_pow2_single_fixed_longlong(long long *A, long long *B, int n) {
    int x = threadIdx.x + blockIdx.x * blockDim.x;	// Column Address

    if(x < n) {
        B[x] = FixedLongLong::mul(A[x], A[x]);
    }
}


__global__ void mat_div_single_fixed_longlong(long long *A, long long *B, long long e, int n) {
    int x = threadIdx.x + blockIdx.x * blockDim.x;	// Column Address

    if(x < n && B[x] != 0) {
        B[x] = FixedLongLong::div(A[x], e);
    }
}

__global__ void mat_sqrt_fixed_longlong(long long *A, long long *B, int n) {
    int x = threadIdx.x + blockIdx.x * blockDim.x;	// Column Address

    if(x < n) {
        B[x] = FixedLongLong::sqrt(A[x]);
    }
}

__global__ void mat_exp_fixed_longlong(long long *A, long long *B, int n) {
    int x = threadIdx.x + blockIdx.x * blockDim.x;	// Column Address

    if (x < n) {
        B[x] = FixedLongLong::exp(A[x]);
    }
}

__global__ void conv2dImplFixedLongLong_kernel(
    long long* inp, long long* kernel, long long* bias, long long* out, // data io
    int kernel_size, int in_channel, int out_channel, // kernel properties
    int in_w, int in_h, int out_w, int out_h, // spatial size of inp,
    int padding, int stride_h, int stride_w // padding mode, one of 'valid': 0 or 'same': 1
)
{
    int out_row = blockIdx.y * blockDim.y + threadIdx.y; 
    int out_col = blockIdx.x * blockDim.x + threadIdx.x;
    int out_c = blockIdx.z * blockDim.z + threadIdx.z;

    if (out_row < out_h && out_col < out_w)
    {
        const int in_row_offset = out_row * stride_h;
        const int in_col_offset = out_col * stride_w;
        int kernel_offset = out_c;
        
        long long sum = bias[out_c];

        for (int i = 0, in_row = in_row_offset * in_w
            ; i < kernel_size
            ; ++i, in_row += in_w
        )
        {
            for (int j = 0, inp_offset = (in_row + in_col_offset) * in_channel
                ; j < kernel_size
                ; ++j, inp_offset += in_channel
            )
            {
                for (int c = 0
                    ; c < in_channel
                    ; ++c, kernel_offset += out_channel
                )
                {
                    sum += FixedLongLong::mul(inp[inp_offset + c], kernel[kernel_offset]);
                }
            }
        }

        out[(out_row * out_w + out_col) * out_channel + out_c] = sum;
    }
}

__global__ void depthwise_conv2d_kernel(
    long long* inp, long long* kernel, long long* bias, long long* out, // data io
    int in_h, int in_w, int in_channel, 
    int kernel_size_h, int kernel_size_w, // kernel properties
    int out_h, int out_w, // spatial size of output,
    int padding, int stride_h, int stride_w // padding mode, one of 'valid': 0 or 'same': 1
)
{
    int out_row = blockIdx.y * blockDim.y + threadIdx.y; 
    int out_col = blockIdx.x * blockDim.x + threadIdx.x;
    int out_c = blockIdx.z * blockDim.z + threadIdx.z;

    if (out_row < out_h && out_col < out_w)
    {
        const int in_row_offset = out_row * stride_h;
        const int in_col_offset = out_col * stride_w;
        long long res = bias[out_c];

        for (int i = 0; i < kernel_size_h; ++i)
        {
            for (int j = 0; j < kernel_size_w; ++j)
            {
                res += FixedLongLong::mul(
                    inp[((in_row_offset + i) * in_w + in_col_offset + j) * in_channel + out_c],
                    kernel[(i * kernel_size_w + j) * in_channel + out_c]
                );
            }
        }

        out[(out_row * out_w + out_col) * in_channel + out_c] = res;
    }
}