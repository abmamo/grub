# Base image
FROM node:13.12.0-alpine
# Set working directory
WORKDIR /app
# Add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH
# Copy dependency file
COPY package.json ./
# Install app dependencies
RUN npm install --silent
RUN npm install react-scripts@3.4.0 -g --silent
# Copy source files
COPY . ./
# Start React app
CMD ["npm", "start"]