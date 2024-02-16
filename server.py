from flask import Flask

app = Flask(__name__)

server_name = 'Server_1'

@app.route('/')
def hello():
    return server_name


if __name__ == '__main__':
    app.run()