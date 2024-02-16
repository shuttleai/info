# ShuttleAI API

## Introduction

Are you a developer interested in creating a AI bot/site, but you find GPT-4 to be too expensive or you're on a tight budget? Look no further!
ShuttleAI API provides `gpt-4` and 100+ other models for free with a super **stable** and **free** api!
## Getting Started

To obtain a free ShuttleAI API key, follow these steps:

1. Join [our Discord server](https://discord.gg/shuttleai).
2. Run `/getkey` command in the Discord server.
3. Verify yourself.

You can view ShuttleAI API docs for better instructions on how to interact with our API at https://docs.shuttleai.app/

## ShuttleAI Models:
```json
{
    "data": [
        {
            "id": "gpt-4-turbo-preview",
            "object": "proxy",
            "proxy_to": "gpt-4-0125-preview",
            "owned_by": "openai",
            "tokens": 128000,
            "info": "Always proxies to latest gpt-4-turbo-preview model.",
            "premium": true,
            "cost": 4,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-4-0125-preview",
            "object": "model",
            "owned_by": "openai",
            "tokens": 128000,
            "info": "New (laziness reduced) GPT-4-Turbo-Preview 128k",
            "premium": true,
            "cost": 4,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-4-1106-preview",
            "object": "model",
            "owned_by": "openai",
            "tokens": 128000,
            "info": "GPT-4-Turbo-Preview 128k",
            "premium": true,
            "cost": 4,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-4-vision-preview",
            "object": "model",
            "owned_by": "openai",
            "tokens": 128000,
            "info": "GPT-4-Vision-Preview 128k",
            "premium": true,
            "cost": 4,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-4",
            "object": "model",
            "owned_by": "openai",
            "tokens": 8192,
            "info": "Plain gpt-4, directly from openai enterprise.",
            "cost": 2,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-4-bing",
            "object": "model",
            "owned_by": "bing/openai",
            "tokens": 29000,
            "info": "Bing GPT-4 Turbo. Bing web search is enabled by default. Send `internet: false` to disable.",
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-4-32k",
            "object": "model",
            "owned_by": "openai",
            "tokens": 32000,
            "info": "Free provider pool of real GPT-4-32k, falls back to Bing GPT-4 32k if free pool is out. Premium pool found @ 32k-0613.",
            "cost": 3,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-4-32k-0613",
            "object": "model",
            "owned_by": "openai",
            "tokens": 32768,
            "premium": true,
            "cost": 3,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-4-0613",
            "object": "model",
            "owned_by": "openai",
            "tokens": 8192,
            "cost": 2,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-4-study",
            "object": "model",
            "owned_by": "openai",
            "tokens": 8192,
            "info": "Raw GPT-4; great for students.",
            "premium": true,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-3.5-turbo",
            "object": "model",
            "owned_by": "openai",
            "tokens": 4097,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-3.5-turbo-0125",
            "object": "model",
            "owned_by": "openai",
            "tokens": 16385,
            "premium": true,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-3.5-turbo-1106",
            "object": "model",
            "owned_by": "openai",
            "tokens": 16385,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-3.5-turbo-0613",
            "object": "model",
            "owned_by": "openai",
            "tokens": 4097,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-3.5-turbo-0301",
            "object": "model",
            "owned_by": "openai",
            "tokens": 4097,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-3.5-turbo-16k",
            "object": "model",
            "owned_by": "openai",
            "tokens": 16384,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-3.5-turbo-16k-0613",
            "object": "proxy",
            "proxy_to": "gpt-3.5-turbo-1106",
            "owned_by": "openai",
            "tokens": 16385,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "chatgpt",
            "object": "model",
            "owned_by": "openai",
            "tokens": 8192,
            "info": "official chatgpt (GPT 3.5) aka openai's fastest and most capable free-to-use model aka text-davinci-002-render-sha",
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "code-gpt",
            "object": "model",
            "owned_by": "shuttle/openai",
            "tokens": 16385,
            "info": "GPT-3.5-Turbo-1106 with added Real-Time Live Code Execution/Interpretation",
            "premium": true,
            "file_upload": true,
            "cost": 3,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "internet-gpt",
            "object": "model",
            "owned_by": "shuttle/openai",
            "tokens": 16385,
            "info": "GPT-3.5-Turbo-1106 with added Real-Time Live Web Browsing",
            "premium": true,
            "file_upload": true,
            "cost": 2,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "claude-instant",
            "object": "model",
            "owned_by": "anthropic",
            "tokens": 25000,
            "file_upload": true,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "claude-instant-100k",
            "object": "model",
            "owned_by": "anthropic",
            "tokens": 100000,
            "premium": true,
            "file_upload": true,
            "cost": 3,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "claude-2.0",
            "object": "model",
            "owned_by": "anthropic",
            "tokens": 100000,
            "premium": true,
            "cost": 4,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "claude-2.1",
            "object": "model",
            "owned_by": "anthropic",
            "tokens": 200000,
            "premium": true,
            "cost": 4,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "palm-2",
            "object": "model",
            "owned_by": "google",
            "tokens": 8192,
            "cost": 2,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gemini-pro",
            "object": "model",
            "owned_by": "google",
            "tokens": 30720,
            "info": "very jailbreakable, works best when sending approval assistant message after jailbreak.",
            "premium": true,
            "cost": 2,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gemini-pro-vision",
            "object": "model",
            "owned_by": "google",
            "tokens": 12288,
            "info": "Gemini Pro Vision, pass an `image` variable with your image (url, bytes, base64).",
            "premium": true,
            "cost": 3,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "mistral-7b",
            "object": "model",
            "owned_by": "mistralai",
            "tokens": 4096,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "mistral-medium",
            "object": "model",
            "owned_by": "mistralai",
            "tokens": 32768,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "mixtral-8x7b",
            "object": "model",
            "owned_by": "mistralai",
            "tokens": 32768,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "dolphin-mixtral-8x7b",
            "object": "model",
            "owned_by": "cognitivecomputations",
            "tokens": 4096,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "pplx-70b-online",
            "object": "model",
            "owned_by": "perplexitylabs",
            "tokens": 32768,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "pplx-70b-chat",
            "object": "model",
            "owned_by": "perplexitylabs",
            "tokens": 32768,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "pplx-7b-online",
            "object": "model",
            "owned_by": "perplexitylabs",
            "tokens": 32768,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "pplx-7b-chat",
            "object": "model",
            "owned_by": "perplexitylabs",
            "tokens": 32768,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "deepseek-coder",
            "object": "model",
            "owned_by": "deepseek",
            "tokens": 4096,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "deepseek-chat",
            "object": "model",
            "owned_by": "deepseek",
            "tokens": 4096,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "airoboros-70b",
            "object": "model",
            "owned_by": "jondurbin",
            "tokens": 6105,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "falcon-180b",
            "object": "model",
            "owned_by": "tiiuae",
            "tokens": 6105,
            "maintenance": true,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "llama-2-70b-chat",
            "object": "model",
            "owned_by": "metaai",
            "tokens": 6105,
            "info": "Llama 2 70b Chat Hf",
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "llama-2-13b-chat",
            "object": "model",
            "owned_by": "metaai",
            "tokens": 2048,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "llama-2-7b-chat",
            "object": "model",
            "owned_by": "metaai",
            "tokens": 2048,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "code-llama-70b",
            "object": "model",
            "owned_by": "metaai",
            "tokens": 6105,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "code-llama-34b",
            "object": "model",
            "owned_by": "metaai",
            "tokens": 6105,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "llava-13b",
            "object": "model",
            "owned_by": "metaai",
            "tokens": 2048,
            "info": "Llava 13b, either send a 'image' paramater with your model and messages or send an image in the last message.",
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "sdxl",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 2,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "sdxl-turbo",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 2,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "sdxl-emoji",
            "object": "model",
            "owned_by": "stabilityai/fofr",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 2,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "dreamshaper-xl",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 2,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "juggernaut-xl",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 2,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "dynavision-xl",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 2,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "realism-engine-xl",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 2,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "sdxl-inpainting",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 4,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "turbovision-xl",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 2,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "kandinsky-2.2",
            "object": "model",
            "owned_by": "sberbank",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "kandinsky-2",
            "object": "model",
            "owned_by": "sberbank",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "stable-diffusion 2.1",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "stable-diffusion 1.5",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "dall-e-3",
            "object": "model",
            "owned_by": "openai",
            "info": "OpenAI's DALL-E 3. Does support the `n` paramter.",
            "max_images": 4,
            "multiple_of": 1,
            "premium": true,
            "cost": 5,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "dalle3",
            "object": "model",
            "owned_by": "bing/openai",
            "info": "ShuttleAI does not support changing the `n` paramater for dalle3 generations. However much Bing gens is how much you will get.",
            "max_images": 4,
            "multiple_of": 4,
            "cost": 10,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "midjourney",
            "object": "model",
            "owned_by": "davidholz",
            "max_images": 4,
            "multiple_of": 4,
            "premium": true,
            "cost": 369,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "dreamshaper-v8",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "rev-animated",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "anything-v5",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "absolutereality-v1.8.1",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "realisticvision-v5",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "timeless-v1",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "portrait-v1",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "analog",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "anything-v3",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "anything-v4",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "abyssorangemix",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "deliberate",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "dreamlike-v1",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "dreamlike-v2",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "dreamshaper-5",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "dreamshaper-6",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "elldrethvividmix",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "lyriel-v15",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "lyriel-v16",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "mechamix",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "meinamix",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "realisticvs-v14",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "realisticvs-v20",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "riffusion",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "sd-v14",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "sd-v15",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "sbp",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "theallysmix",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "openjourney",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "icbinp",
            "object": "model",
            "owned_by": "seco",
            "info": "ICBINP aka ICantBelieveItsNotPhotography by SECO",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "latent-consistency-model",
            "object": "model",
            "owned_by": "LCM",
            "info": "LCMs: The next generation of generative models after Latent Diffusion Models (LDMs).",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "deepfloyd-if",
            "object": "model",
            "owned_by": "deepfloyd",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "material-diffusion",
            "object": "model",
            "owned_by": "material",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "eleven-labs",
            "object": "model",
            "owned_by": "elevenlabs",
            "tokens": 333,
            "info": "Official Eleven Labs multi-lingual audio generation model. Max limit of 333 chars.",
            "voices": [
                "rachel",
                "clyde",
                "domi",
                "dave",
                "fin",
                "bella",
                "antoni",
                "thomas",
                "charlie",
                "emily",
                "elli",
                "callum",
                "patrick",
                "harry",
                "liam",
                "dorothy",
                "josh",
                "arnold",
                "charlotte",
                "matilda",
                "matthew",
                "james",
                "joseph",
                "jeremy",
                "michael",
                "ethan",
                "gigi",
                "freya",
                "grace",
                "daniel",
                "serena",
                "adam",
                "nicole",
                "jessie",
                "ryan",
                "sam",
                "glinda",
                "giovanni",
                "mimi"
            ],
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/audio/generations"
            ]
        },
        {
            "id": "eleven-labs-2",
            "object": "model",
            "owned_by": "elevenlabs",
            "tokens": 100,
            "info": "Official Eleven Labs multi-lingual v2 audio generation model. Max limit of ~100 chars.",
            "voices": [
                "rachel",
                "clyde",
                "domi",
                "dave",
                "fin",
                "bella",
                "antoni",
                "thomas",
                "charlie",
                "emily",
                "elli",
                "callum",
                "patrick",
                "harry",
                "liam",
                "dorothy",
                "josh",
                "arnold",
                "charlotte",
                "matilda",
                "matthew",
                "james",
                "joseph",
                "jeremy",
                "michael",
                "ethan",
                "gigi",
                "freya",
                "grace",
                "daniel",
                "serena",
                "adam",
                "nicole",
                "jessie",
                "ryan",
                "sam",
                "glinda",
                "giovanni",
                "mimi"
            ],
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/audio/generations"
            ]
        },
        {
            "id": "eleven-labs-999",
            "object": "model",
            "owned_by": "elevenlabs",
            "tokens": 999,
            "info": "Eleven Labs multi-lingual audio generation model. Max limit of 999 chars.",
            "voices": [
                "rachel",
                "clyde",
                "domi",
                "dave",
                "fin",
                "bella",
                "antoni",
                "thomas",
                "charlie",
                "emily",
                "elli",
                "callum",
                "patrick",
                "harry",
                "liam",
                "dorothy",
                "josh",
                "arnold",
                "charlotte",
                "matilda",
                "matthew",
                "james",
                "joseph",
                "jeremy",
                "michael",
                "ethan",
                "gigi",
                "freya",
                "grace",
                "daniel",
                "serena",
                "adam",
                "nicole",
                "jessie",
                "ryan",
                "sam",
                "glinda",
                "giovanni",
                "mimi"
            ],
            "premium": true,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/audio/generations"
            ]
        },
        {
            "id": "tts-1",
            "object": "model",
            "owned_by": "openai",
            "tokens": 4096,
            "voices": [
                "alloy",
                "echo",
                "fable",
                "onyx",
                "nova",
                "shimmer"
            ],
            "premium": true,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/audio/generations"
            ]
        },
        {
            "id": "tts-1-hd",
            "object": "model",
            "owned_by": "openai",
            "tokens": 4096,
            "voices": [
                "alloy",
                "echo",
                "fable",
                "onyx",
                "nova",
                "shimmer"
            ],
            "premium": true,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/audio/generations"
            ]
        },
        {
            "id": "whisper-large",
            "object": "model",
            "owned_by": "openai",
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/audio/transcriptions"
            ]
        },
        {
            "id": "whisper-1",
            "object": "model",
            "owned_by": "openai",
            "premium": true,
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/audio/transcriptions"
            ]
        },
        {
            "id": "text-moderation-stable",
            "object": "proxy",
            "proxy_to": "text-moderation-007",
            "owned_by": "openai",
            "tokens": 100000,
            "info": "100,000 characters input limit per request",
            "cost": 0,
            "created": 1687882410,
            "endpoints": [
                "/v1/moderations"
            ]
        },
        {
            "id": "text-moderation-latest",
            "object": "model",
            "owned_by": "openai",
            "tokens": 100000,
            "info": "100,000 characters input limit per request, might be slightly more accurate at times but can false flag.",
            "cost": 0,
            "created": 1687882410,
            "endpoints": [
                "/v1/moderations"
            ]
        },
        {
            "id": "text-moderation-007",
            "object": "model",
            "owned_by": "openai",
            "tokens": 100000,
            "info": "100,000 characters input limit per request",
            "cost": 0,
            "created": 1687882410,
            "endpoints": [
                "/v1/moderations"
            ]
        },
        {
            "id": "text-embedding-3-large",
            "object": "model",
            "owned_by": "openai",
            "tokens": 3072,
            "info": "also supports multiple inputs at once by using a list of strings as the input",
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/embeddings"
            ]
        },
        {
            "id": "text-embedding-3-small",
            "object": "model",
            "owned_by": "openai",
            "tokens": 1536,
            "info": "also supports multiple inputs at once by using a list of strings as the input",
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/embeddings"
            ]
        },
        {
            "id": "text-embedding-ada-002",
            "object": "model",
            "owned_by": "openai",
            "tokens": 1536,
            "info": "also supports multiple inputs at once by using a list of strings as the input",
            "cost": 1,
            "created": 1687882410,
            "endpoints": [
                "/v1/embeddings"
            ]
        }
    ],
    "object": "list",
    "total": 107
}
```
