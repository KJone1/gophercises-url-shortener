pipeline {
    agent any
    environment {
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
      stage("prep") {
            steps {
              script {
                println 'PREP EXECUTION STARTED'
                goExists = sh(
                  script: 'go -v',
                  returnStatus: true
                  )
                println goExists  
                if (goExists != 0) {
                  sh 'rm -rf /usr/local/go'
                  sh 'curl -L -o /tmp/go1.19.3.linux-amd64.tar.gz https://go.dev/dl/go1.19.3.linux-amd64.tar.gz'
                  sh 'tar -xzf /tmp/go1.19.3.linux-amd64.tar.gz'
                  sh 'rm -f /tmp/go1.19.3.linux-amd64.tar.gz'
                  sh 'export PATH=$PATH:/usr/local/go/bin'
                  sh 'go version'
                }
                pmExists = sh(
                  script: 'podman -v',
                  returnStatus: true
                  )
                println pmExists 
                if (pmExists != 0) {
                  sh 'sudo dnf -y install podman'
                  sh 'podman -v'
                }
              }
            }
        }
        stage("unit-test") {
            steps {
                println 'UNIT TEST EXECUTION STARTED'
                sh 'go test ./tests/...'
            }
        }
        stage("build") {
            steps {
              script {
                println 'BUILD EXECUTION STARTED'
                sh 'go version'
                sh "docker build -f ./build/Dockerfile -t kj/url-short:${BUILD_ID}" .
                sh 'docker images'
              }
            }
        }
    }
}