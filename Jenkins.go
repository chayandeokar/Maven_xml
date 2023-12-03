What is Jenkins Pipeline?
Jenkins Pipeline is a suite of plugins that allows defining delivery pipelines using a domain-specific language (DSL) integrated into Jenkins.

What is the difference between Declarative and Scripted Pipeline Syntax in Jenkins?
Declarative Pipeline provides a more structured and simplified syntax, while Scripted Pipeline allows for more flexibility using Groovy scripting.

How do you define a Jenkins Pipeline?
Pipelines can be defined either using a Jenkinsfile (text file) in source control or directly within the Jenkins job configuration.

What are Stages in a Jenkins Pipeline?
Stages represent a phase in the pipeline, and each stage can contain multiple steps. They help organize and visualize the flow of the pipeline.

Explain the differences between Jenkinsfile and Jenkins Scripted Pipeline.
Jenkinsfile is a text file containing the definition of a Jenkins Pipeline and is usually stored in source control. Scripted Pipeline uses Groovy scripting syntax directly within Jenkins.

How do you parameterize a Jenkins Pipeline?
You can use the parameters block in Declarative Pipeline or properties block in Scripted Pipeline to define parameters.

What is the purpose of the "agent" directive in Jenkins Pipeline?
The agent directive specifies where the entire Pipeline or a specific stage should be executed. It defines the environment for running the pipeline.

How do you handle credentials in Jenkins Pipeline?
Use the withCredentials block to safely handle credentials within a Pipeline, ensuring sensitive information is not exposed.

Explain the differences between node and docker in the Jenkins Pipeline.
node allocates an executor on a Jenkins agent, while docker runs the entire Pipeline, or a specific stage, inside a Docker container.

What is the purpose of the script block in Jenkins Declarative Pipeline?
The script block allows you to use scripted syntax within a Declarative Pipeline where more advanced scripting is needed.

How do you parallelize stages in Jenkins Pipeline?
The parallel directive allows you to execute multiple stages concurrently. It can be applied at the pipeline or stage level.

Explain Continuous Integration and Continuous Deployment (CI/CD) and how Jenkins supports them.
CI/CD is a software engineering practice that aims to automate the software delivery process. Jenkins provides pipelines to implement CI/CD workflows.
