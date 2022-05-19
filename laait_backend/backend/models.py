from django.db import models
import datetime
from pytz import timezone
# Create your models here.

class Notebook(models.Model):
    input = models.TextField()
    output = models.TextField(null=True, blank=True)
    created_date = models.DateTimeField(null=True)
    update_date = models.DateTimeField(null=True)
    
    class Meta:
        verbose_name = "Notebook"
        verbose_name_plural = "Notebooks"