FROM node:alpine

WORKDIR /frontend
RUN mkdir -p /frontend/src
RUN mkdir -p /frontend/public
COPY package*.json /frontend
RUN npm install -g npm@7.7.6
RUN npm install
