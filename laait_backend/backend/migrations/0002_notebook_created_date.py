# Generated by Django 4.0.4 on 2022-05-19 13:23

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('backend', '0001_initial'),
    ]

    operations = [
        migrations.AddField(
            model_name='notebook',
            name='created_date',
            field=models.DateTimeField(null=True),
        ),
    ]
