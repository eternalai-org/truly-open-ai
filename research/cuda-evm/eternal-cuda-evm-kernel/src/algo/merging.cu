#include <operations.cuh>


void __concatenate(long long* inp, long long* out, long long* shapes, long long axis, long long ndims, long long n, uint8_t* error)
{
    long long* out_shape = new long long[ndims];

    if (!estimateConcatenate(shapes, axis, ndims, n, out_shape))
    {
       delete[] out_shape;
       return;
    }

    int out_h = 1, out_w = 1;
    for (int i = 0; i < axis; ++i)
    {
        out_h *= out_shape[i];
    }

    for (int i = axis; i < ndims; ++i)
    {
        out_w *= out_shape[i];
    }

    int magic = out_w / out_shape[axis];
    int i_offset = 0, o_offset = 0;
    long long* starts = new long long[n];

    for (int i = 0; i < n; ++i)
    {
        starts[i] = i_offset;
        i_offset += out_h * shapes[axis + ndims * i] * magic;
    }

    for (int i = 0, cnt; i < out_h; ++i)
    {
        for (int j = 0; j < n; ++j)
        {
            cnt = shapes[axis + ndims * j] * magic;
            memcpy(out + o_offset, inp + starts[j] + i * cnt, cnt << 3);
            o_offset += cnt;
        }
    }



    delete[] out_shape;
    delete[] starts;
}


void __concatenate_dummy(long long** inp, long long* out, long long** shapes, long long axis, long long ndims, long long n, uint8_t* error)
{
    long long* out_shape = new long long[ndims];

    if (!estimateConcatenate_dummy(shapes, axis, ndims, n, out_shape))
    {
       delete[] out_shape;
       return;
    }

    int out_h = 1, out_w = 1;
    for (int i = 0; i < axis; ++i)
    {
        out_h *= out_shape[i];
    }

    for (int i = axis; i < ndims; ++i)
    {
        out_w *= out_shape[i];
    }

    for (int i = 0, cnt, o_offset = 0, magic = out_w / out_shape[axis]
        ; i < out_h; ++i
    )
    {
        for (int j = 0; j < n; ++j)
        {
            cnt = shapes[j][axis] * magic;
            memcpy(out + o_offset, inp[j] + i * cnt, cnt << 3);
            o_offset += cnt;
        }
    }

    delete[] out_shape;
}

uint8_t estimateConcatenate(long long* shapes, long long axis, long long ndims, long long n, long long* out)
{
    if (ndims <= 0 || n <= 0 || axis < 0 || axis >= ndims)
    {
        return ERROR;
    }

    memset(out, 0x00, ndims * sizeof(long long));

    for (int i = 0; i < ndims; ++i)
    {
        if (i == axis)
        {
            for (int j = i; j < ndims * n; j += ndims)
            {
                out[i] += shapes[j];
            }
            continue;
        }

        out[i] = shapes[i];

        // verify
        for (int j = ndims + i; j < ndims * n; j += ndims)
        {
            if (shapes[i] != shapes[j])
            {
                return ERROR;
            }
        }
    }

    return OK;
}

uint8_t estimateConcatenate_dummy(long long** shapes, long long axis, long long ndims, long long n, long long* out)
{
    // @todo: fix this

    if (ndims <= 0 || n <= 0 || axis < 0 || axis >= ndims)
    {
        return ERROR;
    }

    memset(out, 0x00, ndims * sizeof(long long));

    for (int i = 0; i < n; ++i)
    {
        for (int j = 0; j < ndims; ++j)
        {
            if (j == axis)
            {
                out[j] += shapes[i][j];
            }
            else if (out[j] == 0)
            {
                out[j] = shapes[i][j];
            }
            else if (out[j] != shapes[i][j])
            {
                return ERROR;
            }
        }
    }

    return OK;
}
