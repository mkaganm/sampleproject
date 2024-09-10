pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "sampleproject:latest"
    }

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/mkaganm/sampleproject.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    sh 'docker build -t $DOCKER_IMAGE .'
                }
            }
        }

        stage('Load Environment Variables') {
            steps {
                script {
                    def envFile = readFile('.env').trim()
                    envFile.split('\n').each { line ->
                        def parts = line.split('=')
                        if (parts.length == 2) {
                            env[parts[0]] = parts[1]
                        }
                    }
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    sh """
                    if [ \$(docker ps -q -f name=sampleproject) ]; then
                        docker stop sampleproject
                        docker rm sampleproject
                    fi
                    """
                    sh "docker run -d --name sampleproject -p ${env.PORT}:${env.PORT} $DOCKER_IMAGE"
                }
            }
        }
    }
}
