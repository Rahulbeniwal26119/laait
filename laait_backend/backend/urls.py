from django.urls import path
from .views import \
    notebook_view, get_history_view


urlpatterns = [
    path('output/', notebook_view, name='notebook_view'),
    path('history/', get_history_view, name='get_history_view')
]