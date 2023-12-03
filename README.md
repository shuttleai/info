# Shuttle API

## Introduction

Are you a developer interested in creating a AI bot/site, but you find GPT-4 to be too expensive or you're on a tight budget? Look no further!
Shuttle API provides `gpt-4` and `dalle-3` all for free with a super **stable** and **free** api!
## Getting Started

To obtain a free Shuttle API key, follow these steps:

1. Join [our Discord server](https://discord.gg/shuttleai).
2. Run `/getkey` command in the Discord server.

Please make sure you have the necessary knowledge to use the API effectively.

## Our Models:
```json
{
    "data": [
        {
            "id": "internet-gpt",
            "object": "model",
            "owned_by": "shuttle/openai",
            "tokens": 16385,
            "info": "GPT-3.5-Turbo-1106 with added Real-Time Live Web Browsing",
            "premium": true,
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
            "maintenance": true,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "code-gpt",
            "object": "model",
            "owned_by": "shuttle/openai",
            "tokens": 16385,
            "info": "GPT 3.5 Turbo 16K with the automatic ability to interpret and run code from a variety of languages including modules/packages. Basically Code Interpreter from GPT Plus.",
            "maintenance": true,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "image-gpt",
            "object": "model",
            "owned_by": "shuttle/openai",
            "tokens": 16385,
            "info": "GPT-3.5-Turbo-1106 Chat with added automatic Image Generations via direct chat",
            "maintenance": true,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-3.5-vision",
            "object": "model",
            "owned_by": "shuttle/openai",
            "tokens": 4096,
            "info": "GPT-3.5 with added automatic image analyzations/descriptions. Not bad, try it.",
            "maintenance": true,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-4",
            "object": "model",
            "owned_by": "bing/openai",
            "tokens": 32768,
            "info": "Full-Stack Bing Chat. Includes Bing internet, optional citations, as well as optional live-per-prompt AI suggestions.",
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-4-32k",
            "object": "proxy",
            "proxy_to": "gpt-4",
            "owned_by": "bing/openai",
            "tokens": 32768,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-4-0613",
            "object": "model",
            "owned_by": "openai",
            "tokens": 8192,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-4-1106-preview",
            "object": "model",
            "owned_by": "openai",
            "tokens": 128000,
            "info": "GPT-4-Turbo-128K Preview",
            "premium": true,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-4-vision-preview",
            "object": "model",
            "owned_by": "openai",
            "tokens": 4096,
            "info": "GPT-4-Turbo-Vision Preview",
            "premium": true,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-3.5-turbo-1106",
            "object": "model",
            "owned_by": "openai",
            "tokens": 16385,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-3.5-turbo",
            "object": "model",
            "owned_by": "openai",
            "tokens": 8192,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-3.5-turbo-16k",
            "object": "model",
            "owned_by": "openai",
            "tokens": 16384,
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
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-3.5-turbo-0301",
            "object": "model",
            "owned_by": "openai",
            "tokens": 4097,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-3.5-turbo-0613",
            "object": "model",
            "owned_by": "openai",
            "tokens": 4097,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "chatgpt",
            "object": "model",
            "owned_by": "openai",
            "tokens": 16385,
            "info": "official chatgpt (GPT 3.5) aka openai's fastest and most capable free-to-use model aka text-davinci-002-render-sha",
            "maintenance": true,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "claude-instant",
            "object": "model",
            "owned_by": "anthropic",
            "tokens": 25000,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "claude-instant-100k",
            "object": "model",
            "owned_by": "anthropic",
            "tokens": 100000,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "claude-2-100k",
            "object": "model",
            "owned_by": "anthropic",
            "tokens": 100000,
            "premium": true,
            "maintenance": true,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "palm-2",
            "object": "model",
            "owned_by": "google",
            "tokens": 8192,
            "maintenance": true,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "code-llama-34b",
            "object": "model",
            "owned_by": "metaai",
            "tokens": 2048,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "llama-70b",
            "object": "model",
            "owned_by": "metaai",
            "tokens": 2048,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "llava-13b",
            "object": "model",
            "owned_by": "metaai",
            "tokens": 2048,
            "premium": true,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "mistral-7b",
            "object": "model",
            "owned_by": "mistralai",
            "tokens": 4096,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "airoboros-70b",
            "object": "model",
            "owned_by": "jondurbin",
            "tokens": 4096,
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "dalle3",
            "object": "model",
            "owned_by": "bing/openai",
            "info": "ShuttleAI does not support changing the `n` paramater for dalle3 generations. However much Bing gens is how much you will get.",
            "max_images": 4,
            "multiple_of": 4,
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
            "maintenance": true,
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "sdxl",
            "object": "model",
            "owned_by": "stabilityai",
            "max_images": 4,
            "multiple_of": 1,
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
            "endpoints": [
                "/v1/images/generations"
            ]
        },
        {
            "id": "eleven-labs",
            "object": "model",
            "owned_by": "elevenlabs",
            "tokens": 333,
            "info": "Official Eleven Labs multi-lingual audio generation model. Max limit of 333 chars. Make a ticket with premium for higher limits.",
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
            "endpoints": [
                "/v1/audio/generations"
            ]
        },
        {
            "id": "text-moderation-006",
            "object": "model",
            "owned_by": "openai",
            "tokens": 16385,
            "endpoints": [
                "/v1/chat/completions",
                "/v1/moderations"
            ]
        }
    ],
    "object": "list",
    "total": 68
}
```
