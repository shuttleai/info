const fetch = require('node-fetch');
const readline = require('readline');

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

rl.question('Enter your message: ', function(message) {
  const data = {
    "action": "_ask",
    "model": "gpt-3.5-turbo-0301",
    "jailbreak": "default",
    "meta": {
        "content": {
            "internet_access": false,
            "conversation": [],
            "parts": [
                {"role": "system", "content": "You are GPT-3.5 Turbo, a large language model trained by OpenAI. Knowledge cutoff: 2021-09-01. Current date: 2023-05-10."},
                {"role": "user", "content": message}
            ]
        }
    }
  };
  
  const headers = {
    'Accept': 'text/event-stream',
    'Accept-Encoding': 'gzip, deflate',
    'Accept-Language': 'en-US,en;q=0.9',
    'Connection': 'keep-alive',
    'Content-type': 'application/json',
    'Host': 'shuttleproxy.com:6999',
    'Origin': 'http://shuttleproxy.com:6999',
    'Referer': 'http://shuttleproxy.com:6999/chat/',
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36',
  };
  
  fetch('http://shuttleproxy.com:6999/backend-api/v2/conversation', {
    method: 'POST',
    body: JSON.stringify(data),
    headers: headers
  })
  .then(response => response.text())
  .then(data => {
    console.log(data);
    rl.close();
  })
  .catch((error) => {
    console.error('Error:', error);
  });
});
