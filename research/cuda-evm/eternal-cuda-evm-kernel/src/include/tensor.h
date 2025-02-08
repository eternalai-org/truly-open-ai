#ifndef __TENSOR_H__
#define __TENSOR_H__

#include <vector>
#include <stdint.h>
#include <cstring>
#include <assert.h>
#include <memory>
#include <numeric>
#include <iostream>
#include <helpers.cuh>

class  TensorWrapper {
    std::vector<uint64_t> mshape;
    int64_t* mdata;

public:
    TensorWrapper() : mshape({}), mdata(nullptr) {}

    TensorWrapper(const std::vector<uint64_t>& shape, const int64_t* data) : mshape(shape) { 
        uint64_t prod = 1;

        for (const auto& dim : mshape) {
            prod *= dim;
        }

        mdata = new int64_t[prod];
        memcpy(mdata, data, prod * sizeof(int64_t));
    }

    TensorWrapper(const TensorWrapper& other) {
        mshape = other.mshape;
        mdata = other.mdata;
    }

    TensorWrapper& assign(int64_t* data, const std::vector<uint64_t>& shape)
    {
        this->mdata = data;
        this->mshape = shape;
        return *this;
    }

    int dims()
    {
        return mshape.size();
    }

    std::vector<uint64_t> shape() const 
    {
        return mshape;
    }

    uint64_t size() const
    {
        return std::accumulate(mshape.begin(), mshape.end(), 1, std::multiplies<uint64_t>());
    }

    int64_t* data() const {
        return mdata;
    }

    TensorWrapper& reshape(std::vector<uint64_t> shape)
    {
        uint64_t prod_1 = 1, prod_2 = 1;

        for (const auto& dim : mshape) {
            prod_1 *= dim;
        }

        for (const auto& dim : shape) {
            prod_2 *= dim;
        }

        assert(prod_1 == prod_2);

        mshape = shape;
        return *this;
    }
};

std::ostream &operator << (std::ostream &s, const TensorWrapper& a) {
    s << "Tensor[" << a.shape() << "]: ";
    int prod = 1;
    for (const auto& x: a.shape())
    {
        prod *= x;
    }

    const int64_t* ref = a.data();

    for (int i = 0; i < prod; ++i)
    {
        s << ref[i] << " ";
    }

    return s;
}

#endif