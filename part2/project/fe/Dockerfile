# syntax=docker/dockerfile:1
FROM node:16.14.0-alpine as builder 
ARG VITE_API_URL=http://localhost:8080
WORKDIR /app
COPY . .
RUN echo $VITE_API_URL
RUN yarn install && yarn run build

FROM nginx:stable-alpine
WORKDIR /var/www/html
EXPOSE 80
COPY --from=builder /app/dist .
COPY ./default.conf /etc/nginx/conf.d/
