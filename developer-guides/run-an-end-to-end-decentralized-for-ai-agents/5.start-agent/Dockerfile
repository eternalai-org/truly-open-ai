# Use a base image with the ability to install specific Node.js version
FROM node:18

# Set the HOME environment variable explicitly
ENV HOME=/root

# Install Node Version Manager (NVM)
RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash

# Set NVM_DIR directly to the location where NVM installs Node.js
ENV NVM_DIR=/root/.nvm

# Install Node.js version 23.3.0 using NVM
RUN . $NVM_DIR/nvm.sh && \
    nvm install 23.3.0 && \
    nvm use 23.3.0 && \
    nvm alias default 23.3.0

# Make sure the correct node version is available in the PATH
# ENV PATH=$PATH:$HOME/.nvm/versions/node/v23.3.0/bin
ENV PATH=$NVM_DIR/versions/node/v23.3.0/bin:$PATH

# Set TERM environment variable to avoid warning
ENV TERM=xterm-256color

# Set the working directory inside the container
WORKDIR /app

COPY . /app

# Get Eliza code 
RUN git clone https://github.com/ai16z/eliza.git

COPY config.json  /app/eliza/agents/config.json

# Change the working directory to the desired folder
WORKDIR /app/eliza

# Install dependencies using pnpm and build the project
RUN npm install -g pnpm
RUN pnpm clear 
RUN pnpm install --no-frozen-lockfile
RUN pnpm build


CMD ["sh", "-c", "pnpm start --characters=/app/eliza/agents/config.json"]
