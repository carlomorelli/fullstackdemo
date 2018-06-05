FROM node:8.11-alpine

# install chrome for protractor tests
#RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add -
#RUN sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list'
#RUN apt-get update && apt-get install -yq google-chrome-stable

RUN mkdir /frontend
WORKDIR /frontend

ENV PATH /frontend/node_modules/.bin:$PATH

ADD . /frontend/ 
RUN npm install 
#RUN ng test

EXPOSE 4200

CMD [ "ng", "serve", "--host", "0.0.0.0" ]
