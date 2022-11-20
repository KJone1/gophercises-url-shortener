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
                  script: 'go version',
                  returnStatus: true
                  )
                println goExists  
                if (goExists != 0) {
                  sh 'curl -L -o /tmp/go1.19.3.linux-amd64.tar.gz https://go.dev/dl/go1.19.3.linux-amd64.tar.gz'
                  sh 'tar -C /var/jenkins_home/ -xzf /tmp/go1.19.3.linux-amd64.tar.gz'
                  sh 'rm -f /tmp/go1.19.3.linux-amd64.tar.gz'
                  sh 'export PATH=$PATH:/var/jenkins_home/go/bin'
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