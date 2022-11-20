pipeline {
    agent any
    environment {
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage("unit-test") {
            steps {
                echo 'UNIT TEST EXECUTION STARTED'
                sh 'go test ./tests/...'
            }
        }
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
                sh 'docker build -f ./build/Dockerfile -t kj/url-short:${BUILD_ID}' .
                sh 'docker images'
            }
        }
    }
}