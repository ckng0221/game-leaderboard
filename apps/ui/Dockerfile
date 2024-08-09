FROM node:20-alpine as build

ARG VITE_API_BASE_URL 
ENV VITE_API_BASE_URL=$VITE_API_BASE_URL

WORKDIR /usr/app 

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM nginx:alpine as runtime

COPY --from=build /usr/app/dist/ /usr/share/nginx/html

# RUN rm /etc/nginx/nginx.conf /etc/nginx/conf.d/default.conf
RUN rm /etc/nginx/nginx.conf 

COPY nginx.conf /etc/nginx

# docker build . -t game-leaderboard-ui
# docker run --name game-leaderboard-ui -p 80:80 -d game-leaderboard-ui
# docker run -p 80:80 -d game-leaderboard-ui