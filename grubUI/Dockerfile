# Base image
FROM node:13.12.0-alpine
# Set working directory
WORKDIR /app
# Add node moules to path
ENV PATH /app/node_modules/.bin:$PATH
# Copy dependency file
COPY yarn.lock ./
# Install app dependencies
RUN npm install react-scripts -g --silent
RUN yarn install
# Copy source files
COPY . ./
# Start React app
CMD ["yarn", "start"]