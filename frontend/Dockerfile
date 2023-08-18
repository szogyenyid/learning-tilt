FROM node:20.5.1-alpine

WORKDIR /src

ADD package.json package.json
RUN npm install -dd

ADD . /src

ENTRYPOINT npm start