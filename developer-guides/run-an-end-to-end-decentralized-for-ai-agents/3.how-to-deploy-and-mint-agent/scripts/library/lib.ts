import fs from "fs";
import fsPromises from "fs/promises";
import path from "path";

export const WORKERHUB_INFERENCE_DIR = "migrate-data/inference";
export const WORKERHUB_ASSIGNMENT_DIR = "migrate-data/assignment";
export const WORKERHUB_MINER_DIR = "migrate-data/miner";
export const WORKERHUB_UNSTAKE_REQUEST_DIR = "migrate-data/unstake_request";

export const WORKERHUB_HYBRID_MODELS_PATH = "migrate-data/newHybridModelsData.json";

export const MIGRATION_PROGRESS_DIR = "migrate-progress";


export function stringifyJSON(obj: any) {
    return JSON.stringify(obj, (_, v) => typeof v === 'bigint' ? v.toString() : v);
}

export function saveFile(dir: string, name: string, content: string) {
    if (!fs.existsSync(dir)) {
        fs.mkdirSync(dir, { recursive: true });
    }
    fs.writeFileSync(path.join(dir, name), content);
}

export async function getJsonPathFromDir(dir: string): Promise<string[]> {
    const files = await fs.promises.readdir(dir, {withFileTypes: true});
    return files
        .filter(item => !item.isDirectory() && item.name.endsWith(".json"))
        .map(item => item.name);
}

export function ceilDiv(a: number, b: number) {
    return Math.ceil(a / b);
}

export async function readJSONFromFile(path: string) {
    return JSON.parse((await fsPromises.readFile(path)).toString());
}