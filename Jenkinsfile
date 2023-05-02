pipeline {
	agent any
	stages{
		stage('Build')	{
			steps{
				echo 'Building'
				go build .
				echo 'Test must print Hello World'
				./try-jenkins
			}
		}
		stage('Test')	{
			steps {
				echo 'Testing'
				go test
			}
		}
	}
}
