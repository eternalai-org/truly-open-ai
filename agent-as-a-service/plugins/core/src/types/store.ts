export interface Tool {
    headers: Record<string, string>;
    method: string;
    label: string;
    executor: string;
    name: string;
    description: string;
    params: Array<Object>;
}

export interface StoreMission {
    name: string;
    prompt: string;
    price: number;
    tool_list: Array<Tool>
}

export interface Store {
    info: {
        owner_address: string;
        name: string;
        description: string;
        authen_url: string;
        type: string;
        icon: string;
    }
    missions: Array<StoreMission>;
}