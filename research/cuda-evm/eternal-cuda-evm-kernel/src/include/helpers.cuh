#ifndef __UTILITIES_H__
#define __UTILITIES_H__

#include <stdio.h>
#include <cuda.h>
#include <cuda_runtime.h>
#include <vector>
#include <iostream>
#include <fixedlonglong32x32.cuh>

template<class T>
class array_view {
public:
    T *begin, *end;
    array_view(T *begin, T *end) : begin(begin), end(end) {}
    array_view(T *begin, int size) : begin(begin), end(begin + size) {}
    array_view(const std::vector<T> &v) : begin(v.data()), end(v.data() + v.size()) {}
};

template <class T> 
std::ostream &operator << (std::ostream &s, const std::vector<T> &a) {
	s << "[";

    for (const auto& x: a)
    {
        s << x << " ";
    }

	return s << "]";
}

template <class T> 
std::ostream &operator << (std::ostream &s, const array_view<T>& a) {
	for (T* it = a.begin; it != a.end; ++it)
    {
        s << *it;

        if (it != a.end - 1)
        {
            s << ", ";
        }
    }
    
    return s;
}

void printmat3d(long long* mat, int h, int w, int c);
std::ostream &operator << (std::ostream &s, const cudaError_t &a);
std::ostream &operator << (std::ostream &s, const FixedLongLong::FixedLongLongType &a);
bool cuda_fmt_error(const cudaError_t &a);

#endif // __UTILITIES_H__