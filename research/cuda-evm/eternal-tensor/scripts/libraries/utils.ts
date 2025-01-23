import { ethers } from "ethers";
import { RandomSeed } from "random-seed";

export function enumElementCount(enumName: any): number {
    let count = 0
    for(let item in enumName) {
        if(isNaN(Number(item))) count++;
    }
    return count
}

export function recursiveToString(arr: any): any {
    return arr.map((val: any) => Array.isArray(val) ? recursiveToString(val) : val.toString());
}

export function recursiveToFloat(arr: any): any {
    return arr.map((val: any) => Array.isArray(val) ? recursiveToFloat(val) : toFloat(val));
}

export function recursiveFromFloat(arr: any): any {
    return arr.map((val: any) => Array.isArray(val) ? recursiveFromFloat(val) : fromFloat(val));
}

export async function measureTime(f: any): Promise<any> {
  const start = Date.now();
  const ret = await f();
  const end = Date.now();
  console.log(`Execution time: ${(end - start) / 1000.0} s`);
  return ret
}

export function pixelsToImage(pixels: ethers.BigNumber[], h: number, w: number, c: number): ethers.BigNumber[][][] {
    let ptr = 0;
    let img: ethers.BigNumber[][][] = [];
    for(let i = 0; i < h; ++i) {
        img.push([]);
        for(let j = 0; j < w; ++j) {
            img[i].push([]);
            for(let k = 0; k < c; ++k) {
                img[i][j].push(pixels[ptr]);
                ++ptr;
            }
        }
    }
    return img;
}

export function fromFloat(num: any) {
    return ethers.BigNumber.from(String(Math.trunc(num * Math.pow(2, 32))));
}

export function fromInt(num: any) {
    return ethers.BigNumber.from(num).mul(ethers.BigNumber.from(2).pow(ethers.BigNumber.from(32)));
}

export function toInt(num: any) {
    return ethers.BigNumber.from(num).div(ethers.BigNumber.from(2).pow(ethers.BigNumber.from(32)));
}

export function toFloat(num: any) {
    return ethers.BigNumber.from(num).toNumber() / Math.pow(2, 32);
}

export function isBigNumberArrayEqual(arr1: any[], arr2: any[]): boolean {
    if (arr1.length != arr2.length) return false;
    for(let i = 0; i < arr1.length; ++i) {
        if (Array.isArray(arr1[i])) {
            if (!Array.isArray(arr2[i])) return false;
            if (!isBigNumberArrayEqual(arr1[i], arr2[i])) return false;
        } else {
            if (!arr1[i].eq(arr2[i])) return false;
        }
    }
    return true;
}

export function randomFloatArray(randomizer: RandomSeed, shape: number[], min: number, max: number): any {
    if (shape.length === 0) {
        return randomizer.floatBetween(min, max);
    }
    const arr = [];
    for(let i = 0; i < shape[0]; ++i) {
        arr.push(randomFloatArray(randomizer, shape.slice(1), min, max));
    }
    return arr;
}

export function isFloatArrayEqual(arr1: any[], arr2: any[], eps: number) {
    if (arr1.length != arr2.length) return false;
    for(let i = 0; i < arr1.length; ++i) {
        if (Array.isArray(arr1[i])) {
            if (!Array.isArray(arr2[i])) return false;
            if (!isFloatArrayEqual(arr1[i], arr2[i], eps)) return false;
        } else {
            if (Math.abs(arr1[i] - arr2[i]) > eps) return false;
        }
    }
    return true;
}

export function encodeBigNum64(num: ethers.BigNumber): ethers.BigNumber {
    return num.gte(0) ? num : (ethers.constants.Two.pow(64).add(num));
}

export function decodeBigNum64(num: ethers.BigNumber): ethers.BigNumber {
    return num.lt(ethers.constants.Two.pow(63)) ? num : (num.sub(ethers.constants.Two.pow(64)));
}

export function encodeData(data: ethers.BigNumber[]): ethers.BigNumber[] {
    const encoded = []
    for(let i = 0; i < data.length; i += 4) {
        let num = ethers.constants.Zero;
        for(let j = 0; j < 4; ++j) {
            if (i + j < data.length) {                
                num = num.or(encodeBigNum64(data[i+j]).shl(64 * (3 - j)));
            }
        }
        encoded.push(num);
    }
    return encoded;
}

export function decodeData(encoded: ethers.BigNumber[]): ethers.BigNumber[] {
    const data = []
    for(let i = 0; i < encoded.length; ++i) {
        for(let j = 0; j < 4; ++j) {
            const num = encoded[i].shr(64 * (3 - j)).and(ethers.constants.Two.pow(64).sub(1)); // (x >> (64 * j)) & (2^64 - 1)
            data.push(decodeBigNum64(num));
        }
    }
    return data;
}

export function deflatten(data: ethers.BigNumber[], shape: number[]): any[] {
    let idx = 0;
    const solve = (data: ethers.BigNumber[], shape: number[]): any => {
        if (shape.length === 0) return data[idx++];
        const res = [];
        for(let i = 0; i < shape[0]; ++i) {
            res.push(solve(data, shape.slice(1)));
        }
        return res;
    };

    return solve(data, shape);
}
