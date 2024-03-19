# ShuttleAI API

## Introduction

Are you a developer keen on integrating AI capabilities into your projects, but find the cost of GPT-4 prohibitive or are operating within budget constraints? Look no further!

ShuttleAI API offers access to `gpt-4`, `dall-e-3`, and over 100 other models for **FREE**, along with a highly **efficient** and **reliable** API service!

## Getting Started

To obtain your complimentary ShuttleAI API key, simply follow these steps:

1. Register on [Our Platform](https://shuttleai.app).
2. Proceed to the [Dashboard](https://shuttleai.app/keys) to generate your unique key.

For comprehensive instructions on utilizing our API effectively, refer to the ShuttleAI API documentation available at [https://docs.shuttleai.app/](https://docs.shuttleai.app/).

## ShuttleAI Models:
```json
{
    "object": "list",
    "data": [
        {
            "id": "shuttle-turbo",
            "object": "model",
            "owned_by": "shuttleai/openai",
            "tokens": 32768,
            "info": "Currently points to shuttle-1",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "shuttle-1",
            "object": "model",
            "owned_by": "shuttleai/openai",
            "tokens": 32768,
            "info": "Shuttles latest chat completion model with built in support for image recognition & web access released on February 17th, 2024.",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "shuttle-tools",
            "object": "model",
            "owned_by": "shuttleai",
            "tokens": 32768,
            "info": "A function/tool calling model. OpenAI format as always.",
            "premium": true,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
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
            "endpoint": "/v1/chat/completions"
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
            "endpoint": "/v1/chat/completions"
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
            "endpoint": "/v1/chat/completions"
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
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "gpt-4-32k",
            "object": "model",
            "owned_by": "openai",
            "tokens": 32000,
            "info": "Free provider pool of real GPT-4-32k, falls back to Bing GPT-4 32k if free pool is out. Premium pool found @ 32k-0613.",
            "cost": 4,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "gpt-4-32k-0613",
            "object": "model",
            "owned_by": "openai",
            "tokens": 32768,
            "premium": true,
            "cost": 3,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "gpt-4-0613",
            "object": "model",
            "owned_by": "openai",
            "tokens": 8192,
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "gpt-4",
            "object": "model",
            "owned_by": "openai",
            "tokens": 8192,
            "info": "GPT-4 is a large multimodal model by OpenAI.",
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "gpt-4-bing-turbo",
            "object": "model",
            "owned_by": "openai",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "gpt-4-bing",
            "object": "model",
            "owned_by": "openai",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "gpt-3.5-turbo",
            "object": "model",
            "owned_by": "openai",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "gpt-3.5-turbo-0125",
            "object": "model",
            "owned_by": "openai",
            "tokens": 16385,
            "premium": true,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "gpt-3.5-turbo-1106",
            "object": "model",
            "owned_by": "openai",
            "tokens": 16385,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "gpt-3.5-turbo-0613",
            "object": "model",
            "owned_by": "openai",
            "tokens": 4097,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "gpt-3.5-turbo-0301",
            "object": "model",
            "owned_by": "openai",
            "tokens": 4097,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "gpt-3.5-turbo-16k",
            "object": "model",
            "owned_by": "openai",
            "tokens": 16384,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "gpt-3.5-turbo-16k-0613",
            "object": "proxy",
            "proxy_to": "gpt-3.5-turbo-1106",
            "owned_by": "openai",
            "tokens": 16385,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "chatgpt",
            "object": "model",
            "owned_by": "openai",
            "tokens": 8192,
            "info": "official chatgpt (GPT 3.5) aka openai's fastest and most capable free-to-use model aka text-davinci-002-render-sha",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "claude-3-opus",
            "object": "model",
            "owned_by": "anthropic",
            "tokens": 200000,
            "premium": true,
            "cost": 9,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "claude-3-sonnet",
            "object": "model",
            "owned_by": "anthropic",
            "tokens": 200000,
            "premium": true,
            "cost": 7,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "claude-3-haiku-200k",
            "object": "model",
            "owned_by": "anthropic",
            "tokens": 200000,
            "premium": true,
            "cost": 5,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "claude-3-haiku",
            "object": "model",
            "owned_by": "anthropic",
            "tokens": 72000,
            "info": "Context window has been shortened to optimize for speed and cost. For longer context messages, please try Claude-3-Haiku-200k.",
            "premium": true,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "claude-2.1",
            "object": "model",
            "owned_by": "anthropic",
            "tokens": 200000,
            "premium": true,
            "cost": 4,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "claude-2.0",
            "object": "model",
            "owned_by": "anthropic",
            "tokens": 100000,
            "premium": true,
            "cost": 4,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
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
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "claude-instant",
            "object": "model",
            "owned_by": "anthropic",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "gemini-1.5-pro",
            "object": "model",
            "owned_by": "google",
            "tokens": 1000000,
            "premium": true,
            "cost": 9,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "gemini-pro",
            "object": "model",
            "owned_by": "google",
            "tokens": 30724,
            "premium": true,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "gemini-pro-vision",
            "object": "model",
            "owned_by": "google",
            "tokens": 30724,
            "premium": true,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "palm-2",
            "object": "model",
            "owned_by": "google",
            "tokens": 8192,
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "mistral-large",
            "object": "model",
            "owned_by": "mistralai",
            "info": "Mistral's Latest and Largest Model.",
            "cost": 4,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "mistral-next",
            "object": "model",
            "owned_by": "mistralai",
            "info": "Prototype model with extra concision.",
            "premium": true,
            "cost": 4,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "mistral-medium",
            "object": "model",
            "owned_by": "mistralai",
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "mistral-small",
            "object": "model",
            "owned_by": "mistralai",
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "mistral-7b",
            "object": "model",
            "owned_by": "mistralai",
            "tokens": 4096,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "mixtral-8x7b",
            "object": "model",
            "owned_by": "mistralai",
            "tokens": 32768,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "dolphin-mixtral-8x7b",
            "object": "model",
            "owned_by": "cognitivecomputations",
            "tokens": 4096,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "chronos-hermes-13b",
            "object": "model",
            "owned_by": "chronos/hermes",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "pplx-70b-online",
            "object": "model",
            "owned_by": "perplexitylabs",
            "tokens": 32768,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "pplx-70b-chat",
            "object": "model",
            "owned_by": "perplexitylabs",
            "tokens": 32768,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "pplx-7b-online",
            "object": "model",
            "owned_by": "perplexitylabs",
            "tokens": 32768,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "pplx-7b-chat",
            "object": "model",
            "owned_by": "perplexitylabs",
            "tokens": 32768,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "nous-capybara-34b",
            "object": "model",
            "owned_by": "nous",
            "tokens": 200000,
            "cost": 4,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "yi-34b-chat",
            "object": "model",
            "owned_by": "yi",
            "tokens": 200000,
            "cost": 4,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "starcoder-16b",
            "object": "model",
            "owned_by": "star",
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "starcoder-7b",
            "object": "model",
            "owned_by": "star",
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "deepseek-coder",
            "object": "model",
            "owned_by": "deepseek",
            "tokens": 4096,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "deepseek-chat",
            "object": "model",
            "owned_by": "deepseek",
            "tokens": 4096,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "airoboros-2-70b",
            "object": "model",
            "owned_by": "jondurbin",
            "tokens": 6105,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "airoboros-70b",
            "object": "model",
            "owned_by": "jondurbin",
            "tokens": 6105,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "falcon-180b",
            "object": "model",
            "owned_by": "tiiuae",
            "tokens": 6105,
            "maintenance": true,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "phind-34b",
            "object": "model",
            "owned_by": "phind",
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "phind-code-llama-v2-34b",
            "object": "model",
            "owned_by": "phind",
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "code-llama-70b",
            "object": "model",
            "owned_by": "metaai",
            "tokens": 6105,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "code-llama-34b",
            "object": "model",
            "owned_by": "metaai",
            "tokens": 6105,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "llama-2-70b-chat",
            "object": "model",
            "owned_by": "metaai",
            "tokens": 6105,
            "info": "Llama 2 70b Chat Hf",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "llama-2-13b-chat",
            "object": "model",
            "owned_by": "metaai",
            "tokens": 2048,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "llama-2-7b-chat",
            "object": "model",
            "owned_by": "metaai",
            "tokens": 2048,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "llama-summarize",
            "object": "model",
            "owned_by": "metaai",
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "llava-13b",
            "object": "model",
            "owned_by": "metaai",
            "tokens": 2048,
            "info": "Llava 13b, either send a 'image' paramater with your model and messages or send an image in the last message.",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/chat/completions"
        },
        {
            "id": "midjourney",
            "object": "model",
            "owned_by": "midjourney",
            "max_images": 4,
            "multiple_of": 4,
            "premium": true,
            "cost": 369,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "dall-e-3",
            "object": "model",
            "owned_by": "openai",
            "cost": 4,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "dall-e-3-premium",
            "object": "model",
            "owned_by": "openai",
            "premium": true,
            "cost": 4,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "sdxl",
            "object": "model",
            "owned_by": "stabilityai",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "sdxl-turbo",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "sdxl-emoji",
            "object": "model",
            "owned_by": "stabilityai/fofr",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "dreamshaper-xl",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "juggernaut-xl",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "dynavision-xl",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "realism-engine-xl",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "sdxl-inpainting",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 4,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "turbovision-xl",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "kandinsky-2.2",
            "object": "model",
            "owned_by": "sberbank",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "kandinsky-2",
            "object": "model",
            "owned_by": "sberbank",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "stable-diffusion 2.1",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "stable-diffusion 1.5",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "playground-v2.5",
            "object": "model",
            "owned_by": "playgroundai",
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "dreamshaper-v8",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "rev-animated",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "anything-v5",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "absolutereality-v1.8.1",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "realisticvision-v5",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "timeless-v1",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "portrait-v1",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "analog",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "anything-v3",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "anything-v4",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "abyssorangemix",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "deliberate",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "dreamlike-v1",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "dreamlike-v2",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "dreamshaper-5",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "dreamshaper-6",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "elldrethvividmix",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "lyriel-v15",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "lyriel-v16",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "mechamix",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "meinamix",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "realisticvs-v14",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "realisticvs-v20",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "riffusion",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "sd-v14",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "sd-v15",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "sbp",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "theallysmix",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "openjourney",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
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
            "endpoint": "/v1/images/generations"
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
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "deepfloyd-if",
            "object": "model",
            "owned_by": "deepfloyd",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "material-diffusion",
            "object": "model",
            "owned_by": "material",
            "max_images": 4,
            "multiple_of": 1,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/images/generations"
        },
        {
            "id": "eleven-labs",
            "object": "model",
            "owned_by": "elevenlabs",
            "voices": "https://api.shuttleai.app/v1/voices/eleven-labs",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/audio/speech"
        },
        {
            "id": "eleven-labs-2",
            "object": "model",
            "owned_by": "elevenlabs",
            "voices": "https://api.shuttleai.app/v1/voices/eleven-labs",
            "premium": true,
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/audio/speech"
        },
        {
            "id": "eleven-labs-999",
            "object": "model",
            "owned_by": "elevenlabs",
            "voices": "https://api.shuttleai.app/v1/voices/eleven-labs",
            "premium": true,
            "cost": 3,
            "created": 1687882410,
            "endpoint": "/v1/audio/speech"
        },
        {
            "id": "speechify",
            "object": "model",
            "owned_by": "speechify",
            "tokens": 3000,
            "voices": "https://api.shuttleai.app/v1/voices/speechify",
            "cost": 3,
            "created": 1687882410,
            "endpoint": "/v1/audio/speech"
        },
        {
            "id": "tts-1",
            "object": "model",
            "owned_by": "openai",
            "tokens": 4096,
            "voices": "https://api.shuttleai.app/v1/voices/openai-tts",
            "premium": true,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/audio/speech"
        },
        {
            "id": "tts-1-hd",
            "object": "model",
            "owned_by": "openai",
            "tokens": 4096,
            "voices": "https://api.shuttleai.app/v1/voices/openai-tts",
            "premium": true,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/audio/speech"
        },
        {
            "id": "whisper-large",
            "object": "model",
            "owned_by": "openai",
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/audio/transcriptions"
        },
        {
            "id": "whisper-1",
            "object": "model",
            "owned_by": "openai",
            "premium": true,
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/audio/transcriptions"
        },
        {
            "id": "text-embedding-3-large",
            "object": "model",
            "owned_by": "openai",
            "tokens": 3072,
            "info": "also supports multiple inputs at once by using a list of strings as the input",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/embeddings"
        },
        {
            "id": "text-embedding-3-small",
            "object": "model",
            "owned_by": "openai",
            "tokens": 1536,
            "info": "also supports multiple inputs at once by using a list of strings as the input",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/embeddings"
        },
        {
            "id": "text-embedding-ada-002",
            "object": "model",
            "owned_by": "openai",
            "tokens": 1536,
            "info": "also supports multiple inputs at once by using a list of strings as the input",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/embeddings"
        },
        {
            "id": "text-moderation-latest",
            "object": "model",
            "owned_by": "openai",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/moderations"
        },
        {
            "id": "text-moderation-stable",
            "object": "model",
            "owned_by": "openai",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/moderations"
        },
        {
            "id": "insult-1",
            "object": "model",
            "owned_by": "shuttleai",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/insult"
        },
        {
            "id": "joke-1",
            "object": "model",
            "owned_by": "shuttleai",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/joke"
        },
        {
            "id": "py-minify-1",
            "object": "model",
            "owned_by": "shuttleai",
            "info": "Instant Python Code Minification",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/code/minify"
        },
        {
            "id": "search-google",
            "object": "model",
            "owned_by": "google",
            "info": "Includes web and image search, as well as support for many languages.",
            "cost": 2,
            "created": 1687882410,
            "endpoint": "/v1/web-search"
        },
        {
            "id": "search-ddg",
            "object": "model",
            "owned_by": "duckduckgo",
            "cost": 1,
            "created": 1687882410,
            "endpoint": "/v1/web-search"
        }
    ],
    "total": 131
}
```
