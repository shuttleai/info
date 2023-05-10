import requests
import json

message = input("Enter your message: ")

data = {
    "action": "_ask", 
    "model": "gpt-3.5-turbo-0301",
    "jailbreak": "default",  
    "meta": {
        "content": {
            "internet_access": False, 
            "conversation": [],  
            "parts": [
                {"role": "system", "content": "You are GPT-3.5 Turbo, a large language model trained by OpenAI. Knowledge cutoff: 2021-09-01. Current date: 2023-05-10."},
                {"role": "user", "content": message}
            ]  
        }
    }
}

headers = {
    'Accept': 'text/event-stream',
    'Accept-Encoding': 'gzip, deflate',
    'Accept-Language': 'en-US,en;q=0.9',
    'Connection': 'keep-alive',
    'Content-type': 'application/json',
    'Host': '64.44.148.247:6999',
    'Origin': 'http://shuttleproxy.com:6999',
    'Referer': 'http://shuttleproxy.com:6999/chat/',
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36',
}

response = requests.post('http://shuttleproxy.com:6999/backend-api/v2/conversation', data=json.dumps(data), headers=headers)

if response.status_code == 200:
    for line in response.iter_lines():
        if line:
            decoded_line = line.decode('utf-8')
            print(decoded_line)
else:
    print(f"Request failed with status code: {response.status_code}")
    print(response.text)  
