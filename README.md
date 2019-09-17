GCP Vision Api
- Set environment variables:
  - export GOOGLE_APPLICATION_CREDENTIALS="[FILE_NAME].json"

AWS medical comprehend Api
- Set aws credentials into a .aws folder

Run the project on port 8081:
- docker build -t appmed .
- docker run -p 8081:8081 appmed

Identifies medical boxes that are writen in spanish:
 - Loperamida
 - Ibuprofeno
 
Note: Somethimes identifies handwritten letter.