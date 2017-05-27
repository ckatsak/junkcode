from flask import Flask

cool_app = Flask(__name__)


@cool_app.route('/', methods=['GET'])
def yo_index():
    """
    """
    return "Yo yo yo, welcome"


if __name__ == "__main__":
    cool_app.run(host='0.0.0.0', debug=True)
