pipeline {
  agent {
        kubernetes {
      yaml """
apiVersion: v1
kind: Pod
metadata:
  labels:
    test: build
spec:
  containers:
  - name: maven
    image: registry.cn-hangzhou.aliyuncs.com/haozheyu/maven:ora8u201.mvn3.2.5
    command:
    - cat
    tty: true
    imagePullPolicy: "IfNotPresent"
    volumeMounts:
      - mountPath: "/etc/localtime"
        name: "volume-2"
        readOnly: false
      - mountPath: "/root/.m2/"
        name: "volume-maven-repo"
        readOnly: false
      - mountPath: "/etc/hosts"
        name: "volume-hosts"
        readOnly: false   
    env:
      - name: "LANGUAGE"
        value: "en_US:en"
      - name: "LC_ALL"
        value: "en_US.UTF-8"
      - name: "LANG"
        value: "en_US.UTF-8"
  - name: "docker"
    image: "registry.cn-beijing.aliyuncs.com/citools/docker:19.03.9-git"
    imagePullPolicy: "IfNotPresent"
    tty: true
    volumeMounts:
      - mountPath: "/etc/localtime"
        name: "volume-2"
        readOnly: false
      - mountPath: "/var/run/docker.sock"
        name: "volume-docker"
        readOnly: false
      - mountPath: "/etc/hosts"
        name: "volume-hosts"
        readOnly: false    
        
  - name: "kubectl"
    image: "registry.cn-beijing.aliyuncs.com/citools/kubectl:self-1.17"
    imagePullPolicy: "IfNotPresent"
    tty: true
    volumeMounts:
      - mountPath: "/etc/localtime"
        name: "volume-2"
        readOnly: false
      - mountPath: "/var/run/docker.sock"
        name: "volume-docker"        
      - mountPath: "/mnt/.kube/"
        name: "volume-kubeconfig"
        readOnly: false
      - mountPath: "/etc/hosts"
        name: "volume-hosts"
        readOnly: false    
    
  securityContext: {}
  nodeSelector:
    kubernetes.io/hostname: "192.168.0.3"
  restartPolicy: "Never"
   
  volumes:
    - hostPath:
        path: "/var/run/docker.sock"
      name: "volume-docker"
    - hostPath:
        path: "/usr/share/zoneinfo/Asia/Shanghai"
      name: "volume-2"
    - hostPath:
        path: "/etc/hosts"
      name: "volume-hosts"
    - name: "volume-maven-repo"
      hostPath:
        path: "/opt/m2"
    - name: "volume-kubeconfig"
      secret:
        secretName: "multi-kube-config"
"""
    }
  }
  

  stages {
    stage('run Maven') {
      steps {
        container('maven') {
          sh 'mvn -version'
        }
      }
    }

    stage('docker build') {
      steps {
        container('docker') {
          sh 'docker ps'
        }
        container('docker') {
          sh 'docker info'
        }
      }
    }

    stage('kubectl set image') {
      steps {
        container('kubectl') {
          sh 'kubectl get pod'
        }
      }
    }

    stage('test') {
      steps {
        echo 'test'
      }
    }
  }
  post {
    success {
        echo "-------success---------------"
        dingtalk (
            robot: 'd211837b-0ebb-48e0-b7fe-00fc0a44ae3b',
            type: 'MARKDOWN',
            text: [
                "# ${env.JOB_NAME}\n",
                "${currentBuild.result}\n",
                "第${env.BUILD_ID}次构建\n",
                "构建->${currentBuild.duration}毫秒  \n",
                '',
                '---',
                "[build is a link](${env.BUILD_URL})"
            ]
        )
    }

    failure {
        echo "-------failure1---------------"
        dingtalk (
            robot: 'd211837b-0ebb-48e0-b7fe-00fc0a44ae3b',
            type: 'MARKDOWN',
            text: [
                "# ${env.JOB_NAME}\n",
                "${currentBuild.result}\n",
                "第${env.BUILD_ID}次构建\n",
                "构建->${currentBuild.duration}毫秒  \n",
                '',
                '---',
                "[build is a link](${env.BUILD_URL})"
            ]
        )
    }
  }
}
