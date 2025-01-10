import { createOpenAI } from "@ai-sdk/openai";
import {
    generateText as aiGenerateText,
} from "ai";
import { GenerateTextOptions } from "../types";

const generationTextOpenAi = async (params: GenerateTextOptions): Promise<string> => {
    const {
        apiKey,
        baseURL,
        model,
        prompt,
        systemContent,
        modelConfig
    } = params;

    const openai = createOpenAI({ apiKey, baseURL });
    const { text: openaiResponse } = await aiGenerateText({
        model: openai.languageModel(model),
        system: systemContent,
        prompt,
        ...(modelConfig || {}),
    });

    return openaiResponse;
}

export {
    generationTextOpenAi,
}
