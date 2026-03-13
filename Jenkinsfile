pipeline {
    agent any

    environment {
        APP_NAME = "my-git-service"
        VERSION = "1.9.2"
        IMAGE_TAG = "${APP_NAME}:${VERSION}"
    }

    stages {
        stage('Build') {
            steps {
                sh 'docker build --no-cache -t ${IMAGE_TAG} .'
                sh 'docker save -o ${APP_NAME}_${VERSION}.tar.gz ${IMAGE_TAG}'
            }
        }
    }

    post {
        always {
            sh 'docker image prune -f'
            archiveArtifacts artifacts: '**/*.tar.gz'
            sh 'rm -rf ./build'
            sh 'rm -rf ./src'
            sh 'rm -rf ./.dockerignore'
            sh 'rm -rf ./Dockerfile'
            sh 'rm -rf ./go.mod'
            sh 'rm -rf ./go.sum'
        }
    }
}