pipeline {
  agent {
    node {
      label 'jenkins_dep'
    }

  }
  stages {
    stage('clon') {
      steps {
        echo 'clon'
      }
    }

    stage('test') {
      steps {
        echo 'test'
      }
    }

    stage('build') {
      steps {
        echo 'build'
      }
    }

    stage('push') {
      steps {
        echo 'push'
      }
    }

    stage('dep') {
      steps {
        echo 'dep'
      }
    }

  }
  post {
    always {
      echo '------------'
    }

    success {
      echo '------------'
    }

    failure {
      echo '------------'
    }

    unstable {
      echo '------------'
    }

  }
}