# Generated by Django 3.1.1 on 2021-04-04 01:57

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('qcard', '0005_auto_20210404_0956'),
    ]

    operations = [
        migrations.AlterField(
            model_name='user',
            name='birthday',
            field=models.DateTimeField(blank=True, null=True, verbose_name='club'),
        ),
    ]