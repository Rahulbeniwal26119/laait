from collections import OrderedDict
from datetime import datetime

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