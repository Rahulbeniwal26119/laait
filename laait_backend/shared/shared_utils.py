from rest_framework.response import Response

def log_and_respond(data=None, status=None, message=None, exception=None):
    return Response(data={
        "data" : data,
        "message" : message,
    }, status=status, exception=exception)