import requests
import argparse

def parse_option():
    parser = argparse.ArgumentParser(description="Health Check")
    parser.add_argument("-H", type=str, default="localhost", help="Host of the daemon")
    parser.add_argument("-P", type=int, default=8080, help="Port of the daemon")
    parser.add_argument("--timeout", type=int, default=30, help="Timeout for the request")
    return parser.parse_args()

if __name__ == "__main__":
    opt = parse_option()
    url = f"http://{opt.H}:{opt.P}/health"
    resp = requests.get(url, timeout=opt.timeout)

    if resp.status_code == 200:
        print("Daemon is healthy")
        
    else:
        print(f"Daemon is not healthy, status code: {resp.status_code}")
        print(resp.text)