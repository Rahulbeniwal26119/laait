from collections import OrderedDict
from datetime import datetime
from shared.shared_utils import check_error_message
from backend.models import Notebook as NB
from pytz import timezone

def generate_key_from_datetime(datetime):
    format="%d%m%Y%H%M%S"
    return datetime.strftime(format)

def get_time_format():
    return "%d/%m/%Y %H:%M:%S"

def prepare_data_for_history_view(query_set_data):
    data = {}
    for data_item in query_set_data:
        key = generate_key_from_datetime(data_item.created_date)
        data[key] = {
            "input": data_item.input,
            "output": data_item.output,
            "update_date": data_item.update_date.strftime(get_time_format())
        }

    data = OrderedDict({ k : v for k,v in sorted(data.items(), 
    key=lambda item:item[1]['update_date'])})
    return data


def get_current_dtm():
    return datetime.now(timezone("Asia/Kolkata"))


def check_if_python_code(code):
    return code.startswith("Python{") or code.startswith("python{") \
        or code.startswith("python {") or code.startswith("Python {")


def execute_python_code(code):
    if check_if_python_code(code):
        code = code.split("\n")[1:-1]
        code = "\n".join(code)
        output = exec(code)
        return output

def check_if_python_code_interpreted():
    with open('output_.txt', 'r') as f:
        output = f.readlines()
    if not output:
        return False 



def print_output(file_path):
    response_code = ''
    output = []
    with open('output_.txt', 'r') as f:
        output = f.readlines()
    if not output:
        output = 'No output. It is under developement'
        response_code = 'DEVEL'

    output = ''.join(output)
    if output:
        if check_error_message(output):
            response_code = 'ERROR'
        else:
            response_code = 'SUCCESS'
    return output, response_code


def create_notebook_object(input_code, output):
    return NB.objects.create(
        input=input_code, 
        output=output, 
        created_date=get_current_dtm(), 
        update_date=get_current_dtm(),
    )