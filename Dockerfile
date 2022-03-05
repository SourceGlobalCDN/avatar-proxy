FROM node:16-alpine
WORKDIR /home/web/
ENV PORT 9000

LABEL org.opencontainers.image.authors="AH-dark"
LABEL org.opencontainers.image.source="https://github.com/SourceGlobalCDN/gravatar-proxy"

COPY ./ /home/web

EXPOSE 9000

# Update
RUN npm install -g npm@latest

# Install dependencies
RUN chmod -Rf 777 /home/web
RUN cd /home/web/ && npm install

ENTRYPOINT node /home/web/app.js
