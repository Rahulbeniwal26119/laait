from django.urls import path
from .views import notebook_view
urlpatterns = [
    path('output/', notebook_view, name='notebook_view'),
]