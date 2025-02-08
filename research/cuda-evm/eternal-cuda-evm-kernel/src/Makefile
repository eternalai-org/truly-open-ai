cu_headers = # -I/usr/local/cuda/include
HEADERS = -Iinclude $(cu_headers)
OPTIMIZATION = -O2
COMPILE_FLAGS = -Xcompiler="-fpermissive"
CPP_VERSION = -std=c++17
_CPP_SYMBOLS = -DLOGGING_DEBUG -DLOGGING_VERBOSE

ifeq ($(debug), 1)
	CPP_SYMBOLS=-DLOGGING_DEBUG -DLOGGING_VERBOSE
	
else
	CPP_SYMBOLS=-DNDEBUG
endif 

# CPP_SYMBOLS=-DLOGGING_DEBUG -DLOGGING_VERBOSE

kernels.o: algo/kernels.cu
	nvcc --compiler-options -fPIC -c -o kernels.o  algo/kernels.cu \
	$(HEADERS) $(OPTIMIZATION) $(COMPILE_FLAGS) $(CPP_VERSION) $(CPP_SYMBOLS)

matrixmul.o: algo/matrixmul.cu
	nvcc --compiler-options -fPIC -c -o matrixmul.o  algo/matrixmul.cu \
	$(HEADERS) $(OPTIMIZATION) $(COMPILE_FLAGS) $(CPP_VERSION) $(CPP_SYMBOLS)

element_wise_operator.o:  algo/element_wise_operator.cu
	nvcc --compiler-options -fPIC -c -o element_wise_operator.o  algo/element_wise_operator.cu \
	$(HEADERS) $(OPTIMIZATION) $(COMPILE_FLAGS) $(CPP_VERSION) $(CPP_SYMBOLS)

transforms.o:  algo/transforms.cu
	nvcc --compiler-options -fPIC -c -o transforms.o  algo/transforms.cu \
	$(HEADERS) $(OPTIMIZATION) $(COMPILE_FLAGS) $(CPP_VERSION) $(CPP_SYMBOLS)

normalization.o:  algo/normalization.cu
	nvcc --compiler-options -fPIC -c -o normalization.o  algo/normalization.cu \
	$(HEADERS) $(OPTIMIZATION) $(COMPILE_FLAGS) $(CPP_VERSION) $(CPP_SYMBOLS)

conv.o:  algo/conv.cu
	nvcc --compiler-options -fPIC -c -o conv.o  algo/conv.cu \
	$(HEADERS) $(OPTIMIZATION) $(COMPILE_FLAGS) $(CPP_VERSION) $(CPP_SYMBOLS)

pooling.o:  algo/pooling.cu
	nvcc --compiler-options -fPIC -c -o pooling.o  algo/pooling.cu \
	$(HEADERS) $(OPTIMIZATION) $(COMPILE_FLAGS) $(CPP_VERSION) $(CPP_SYMBOLS)

activations.o:  algo/activations.cu
	nvcc --compiler-options -fPIC -c -o activations.o  algo/activations.cu \
	$(HEADERS) $(OPTIMIZATION) $(COMPILE_FLAGS) $(CPP_VERSION) $(CPP_SYMBOLS)

merging.o:  algo/merging.cu
	nvcc --compiler-options -fPIC -c -o merging.o  algo/merging.cu \
	$(HEADERS) $(OPTIMIZATION) $(COMPILE_FLAGS) $(CPP_VERSION) $(CPP_SYMBOLS)

reduction.o:  algo/reduction.cu
	nvcc --compiler-options -fPIC -c -o reduction.o  algo/reduction.cu \
	$(HEADERS) $(OPTIMIZATION) $(COMPILE_FLAGS) $(CPP_VERSION) $(CPP_SYMBOLS)

wrapper.o: deprecated/algo/wrapper.cu
	nvcc --compiler-options -fPIC -c -o wrapper.o deprecated/algo/wrapper.cu \
	$(HEADERS) $(OPTIMIZATION) $(COMPILE_FLAGS) $(CPP_VERSION) $(CPP_SYMBOLS)

computelib.o: computelib.cu
	nvcc --compiler-options -fPIC -c -o computelib.o computelib.cu \
	$(HEADERS) $(OPTIMIZATION) $(COMPILE_FLAGS) $(CPP_VERSION) $(CPP_SYMBOLS)

helpers.o: utils/helpers.cu
	nvcc --compiler-options -fPIC -c -o helpers.o utils/helpers.cu \
	$(HEADERS) $(OPTIMIZATION) $(COMPILE_FLAGS) $(CPP_VERSION) $(CPP_SYMBOLS)

libcomputelib.so: matrixmul.o element_wise_operator.o normalization.o conv.o pooling.o \
	activations.o merging.o kernels.o reduction.o wrapper.o computelib.o helpers.o transforms.o
	nvcc --ptxas-options=-v -o libcomputelib.so --shared matrixmul.o element_wise_operator.o normalization.o \
	conv.o pooling.o activations.o merging.o kernels.o reduction.o wrapper.o computelib.o helpers.o transforms.o \
	$(HEADERS) $(OPTIMIZATION) $(COMPILE_FLAGS) $(CPP_VERSION) $(CPP_SYMBOLS)


main: main.cu libcomputelib.so helpers.o
	nvcc -o main main.cu helpers.o -L. -lcomputelib 
	$(HEADERS) $(OPTIMIZATION) $(CPP_SYMBOLS)

exec: main
	LD_LIBRARY_PATH=$(LD_LIBRARY_PATH):$(pwd) ./main

test: libcomputelib.so
	python ../test/run.py

debug: libcomputelib.so
	python ../test/debug.py


release: clean libcomputelib.so
	if [ -d "build" ]; then rm -rf build; fi
	mkdir build
	mkdir build/include
	mv libcomputelib.so build
	cp include/operations.cuh build/include
	cp include/computelib.h build/include

git_release: release
	rm -rf ../release
	mv build ../release
	git add ../release
	git commit -m "[automatic] Release $(shell date)"
	git push

clean:
	rm -f *.o *.so
