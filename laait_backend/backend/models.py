from typing_extensions import Required
from django.db import models

# Create your models here.

class Notebook(models.Model):
    input = models.TextField()
    output = models.TextField(null=True, blank=True)

    class Meta:
        verbose_name = "Notebook"
        verbose_name_plural = "Notebooks"