#ifndef __FIXED_LL__
#define __FIXED_LL__

#include <cuda.h>
#include <assert.h>
#include <cstring>
#include <iostream>
#include <math.h>

namespace FixedLongLong {

class FixedLongLongType {
public:
    long long value;
    FixedLongLongType(long long value) : value(value) {}
    FixedLongLongType() : value(0) {}
};

    static const long long MIN = LLONG_MIN;
    static const long long MAX = LLONG_MAX;
    static const long long ONE = 1LL << 32;
    static const long long MAX_SQRT = 199032864766430;

    inline __device__ __host__ long long fromInt(const int &x) {
        return (long long) x << 32;
    }

    inline __device__ __host__ int toInt(const long long &x) {
        return x >> 32;
    }

    inline __device__ __host__ long long *fromHex(const char *bytes, int n) {
        long long *decimals = new long long[n];
        for (int i = 0, p = 0; i < n; ++i, p += 8) {
            decimals[i] = (long long) bytes[p]
                          | ((long long) bytes[p + 1] << 8)
                          | ((long long) bytes[p + 2] << 16)
                          | ((long long) bytes[p + 3] << 24)
                          | ((long long) bytes[p + 4] << 32)
                          | ((long long) bytes[p + 5] << 40)
                          | ((long long) bytes[p + 6] << 48)
                          | ((long long) (bytes[p + 7] & 127) << 56);
            if (bytes[p + 7] & 128) decimals[i] = -decimals[i];
        }
        return decimals;
    }

    inline __device__ __host__ char *toHex(const long long &decimal) {
        char *hex = new char[8];
        hex[0] = decimal & 255;
        hex[1] = (decimal >> 8) & 255;
        hex[2] = (decimal >> 16) & 255;
        hex[3] = (decimal >> 24) & 255;
        hex[4] = (decimal >> 32) & 255;
        hex[5] = (decimal >> 40) & 255;
        hex[6] = (decimal >> 48) & 255;
        hex[7] = (decimal >> 56) & 127;
        if (decimal < 0) hex[7] |= 128;
        return hex;
    }

    inline __device__ __host__ char *toHex(const long long *decimals, int n) {
        char *hex = new char[n << 3];
        for (int i = 0, p = 0; i < n; ++i, p += 8) {
            long long decimal = decimals[i];
            hex[p] = decimal & 255;
            hex[p + 1] = (decimal >> 8) & 255;
            hex[p + 2] = (decimal >> 16) & 255;
            hex[p + 3] = (decimal >> 24) & 255;
            hex[p + 4] = (decimal >> 32) & 255;
            hex[p + 5] = (decimal >> 40) & 255;
            hex[p + 6] = (decimal >> 48) & 255;
            hex[p + 7] = (decimal >> 56) & 127;
            if (decimal < 0) hex[p + 7] |= 128;
        }
        return hex;
    }

    inline __device__ __host__ long long add(const long long &x, const long long &y) {
        assert((y >= 0 && MAX - y >= x) || (y < 0 && MIN - y <= x));
        return x + y;
    }

    inline __device__ __host__ long long sub(const long long &x, const long long &y) {
        assert((y >= 0 && MIN + y <= x) || (y < 0 && MAX + y >= x));
        return x - y;
    }

    inline __device__ __host__ long long mul(const long long &x, const long long &y) {
        bool xSign = x > 0;
        bool ySign = y > 0;
        unsigned long long X = x == MIN ? 9223372036854775808 : abs(x);
        unsigned long long Y = y == MIN ? 9223372036854775808 : abs(y);

        unsigned long long xI = X >> 32;
        unsigned long long xF = X & UINT_MAX;
        unsigned long long yI = Y >> 32;
        unsigned long long yF = Y & UINT_MAX;

        unsigned long long xIy = xI*Y;
        if (xI != 0) assert(xIy / xI == Y);
        unsigned long long xFyI = xF*yI;
        unsigned long long xFyF = xF*yF >> 32;

        if (xSign == ySign) 
        {
            assert(xIy <= MAX);
            unsigned long long result = xIy;
            assert(MAX - result >= xFyI);
            result += xFyI;
            assert(MAX - result >= xFyF);
            result += xFyF;

            return result;
        } 
        else 
        {
            assert(xIy <= 9223372036854775808);
            unsigned long long result = xIy;
            assert(9223372036854775808 - result >= xFyI);
            result += xFyI;
            assert(9223372036854775808 - result >= xFyF);
            result += xFyF;

            return -((long long) result);
        }
    }

    inline __device__ __host__ long long reciprocal(const long long &x) {
        assert(x != 0);

        bool xSign = x > 0;
        unsigned long long X = x == MIN ? 9223372036854775808 : abs(x);
        unsigned long long result = ULLONG_MAX / X;

        if (X > 1 && ((X & (~X + 1)) == X))
        {
            ++result;
        }

        if (xSign) 
        {
            assert(result < 9223372036854775808);
            return result;
        } 
        else 
        {
            assert(result <= 9223372036854775808);
            return -(long long) result;
        }
    }

    inline __device__ __host__ long long div(const long long &x, const long long &y) {
        assert(y != 0);

        if (y == ONE) 
        {
            return x;
        }

        return mul(x, reciprocal(y));
    }

    inline __device__ __host__ long long sqrt(const long long& x) {
        assert(x >= 0);

        long long l = 1;
        long long r = MAX_SQRT;
        long long res = 0;

        while (l <= r) {
            long long mid = (l + r) >> 1;
            if (mul(mid, mid) <= x) 
            {
                res = mid;
                l = mid + 1;
            } 
            else 
            {
                r = mid - 1;
            }
        }

        return res;
    }

    // this function is indeterministic
    inline __device__ __host__ long long exp_2(const long long& x)
    {
        return (long long)(pow(2, 1.0f * x / (1LL << 32)) * (1LL << 32));
    }

    // this function is indeterministic
    inline __device__ __host__ long long exp(const long long& x)
    {
        return (long long)(expf(1.0f * x / (1LL << 32)) * (1LL << 32));
    }
}

#endif