FROM node:alpine

RUN npm install npm-cli-login -g

ADD build/built-check /opt/resource/check
ADD build/built-out /opt/resource/out
ADD build/built-in /opt/resource/in
