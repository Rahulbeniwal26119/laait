import os
from urllib import response
from backend.models import Notebook as NB 
from rest_framework.decorators import api_view
from backend.forms import NoteBookForm
from shared.shared_utils import log_and_respond, check_error_message
from rest_framework import status 
from os import system 
import simplejson as json 

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
        

        with open('read_text.txt', 'w') as f:
            f.write(input_code)
        

        """
        Return take output of laait and return output json 
        """
        # append the location of excutable 
        os.environ["PATH"]+=f":{os.path.join(os.path.abspath('.'))}"
        
        # check if output is exist then remove this 
        output_file_abs_path = os.path.join(os.path.abspath('.') , "output_.txt")
        
        print(output_file_abs_path)
        if os.path.exists(output_file_abs_path):
            os.remove(output_file_abs_path)

        system('laait_nb_interpreter evaluator notebook')

        # if command is successful it create a output file in 
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

        NB.objects.create(input=input_code, output=output)
        data = {
            "output" : output,
            "response_code": response_code
        } 
        return log_and_respond(data = data, 
        status = status.HTTP_200_OK,message="Interpreted Successfully")
    except Exception as e:
        return log_and_respond(
            data = None,
            message = str(e),
            status=  status.HTTP_500_INTERNAL_SERVER_ERROR,
            exception=e)
