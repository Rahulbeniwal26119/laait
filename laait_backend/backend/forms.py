from typing_extensions import Required
from django import forms 


class NoteBookForm(forms.Form):
    input = forms.TextInput()