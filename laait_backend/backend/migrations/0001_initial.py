# Generated by Django 4.0.3 on 2022-03-12 17:12

from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='Notebook',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('input', models.TextField()),
                ('output', models.TextField(blank=True, null=True)),
            ],
            options={
                'verbose_name': 'Notebook',
                'verbose_name_plural': 'Notebooks',
            },
        ),
    ]
