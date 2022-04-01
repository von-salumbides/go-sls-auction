pipeline {
    agent any
    stages {
        stage('DEPLOY'){
            steps {
                println('hello')
            }
        }
    }
    post {
        always {
            cleanWs()
        }
    }
}