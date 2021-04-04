FROM python:3.8.2
ENV PYTHONUNBUFFERED 1
RUN mkdir /migration_tool
WORKDIR /migration_tool
COPY . /migration_tool
RUN apt update
RUN python3 -m pip install --upgrade pip
RUN apt install -y python3-setuptools
RUN pip3 install --upgrade setuptools
RUN pip3 install -r requirements.txt
