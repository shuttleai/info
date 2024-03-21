# ShuttleAI API

## Introduction

Are you a developer keen on integrating AI capabilities into your projects, but find the cost of GPT-4 prohibitive or are operating within budget constraints? Look no further!

ShuttleAI API offers **FREE** access to `gpt-4`, `dall-e-3`, and over 100+ other models directly within the API, providing a highly **efficient** and **reliable** service!

## Getting Started

To obtain your complimentary ShuttleAI API key, simply follow these steps:

1. Register on [Our Platform](https://shuttleai.app).
2. Proceed to the [Dashboard](https://shuttleai.app/keys) to generate your unique key.

For comprehensive instructions on utilizing our API effectively, refer to the ShuttleAI API documentation available at [https://docs.shuttleai.app/](https://docs.shuttleai.app/).

## ShuttleAI Models (131 Total):
[api.shuttleai.app/v1/models](https://api.shuttleai.app/v1/models)
| id | owned_by | tokens | info | cost | endpoint | maintenance | premium |
| -- | -------- | ------ | ---- | ---- | -------- | ----------- | ------- |
| shuttle-turbo | openai/shuttleai | 32768 | Currently points to shuttle-1 | 1 | /v1/chat/completions | False | False |
| shuttle-1 | openai/shuttleai | 32768 | Shuttles latest chat completion model with built in support for image recognition & web access released on February 17th, 2024. | 1 | /v1/chat/completions | False | False |
| shuttle-tools | shuttleai | 32768 | A function/tool calling model. OpenAI format as always. | 1 | /v1/chat/completions | False | True |
| gpt-4-turbo-preview | openai | 128000 | Always proxies to latest gpt-4-turbo-preview model. | 4 | /v1/chat/completions | False | True |
| gpt-4-0125-preview | openai | 128000 | New (laziness reduced) GPT-4-Turbo-Preview 128k | 4 | /v1/chat/completions | False | True |
| gpt-4-1106-preview | openai | 128000 | GPT-4-Turbo-Preview 128k | 4 | /v1/chat/completions | False | True |
| gpt-4-vision-preview | openai | 128000 | GPT-4-Vision-Preview 128k | 4 | /v1/chat/completions | False | True |
| gpt-4-32k | openai | 32000 | Free provider pool of real GPT-4-32k, falls back to Bing GPT-4 32k if free pool is out. Premium pool found @ 32k-0613. | 4 | /v1/chat/completions | False | False |
| gpt-4-32k-0613 | openai | 32768 | N/A | 3 | /v1/chat/completions | False | True |
| gpt-4-0613 | openai | 8192 | N/A | 2 | /v1/chat/completions | False | False |
| gpt-4 | openai | 8192 | GPT-4 is a large multimodal model by OpenAI. | 2 | /v1/chat/completions | False | False |
| gpt-4-bing-turbo | openai/bing | - | N/A | 1 | /v1/chat/completions | False | False |
| gpt-4-bing | openai/bing | - | N/A | 1 | /v1/chat/completions | False | False |
| gpt-3.5-turbo | openai | - | N/A | 1 | /v1/chat/completions | False | False |
| gpt-3.5-turbo-0125 | openai | 16385 | N/A | 1 | /v1/chat/completions | False | True |
| gpt-3.5-turbo-1106 | openai | 16385 | N/A | 1 | /v1/chat/completions | False | False |
| gpt-3.5-turbo-0613 | openai | 4097 | N/A | 1 | /v1/chat/completions | False | False |
| gpt-3.5-turbo-0301 | openai | 4097 | N/A | 1 | /v1/chat/completions | False | False |
| gpt-3.5-turbo-16k | openai | 16384 | N/A | 1 | /v1/chat/completions | False | False |
| gpt-3.5-turbo-16k-0613 | openai | 16385 | N/A | 1 | /v1/chat/completions | False | False |
| chatgpt | openai | 8192 | official chatgpt (GPT 3.5) aka openai's fastest and most capable free-to-use model aka text-davinci-002-render-sha | 1 | /v1/chat/completions | False | False |
| claude-3-opus | anthropic | 200000 | N/A | 9 | /v1/chat/completions | False | True |
| claude-3-sonnet | anthropic | 200000 | N/A | 7 | /v1/chat/completions | False | True |
| claude-3-haiku-200k | anthropic | 200000 | N/A | 5 | /v1/chat/completions | False | True |
| claude-3-haiku | anthropic | 72000 | Context window has been shortened to optimize for speed and cost. For longer context messages, please try Claude-3-Haiku-200k. | 1 | /v1/chat/completions | False | True |
| claude-2.1 | anthropic | 200000 | N/A | 4 | /v1/chat/completions | False | True |
| claude-2.0 | anthropic | 100000 | N/A | 4 | /v1/chat/completions | False | True |
| claude-instant-100k | anthropic | 100000 | N/A | 3 | /v1/chat/completions | False | True |
| claude-instant | anthropic | - | N/A | 1 | /v1/chat/completions | False | False |
| gemini-1.5-pro | google | 1000000 | N/A | 9 | /v1/chat/completions | False | True |
| gemini-pro | google | 30724 | N/A | 1 | /v1/chat/completions | False | True |
| gemini-pro-vision | google | 30724 | N/A | 1 | /v1/chat/completions | False | True |
| palm-2 | google | 8192 | N/A | 2 | /v1/chat/completions | False | False |
| mistral-large | mistralai | - | Mistral's Latest and Largest Model. | 4 | /v1/chat/completions | False | False |
| mistral-next | mistralai | - | Prototype model with extra concision. | 4 | /v1/chat/completions | False | True |
| mistral-medium | mistralai | - | N/A | 2 | /v1/chat/completions | False | False |
| mistral-small | mistralai | - | N/A | 2 | /v1/chat/completions | False | False |
| mistral-7b | mistralai | 4096 | N/A | 1 | /v1/chat/completions | False | False |
| mixtral-8x7b | mistralai | 32768 | N/A | 1 | /v1/chat/completions | False | False |
| dolphin-mixtral-8x7b | cognitivecomputations | 4096 | N/A | 1 | /v1/chat/completions | False | False |
| chronos-hermes-13b | chronos/hermes | - | N/A | 1 | /v1/chat/completions | False | False |
| pplx-70b-online | perplexitylabs | 32768 | N/A | 1 | /v1/chat/completions | False | False |
| pplx-70b-chat | perplexitylabs | 32768 | N/A | 1 | /v1/chat/completions | False | False |
| pplx-7b-online | perplexitylabs | 32768 | N/A | 1 | /v1/chat/completions | False | False |
| pplx-7b-chat | perplexitylabs | 32768 | N/A | 1 | /v1/chat/completions | False | False |
| nous-capybara-34b | nous | 200000 | N/A | 4 | /v1/chat/completions | False | False |
| yi-34b-chat | yi | 200000 | N/A | 4 | /v1/chat/completions | False | False |
| starcoder-16b | star | - | N/A | 2 | /v1/chat/completions | False | False |
| starcoder-7b | star | - | N/A | 2 | /v1/chat/completions | False | False |
| deepseek-coder | deepseek | 4096 | N/A | 1 | /v1/chat/completions | False | False |
| deepseek-chat | deepseek | 4096 | N/A | 1 | /v1/chat/completions | False | False |
| airoboros-2-70b | jondurbin | 6105 | N/A | 1 | /v1/chat/completions | False | False |
| airoboros-70b | jondurbin | 6105 | N/A | 1 | /v1/chat/completions | False | False |
| falcon-180b | tiiuae | 6105 | N/A | 1 | /v1/chat/completions | True | False |
| phind-34b | phind | - | N/A | 2 | /v1/chat/completions | False | False |
| phind-code-llama-v2-34b | phind | - | N/A | 2 | /v1/chat/completions | False | False |
| code-llama-70b | metaai | 6105 | N/A | 1 | /v1/chat/completions | False | False |
| code-llama-34b | metaai | 6105 | N/A | 1 | /v1/chat/completions | False | False |
| llama-2-70b-chat | metaai | 6105 | Llama 2 70b Chat Hf | 1 | /v1/chat/completions | False | False |
| llama-2-13b-chat | metaai | 2048 | N/A | 1 | /v1/chat/completions | False | False |
| llama-2-7b-chat | metaai | 2048 | N/A | 1 | /v1/chat/completions | False | False |
| llama-summarize | metaai | - | N/A | 2 | /v1/chat/completions | False | False |
| llava-13b | metaai | 2048 | Llava 13b, either send a 'image' paramater with your model and messages or send an image in the last message. | 1 | /v1/chat/completions | False | False |
| midjourney | midjourney | - | N/A | 369 | /v1/images/generations | False | True |
| dall-e-3 | openai | - | N/A | 4 | /v1/images/generations | False | False |
| dall-e-3-premium | openai | - | N/A | 4 | /v1/images/generations | False | True |
| dall-e-2 | openai | - | N/A | 3 | /v1/images/generations | False | False |
| sdxl | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| sdxl-turbo | stabilityai | - | N/A | 2 | /v1/images/generations | False | False |
| sdxl-emoji | stabilityai/fofr | - | N/A | 2 | /v1/images/generations | False | False |
| dreamshaper-xl | stabilityai | - | N/A | 2 | /v1/images/generations | False | False |
| juggernaut-xl | stabilityai | - | N/A | 2 | /v1/images/generations | False | False |
| dynavision-xl | stabilityai | - | N/A | 2 | /v1/images/generations | False | False |
| realism-engine-xl | stabilityai | - | N/A | 2 | /v1/images/generations | False | False |
| sdxl-inpainting | stabilityai | - | N/A | 4 | /v1/images/generations | False | False |
| turbovision-xl | stabilityai | - | N/A | 2 | /v1/images/generations | False | False |
| kandinsky-2.2 | sberbank | - | N/A | 1 | /v1/images/generations | False | False |
| kandinsky-2 | sberbank | - | N/A | 1 | /v1/images/generations | False | False |
| stable-diffusion 2.1 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| stable-diffusion 1.5 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| playground-v2.5 | playgroundai | - | N/A | 2 | /v1/images/generations | False | False |
| dreamshaper-v8 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| rev-animated | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| anything-v5 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| absolutereality-v1.8.1 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| realisticvision-v5 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| timeless-v1 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| portrait-v1 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| analog | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| anything-v3 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| anything-v4 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| abyssorangemix | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| deliberate | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| dreamlike-v1 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| dreamlike-v2 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| dreamshaper-5 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| dreamshaper-6 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| elldrethvividmix | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| lyriel-v15 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| lyriel-v16 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| mechamix | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| meinamix | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| realisticvs-v14 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| realisticvs-v20 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| riffusion | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| sd-v14 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| sd-v15 | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| sbp | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| theallysmix | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| openjourney | stabilityai | - | N/A | 1 | /v1/images/generations | False | False |
| icbinp | seco | - | ICBINP aka ICantBelieveItsNotPhotography by SECO | 1 | /v1/images/generations | False | False |
| latent-consistency-model | LCM | - | LCMs: The next generation of generative models after Latent Diffusion Models (LDMs). | 1 | /v1/images/generations | False | False |
| deepfloyd-if | deepfloyd | - | N/A | 1 | /v1/images/generations | False | False |
| material-diffusion | material | - | N/A | 1 | /v1/images/generations | False | False |
| eleven-labs | elevenlabs | - | N/A | 1 | /v1/audio/speech | False | False |
| eleven-labs-2 | elevenlabs | - | N/A | 2 | /v1/audio/speech | False | True |
| eleven-labs-999 | elevenlabs | - | N/A | 3 | /v1/audio/speech | False | True |
| speechify | speechify | 3000 | N/A | 3 | /v1/audio/speech | False | False |
| tts-1 | openai | 4096 | N/A | 1 | /v1/audio/speech | False | True |
| tts-1-hd | openai | 4096 | N/A | 1 | /v1/audio/speech | False | True |
| whisper-large | openai | - | N/A | 2 | /v1/audio/transcriptions | False | False |
| whisper-1 | openai | - | N/A | 1 | /v1/audio/transcriptions | False | True |
| text-embedding-3-large | openai | 3072 | also supports multiple inputs at once by using a list of strings as the input | 1 | /v1/embeddings | False | False |
| text-embedding-3-small | openai | 1536 | also supports multiple inputs at once by using a list of strings as the input | 1 | /v1/embeddings | False | False |
| text-embedding-ada-002 | openai | 1536 | also supports multiple inputs at once by using a list of strings as the input | 1 | /v1/embeddings | False | False |
| text-moderation-latest | openai | - | N/A | 1 | /v1/moderations | False | False |
| text-moderation-stable | openai | - | N/A | 1 | /v1/moderations | False | False |
| insult-1 | shuttleai | - | N/A | 1 | /v1/insult | False | False |
| joke-1 | shuttleai | - | N/A | 1 | /v1/joke | False | False |
| py-minify-1 | shuttleai | - | Instant Python Code Minification | 1 | /v1/code/minify | False | False |
| search-google | google | - | Includes web and image search, as well as support for many languages. | 2 | /v1/web-search | False | False |
| search-ddg | duckduckgo | - | N/A | 1 | /v1/web-search | False | False |
