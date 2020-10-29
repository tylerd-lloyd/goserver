import requests
import urllib
import sys

global host

def get_data():
    params = {'title': 'valid app 1'}
    p = urllib.parse.urlencode(params)
    res = requests.get(f'{host}/api/metadata?{p}')
    assert(res.ok)
    print(f'Response: {res.status_code}')
    print(bytes.decode(res.content))

if __name__ == "__main__":
    if len(sys.argv) > 1:
        port = sys.argv[1]
    else:
        port = "50001"
    
    host = f'http://localhost:{port}'
    get_data()