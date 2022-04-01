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
                script {
                    if ( SLS_ACTION == 'deploy' ) {
                    sh "make deploy"
                    } else if ( SLS_ACTION == 'remove' ){
                        sh "make remove"
                    } else {
                        error("Build Failed, ${SLS_ACTION} is not defined")
                    }
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