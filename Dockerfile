FROM node:14

ENV MODE=prod

WORKDIR /app
COPY package.json /app

RUN npm install

COPY . /app

CMD npm run start:${MODE}

EXPOSE 8080