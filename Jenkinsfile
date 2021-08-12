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
      dingtalk (
          robot: '67e78387-1051-4cff-b1fb-06d3d1e24234',
          type: 'MARKDOWN',
          at: [],
          atAll: false,
          title: "success: ${JOB_NAME}",
          text: ["- 成功构建:${JOB_NAME}项目!\n- 分支:${branch}\n- 数据数据初始化:${iDb}\n- 持续时间:${currentBuild.durationString}\n- 任务:#${BUILD_ID}"],
          messageUrl: 'http://www.baidu.com',
          picUrl: '',
          singleTitle: '',
          btns: [],
          hideAvatar: false
      )
    }

    success {
      echo '------success------'
//       dingtalk (
//                 robot: '67e78387-1051-4cff-b1fb-06d3d1e24234',
//                 type:'MARKDOWN',
//                 atAll: false,
//                 title: "success: ${JOB_NAME}",
//                 messageUrl: 'http://www.baidu.com',
//                 text: ["- 成功构建:${JOB_NAME}项目!\n- 分支:${branch}\n- 数据数据初始化:${iDb}\n- 持续时间:${currentBuild.durationString}\n- 任务:#${BUILD_ID}"],
//             )
    }

    failure {
      echo '------failure------'
//       dingtalk (
//           robot: '67e78387-1051-4cff-b1fb-06d3d1e24234',
//           type:'MARKDOWN',
//           atAll: false,
//           title: "success: ${JOB_NAME}",
//           messageUrl: 'http://www.baidu.com',
//           text: ["- 失败构建:${JOB_NAME}项目!\n- 分支:${branch}\n- 数据数据初始化:${iDb}\n- 持续时间:${currentBuild.durationString}\n- 任务:#${BUILD_ID}"],
//       )
    }

    unstable {
      echo '-----unstable-------'
    }

  }
}
