import { ethers } from "ethers";
import { decodeData, deflatten, encodeData, recursiveFromFloat, recursiveToFloat } from "./utils";

const abic = ethers.utils.defaultAbiCoder;

function getABIType(dimCount: number): string {
    if (dimCount == 1) return "int64[]";        
    if (dimCount == 2) return "int64[][]";        
    if (dimCount == 3) return "int64[][][]";        
    if (dimCount == 4) return "int64[][][][]";        
    throw new Error("Number of dimension not supported");
}

export class TensorData {
  data: string;
  dim: ethers.BigNumber[];

  constructor(data: string, dim: ethers.BigNumber[]) {
    this.data = data;
    this.dim = dim;
  }

  static fromArray(arr: any): TensorData {
    const dim = [];
    for(let x = arr; Array.isArray(x); x = x[0]) {
      dim.push(ethers.BigNumber.from(x.length));
    }
    const data = abic.encode([getABIType(dim.length)], [arr]);
    return new TensorData(data, dim);
  }

  encodeData(): string {
    return abic.encode(
      ['bytes', 'uint[]'],
      [this.data, this.dim],
    );
  }

  static decodeData(data: string): TensorData {
    const res = abic.decode(
      ['bytes', 'uint[]'],
      data
    )
    return new TensorData(res[0], res[1]);
  }    
}

export class Tensor {
  data: ethers.BigNumber[];
  shapes: ethers.BigNumber[];

  constructor(data: ethers.BigNumber[], shapes: ethers.BigNumber[]) {
    this.data = data;
    this.shapes = shapes;
  }

  static fromFloatArray(arr: any): Tensor {
    const shapes = [];
    for(let x = arr; Array.isArray(x); x = x[0]) {
      shapes.push(ethers.BigNumber.from(x.length));
    }
    const data = encodeData(recursiveFromFloat(arr).flat(Infinity));
    return new Tensor(data, shapes);
  }

  toFloatArray(): any {
    return recursiveToFloat(deflatten(decodeData(this.data), this.shapes.map(x => x.toNumber())))
  }
}
