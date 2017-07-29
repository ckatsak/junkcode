"""
To run the app through uWSGI, issue this command:
    $ uwsgi --socket 0.0.0.0:5000 --protocol=http -w wsgi:cool_app
or:
    $ uwsgi --socket 0.0.0.0:5000 --protocol=http -w wsgi:cool_app --master --enable-threads

To run the app through gunicorn, issue this command:
    $ gunicorn --bind 0.0.0.0:5000 wsgi:cool_app

To run a client to test the app, issue this command on a separate terminal
while the server is up and running:
    $ curl localhost:5000
"""

from myapp import cool_app


if __name__ == "__main__":
    cool_app.run()
