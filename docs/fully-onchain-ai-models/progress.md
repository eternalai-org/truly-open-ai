# Progress

## The long term

Our roadmap is centered around a singular goal: to preserve as many AI models on-chain as possible. To achieve this, we basically need to solve three key Crypto x AI problems over time:

1. **On-chain model size limit:** the bigger, the better
2. **On-chain AI computation:** the faster, the better
3. **On-chain model architectures:** the more supported, the better

### 2023: Minimum viable protocol with 1 MB model size

In our first year, we focused on building the foundational components for Eternal AI and delivering a minimum viable protocol.

:white\_check\_mark: **On-chain model size limit:** 1 MB

:white\_check\_mark: **On-chain AI computation:** CPU

:white\_check\_mark: **On-chain model architectures:** Perceptrons

### 2024: Production-grade infrastructure with 100 MB model size

In our second year, we aim to scale Eternal AI from a minimum viable protocol to a production-grade computing infrastructure.

* [ ] **On-chain model size limit:** 100 MB
* [ ] **On-chain AI computation:** GPU with CUDA
* [ ] **On-chain model architectures:** CNN, RNN,  LSTM, Autoencoders, Transformers

### 2025: Scale to 1 GB model size

In our third year, we plan to further scale our capabilities.

* [ ] **On-chain model size:** 1 GB
* [ ] **On-chain AI computation:** TBD
* [ ] **On-chain model architectures:** TBD

## The short term

Our immediate goal is to scale the model size limit from 1 MB to 100 MB. Achieving this milestone will significantly enhance the ability of AI developers to deploy more sophisticated models on-chain.

**At the deep learning library level**, we are expanding the Eternal AI Toolkit by implementing additional smart contracts to support a wider variety of neural network layers and architectures.

**At the computation level**, we are enhancing the VM by integrating additional CUDA instructions, thereby increasing the capacity for more complex on-chain computations.

<table data-full-width="false"><thead><tr><th>When</th><th width="250">On-chain layers to support</th><th width="209">Model examples</th><th># Parameters</th></tr></thead><tbody><tr><td>Jun 2024</td><td>Activation</td><td>Lenet5</td><td>0.06M</td></tr><tr><td>Jul 2024</td><td><p>Concatenate</p><p>AveragePooling2D</p><p>BatchNormalization</p><p>Activation</p><p>ZeroPadding2D</p><p>GlobalAveragePooling2D</p></td><td>DenseNet121</td><td>7.2M</td></tr><tr><td>Aug 2024</td><td><p>BatchNormalization</p><p>DepthwiseConv2D</p><p>ZeroPadding2D</p><p>Reshape</p><p>Activation</p><p>GlobalAveragePooling2D</p><p>Dropout</p></td><td>MobileNet</td><td>V1: 4.2M<br>V2: 3.4M<br>V3small: 2.9M<br>V3large: 5.4M</td></tr><tr><td>Aug 2024</td><td><p>Normalization</p><p>ZeroPadding2D</p><p>Activation</p><p>DepthwiseConv2D</p><p>GlobalAveragePooling2D</p><p>Reshape</p><p>Multiply</p><p>Dropout</p></td><td>EfficientNet</td><td>B0: 5.3M<br>B1: 7.8M</td></tr><tr><td>Aug 2024</td><td>Activation<br>SeparableConv2D<br>Cropping2D</td><td>NASNetMobile</td><td>5.6M</td></tr><tr><td>TBD</td><td></td><td>TinyStories</td><td>5M</td></tr><tr><td>TBD</td><td></td><td>Tokenizer</td><td></td></tr><tr><td>TBD</td><td></td><td>Squeezenet</td><td></td></tr><tr><td>TBD</td><td></td><td>GPT-2 (S)</td><td>200M</td></tr><tr><td>TBD</td><td></td><td>GPT-2 (XL)</td><td>1500M</td></tr><tr><td>TBD</td><td></td><td>Stable Diffusion</td><td></td></tr></tbody></table>
