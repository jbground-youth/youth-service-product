#FROM node:lts-alpine as builder
#WORKDIR /data
#COPY package*.json ./
#RUN npm install
#COPY . .
#RUN npm run build
#
#FROM nginx:stable-alpine
#RUN rm -rf /etc/nginx/conf.d/default/conf
#COPY --from=builder /data/nginx/default.conf /etc/nginx/conf.d/default.conf
#
#RUN rm -rf /usr/share/nginx/html/*
#COPY --from=builder /data/dist /usr/share/nginx/html
#
#EXPOSE 8080
#ENTRYPOINT ["nginx", "-g", "daemon off;"]