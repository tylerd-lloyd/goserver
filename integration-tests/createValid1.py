import pathlib
import requests
import sys

global host

def post_data():
    data = None
    with open(pathlib.Path(__file__).parent / 'valid1.yaml') as fd:
        data = fd.read()

    res = requests.post(f'{host}/api/metadata', data = data)
    assert(res.ok)
    print(f'Response: {res.status_code} OK')
    print(bytes.decode(res.content))

if __name__ == "__main__":
    if len(sys.argv) > 1:
        port = sys.argv[1]
    else:
        port = "50001"
    
    host = f'http://localhost:{port}'
    post_data()