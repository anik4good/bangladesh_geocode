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

     tage('Deploy to K3s') {
            steps {
                withCredentials([file(credentialsId: 'k3s-kubeconfig', variable: 'KUBECONFIG_FILE')]) {
                    sh """
                        export KUBECONFIG=$KUBECONFIG_FILE
                        # Apply everything in k3s_deploy
                        kubectl apply -f k3s_deploy/ -n default
                        # Force pods to refresh latest image
                        kubectl rollout restart deployment app -n default
                    """
                }
            }

    post {
        always {
            cleanWs()
        }
    }
}
