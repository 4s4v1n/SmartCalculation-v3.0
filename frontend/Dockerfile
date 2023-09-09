FROM node:latest as build

WORKDIR /root

COPY ./package*.json ./
COPY ./ ./

RUN npm install
RUN npm run build -o app

FROM nginx:alpine as prod

EXPOSE 3000

COPY --from=build /root/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
