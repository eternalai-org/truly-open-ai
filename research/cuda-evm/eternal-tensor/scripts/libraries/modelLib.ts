import { ethers } from "ethers";
import { fromFloat, enumElementCount } from "./utils";
import { RandomSeed } from "random-seed";

export enum LayerType {
    InputLayer,
    Dense,
    Flatten,
    Rescaling,
    MaxPooling2D,
    Conv2D,
    Embedding,
    SimpleRNN,
    LSTM,
    Softmax,
    Sigmoid,
    ReLU,
    Linear,
};

export enum InputType {
    Scalar,
    Tensor1D,
    Tensor2D,
    Tensor3D
};

export enum Activation {
    leakyrelu,
    linear,
    relu,
    sigmoid,
    tanh,
    softmax,
};

export enum Padding {
    valid,
    same
}

export enum ZeroPaddingFormat {
    ChannelsFirst,
    ChannelsLast
}

export const MaxLayerType = enumElementCount(LayerType);

class InputNode {
    name!: string;
    idx!: number;
    shape!: (number | null)[];
}

class InboundNodeData {
    args!: InputNode[];
    kwargs!: any;
}

class LayerData {
    idx!: number;
    name!: string;
    class_name!: string;
    layer_config!: any;
    inbound_nodes!: InboundNodeData[];
}

class LayerConfig {
    layerType: LayerType;
    address: String;
    inputIndices: ethers.BigNumber[];

    constructor(layerType: LayerType, address: String, inputIndices: ethers.BigNumber[]) {
        this.layerType = layerType;
        this.address = address;
        this.inputIndices = inputIndices;
    }
}

export function getLayerType(name: string): number {
    const layer = LayerType[name as keyof typeof LayerType];
    return (layer === undefined) ? -1 : layer;
}

export function getLayerName(type: number): string {
    return LayerType[type];
}

export function getActivationType(name: string): number {
    const activation = Activation[name as keyof typeof Activation];
    return (activation === undefined) ? -1 : activation;
}

export function getPaddingType(name: string): number {
    const padding = Padding[name as keyof typeof Padding];
    return (padding === undefined) ? -1 : padding;
}

export function getConvSize(
    dim: number[],
    size: number[],
    stride: number[],
    padding: string,
): { 
    out: number[], 
    pad: number[] 
} {
    const out = [], pad = [];
    for(let i = 0; i < 2; ++i) {
        if (padding == "same") {
            out.push((dim[i] + stride[i] - 1) / stride[i]);
            const total_pad = (dim[i] % stride[i] == 0) ? Math.max(size[i] - stride[i], 0) : Math.max(size[i] - dim[i] % stride[i], 0);
            pad.push(total_pad / 2);
        } else if (padding == "valid") {
            // TODO: What if dim[i] < size[i]
            out.push((dim[i] - size[i]) / stride[i] + 1);
            pad.push(0);
        }
    }
    return { out, pad };
}

export function getLayerConfigNew(
    layers: LayerData[], 
): {
    layerConfigs: LayerConfig[],
    totalWeights: number,
} {
    const abic = ethers.utils.defaultAbiCoder;

    let layerConfigs = [];
    let totalWeights = 0;
    for (let i = 0; i < layers.length; i++) {
        const layer = layers[i];
        
        let configData: String = "";        
        const layerType = getLayerType(layer.class_name);
        console.log(`Layer ${i}: ${layer.class_name}, type: ${layerType}`);
        const inputIndices: number[] = layer.inbound_nodes.map((node: InboundNodeData) => node.args[0].idx);
        
        if (layer.class_name === 'Dense') {
            const inputNode = layer.inbound_nodes[0].args[0];
            const inputUnits = inputNode.shape[1]!;
            const units = layer.layer_config.units;
            let activationFn: number = getActivationType(layer.layer_config.activation);
            configData = abic.encode(["uint8", "uint256", "uint256"], [activationFn, ethers.BigNumber.from(units), ethers.BigNumber.from(inputUnits)]);
            totalWeights += inputUnits * units + units;
        } else if (layer.class_name === 'Flatten') {
            configData = abic.encode([], []);
        } else if (layer.class_name === 'Rescaling') {
            const n1 = fromFloat(layer.layer_config.scale);
            const n2 = fromFloat(layer.layer_config.offset);
            configData = abic.encode(["int64", "int64"], [n1, n2]);
        } else if (layer.class_name === 'Softmax'){
            configData = abic.encode([], []);
        } else if (layer.class_name === 'ReLU'){
            const max_value = fromFloat(layer.layer_config.max_value);
            const negative_slope = fromFloat(layer.layer_config.negative_slope);
            const threshold = fromFloat(layer.layer_config.threshold);
            configData = abic.encode(["int64", "int64", "int64"], [max_value, negative_slope, threshold]);
        } else if (layer.class_name === 'Sigmoid'){
            configData = abic.encode([], []);
        } else if (layer.class_name === 'Linear'){
            configData = abic.encode([], []);
        } else if (layer.class_name === 'InputLayer') {            
            let dim: (number | null)[] = layer.layer_config.batch_input_shape;
            let pos = dim.lastIndexOf(null);
            dim = dim.slice(pos + 1);
            console.log("inputDim: ", dim);
            if (dim.length == 0) {
                configData = abic.encode(["uint8"], [InputType.Scalar]);
            } else if (dim.length == 1) {
                const n = ethers.BigNumber.from(dim[0]);
                configData = abic.encode(["uint8", "uint[]"], [InputType.Tensor1D, [n]]);
            } else if (dim.length == 2) {
                const n = ethers.BigNumber.from(dim[0]);
                const m = ethers.BigNumber.from(dim[1]);
                configData = abic.encode(["uint8", "uint[]"], [InputType.Tensor2D, [n, m]]);
            } else if (dim.length == 3) {
                const n = ethers.BigNumber.from(dim[0]);
                const m = ethers.BigNumber.from(dim[1]);
                const p = ethers.BigNumber.from(dim[2]);
                configData = abic.encode(["uint8", "uint[]"], [InputType.Tensor3D, [n, m, p]]);
            }
        } else if (layer.class_name === 'MaxPooling2D') {
            const f_w = layer.layer_config.pool_size[0];
            const f_h = layer.layer_config.pool_size[1];
            const s_w = layer.layer_config.strides[0];
            const s_h = layer.layer_config.strides[1];
            const padding = layer.layer_config.padding;
    
            configData = abic.encode(["uint[2]", "uint[2]", "uint8"], [
                [ethers.BigNumber.from(f_w), ethers.BigNumber.from(f_h)],
                [ethers.BigNumber.from(s_w), ethers.BigNumber.from(s_h)],
                getPaddingType(padding),
            ]);
        } else if (layer.class_name === 'Conv2D') {
            const inputNode = layer.inbound_nodes[0].args[0];
            const inputFilters = inputNode.shape[3]!;
            const outputFilters = layer.layer_config.filters;
            const f_w = layer.layer_config.kernel_size[0];
            const f_h = layer.layer_config.kernel_size[1];
            const s_w = layer.layer_config.strides[0];
            const s_h = layer.layer_config.strides[1];
            const padding = layer.layer_config.padding;
            let activationFn: number = getActivationType(layer.layer_config.activation);

            configData = abic.encode(["uint8", "uint", "uint", "uint[2]", "uint[2]", "uint8"], [
                activationFn,
                ethers.BigNumber.from(inputFilters),
                ethers.BigNumber.from(outputFilters),
                [ethers.BigNumber.from(f_w), ethers.BigNumber.from(f_h)],
                [ethers.BigNumber.from(s_w), ethers.BigNumber.from(s_h)],
                getPaddingType(padding),
            ]);
            totalWeights += f_w * f_h * inputFilters * outputFilters + outputFilters;
        } else if (layer.class_name === 'Embedding') {
            let inputDim = layer.layer_config.input_dim;
            let outputDim = layer.layer_config.output_dim;
            configData = abic.encode(["uint256", "uint256"], [ethers.BigNumber.from(inputDim), ethers.BigNumber.from(outputDim)]);
            totalWeights += inputDim * outputDim;
        } else if (layer.class_name === 'SimpleRNN') {
            const inputNode = layer.inbound_nodes[0].args[0];
            const inputUnits = inputNode.shape[2]!;
            const units = layer.layer_config.units;
            const activationFn: number = getActivationType(layer.layer_config.activation);
            configData = abic.encode(["uint8", "uint256"], [activationFn, ethers.BigNumber.from(units)]);
            totalWeights += inputUnits * units + units * units + units;
        } else if (layer.class_name === 'LSTM') {
            const inputNode = layer.inbound_nodes[0].args[0];
            const inputUnits = inputNode.shape[2]!;
            const units = layer.layer_config.units;
            const activationFn: number = getActivationType(layer.layer_config.activation);
            const recActivationFn: number = getActivationType(layer.layer_config.recurrent_activation);
            configData = abic.encode(["uint8", "uint8", "uint256", "uint256"], [activationFn, recActivationFn, ethers.BigNumber.from(units), ethers.BigNumber.from(inputUnits)]);
            totalWeights += inputUnits * units * 4 + units * units * 4 + units * 4;
        }

        layerConfigs.push(new LayerConfig(
            layerType,
            configData,
            inputIndices.map(x => ethers.BigNumber.from(x)),
        ));
    }

    return {
        layerConfigs,
        totalWeights
    };
}

export async function uploadModelWeights(model: ethers.Contract, weights: ethers.BigNumber[], maxlen: number) {    
    const weightStr = JSON.stringify(weights);
    console.log(`Weights size: ${weights.length}, total length: ${weightStr.length}`);
    let txIdx = 0;
    while (weights.length > 0) {
        const weightsToUpload = weights.splice(0, maxlen);
        const appendWeightTx = await model.appendWeights(weightsToUpload);
        console.log(`Append weights #${txIdx}`);
        const receipt = await appendWeightTx.wait();
        console.log(`tx ${appendWeightTx.hash}, gas used: ${receipt.gasUsed}`);
        txIdx += 1;
    }
}

export function getRandomActivation(randomizer: RandomSeed) {
    const activations = [
        Activation.relu,
        Activation.sigmoid,
        Activation.tanh,
    ]
    const x = randomizer.intBetween(0, 2);
    return activations[x];
}
