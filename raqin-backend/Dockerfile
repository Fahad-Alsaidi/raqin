FROM strapi/base
WORKDIR /app
COPY package*.json yarn.lock ./
RUN yarn
COPY . .
EXPOSE 1337
CMD ["yarn", "develop"]