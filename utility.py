from contextlib import contextmanager
import errno
import os
import socket
import sys

from jujupy import until_timeout


@contextmanager
def scoped_environ():
    old_environ = dict(os.environ)
    try:
        yield
    finally:
        os.environ.clear()
        os.environ.update(old_environ)


def wait_for_port(host, port, closed=False, timeout=30):
    for remaining in until_timeout(timeout):
        conn = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        conn.settimeout(max(remaining, 5))
        try:
            conn.connect((host, port))
        except socket.timeout:
            if closed:
                return
        except socket.error as e:
            if e.errno != errno.ECONNREFUSED:
                raise
            if closed:
                return
        except Exception as e:
            print_now('Unexpected %r: %s' % (type(e), e))
            raise
        else:
            conn.close()
            if not closed:
                return
            sleep(1)
    raise Exception('Timed out waiting for port.')


def print_now(string):
    print(string)
    sys.stdout.flush()
