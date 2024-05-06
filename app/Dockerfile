FROM node:21.7.1

RUN useradd -ms /bin/sh -u 1001 app
USER app

WORKDIR /usr/server/app

COPY ./package.json ./
RUN yarn install --prod

COPY ./ .

RUN yarn build
ENV NODE_ENV=production

CMD ["yarn","start"]
