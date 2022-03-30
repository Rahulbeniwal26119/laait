from rest_framework.response import Response
from fnmatch import fnmatch as bash_re

def log_and_respond(data=None, status=None, message=None, exception=None):
    return Response(data={
        "data" : data,
        "message" : message,
    }, status=status, exception=exception)


def check_error_message(msg):
    pattern = "*ERROR*"
    return bash_re(msg, pattern)