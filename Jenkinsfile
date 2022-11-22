pipeline {
    agent any
    environment {
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
        PATH="$PATH:/home/jenkins/go/bin"
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
                if (goExists != 0) {
                  sh 'curl -L -o /tmp/go1.19.3.linux-amd64.tar.gz https://go.dev/dl/go1.19.3.linux-amd64.tar.gz'
                  sh 'tar -C /home/jenkins -xzf /tmp/go1.19.3.linux-amd64.tar.gz'
                  sh 'rm -f /tmp/go1.19.3.linux-amd64.tar.gz'
                  sh 'export PATH=$PATH:/home/jenkins/go/bin'
                }
                blExists = sh(
                  script: 'buildah -v',
                  returnStatus: true
                  )
                if (blExists != 0) {
                  sh 'sudo yum -y install buildah'
                  sh 'buildah -v'
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
                sh "sudo buildah bud -f ./build/Dockerfile -t kj/url-short:${BUILD_ID}"
                sh 'sudo buildah images'
              }
            }
        }
    }
}