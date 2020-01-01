FROM node:12-stretch as builder

WORKDIR /app

RUN apt-get update && apt-get upgrade -y

COPY ["package.json", "package-lock.json", "./"]

RUN npm ci

COPY . .

RUN npm run build

FROM nginx:1.17

RUN apt-get update && apt-get upgrade -y
COPY --from=builder /app/build              /usr/share/nginx/html
COPY --from=builder /app/nginx-ssl.conf     /etc/nginx/conf.d/app.conf
