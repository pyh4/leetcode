pipeline {
    agent { docker { image 'golang:1.17.5-alpine' } }
    stages {
        stage('test') {
            steps {
                sh 'go version'
                sh 'go test'
                sh 'echo "test complete"'
            }
        }
        stage('build') {
            steps {
                sh 'go build'
                sh 'chmod 666 main'
                sh 'echo "build complete"'
            }
        }
        stage('deploy') {
            steps {
                sh "./leetcode"
                sh 'echo "deploy complete"'
            }
        }
    }
}