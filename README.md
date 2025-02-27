# Executus
MP3/Wav web client

Create Docker image
docker build -t executus-app .

Run Docker container
docker run -p 8081:8081 -e ClientPort=:8081 executus-app