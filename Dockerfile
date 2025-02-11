FROM --platform=linux/amd64 node:22-bookworm-slim
LABEL authors="z0ff"

# Enable pnpm
#ENV PNPM_HOME="/pnpm"
#ENV PATH="$PNPM_HOME:$PATH"
#RUN corepack enable

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
#RUN pnpm install
RUN npm install

# Copy the source code
COPY . .

# Start the application
#CMD ["pnpm", "dev"]
CMD ["npm", "run", "dev"]