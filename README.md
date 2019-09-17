# medicalRecognition

This project is about upload photos of medicament boxes and then it retrieves data about the medicament.

Identifies medical boxes that are writen in spanish:
 - Loperamida
 - Ibuprofeno
 
Note: Somethimes identifies handwritten letter.

## Installation
GCP Vision Api
- Create a cred.json and save your GCP credentials.
- Set environment variables:
  ```bash
  export GOOGLE_APPLICATION_CREDENTIALS="cred.json"
  ```
AWS medical comprehend Api
- Set aws credentials file and config file into the .aws folder

## Usage
Run the project on port 8081:
 ```bash
  docker build -t appmed .
  docker run -p 8081:8081 appmed
 ```
## License
[MIT](https://choosealicense.com/licenses/mit/)