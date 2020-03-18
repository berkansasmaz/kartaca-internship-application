FROM node:12.15
WORKDIR /
COPY package*.json ./
RUN npm install
RUN npm cache --force clean && npm install --force node-sass 


COPY . .
EXPOSE 8080
CMD [ "npm", "run", "dev" ]