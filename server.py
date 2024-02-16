from flask import Flask
import sys

app = Flask(__name__)

server_name = sys.argv[1]
port_number = sys.argv[2]

@app.route('/')
def hello():
    return server_name


if __name__ == '__main__':
    app.run(port=port_number)