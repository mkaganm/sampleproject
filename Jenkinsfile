pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/kullaniciAdi/projeAdi.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    sh 'docker build -t myapp:latest .'
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
                    if [ \$(docker ps -q -f name=myapp) ]; then
                        docker stop myapp
                        docker rm myapp
                    fi
                    """
                    sh "docker run -d --name myapp -p ${env.PORT}:${env.PORT} myapp:latest"
                }
            }
        }
    }
}
