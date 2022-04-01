pipeline {
    agent any
    stages {
        stage('Switch Environment') {
            steps {
                script {
                    switch(DEPLOY_ENV) {
                        case "dev":
                            AWS_ROLE = ""
                        break
                            error("Build Failed for ${DEPLOY_ENV}. No match found.")
                    }
                }
            }
        }
        stage('DEPLOY'){
            steps {
                println('Deploying')
                sh "aws s3 ls"
            }
        }
    }
    post {
        always {
            cleanWs()
        }
    }
}