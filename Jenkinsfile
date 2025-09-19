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

        stage('Build Docker Image') {
            steps {
                script {
                    sh "docker build -t $IMAGE_NAME:latest ."
                }
            }
        }

        stage('Login to Docker Hub & Push') {
            steps {
                withCredentials([usernamePassword(credentialsId: "${DOCKER_CREDENTIALS_ID}", usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                    sh """
                        echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
                        docker push $IMAGE_NAME:latest
                        docker logout
                    """
                }
            }
        }

        stage('Deploy to K3s') {
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
        }

        stage('Verify Deployment') {
            steps {
                withCredentials([file(credentialsId: 'k3s-kubeconfig', variable: 'KUBECONFIG_FILE')]) {
                    sh """
                        export KUBECONFIG=$KUBECONFIG_FILE
                        echo "⏳ Waiting for pods to become ready..."
                        kubectl rollout status deployment/app -n default --timeout=120s

                        echo "📦 Pods after deploy:"
                        kubectl get pods -n default -l io.kompose.service=app -o wide

                        echo "🌐 Testing API inside cluster..."
                        kubectl run curl-test --rm -i --restart=Never --image=curlimages/curl:latest -n default \
                          -- curl -s http://app:1552/api/divisions | head -n 5
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
