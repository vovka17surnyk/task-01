pipeline {
	environment {
	  imagename = "vovka7surnyk/task-01-github"
	  registryCredential = 'dockerhub_id'
	  dockerImage = ''
	}
	agent {label 'agent2'}
	stages {
		
	   stage("Git clone"){
		  
		  steps{
			  checkout([$class: 'GitSCM',
				  branches: [[name: "$REF"]],
				  userRemoteConfigs: [[url: 'https://github.com/vovka17surnyk/task-01-github.git']]])
  
		  }}
	  
	  
	  stage("Test") {
		  tools {go '1.19.2'}
		  steps{
			  
			  sh "gofmt -e 2main.go"
		 }}
		 
		 
	  stage('Building image') {
		steps{
		  script {
			dockerImage = docker.build imagename
		  }
		}
	  }
	  stage('Deploy Image') {
		steps{
		  script {
			docker.withRegistry( '', registryCredential ) {
			  dockerImage.push("$BUILD_NUMBER")
			   dockerImage.push('latest')
  
			}
		  }
		}
	  }
	  stage('Remove Unused docker image') {
		steps{
		  sh "docker rmi $imagename:$BUILD_NUMBER"
		   sh "docker rmi $imagename:latest"
  
		}
	  }
	}
   
	post {
		  success {
			  slackSend color: 'good',
						message: "The pipeline ${currentBuild.fullDisplayName} completed successfully.",
						baseUrl: "https://itoutposts.slack.com/services/hooks/jenkins-ci",
						channel: "#devops-internship-alerts",
						tokenCredentialId: 'SLACK_TOKEN'
		  }
		  failure {
			   slackSend color: '#FA2C00',
						message: "${custom_msg()}",
						baseUrl: "https://itoutposts.slack.com/services/hooks/jenkins-ci",
						channel: "#devops-internship-alerts",
						tokenCredentialId: 'SLACK_TOKEN'
		  }
	  }
  }
  
  def custom_msg()
  {
	def JENKINS_URL= "https://jenkins.itoutposts.com"
	def JOB_NAME = env.JOB_NAME
	def BUILD_ID= env.BUILD_ID
	def JENKINS_LOG= " FAILED: Job [${env.JOB_NAME}] Logs path: ${JENKINS_URL}/job/${JOB_NAME}/${BUILD_ID}/consoleText"
	return JENKINS_LOG
  STARTED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})
  }