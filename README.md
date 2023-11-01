# Shuttle API

## Introduction

Are you a developer interested in creating a AI bot/site, but you find GPT-4 to be too expensive or you're on a tight budget? Look no further!
Shuttle API provides `gpt-4` and `dalle-3` all for free with a super **stable** and **free** api!
## Getting Started

To obtain a free Shuttle API key, follow these steps:

1. Join [our Discord server](https://discord.gg/XkBUaxYEWZ).
2. Run `/getkey` command in the Discord server.

Please make sure you have the necessary knowledge to use the API effectively.

## Our Models:
```json
{
    "data": [
        {
            "id": "gpt-4",
            "object": "proxy",
            "proxy_to": "gpt-4-0613",
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
            ],
            "public": true,
            "permission": [],
            "tags": [],
            "group": "default-gpt-4"
        },
        {
            "id": "gpt-4-32k",
            "object": "proxy",
            "proxy_to": "gpt-4-32k-0613",
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-4-32k-0613",
            "object": "model",
            "owned_by": "openai",
            "tokens": 32768,
            "endpoints": [
                "/v1/chat/completions"
            ],
            "public": false,
            "permission": [],
            "tags": [],
            "group": "default-gpt-4-32k"
        },
        {
            "id": "gpt-3.5-turbo",
            "object": "proxy",
            "proxy_to": "gpt-3.5-turbo-0613",
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
            ],
            "public": true,
            "permission": [],
            "tags": [],
            "group": "default-gpt-3.5-turbo"
        },
        {
            "id": "gpt-3.5-turbo-16k",
            "object": "proxy",
            "proxy_to": "gpt-3.5-turbo-16k-0613",
            "endpoints": [
                "/v1/chat/completions"
            ]
        },
        {
            "id": "gpt-3.5-turbo-16k-0613",
            "object": "model",
            "owned_by": "openai",
            "tokens": 16384,
            "endpoints": [
                "/v1/chat/completions"
            ],
            "public": true,
            "permission": [],
            "tags": [],
            "group": "default-gpt-3.5-turbo-16k"
        },
        {
            "id": "sdxl",
            "object": "model",
            "owned_by": "stabilityai",
            "endpoints": [
                "/v1/images/generations"
            ],
            "public": true,
            "max_images": 5,
            "multiple_of": 1,
            "permission": [],
            "tags": [],
            "group": "default-sdxl"
        },
        {
            "id": "kandinsky-2.2",
            "object": "model",
            "owned_by": "sberbank",
            "endpoints": [
                "/v1/images/generations"
            ],
            "public": true,
            "max_images": 10,
            "multiple_of": 1,
            "permission": [],
            "tags": [],
            "group": "default-kandinsky-2.2"
        },
        {
            "id": "kandinsky-2",
            "object": "model",
            "owned_by": "sberbank",
            "endpoints": [
                "/v1/images/generations"
            ],
            "public": true,
            "max_images": 10,
            "multiple_of": 1,
            "permission": [],
            "tags": [],
            "group": "default-kandinsky-2"
        },
        {
            "id": "dall-e",
            "object": "model",
            "owned_by": "openai",
            "endpoints": [
                "/v1/images/generations"
            ],
            "public": true,
            "max_images": 10,
            "multiple_of": 1,
            "permission": [],
            "tags": [],
            "group": "default-dall-e"
        },
        {
            "id": "stable-diffusion 2.1",
            "object": "model",
            "owned_by": "stabilityai",
            "endpoints": [
                "/v1/images/generations"
            ],
            "public": true,
            "max_images": 10,
            "multiple_of": 1,
            "permission": [],
            "tags": [],
            "group": "default-stable-diffusion-2.1"
        },
        {
            "id": "stable-diffusion 1.5",
            "object": "model",
            "owned_by": "stabilityai",
            "endpoints": [
                "/v1/images/generations"
            ],
            "public": true,
            "max_images": 10,
            "multiple_of": 1,
            "permission": [],
            "tags": [],
            "group": "default-stable-diffusion-1.5"
        }
    ],
    "object": "list"
}```
