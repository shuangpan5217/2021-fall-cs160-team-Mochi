FROM node:alpine3.11
WORKDIR /frontend
COPY frontend/package.json /frontend
RUN npm install --production
COPY frontend /frontend
CMD [ "npm", "start" ]