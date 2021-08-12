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
      echo '------always------'
    }

    success {
      echo '------success------'
      dingtalk (
                robot: 'alert-msg',
                type:'MARKDOWN',
                atAll: false,
                title: "success: ${JOB_NAME}",
                messageUrl: 'http://www.baidu.com',
                text: ["- 成功构建:${JOB_NAME}项目!\n- 分支:${branch}\n- 数据数据初始化:${iDb}\n- 持续时间:${currentBuild.durationString}\n- 任务:#${BUILD_ID}"],
            )
    }

    failure {
      echo '------failure------'
      dingtalk (
          robot: 'alert-msg',
          type:'MARKDOWN',
          atAll: false,
          title: "success: ${JOB_NAME}",
          messageUrl: 'http://www.baidu.com',
          text: ["- 失败构建:${JOB_NAME}项目!\n- 分支:${branch}\n- 数据数据初始化:${iDb}\n- 持续时间:${currentBuild.durationString}\n- 任务:#${BUILD_ID}"],
      )
    }

    unstable {
      echo '-----unstable-------'
    }

  }
}
