# Generated by Django 4.0.4 on 2022-05-19 13:55

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('backend', '0002_notebook_created_date'),
    ]

    operations = [
        migrations.AddField(
            model_name='notebook',
            name='update_date',
            field=models.DateTimeField(null=True),
        ),
    ]
