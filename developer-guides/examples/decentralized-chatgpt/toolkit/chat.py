import requests
import argparse
import json

def parse_option():
    parser = argparse.ArgumentParser(description="Chat with your eternal")
    parser.add_argument("-H", type=str, default="localhost", help="Host of the daemon")
    parser.add_argument("-P", type=int, default=8080, help="Port of the daemon")
    return parser.parse_args()

if __name__ == "__main__":
    opt = parse_option()
    base_url = f"http://{opt.H}:{opt.P}"
    session_info_resp = requests.post(f"{base_url}/api/v1/init-chat")

    if session_info_resp.status_code != 200:
        raise ValueError("Failed to initialize chat session")

    session_info = session_info_resp.json()    
    
    session_id = session_info["session_id"]
    print(f"Session ID: {session_id}")
    print("Greetings! I am your eternal assistant. How can I help you today?")

    try:
        while True:
            message = input("> You: ")

            response = requests.post(
                f"{base_url}/api/v1/chat/{session_id}", 
                params={"message": message}
            )
            
            response_json = response.json()
            response_message = response_json.get("response")
            print(f"> Eternal: {response_message}")
    except KeyboardInterrupt:
        history_resp = requests.get(f"{base_url}/api/v1/chat/{session_id}/history")
        
        if history_resp.status_code == 200:
            history = history_resp.json()

            with open("chat_history.json", "w") as f:
                json.dump(history, f, indent=4)

        requests.get(f"{base_url}/api/v1/deinit-chat/{session_id}")
        