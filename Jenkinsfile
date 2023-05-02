pipeline {
	agent any
	stages{
		stage('Build')	{
			steps{
				echo 'Building'
				sh 'go build .'
				echo 'Test must print Hello World'
				sh './try-jenkins'
			}
		}
		stage('Test')	{
			steps {
				echo 'Testing'
				sh 'go test'
			}
		}
	}
}
