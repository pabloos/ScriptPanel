FROM node:8

WORKDIR /usr/src/app

COPY package.json ./

RUN npm i -D --save-dev puppeteer jest jest-cli

CMD sleep 40 ; npm test