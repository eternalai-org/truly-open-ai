function toSnakeCase(str: string): string {
    return str.replace(/[A-Z]/g, (letter) => `_${letter.toLowerCase()}`);
}

export const cloneAndSnakeCaseFields = <T extends Record<string, any>>(obj: T): T => {
    if (obj == null || obj == undefined) {
        return null
    }
    const transformKeysToSnakeCase = (input: any): any => {
        if (input && typeof input === "object" && !Array.isArray(input)) {
            return Object.entries(input).reduce((acc, [key, value]) => {
                const snakeCaseKey = toSnakeCase(key);
                acc[snakeCaseKey] = transformKeysToSnakeCase(value);
                return acc;
            }, {} as any);
        } else if (Array.isArray(input)) {
            return input.map(item => transformKeysToSnakeCase(item));
        } else {
            return input;
        }
    };
    return transformKeysToSnakeCase(obj);
}