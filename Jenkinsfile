env.PROJ_DIR="${JENKINS_HOME}/workspace"        //jenkins  workspace
env.PROJ_DIR_HOSTPATH="/opt/data/jenkinsdata/workspace"
env.PROJ_URL="http://gitlab.oopspy.com:7166/go/vAdmin.git"  // 项目的http地址
//env.PROJ_NAME="*********" // 项目名
env.LANGUAGE="golang" //基础镜像
env.TAGS="1.15-alpine"     //基础镜像TAG
env.HARBOR="registry.cn-beijing.aliyuncs.com/k8s22333"  //镜像仓库的URL
env.INAME="vadmin"  //制作的镜像名(自定义时不要使用大写字母)
//    String tag='latest' //制作的镜像TAG

node {
    withEnv(["GOPATH=$WORKSPACE"]) {
        stage('Init gopath') {
            sh 'mkdir -p $GOPATH/{bin,pkg,src}'  // go运行环境目录
            sh 'mkdir -p $GOPATH/src/$JOB_NAME'
        }
    }
  
  wrap([$class: 'BuildUser']) {
    def user = env.BUILD_USER_ID
    def email = env.BUILD_USER_EMAIL
    println user
  }
  
  //def env = ["JAVA_HOME=${tool 'jdk1.8.0_131'}", "PATH+MAVEN=${tool 'maven_3.1.1'}/bin:${env.JAVA_HOME}/bin", "PATH+GRADLE=${tool 'gradle_4.1'}/bin:${env.JAVA_HOME}/bin" ]

  //def err = null
  //currentBuild.result = "SUCCESS"
}

def InPutCheckBranch(){
    timeout(time:10, unit:'SECONDS') {
        INPUT = input message: "两次输入不一致! 请重新输入你想构建的代码分支",
        ok: "do it",
        // submitter: "zhahuangbo",
        parameters: [
        string(name: 'branch', defaultValue: '', description: 'input your branch'),
        string(name: 'branch_confirm', defaultValue: '', description: 'input your branch again'),
        ]
    }
    env.branch="${INPUT.branch}"
}

class MyChange {
   String author;
   String msg;
}
@NonCPS
def GetChanges(){
    def changeLogSets = currentBuild.changeSets
    for (int i = 0; i < changeLogSets.size(); i++)
    {
        def entries = changeLogSets[0].items
        for (int j = 0; j < entries.length; j++)
        {
            def entry = entries[0]
            def change = new MyChange()
            change.author = entry.author
            change.msg = entry.msg
            return change
        }
    }
}

pipeline {
    agent any
    
    //triggers {
    //    cron('H 4/* 0 0 1-5')
    //}
    
    //tools {
        //工具名称必须在Jenkins 管理Jenkins → 全局工具配置中预配置。
    //    maven 'apache-maven-3.0.1'
    //}
    
    stages {
        
        stage('Example') {
            steps {
                echo 'Hello World'

                script {
                    def browsers = ['chrome', 'firefox']
                    for (int i = 0; i < browsers.size(); ++i) {
                        echo "Testing the ${browsers[i]} browser"
                    }
                }
            }
        }
        
        // in the case of launching a multibranch job you have to specify the branch to build like this
        //stage ('Invoke_pipeline') {
        //    steps {
        //        build job: 'pipeline1/master', parameters: [
        //        string(name: 'param1', value: "value1")
        //    ]
        //    }
        //}
        
        stage('Get code') {
            steps {
            checkout([                      // git repo
                $class: 'GitSCM',
                branches: [[name: '*/master']],
                doGenerateSubmoduleConfigurations: false,
                extensions: [[$class: 'RelativeTargetDirectory', relativeTargetDir: './']],
                submoduleCfg: [],
                userRemoteConfigs: [[
                    credentialsId: 'vugitlab',
                    url: 'http://gitlab.oopspy.com:7166/go/vAdmin.git'
                ]]
            ])
            }
        }
        
        stage('Sonar scan') {
            steps {
                // groovy 脚本
                sh "echo ${WORKSPACE} "
                sh "echo ${JOB_NAME} "
                //sh "echo ${deploy_step} "
                //sh "echo ${env.JOB_NAME} "
                
                // groovy 脚本
                script {
                switch(SONAR_SCAN) {
                case "TRUE":
                    sh "echo \u001B[31m不扫描\u001B[0m ！！！"     
                default:
                    sh "echo 扫描"
                    sh "/var/jenkins_home/sonar-scanner-4.6.0.2311/bin/sonar-scanner \
                        -Dsonar.projectKey=vadmin \
                        -Dsonar.sources=. \
                        -Dsonar.host.url=http://gitlab.oopspy.com:9000 \
                        -Dsonar.login=admin \
                        -Dsonar.password=sftf0423@"
                    //sh "sh  /srv/deploy_scripts/do_sonar.sh ${WORKSPACE} ${JOB_NAME}"
                           
                }
                }
            }

        }
        
        stage('Get commit_msg') {
            steps {
                script {
                    //env.GIT_COMMIT_MSG = sh (script: 'git log -1 --pretty=%B ${GIT_COMMIT}', returnStdout: true).trim()
                    env.commitmsg = sh(script: 'git log -1 --pretty=%B ${GIT_COMMIT}', returnStdout: true).trim()
                    env.tagfull = sh(script: 'git rev-parse HEAD', returnStdout: true).trim()
                    env.tag = sh(script: 'git rev-parse --short HEAD', returnStdout: true).trim()
                    env.GIT_COMMIT_EMAIL = sh(script: "git --no-pager show -s --format=\'%ae\'", returnStdout: true).trim()
                    env.BUILD_FULL = sh (script: "git log -1 --pretty=%B | grep '\\[jenkins-full]'",returnStatus: true) == 0
                }   
            }
        }

        stage('Docker Build') {           //构建镜像,ssh到远程主机执行命令，该命令要用""，Dodckerfile顶格
            steps {
            withEnv(["IMAGE_TAG=${tag}"]) {
            sh '''
                  ssh -i /var/jenkins_home/id_rsa  root@172.19.142.17 "cd $PROJ_DIR_HOSTPATH/$JOB_NAME/;
                  cat << EOF > Dockerfile
FROM ${LANGUAGE}:${TAGS} as builder
WORKDIR /go/src/$JOB_NAME
COPY . .
#ENV GOPROXY https://goproxy.io
ENV GOPROXY https://mirrors.aliyun.com/goproxy/
ENV GO111MODULE on
RUN go build -o $JOB_NAME
FROM alpine:latest as prod
WORKDIR /root/
EXPOSE 8000
COPY --from=builder /go/src/$JOB_NAME/$JOB_NAME .
COPY --from=builder /go/src/$JOB_NAME/config ./config
CMD [\\"./$JOB_NAME\\", \\"server\\", \\"-c\\", \\"config/settings.yml\\", \\"-p\\", \\"8000\\", \\"-m\\", \\"dev\\"]
EOF
            docker build -t ${HARBOR}/${INAME}:$IMAGE_TAG ."
              '''
           }
        }
        }
        
        stage('Image Push') {              //上传镜像到远程镜像仓库
            steps {
            withEnv(["IMAGE_TAG=${tag}"]) {    
            sh '''
            ssh -i /var/jenkins_home/id_rsa  root@172.19.142.17 "docker push ${HARBOR}/${INAME}:$IMAGE_TAG;
            "
            '''
        }
            }
        }


        stage('Measurement Data') {              //效能度量数据收集

            steps {
                sh "echo \u001B[31m不收集效能度量数据\u001B[0m ！！！"
            //sh '''
            //ssh -i /var/jenkins_home/id_rsa  root@172.19.142.17 "cd $PROJ_DIR_HOSTPATH/$JOB_NAME/; sh /opt/data/jenkinsdata/cicdinfo.sh-v5.0 $JOB_NAME"
            //'''
            }
        }

       
       //部署阶段
       
       stage('Image Pull') {              //从远程镜像仓库拉取镜像
            steps {
            withEnv(["IMAGE_TAG=${tag}"]) {
            sh '''
            ssh -i /var/jenkins_home/id_rsa  root@172.19.142.17 "docker pull ${HARBOR}/${INAME}:$IMAGE_TAG"
            '''
       }
            }
       }
       
       stage('Remove old container') {              //删除原有容器
            steps {
            script {
            try{
                sh '''
                ssh -i /var/jenkins_home/id_rsa  root@172.19.142.17 "docker rm -f $JOB_NAME"
                '''
            }catch (error){
            }finally{
                echo "remove old container success"
            }
            }
            }
       }
        
        stage("Image Deploy") {
            steps {
                withEnv(["IMAGE_TAG=${tag}"]) {
                timeout(time: 1, unit: 'MINUTES') {
                script {
                    //env.DEPLOY_ENV = input message: '选择部署的环境', ok: 'deploy',
                    //    parameters: [choice(name: 'DEPLOY_ENV', choices: ['dev', 'test', 'pre', 'prd'], description: '选择部署环境')]

                        //switch("${env.DEPLOY_ENV}"){
                        switch("${environment}"){
                            case 'dev':
                                println('deploy dev env')
                                sh '''
                                ssh -i /var/jenkins_home/id_rsa  root@172.19.142.17 "docker run -d  -p 8000:8000 --name=$JOB_NAME ${HARBOR}/${INAME}:${IMAGE_TAG}"
                                '''
                                echo '记录线上版本'
                                sh "echo > /var/jenkins_home/onlineRevision/${JOB_NAME}.${tagfull}.`date '+%Y%m%d-%H:%M:%y'`.revision"
                                break;

                            case 'test':
                                println('deploy test env')
                                break;

                            case 'pre':
                                println('deploy pre env')
                                break;
                            
                            case 'prd':
                                println('deploy prd env')
                                break;
                            
                            default:
                                println('error env')

                        }
                    }
                }
            }
        }
        }       
    }
    
    post { //stages所有任务执行后触发post
    
        always { //发送钉钉通知
           echo "构建报告已发"
           //echo 'Dingtalk Notification'
           //sh "python3 sonar.py" 
        } 
          
        failure { //构建失败通知
            echo "构建失败"
            //dingTalk accessToken:'（钉钉通知密钥）',imageUrl: '', jenkinsUrl: 'http://go.oopspy.com:7176/', message: '代码部署失败', notifyPeople: 'phone'
        }
 
        success { //构建成功通知
            echo "构建成功"
            //dingTalk accessToken:'（钉钉通知密钥）',imageUrl: '', jenkinsUrl: 'http://go.oopspy.com:7176/', message: '代码部署成功', notifyPeople: 'phone'
        }
    
    }

}