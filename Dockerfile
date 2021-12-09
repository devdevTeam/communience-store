FROM node:16-alpine

LABEL maintainer=shunsuke

ENV HOST 0.0.0.0

WORKDIR /app
# ADD . /app

# デプロイ
COPY . /app

RUN npm install

# CMD PORT=8080 npm run dev
CMD npm run dev
