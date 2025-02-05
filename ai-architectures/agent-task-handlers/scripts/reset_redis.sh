docker stop /my-redis-stack
docker rm /my-redis-stack
docker run -d --name my-redis-stack -p 6379:6379  redis/redis-stack-server:latest