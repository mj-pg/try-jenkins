pipeline {
	agent any
	stages{
		stage('Build')	{
			echo 'Building'
			go build .
			echo 'Test must print Hello World'
			./try-jenkins
		}
		stage('Test')	{
			echo 'Testing'
			go test
		}
	}
}
