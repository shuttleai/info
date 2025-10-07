# ShuttleAI — FREE & Unified AI API

**One API for all your AI needs.**  
Access GPT-5, Claude, GPT-OSS, and more.  
Free tier for testing, cheap paid plans for unlimited AI power.

[Website](https://shuttleai.com) • [Docs](https://docs.shuttleai.com) • [Discord](https://discord.gg/shuttleai)

---

## What is ShuttleAI?

ShuttleAI is a unified platform that gives developers access to the best AI models — all through one simple API.  
Use it for text, image, audio, or search tasks without switching providers.

Built for:
- Developers who want free AI access  
- Startups needing predictable pricing  
- Teams tired of expensive token-based APIs  

---

## Pricing

Simple, transparent plans with no hidden fees.

| Plan | Monthly Price | Features |
|------|----------------|-----------|
| **Free** | $0/month | Access core models (GPT-5 Mini, GPT-OSS 20B/120B, Shuttle 3.5, etc) — 8K context, 2 req/min |
| **Basic** | $10/month | Access GPT-5, GPT-4o, Claude 3.5 Haiku — 16K context, 10 req/min |
| **Premium** | $25/month | Access Claude Sonnet 4 / 4.5 / 3.7 and Claude Opus 4 — 36K context, 30 req/min |
| **Scale** | $75/month | Unlock all models + beta access — 64K context, 80 req/min |
| **Enterprise** | Custom | Dedicated support and higher limits |

Credit card required for verification.

---

## Available Models

| Model | Type | Plan |
|--------|------|------|
| Shuttle 3.5 | Text Generation | Free |
| Shuttle STT | Audio Transcription | Free |
| GPT-5 | Text Generation | Basic |
| GPT-5 Mini | Text Generation | Free |
| GPT-4o | Text Generation | Basic |
| GPT-4.1 | Text Generation | Basic |
| GPT-OSS 20B | Text Generation | Free |
| GPT-OSS 120B | Text Generation | Free |
| DeepSeek R1 Distill LLaMA 70B | Text Generation | Free |
| Claude 3.5 Haiku | Text Generation | Basic |
| Claude Sonnet 4 | Text Generation | Premium |
| Claude Sonnet 4.5 | Text Generation | Premium |
| Claude 3.7 Sonnet | Text Generation | Premium |
| Claude Opus 4 | Text Generation | Premium |
| Lucid Origin | Image Generation | Free |

See full list → [shuttleai.com/models](https://shuttleai.com/models)

---

## Getting Started

1. Sign up at [shuttleai.com](https://shuttleai.com)  
2. Go to your [Dashboard](https://shuttleai.com/keys)  
3. Generate your API key  
4. Start using it with your favorite SDK or cURL

Example (Node.js):

```js
import axios from "axios";

const response = await axios.post(
  "https://api.shuttleai.com/v1/chat/completions",
  {
    model: "gpt-oss-20b",
    messages: [{ role: "user", content: "Explain quantum computing simply." }],
  },
  {
    headers: { Authorization: `Bearer YOUR_API_KEY` },
  }
);

console.log(response.data);
```

---

## Why ShuttleAI

- Free forever plan for testing and hobby use  
- Simple upgrade path for unlimited usage  
- Fast, reliable infrastructure  
- Works with top models from OpenAI, Anthropic, Google, and more  
- Easy-to-use documentation and examples  

---

## Community

Join our [Discord](https://discord.gg/shuttleai) to connect with other developers, get help, and stay updated on new releases.

---

**ShuttleAI — One API. Infinite Power.**  
© 2025 ShuttleAI Inc. • support@shuttleai.com
