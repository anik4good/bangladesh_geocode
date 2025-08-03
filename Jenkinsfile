pipeline {
    agent any

    environment {
        COMPOSE_FILE = 'docker-compose.yml'
    }

    stages {
        stage('Clone Repository') {
            steps {
                checkout scm
            }
        }

        stage('Build & Deploy') {
            steps {
                script {
                    // Stop existing containers
                    sh 'docker-compose down || true'

                    // Pull latest images (optional)
                    sh 'docker-compose pull || true'

                    // Rebuild and start in detached mode
                    sh 'docker-compose up --build -d'
                }
            }
        }
    }

    post {
        success {
            echo '✅ Deployment successful.'
        }
        failure {
            echo '❌ Deployment failed.'
        }
    }
}
