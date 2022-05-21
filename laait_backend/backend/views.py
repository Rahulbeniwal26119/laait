from datetime import datetime, timedelta
from inspect import trace
from pytz import timezone 
import os
from urllib import response
from backend.models import Notebook as NB 
from rest_framework.decorators import api_view
from backend.forms import NoteBookForm
from shared.shared_utils import log_and_respond, check_error_message
from rest_framework import status 
from os import system 
import simplejson as json 
from backend.notebook_view_helper import \
    prepare_data_for_history_view, get_current_dtm, \
    check_if_python_code, execute_python_code,  print_output,\
    create_notebook_object, check_if_python_code_interpreted
import sys 
import traceback

# Create your views here.

@api_view(["POST"])
def notebook_view(request):
    try:
        input_form_data = NoteBookForm(request.POST)
        if not input_form_data.is_valid():
            return log_and_respond(
                data = None,
                message = "Input cannot be null",
                status=  status.HTTP_400_BAD_REQUEST
            )
        input_code = request.POST.get("input")
        # change system_output to output file 

        output_file_abs_path = os.path.join(os.path.abspath('.') , "output_.txt")

        # check if code is belong to python 
        if check_if_python_code(input_code):
            sys.stdout = open(output_file_abs_path, 'w')
            output = execute_python_code(input_code)
            sys.stdout.close()

        else:
            with open('read_text.txt', 'w') as f:
                f.write(input_code)


            """
            Return take output of laait and return output json 
            """
            # append the location of excutable 
            os.environ["PATH"]+=f":{os.path.join(os.path.abspath('.'))}"

            if os.path.exists(output_file_abs_path):
                os.remove(output_file_abs_path)

            # check if output is exist then remove this 

            system('laait_nb_interpreter evaluator notebook')

        # if command is successful it create a output file in 
        output, response_code = print_output(output_file_abs_path)

        create_notebook_object(
            input_code=input_code, 
            output=output, 
            )

        data = {
            "output" : output,
            "response_code": response_code
        } 

        return log_and_respond(
            data = data, 
            status = status.HTTP_200_OK,
            message="Interpreted Successfully"
        )

    except Exception as e:
        print(traceback.format_exc())
        return log_and_respond(
            data = None,
            message = str(e),
            status=  status.HTTP_500_INTERNAL_SERVER_ERROR,
            exception=e)


@api_view(["GET"])
def get_history_view(request):
    try:
        mins = request.GET.get("mins")
        time_stamp = get_current_dtm()

        if not mins:
            last_date = NB.objects.all().order_by("-update_date").last().update_date
            # remove the time from last date 
            time_stamp = last_date.replace(hour=0, minute=0, second=0, microsecond=0)
        else:
            time_stamp = get_current_dtm() - timedelta(minutes=mins)

        data = NB.objects.filter(update_date__gte=time_stamp)
        response_data = prepare_data_for_history_view(data)

        # update all the records to the current data as these are opended today 
        for record in data:
            record.update_date = get_current_dtm()
            record.save()
            
        return log_and_respond(
            data=response_data,
            status=status.HTTP_200_OK,
            message="Successfully fetched history"
        )
    except Exception as e:
        print(traceback.format_exc())
        return log_and_respond(
            data=None,
            message=str(e),
            status=status.HTTP_500_INTERNAL_SERVER_ERROR,
            exception=e
        )


