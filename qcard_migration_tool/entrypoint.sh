#!/bin/bash


# Apply database migrations
echo "Apply database migrations"
python manage.py makemigrations qcard
python manage.py migrate