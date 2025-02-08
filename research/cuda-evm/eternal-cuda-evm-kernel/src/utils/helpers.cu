#include <stdio.h>
#include <helpers.cuh>

void printmat3d(long long* mat, int h, int w, int c)
{
    for (int i = 0; i < h; ++i)
    {
        for (int j = 0; j < w; ++j)
        {
            printf("[");
            for (int k = 0; k < c; ++k)
            {
                printf("%lf ", 1.0f * mat[i * w * c + j * c + k] / (1LL << 32));
            }
            printf("], ");
        }
        printf("\n");
    }
}


bool cuda_fmt_error(const cudaError_t &a)
{
    if (a != cudaSuccess)
    {

#if defined(LOGGING_DEBUG)
        std::cerr << "CUDA Error: " << a << std::endl;
#endif

        return 1;
    }

    return 0;
}

std::ostream &operator << (std::ostream &s, const cudaError_t &a)
{
#if defined(LOGGING_VERBOSE)
    const char* error = cudaGetErrorString(a);
    s << error;
#else
    if (a != cudaSuccess)
        s << int(a);
#endif // LOGGING_VERBOSE

    return s;
}

std::ostream &operator << (std::ostream &s, const FixedLongLong::FixedLongLongType &a) {
	return s << 1.0f * a.value / FixedLongLong::ONE;
};