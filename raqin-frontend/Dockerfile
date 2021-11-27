FROM node:lts-alpine
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn global add @quasar/cli
RUN yarn
COPY . .
EXPOSE 8080
CMD ["quasar", "dev"]