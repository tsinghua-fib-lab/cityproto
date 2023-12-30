FROM node:18-alpine
ENV NODE_ENV=production
WORKDIR /home/node/app

# Install serve
RUN yarn global add serve

# Copy build files
COPY ./doc ./build

EXPOSE 80

CMD serve build -p 80
