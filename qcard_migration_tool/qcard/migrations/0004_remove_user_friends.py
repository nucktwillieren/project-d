# Generated by Django 3.1.1 on 2021-04-04 23:14

from django.db import migrations


class Migration(migrations.Migration):

    dependencies = [
        ('qcard', '0003_user_friends'),
    ]

    operations = [
        migrations.RemoveField(
            model_name='user',
            name='friends',
        ),
    ]
