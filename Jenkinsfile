pipeline {
    agent any

    environment {
        IMAGE_NAME = "anik4good/bangladesh_geocode"
        DOCKER_CREDENTIALS_ID = "docker-hub-credentials"
    }

    stages {
        stage('Clone Repository') {
            steps {
                git 'https://github.com/anik4good/bangladesh_geocode.git'
            }
        }

        stage('Read Version') {
            steps {
                script {
                    VERSION = readFile('VERSION').trim()
                }
            }
        }

        stage('Build Docker Images') {
            steps {
                script {
                    sh "docker build -t $IMAGE_NAME:$VERSION ."

                }
            }
        }
        stage('Login to Docker Hub & Push') {
            steps {
                withCredentials([usernamePassword(credentialsId: "${DOCKER_CREDENTIALS_ID}", usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    sh """
                        echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
                        docker push $IMAGE_NAME:$VERSION
                        docker logout
                    """
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
